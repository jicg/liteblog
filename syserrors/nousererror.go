package syserrors

type NoUserError struct {
	UnKnowError
}

func (err NoUserError) Error() string {
	return "请登陆系统"
}
func (err NoUserError) Code() int {
	return 1001
}
