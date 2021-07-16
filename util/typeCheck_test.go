package util

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Error_Type_Check_Success(t *testing.T) {

	Convey("Given", t, func() {
		errValue := errors.New("this is error type")

		Convey("When", func() {
			result := IsErrorType(errValue)

			Convey("Then", func() {
				So(result, ShouldBeNil)
			})
		})
	})
}

func Test_Error_Type_Check_Fail(t *testing.T) {

	Convey("Given", t, func() {
		strValue := "this is string type"
		expectedError := "this is not error type"

		Convey("When", func() {
			result := IsErrorType(strValue)

			Convey("Then", func() {
				So(result, ShouldNotBeNil)
				So(result, ShouldBeError, expectedError)
			})
		})
	})
}

func Test_Integer_Type_Check_Success(t *testing.T) {

	Convey("Given", t, func() {
		intValue := 30020

		Convey("When", func() {
			result := IsIntegerType(intValue)

			Convey("Then", func() {
				So(result, ShouldBeNil)
			})
		})
	})
}

func Test_Integer_Type_Check_Fail(t *testing.T) {

	Convey("Given", t, func() {
		strValue := "this is string type"
		expectedError := "this is not integer type"

		Convey("When", func() {
			result := IsIntegerType(strValue)

			Convey("Then", func() {
				So(result, ShouldNotBeNil)
				So(result, ShouldBeError, expectedError)
			})
		})
	})
}

func Test_String_Type_Check_Success(t *testing.T) {

	Convey("Given", t, func() {
		strValue := "this is string type"

		Convey("When", func() {
			result := IsStringType(strValue)

			Convey("Then", func() {
				So(result, ShouldBeNil)
			})
		})
	})
}

func Test_String_Type_Check_Fail(t *testing.T) {

	Convey("Given", t, func() {
		intValue := 30020
		expectedError := "this is not string type"

		Convey("When", func() {
			result := IsStringType(intValue)

			Convey("Then", func() {
				So(result, ShouldNotBeNil)
				So(result, ShouldBeError, expectedError)
			})
		})
	})
}
