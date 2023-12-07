package provinsi

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.WilayahProvinsi, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetByNama(ctx context.Context, prov string) (val []pmbulbi.WilayahProvinsi, err error) {
	err = r.db.
		WithContext(ctx).
		Where("nama_provinsi ILIKE ?", "%"+prov+"%").
		Find(&val).
		Error
	return
}
