package fakultas

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetMAllFakultas(conn *gorm.DB, ctx context.Context) (val []pmbulbi.Fakultas, err error) {
	err = conn.
		WithContext(ctx).
		Find(&val).
		Error
	return
}
