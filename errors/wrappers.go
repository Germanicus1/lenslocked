// Package errors provides wrappers for functions in the std lib errors package.
package errors

import "errors"

// As is a variable that gives access to the As function in the std lib errors
// package.
var As = errors.As

// Is is a variable that gives access to the Is function in the std lib errors
// package.
var Is = errors.Is

