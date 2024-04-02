package custom_errors

import "fmt"

type ErrorHandler struct {
}

func (eh *ErrorHandler) HandleError(err error) {
	fmt.Println(err)
}
