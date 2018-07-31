package syserrors

type HasPraiseError struct {
	UnKnowError
}

func (err HasPraiseError) Error() string {
	return "您已经点过赞!"
}
func (err HasPraiseError) Code() int {
	return 4444
}
