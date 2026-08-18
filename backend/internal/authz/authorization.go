package authz

import (
	"fmt"

	"github.com/google/uuid"
)

// -----------------------------------------------------------------------------
// OpenFGA Types
// -----------------------------------------------------------------------------

const (
	TypeUser       = "user"
	TypeCustomer   = "customer"
	TypeVehicle    = "vehicle"
	TypeInvoice    = "invoice"
	TypePartSearch = "part_search"
)

// -----------------------------------------------------------------------------
// OpenFGA Relations
// -----------------------------------------------------------------------------

const (
	RelationOwner = "owner"
)

// -----------------------------------------------------------------------------
// Object Builders
// -----------------------------------------------------------------------------

func UserObject(
	userID uuid.UUID,
) string {

	return fmt.Sprintf(
		"%s:%s",
		TypeUser,
		userID.String(),
	)
}

func CustomerObject(
	customerID uuid.UUID,
) string {

	return fmt.Sprintf(
		"%s:%s",
		TypeCustomer,
		customerID.String(),
	)
}

func VehicleObject(
	vehicleID uuid.UUID,
) string {

	return fmt.Sprintf(
		"%s:%s",
		TypeVehicle,
		vehicleID.String(),
	)
}

func InvoiceObject(
	invoiceID uuid.UUID,
) string {

	return fmt.Sprintf(
		"%s:%s",
		TypeInvoice,
		invoiceID.String(),
	)
}

func PartSearchObject(
	searchID uuid.UUID,
) string {

	return fmt.Sprintf(
		"%s:%s",
		TypePartSearch,
		searchID.String(),
	)
}

// -----------------------------------------------------------------------------
// Authorization Tuple Helpers
// -----------------------------------------------------------------------------

// OwnerTuple returns the OpenFGA tuple representing ownership.
//
// Example:
//
//	user:54f38ddd-5a39-4bae-8408-29bf42b002d6
//	owner
//	vehicle:3b0449e5-e808-4d25-90a4-eb5ee9031768
func OwnerTuple(
	userID uuid.UUID,
	objectType string,
	objectID uuid.UUID,
) (
	user string,
	relation string,
	object string,
) {

	return UserObject(userID),
		RelationOwner,
		fmt.Sprintf(
			"%s:%s",
			objectType,
			objectID.String(),
		)
}

// VehicleOwnerTuple returns the ownership tuple
// for a vehicle.
func VehicleOwnerTuple(
	userID uuid.UUID,
	vehicleID uuid.UUID,
) (
	user string,
	relation string,
	object string,
) {

	return OwnerTuple(
		userID,
		TypeVehicle,
		vehicleID,
	)
}

// CustomerOwnerTuple returns the ownership tuple
// for a customer.
func CustomerOwnerTuple(
	userID uuid.UUID,
	customerID uuid.UUID,
) (
	user string,
	relation string,
	object string,
) {

	return OwnerTuple(
		userID,
		TypeCustomer,
		customerID,
	)
}

// InvoiceOwnerTuple returns the ownership tuple
// for an invoice.
func InvoiceOwnerTuple(
	userID uuid.UUID,
	invoiceID uuid.UUID,
) (
	user string,
	relation string,
	object string,
) {

	return OwnerTuple(
		userID,
		TypeInvoice,
		invoiceID,
	)
}

// PartSearchRequesterTuple returns the authorization
// tuple for a part-search request.
//
// NOTE:
// Your current OpenFGA model uses "requester" rather
// than "owner" for part_search.
func PartSearchRequesterTuple(
	userID uuid.UUID,
	searchID uuid.UUID,
) (
	user string,
	relation string,
	object string,
) {

	return UserObject(userID),
		"requester",
		PartSearchObject(searchID)
}
