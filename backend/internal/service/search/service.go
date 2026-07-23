package search

import (
	"context"
	"errors"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
	"github.com/megatr0n28/autoparts-pro/backend/internal/provider"
)

type Service struct {
	logger    *zap.Logger
	providers []provider.Provider
}

func New(
	logger *zap.Logger,
	providers ...provider.Provider,
) *Service {

	return &Service{
		logger:    logger,
		providers: providers,
	}
}

func (s *Service) Search(
	ctx context.Context,
	vehicleID uuid.UUID,
	query string,
) ([]dto.PartSearchResponse, error) {

	start := time.Now()

	query =
		strings.TrimSpace(
			query,
		)

	if query == "" {

		return nil,
			errors.New(
				"search query required",
			)
	}

	if vehicleID == uuid.Nil {

		return nil,
			errors.New(
				"invalid vehicle id",
			)
	}

	ctx, cancel :=
		context.WithTimeout(
			ctx,
			5*time.Second,
		)

	defer cancel()

	var (
		results []dto.PartSearchResponse

		mutex = make(chan struct{}, 1)
	)

	group, ctx :=
		errgroup.WithContext(
			ctx,
		)

	for _, p := range s.providers {

		if p == nil {
			continue
		}

		provider := p

		group.Go(func() error {

			providerStart :=
				time.Now()

			select {

			case <-ctx.Done():

				return ctx.Err()

			default:

			}

			parts, err :=
				provider.Search(
					ctx,
					vehicleID,
					query,
				)

			if err != nil {

				s.logger.Warn(
					"provider search failed",
					zap.String(
						"provider",
						provider.Name(),
					),
					zap.String(
						"query",
						query,
					),
					zap.Error(
						err,
					),
				)

				return nil
			}

			mapped :=
				mapProviderParts(
					parts,
				)

			mutex <- struct{}{}

			results =
				append(
					results,
					mapped...,
				)

			<-mutex

			s.logger.Debug(
				"provider search completed",
				zap.String(
					"provider",
					provider.Name(),
				),
				zap.Int(
					"results",
					len(mapped),
				),
				zap.Duration(
					"duration",
					time.Since(
						providerStart,
					),
				),
			)

			return nil
		})
	}

	if err := group.Wait(); err != nil {

		return nil, err
	}

	results =
		dedupe(
			results,
		)

	sort.SliceStable(
		results,
		func(i, j int) bool {

			if results[i].Price ==
				results[j].Price {

				return results[i].PartNumber <
					results[j].PartNumber
			}

			return results[i].Price <
				results[j].Price
		},
	)

	s.logger.Info(
		"parts search completed",
		zap.String(
			"query",
			query,
		),
		zap.String(
			"vehicle_id",
			vehicleID.String(),
		),
		zap.Int(
			"providers",
			len(s.providers),
		),
		zap.Int(
			"results",
			len(results),
		),
		zap.Duration(
			"duration",
			time.Since(start),
		),
	)

	return results, nil
}
