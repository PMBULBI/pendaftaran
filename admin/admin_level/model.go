package admin_level

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetAllLevelAdmin(conn *gorm.DB, ctx context.Context) (val []pmbulbi.AdminLevel, err error) {
	err = conn.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func GetOneLevelAdmin(conn *gorm.DB, ctx context.Context, id string) (val pmbulbi.AdminLevel, err error) {
	err = conn.
		WithContext(ctx).
		Where("id_level = ? ", id).
		First(&val).
		Error
	return
}
