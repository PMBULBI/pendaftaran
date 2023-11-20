package biodata

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func InsertBioJalur(conn *gorm.DB, ctx context.Context, val pmbulbi.BiodataJalur) (err error) {
	err = conn.
		WithContext(ctx).
		Create(&val).
		Error
	return
}
