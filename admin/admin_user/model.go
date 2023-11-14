package admin_user

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetAllAdmin(conn *gorm.DB, ctx context.Context) (val []pmbulbi.Admin, err error) {
	err = conn.
		WithContext(ctx).
		Find(&val).
		Error
	return
}
