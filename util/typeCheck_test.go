package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error_Type_Check_Success(t *testing.T) {

	// given
	errValue := errors.New("this is error type")

	// when
	result := IsErrorType(errValue)

	// then
	assert.Nil(t, result)

}

func Test_Error_Type_Check_Fail(t *testing.T) {

	// given
	strValue := "this is string type"
	expectedError := "this is not error type"

	// when
	result := IsErrorType(strValue)

	// then
	assert.NotNil(t, result)
	assert.EqualError(t, result, expectedError)

}

func Test_Integer_Type_Check_Success(t *testing.T) {

	// given
	intValue := 30020

	// when
	result := IsIntegerType(intValue)

	// then
	assert.Nil(t, result)

}

func Test_Integer_Type_Check_Fail(t *testing.T) {

	// given
	strValue := "this is string type"
	expectedError := "this is not integer type"

	// when
	result := IsIntegerType(strValue)

	// then
	assert.NotNil(t, result)
	assert.EqualError(t, result, expectedError)

}

func Test_String_Type_Check_Success(t *testing.T) {

	// given
	strValue := "this is string type"

	// when
	result := IsStringType(strValue)

	// then
	assert.Nil(t, result)

}

func Test_String_Type_Check_Fail(t *testing.T) {

	// given
	intValue := 30020
	expectedError := "this is not string type"

	// when
	result := IsStringType(intValue)

	// then
	assert.NotNil(t, result)
	assert.EqualError(t, result, expectedError)

}
