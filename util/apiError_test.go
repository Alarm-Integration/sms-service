package util

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error_Handle_No_Error(t *testing.T) {

	// given
	errMsg := "this is no error test %s"

	// when
	err := ErrorHandle(errMsg, nil)

	// then
	assert.Nil(t, err)
}

func Test_Error_Handle_With_Error(t *testing.T) {

	// given
	errMsg := "this is with error %s"
	errValue := errors.New("test")

	// when
	err := ErrorHandle(errMsg, errValue)

	// then
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(errMsg, errValue))
}

func Test_Error_Handle_With_Status_Code(t *testing.T) {

	// given
	errMsg := "this is with status code %s"
	statusCode1 := 400
	statusCode2 := 404

	// when
	err1 := ErrorHandle(errMsg, nil, statusCode1)
	err2 := ErrorHandle(errMsg, nil, statusCode2)

	// then
	assert.NotNil(t, err1)
	assert.EqualError(t, err1, fmt.Sprintf(errMsg, strconv.Itoa(statusCode1)))
	assert.NotNil(t, err2)
	assert.EqualError(t, err2, ErrNotFound.Error())
}
