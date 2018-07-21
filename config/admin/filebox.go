package admin

import (
	"github.com/qor/filebox"
	"github.com/qor/roles"

	"github.com/cheneylew/qorcms/config"
	"github.com/cheneylew/qorcms/config/auth"
)

var Filebox *filebox.Filebox

func init() {
	Filebox = filebox.New(config.Root + "/public/downloads")
	Filebox.SetAuth(auth.AdminAuth{})
	dir := Filebox.AccessDir("/")
	dir.SetPermission(roles.Allow(roles.Read, "admin"))
}
