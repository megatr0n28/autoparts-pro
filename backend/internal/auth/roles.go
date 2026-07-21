package auth

type Role string

const (
	RoleAdmin Role = "admin"

	RoleManager Role = "manager"

	RoleUser Role = "user"
)

func HasRole(
	userRole string,
	required Role,
) bool {

	return userRole == string(required)

}
