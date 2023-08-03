package presenter

type DefaultResponse struct {
	HTTPCode int `json:"-"`
	Status   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}
