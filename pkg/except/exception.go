package except

import "net/http"

type InvalidArgument struct{
	Message string
} 
func (e *InvalidArgument) Error() string {
    return e.Message
}

type SystemError struct{
	Message string
}
func (e *SystemError) Error() string {
    return e.Message
}

func Error2StatusCode(err error)(statusCode int){
	switch err.(type){
		case *InvalidArgument:
			return http.StatusUnprocessableEntity

		case *SystemError:
			return 	http.StatusInternalServerError
			
		default:
			return http.StatusBadRequest
	}
}