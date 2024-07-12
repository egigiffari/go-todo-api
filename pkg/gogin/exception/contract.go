package exception

type Contract interface {
	Code() int
	Messages() map[string][]string
}

type Exception struct {
	code     int
	messages map[string][]string
}

func (e *Exception) Code() int {
	return e.code
}

func (e *Exception) Messages() map[string][]string {
	return e.messages
}
