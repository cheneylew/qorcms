// +build enterprise

package routes

import "github.com/cheneylew/qorcms/config/admin"

func init() {
	Router()
	WildcardRouter.AddHandler(admin.MicroSite)
}
