package rs

// Response `type` to respond the requests
type Response map[string]interface{}

// NewResponse factory to create a new `Response`
func NewResponse(data interface{}) Response {
	return Response{"data": data}
}

// NewErrorResponse factory to create a new `Response` with an error
func NewErrorResponse(code int, description string, detail interface{}) Response {
	err := map[string]interface{}{
		"code":        code,
		"description": description,
	}

	if detail != nil {
		err["detail"] = detail
	}

	return Response{"error": err}
}
