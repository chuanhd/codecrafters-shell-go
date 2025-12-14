package domains

type ExitRequest struct {
	Code int
}

func (er *ExitRequest) Error() string {
	return "exit requested"
}
