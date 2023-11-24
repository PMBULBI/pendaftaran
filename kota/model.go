package kota

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.WilayahKota, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetKotaByIdProvinsiNamaKota(ctx context.Context, id_prov string, kota string) (dest []pmbulbi.WilayahKota, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_provinsi LIKE ? AND nama_kota LIKE ?", "%"+id_prov+"%", "%"+kota+"%").
		Find(&dest).
		Error
	return
}
