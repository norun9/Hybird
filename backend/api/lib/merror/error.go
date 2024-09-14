package merror

type InternalErr string

func (e InternalErr) Error() string {
	return string(e)
}

var (
	ErrDatabase InternalErr = "ErrDatabase"
)
