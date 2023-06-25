package entryPoints

const (
	// ErrPermissionDenied is the error message that should be sent
	// when the user does not have enough permissions to perform the action.
	ErrPermissionDenied = "Permission denied"

	// ErrInvalidUserId is the error message that should be sent when
	// the target's user id is invalid. (contains invalid characters)
	ErrInvalidUserId = "Invalid user-id provided"

	ErrPackNotFound = "User not found"

	// ErrRestoredOnly is the error message that should be sent when user is
	// trying to use a method which can only be used on a target user with `Restored`
	// status, but the target user doesn't have that status currently.
	ErrRestoredOnly = "This feature can only be used on users with Restored status"

	// ErrUserNotRegistered is the error message that should be sent when the
	// target's user is not considered as a registered user.
	ErrUserNotRegistered = "User not registered"

	// ErrNoData is the error message that should be sent when the
	// inspector hasn't provided any data for us.
	ErrNoData = "No raw data provided"

	// ErrInvalidToken is the error message that should be sent when the
	// current user's token is invalid.
	ErrInvalidToken = "Invalid token"

	// ErrInternalServerError is the error message that should be sent when there is nothing
	// wrong from client-side; rather something unexpected had happened on
	// server-side.
	ErrInternalServerError = "Internal server error; incidents have been reported."
)
