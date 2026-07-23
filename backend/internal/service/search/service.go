package search

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

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

	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)
	defer cancel()

	var (
		wg      sync.WaitGroup
		mutex   sync.Mutex
		results []dto.PartSearchResponse
	)

	for _, p := range s.providers {

		wg.Add(1)

		go func(provider provider.Provider) {

			defer wg.Done()

			providerStart := time.Now()

			items, err := provider.Search(
				ctx,
				vehicleID,
				query,
			)

			if err != nil {

				s.logger.Warn(
					"provider search failed",
					zap.String("provider", provider.Name()),
					zap.String("query", query),
					zap.Error(err),
				)

				return
			}

			s.logger.Debug(
				"provider search completed",
				zap.String("provider", provider.Name()),
				zap.String("query", query),
				zap.Int("results", len(items)),
				zap.Duration("duration", time.Since(providerStart)),
			)

			mapped :=
				mapProviderParts(
					items,
				)

			mutex.Lock()

			results = append(
				results,
				mapped...,
			)

			mutex.Unlock()

		}(p)

	}

	wg.Wait()

	results = dedupe(results)

	sort.Slice(
		results,
		func(i, j int) bool {
			return results[i].Price < results[j].Price
		},
	)

	s.logger.Info(
		"parts search completed",
		zap.String("query", query),
		zap.String("vehicle_id", vehicleID.String()),
		zap.Int("providers", len(s.providers)),
		zap.Int("results", len(results)),
		zap.Duration("duration", time.Since(start)),
	)

	return results, nil
}
