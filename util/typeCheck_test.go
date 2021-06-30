package util_test

import (
	"errors"
	"testing"

	"github.com/GreatLaboratory/go-sms/util"
	"github.com/stretchr/testify/assert"
)

func Test_Error_Type_Check_Success(t *testing.T) {

	// given
	errValue := errors.New("this is error type")

	// when
	result := util.IsErrorType(errValue)

	// then
	assert.Nil(t, result)

}

func Test_Error_Type_Check_Fail(t *testing.T) {

	// given
	strValue := "this is string type"
	expectedError := "this is not error type"

	// when
	result := util.IsErrorType(strValue)

	// then
	assert.NotNil(t, result)
	assert.EqualError(t, result, expectedError)

}

func Test_Integer_Type_Check_Success(t *testing.T) {

	// given
	intValue := 30020

	// when
	result := util.IsIntegerType(intValue)

	// then
	assert.Nil(t, result)

}

func Test_Integer_Type_Check_Fail(t *testing.T) {

	// given
	strValue := "this is string type"
	expectedError := "this is not integer type"

	// when
	result := util.IsIntegerType(strValue)

	// then
	assert.NotNil(t, result)
	assert.EqualError(t, result, expectedError)

}

func Test_String_Type_Check_Success(t *testing.T) {

	// given
	strValue := "this is string type"

	// when
	result := util.IsStringType(strValue)

	// then
	assert.Nil(t, result)

}

func Test_String_Type_Check_Fail(t *testing.T) {

	// given
	intValue := 30020
	expectedError := "this is not string type"

	// when
	result := util.IsStringType(intValue)

	// then
	assert.NotNil(t, result)
	assert.EqualError(t, result, expectedError)

}
