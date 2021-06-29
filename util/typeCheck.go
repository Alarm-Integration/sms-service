package util

import "errors"

func IsErrorType(param interface{}) error {
	switch param.(type) {
	case error:
		return nil
	default:
		return errors.New("this is not error type")
	}
}

func IsIntegerType(param interface{}) error {
	switch param.(type) {
	case int:
		return nil
	default:
		return errors.New("this is not integer type")
	}
}
