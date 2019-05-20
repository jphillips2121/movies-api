package controllers

import (
	"testing"

	"github.com/gorilla/mux"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRegister(t *testing.T) {
	Convey("Register routes", t, func() {
		router := mux.NewRouter()
		Register(router)
		So(router.GetRoute("get-movies"), ShouldNotBeNil)
		So(router.GetRoute("get-movie"), ShouldNotBeNil)
		So(router.GetRoute("get-most-comments"), ShouldNotBeNil)
		So(router.GetRoute("get-most-likes"), ShouldNotBeNil)
	})
}
