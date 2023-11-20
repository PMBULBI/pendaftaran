package jalur

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetAllJalur(conn *gorm.DB, ctx context.Context) (val []pmbulbi.JalurPendaftaran, err error) {
	err = conn.
		WithContext(ctx).
		Find(&val).
		Error
	return
}
