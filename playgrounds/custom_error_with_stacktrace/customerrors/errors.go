package customerrors

var (
	ErrUnknown = newCustomError(
		"UNKNOWN_ERROR",
		"an unknown error occurred",
	)
)
