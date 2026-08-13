package authz

import (
	"context"
	"fmt"

	"github.com/openfga/go-sdk/client"
)

type Client struct {
	fga *client.OpenFgaClient
}

type Config struct {
	APIURL               string
	StoreID              string
	AuthorizationModelID string
}

func New(cfg Config) (*Client, error) {

	if cfg.APIURL == "" {
		return nil, fmt.Errorf(
			"OpenFGA API URL is required",
		)
	}

	if cfg.StoreID == "" {
		return nil, fmt.Errorf(
			"OpenFGA store ID is required",
		)
	}

	if cfg.AuthorizationModelID == "" {
		return nil, fmt.Errorf(
			"OpenFGA authorization model ID is required",
		)
	}

	fgaClient, err := client.NewSdkClient(
		&client.ClientConfiguration{
			ApiUrl:               cfg.APIURL,
			StoreId:              cfg.StoreID,
			AuthorizationModelId: cfg.AuthorizationModelID,
		},
	)

	if err != nil {
		return nil, fmt.Errorf(
			"create OpenFGA client: %w",
			err,
		)
	}

	return &Client{
		fga: fgaClient,
	}, nil
}

func (c *Client) Check(
	ctx context.Context,
	user string,
	relation string,
	object string,
) (bool, error) {

	body := client.ClientCheckRequest{
		User:     user,
		Relation: relation,
		Object:   object,
	}

	response, err := c.fga.
		Check(ctx).
		Body(body).
		Execute()

	if err != nil {
		return false, fmt.Errorf(
			"OpenFGA check: %w",
			err,
		)
	}

	if response == nil ||
		response.Allowed == nil {

		return false, nil
	}

	return *response.Allowed, nil
}

func (c *Client) CheckConnection(
	ctx context.Context,
) error {

	_, err := c.fga.
		ListStores(ctx).
		Execute()

	if err != nil {
		return fmt.Errorf(
			"OpenFGA connection failed: %w",
			err,
		)
	}

	return nil
}
