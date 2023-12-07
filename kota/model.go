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
		Where("id_provinsi ILIKE ? AND nama_kota ILIKE ?", "%"+id_prov+"%", "%"+kota+"%").
		Find(&dest).
		Error
	return
}

func (r *Repository) GetKotaByIdProvinsi(ctx context.Context, id_prov string) (dest []pmbulbi.WilayahKota, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_provinsi ILIKE ?", "%"+id_prov+"%").
		Find(&dest).
		Error
	return
}
