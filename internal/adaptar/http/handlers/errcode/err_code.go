package errcode

const (
	CodeInvalidData      = "INVALID_BODY"
	CodeValidationError  = "VALIDATION_ERROR"
	CodeBadRequest       = "BAD_REQUEST"
	CodeMissingParam     = "MISSING_PARAM"
	CodeUnsupportedMedia = "UNSUPPORTED_MEDIA_TYPE"

	CodeUnauthorized       = "UNAUTHORIZED"
	CodeForbidden          = "FORBIDDEN"
	CodeTokenExpired       = "TOKEN_EXPIRED"
	CodeInvalidCredentials = "INVALID_CREDENTIALS"

	CodeNotFound = "NOT_FOUND"
	CodeIsExists = "IS_EXISTS"
	CodeConflict = "CONFLICT"

	CodeInternalError      = "INTERNAL_ERROR"
	CodeDBError            = "DB_ERROR"
	CodeConfigError        = "CONFIG_ERROR"
	CodeServiceUnavailable = "SERVICE_UNAVAILABLE"
	CodeTimeout            = "TIMEOUT"
	CodeNetworkError       = "NETWORK_ERROR"
)
