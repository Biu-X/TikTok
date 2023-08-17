package response

type Code int

const (
	OK Code = iota
	Error
	UnknownError
)

func Msg(code Code) string {
	switch code {
	case OK:
		return "success"
	case Error:
		return "error"
	case UnknownError:
		return "unknown error"
	default:
		return "unknown error"
	}
}
