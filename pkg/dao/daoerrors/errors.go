package daoerrors

import "errors"

// ErrUserNotFound is returned when a user is not found in the DAO.
var ErrUserNotFound = errors.New("user not found")

// ErrPhoneBookNotFound is returned when a phone book is not found in the DAO.
var ErrPhoneBookNotFound = errors.New("phone book not found")
