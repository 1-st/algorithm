package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsNumber(t *testing.T) {
	Convey("basic test", t, func() {
		So(isNumber("1   "),ShouldEqual,true)
		So(isNumber("0"), ShouldEqual, true)
		So(isNumber("0.1"), ShouldEqual, true)
		So(isNumber("abc"), ShouldEqual, false)
		So(isNumber("1 a"), ShouldEqual, false)
		So(isNumber("2e10"), ShouldEqual, true)
		So(isNumber("-90e3"), ShouldEqual, true)
		So(isNumber("1e"), ShouldEqual, false)
		So(isNumber("e3"), ShouldEqual, false)
		So(isNumber("6e-1"), ShouldEqual, true)
		So(isNumber("99e2.5"), ShouldEqual, false)
		So(isNumber("53.5e93"), ShouldEqual, true)
		So(isNumber("--6"), ShouldEqual, false)
		So(isNumber("-+3"), ShouldEqual, false)
		So(isNumber("95a54e53"), ShouldEqual, false)
	})
}
