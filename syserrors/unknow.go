package syserrors

type UnKnowError struct {
	msg    string
	reason error
}

func (err UnKnowError) Error() string {
	if len(err.msg) == 0 {
		return "未知错误"
	} else {
		return err.msg
	}
}
func (err UnKnowError) Code() int {
	return 1000
}

func (err UnKnowError) ReasonError() error {
	return err.reason
}

