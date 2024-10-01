package response

const (
	SuccessCode = 20001 // Success
	ErrorCodeParamInvalid = 20003 // Parameter is invalid
	ErrorCodeInvalidToken = 30001 // Token is invalid
)

var message = map[int]string{
	SuccessCode: "Success",
	ErrorCodeParamInvalid: "Parameter is invalid",
	ErrorCodeInvalidToken: "Token is invalid",
}

