package users

type UserRole struct {
	Name string
}

var (
	UserRoleAdmin = UserRole{Name: "admin"}
	UserRoleUser  = UserRole{Name: "user"}
)
