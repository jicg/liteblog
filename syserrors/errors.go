package syserrors


type Error interface {
	Error() string
	Code() int
	ReasonError() error
}


func NewError(msg string, err2 error) (UnKnowError) {
	err := UnKnowError{}
	err.msg = msg
	err.reason = err2
	return err
}

