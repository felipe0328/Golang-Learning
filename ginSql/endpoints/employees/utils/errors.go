package utils

import "errors"

var (
	ErrMissingEmployeeId   = errors.New("missing employee id")
	ErrMissingEmployeeData = errors.New("missing employee data in payload")
)
