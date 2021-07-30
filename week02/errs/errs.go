package errs

type DataNotFound struct {
	msg string
}

func (d DataNotFound) Error() string {
	d.msg = "数据不存在"
	return d.msg
}