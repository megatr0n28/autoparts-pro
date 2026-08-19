package authz

import (
	"context"
	"fmt"
)

type Authorizer struct {
	client *Client
}

func NewAuthorizer(
	client *Client,
) *Authorizer {
	return &Authorizer{
		client: client,
	}
}

func (a *Authorizer) Check(
	ctx context.Context,
	user string,
	relation string,
	object string,
) (bool, error) {

	if a == nil || a.client == nil {
		return false, fmt.Errorf(
			"authorization client is not configured",
		)
	}

	if user == "" {
		return false, fmt.Errorf(
			"authorization user is required",
		)
	}

	if relation == "" {
		return false, fmt.Errorf(
			"authorization relation is required",
		)
	}

	if object == "" {
		return false, fmt.Errorf(
			"authorization object is required",
		)
	}

	return a.client.Check(
		ctx,
		user,
		relation,
		object,
	)
}

func (a *Authorizer) WriteTuple(
	ctx context.Context,
	user string,
	relation string,
	object string,
) error {

	if a == nil || a.client == nil {
		return fmt.Errorf(
			"authorization client is not configured",
		)
	}

	if user == "" {
		return fmt.Errorf(
			"authorization user is required",
		)
	}

	if relation == "" {
		return fmt.Errorf(
			"authorization relation is required",
		)
	}

	if object == "" {
		return fmt.Errorf(
			"authorization object is required",
		)
	}

	return a.client.WriteTuple(
		ctx,
		user,
		relation,
		object,
	)
}

func (a *Authorizer) EnsureTuple(
	ctx context.Context,
	user string,
	relation string,
	object string,
) error {
	allowed, err := a.Check(ctx, user, relation, object)
	if err != nil {
		return err
	}
	if allowed {
		return nil
	}

	if err := a.WriteTuple(ctx, user, relation, object); err != nil {
		allowed, checkErr := a.Check(ctx, user, relation, object)
		if checkErr == nil && allowed {
			return nil
		}
		return err
	}

	return nil
}

func (a *Authorizer) CheckConnection(
	ctx context.Context,
) error {

	if a == nil || a.client == nil {
		return fmt.Errorf(
			"authorization client is not configured",
		)
	}

	return a.client.CheckConnection(ctx)
}
