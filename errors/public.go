package errors

// Public wraps an error with a custom message to make it public.
// It returns a new error that implements the error interface.
// The returned error can be used to provide a public-facing error message
// while still preserving the original error for debugging or logging purposes.
//
// Parameters:
// - err: The original error to wrap.
// - msg: The custom message to associate with the error.
//
// Returns:
// - error: A new error that wraps the original error and provides a public message.
//
// Example:
//   err := errors.New("database connection failed")
//   publicErr := errors.Public(err, "Failed to connect to the database")
//   fmt.Println(publicErr.Error()) // Output: Failed to connect to the database
//   fmt.Println(publicErr.Public()) // Output: Failed to connect to the database
//   fmt.Println(errors.Unwrap(publicErr)) // Output: database connection failed
func Public(err error, msg string) error {
	return publicError{err, msg}
}

type publicError struct {
	err error
	msg string
}

func (e publicError) Error() string {
	return e.err.Error()
}

func (e publicError) Public() string {
	return e.msg
}

func (e publicError) Unwrap() error {
	return e.err
}

