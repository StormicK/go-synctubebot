package common

import (
	"errors"
	"fmt"
)

const PREFIX = "Domain Error: "
const (
	IsNullOrEmptyText    string = PREFIX + "%s is null or empty!"
	AlreadyExistRoleText string = PREFIX + "User already has role %s!"
)

func IsNullOrEmptyError(str string) error {
	return errors.New(fmt.Sprintf(IsNullOrEmptyText, str))
}

func AlreadyExistRoleError(role string) error {
	return errors.New(fmt.Sprintf(AlreadyExistRoleText, role))
}

func InvalidPasswordError() error {
	return errors.New("Invalid password!")
}
