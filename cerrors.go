package cerrors

import "fmt"

type CustomError struct {
	errs      []error
	message   string
	delimiter string `default:", "`
}

func (cerr CustomError) Error() string {
	n := len(cerr.errs)
	if n == 0 {
		return cerr.message
	}
	message := fmt.Sprintf("%s : %s", cerr.message, cerr.errs[0].Error())
	for i := 1; i < n; i++ {
		message = fmt.Sprintf("%s%s%s", message, cerr.delimiter, cerr.errs[i].Error())
	}
	return message
}

func (cerr *CustomError) HasErrors() bool {
	return len(cerr.errs) > 0
}

func (cerr *CustomError) Append(err error) {
	cerr.errs = append(cerr.errs, err)
}

func New(errors ...error) *CustomError {
	cerr := &CustomError{
		errs: make([]error, 0),
	}
	for _, err := range errors {
		cerr.Append(err)
	}
	return cerr
}

func Newf(msg string, args ...interface{}) *CustomError {
	cerr := New()
	cerr.message = fmt.Sprintf(msg, args...)
	return cerr
}
