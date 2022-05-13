package errors

type IUserError interface {
	error
	UserError() string
}

type UserError struct {
	Err         error
	UserMessage string
}

func (e UserError) Error() string {
	return e.Err.Error()
}

func (e UserError) UserError() string {
	return e.UserMessage
}

func IsUserError(err error) (IUserError, bool) {
	userError, ok := err.(IUserError)
	return userError, ok
}
