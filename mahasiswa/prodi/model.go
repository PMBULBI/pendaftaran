package prodi

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetMAllProdi(conn *gorm.DB, ctx context.Context) (val []pmbulbi.ProgramStudi, err error) {
	err = conn.
		WithContext(ctx).
		Find(&val).
		Error
	return
}
