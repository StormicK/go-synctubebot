package users

import (
	"synctubebot/domain/common"
	"synctubebot/domain/users/events"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int
	Username     string
	Password     string
	Roles        []*UserRole
	domainEvents []common.IBaseEvent
}

func (u *User) ClearEvents() {
	u.domainEvents = []common.IBaseEvent{}
}

func (u *User) GetEvents() []common.IBaseEvent {
	return u.domainEvents
}

func (u *User) AddEvent(event common.IBaseEvent) {
	u.domainEvents = append(u.domainEvents, event)
}

func NewUser(id int, username string, password string) *User {
	var user *User

	if common.IsNullOrEmpty(username) {
		panic(common.IsNullOrEmptyError("username"))
	}

	user = &User{
		Id:       id,
		Username: username,
		Password: password,
		Roles:    []*UserRole{&UserRoleUser},
	}

	user.AddEvent(&events.UserCreated{
		Id:       id,
		Username: username,
	})

	return user
}

func NewAdminUser(id int, username string, password string) *User {
	user := NewUser(id, username, password)
	user.Roles = append(user.Roles, &UserRoleAdmin)

	return user
}

func (u *User) AddUserRole(role *UserRole) {
	if role == nil {
		panic(common.IsNullOrEmptyError("role"))
	}

	for _, roleItem := range u.Roles {
		if roleItem.Name == role.Name {
			panic(common.AlreadyExistRoleError(role.Name))
		}
	}

	u.Roles = append(u.Roles, role)
}

func (u *User) ChangePassword(oldPassword string, newPassword string) {
	bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(oldPassword))
	if oldPassword != u.Password {
		panic(common.InvalidPasswordError())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u.Password = string(hashedPassword)
}
