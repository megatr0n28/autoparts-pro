package auth

type Role string

const (
	RoleAdmin      Role = "admin"
	RoleEditor     Role = "editor"
	RoleViewer     Role = "viewer"
	RoleLegacyUser Role = "user"
)

func CanModifyOwnData(role string) bool {
	return role == string(RoleAdmin) ||
		role == string(RoleEditor) ||
		role == string(RoleLegacyUser)
}

func HasRole(
	userRole string,
	required Role,
) bool {

	return userRole == string(required)

}
