// Description: Wrappers for the errors package.
package errors

import "errors"

// As and Is are wrappers for the errors package. They let us access the As and
// Is functions from the errors package.
var (
	As = errors.As
	Is = errors.Is
	New = errors.New
)

