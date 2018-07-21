// +build enterprise

package migrations

import (
	"github.com/cheneylew/qorcms/config/admin"
)

func init() {
	AutoMigrate(&admin.QorMicroSite{})
}
