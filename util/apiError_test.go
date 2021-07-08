package util

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Error_Handle_No_Error(t *testing.T) {

	Convey("Given", t, func() {
		errMsg := "this is no error test %s"

		Convey("When", func() {
			err := ErrorHandle(errMsg, nil)

			Convey("Then", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func Test_Error_Handle_With_Error(t *testing.T) {

	Convey("Given", t, func() {
		errMsg := "this is with error %s"
		errValue := errors.New("test")

		Convey("When", func() {
			err := ErrorHandle(errMsg, errValue)

			Convey("Then", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldBeError, fmt.Sprintf(errMsg, errValue))
			})
		})
	})
}

func Test_Error_Handle_With_Status_Code(t *testing.T) {

	Convey("Given", t, func() {
		errMsg := "this is with status code %s"
		statusCode1 := 400
		statusCode2 := 404
		statusCode3 := "this is not integer type"
		err := "this is not error type"

		Convey("When", func() {
			err1 := ErrorHandle(errMsg, nil, statusCode1)
			err2 := ErrorHandle(errMsg, nil, statusCode2)
			err3 := ErrorHandle(errMsg, nil, statusCode3)
			err4 := ErrorHandle(errMsg, err)

			Convey("Then", func() {
				So(err1, ShouldNotBeNil)
				So(err1, ShouldBeError, fmt.Sprintf(errMsg, strconv.Itoa(statusCode1)))
				So(err2, ShouldNotBeNil)
				So(err2, ShouldBeError, ErrNotFound.Error())
				So(err3, ShouldNotBeNil)
				So(err3, ShouldBeError, "this is not integer type")
				So(err4, ShouldNotBeNil)
				So(err4, ShouldBeError, "this is not error type")
			})
		})
	})
}
