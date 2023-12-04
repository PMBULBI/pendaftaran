package kecamatan

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.WilayahKecamatan, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetKecamatanByIdKota(ctx context.Context, id_kota string) (dest []pmbulbi.WilayahKecamatan, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_kota LIKE ?", "%"+id_kota+"%").
		Find(&dest).
		Error
	return
}

func (r *Repository) GetKecamatanByIdKotaNamaKecamatan(ctx context.Context, id_kota, kecamatan string) (dest []pmbulbi.WilayahKecamatan, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_kota LIKE ? AND nama_kecamatan LIKE ?", "%"+id_kota+"%", "%"+kecamatan+"%").
		Find(&dest).
		Error
	return
}
