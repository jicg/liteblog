package syserrors


type Error404 struct {
	UnKnowError
}

func (err Error404) Error() string {
	return "非法访问"
}
func (err Error404) Code() int {
	return 1002
}
