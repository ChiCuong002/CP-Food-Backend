package response

const (
	SuccessCode = 20001 // Success
	ErrorCodeParamInvalid = 20003 // Parameter is invalid
	ErrorCodeInvalidToken = 30001 // Token is invalid
	ErrorInternalServer = 50001 // Internal server error
	ErrorUnauthorized = 40101 // Unauthorized
	ErrorNoAuthorizationHeader = 40102 // No authorization
)

var message = map[int]string{
	SuccessCode: "Success",
	ErrorCodeParamInvalid: "Parameter is invalid",
	ErrorCodeInvalidToken: "Token is invalid",
	ErrorInternalServer: "Internal server error",
	ErrorUnauthorized: "Unauthorized",
	ErrorNoAuthorizationHeader: "No Authorization Header Provided",
}

