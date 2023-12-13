package tahun_lulusan

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.TahunLulusan, err error) {
	err = r.db.
		WithContext(ctx).
		Order("tahun DESC").
		Find(&val).
		Error
	return
}
