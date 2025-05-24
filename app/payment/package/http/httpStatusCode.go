package response

const (
	// 2xx Success
	Success   = 2000
	Created   = 2010
	Accepted  = 2020
	NoContent = 2040

	// 4xx Client errors
	ErrorBadRequest           = 4000
	ErrorUnauthorized         = 4010
	ErrorForbidden            = 4030
	ErrorNotFound             = 4040
	ErrorMethodNotAllowed     = 4050
	ErrorRequestTimeout       = 4080
	ErrorConflict             = 4090
	ErrorPayloadTooLarge      = 4130
	ErrorURITooLong           = 4140
	ErrorUnsupportedMediaType = 4150
	ErrorTooManyRequests      = 4290
	ErrorInvalidParams        = 4001

	// 5xx Server errors
	ErrorInternalServer     = 5000
	ErrorNotImplemented     = 5010
	ErrorBadGateway         = 5020
	ErrorServiceUnavailable = 5030
	ErrorGatewayTimeout     = 5040
)

var msg = map[int]string{
	// 2xx Success
	Success:   "Success",
	Created:   "Created",
	Accepted:  "Accepted",
	NoContent: "No Content",

	// 4xx Client errors
	ErrorBadRequest:           "Bad Request",
	ErrorUnauthorized:         "Unauthorized",
	ErrorForbidden:            "Forbidden",
	ErrorNotFound:             "Not Found",
	ErrorMethodNotAllowed:     "Method Not Allowed",
	ErrorRequestTimeout:       "Request Timeout",
	ErrorConflict:             "Conflict",
	ErrorPayloadTooLarge:      "Payload Too Large",
	ErrorURITooLong:           "URI Too Long",
	ErrorUnsupportedMediaType: "Unsupported Media Type",
	ErrorTooManyRequests:      "Too Many Requests",
	ErrorInvalidParams:        "Invalid Parameters",

	// 5xx Server errors
	ErrorInternalServer:     "Internal Server Error",
	ErrorNotImplemented:     "Not Implemented",
	ErrorBadGateway:         "Bad Gateway",
	ErrorServiceUnavailable: "Service Unavailable",
	ErrorGatewayTimeout:     "github.com/bhtoan2204/gateway Timeout",
}
