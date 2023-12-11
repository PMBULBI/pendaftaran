package jalur_tahun

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.JalurTahun, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}
