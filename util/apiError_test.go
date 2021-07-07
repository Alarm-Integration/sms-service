package util

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func Test_Error_Handle_No_Error(t *testing.T) {

	Convey("Given", t, func() {
		errMsg := "this is no error test %s"

		Convey("When", func() {
			err := ErrorHandle(errMsg, nil)

			Convey("Then", func() {
				assert.Nil(t, err)
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
				assert.NotNil(t, err)
				assert.EqualError(t, err, fmt.Sprintf(errMsg, errValue))
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
				assert.NotNil(t, err1)
				assert.EqualError(t, err1, fmt.Sprintf(errMsg, strconv.Itoa(statusCode1)))
				assert.NotNil(t, err2)
				assert.EqualError(t, err2, ErrNotFound.Error())
				assert.NotNil(t, err3)
				assert.EqualError(t, err3, "this is not integer type")
				assert.NotNil(t, err4)
				assert.EqualError(t, err4, "this is not error type")
			})
		})
	})
}
