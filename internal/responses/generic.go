package responses

type ErrorResponse struct {
	Code       int    `json:"code"`
	Identifier int    `json:"identifier"`
	Error      string `json:"error"`
}

type Arrayed struct {
	Data any `json:"data"`
}

type GenericResponse struct {
	Data any `json:"data"`
}

// Create creates a GenericResponse that tends to indicate that something is happening or the stream is still processing
// and is doing alright, this can also be used to return back general data back to any request.
func Create(data any) *GenericResponse {
	return &GenericResponse{Data: data}
}
