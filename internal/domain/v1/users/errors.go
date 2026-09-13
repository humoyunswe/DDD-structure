package users

type ErrNotFound struct{}

func (ErrNotFound) Error() string { return "USER_NOT_FOUND" }

type ErrAlreadyExists struct{}

func (ErrAlreadyExists) Error() string { return "USER_ALREADY_EXISTS" }

type ErrInvalidInput struct {
	Message string
}

func (e ErrInvalidInput) Error() string {
	if e.Message == "" {
		return "INVALID_INPUT"
	}
	return e.Message
}
