package util

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var (
	ErrNotFound = errors.New("not found")
)

func ErrorHandle(errMsg string, err interface{}, statusCode ...interface{}) error {
	if err != nil {
		if typeErr := IsErrorType(err); typeErr != nil {
			return typeErr
		}
		return fmt.Errorf(errMsg, err)
	}
	if len(statusCode) != 0 {
		if typeErr := IsIntegerType(statusCode[0]); typeErr != nil {
			return typeErr
		}
		if statusCode[0] != http.StatusOK {
			if statusCode[0] == http.StatusNotFound {
				return ErrNotFound
			}
			res := (statusCode[0]).(int)
			return fmt.Errorf(errMsg, strconv.Itoa(res))
		}
	}
	return nil
}
