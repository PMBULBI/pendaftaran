package kelurahan

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.WilayahKelurahan, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetKelurahanByIdKecamatan(ctx context.Context, id_kecamatan string) (dest []pmbulbi.WilayahKelurahan, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_kecamatan ILIKE ?", "%"+id_kecamatan+"%").
		Find(&dest).
		Error
	return
}

func (r *Repository) GetKelurahanByIdKecamatanNamaKelurahan(ctx context.Context, id_kecamatan, kelurahan string) (dest []pmbulbi.WilayahKelurahan, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_kecamatan ILIKE ? AND nama_kelurahan ILIKE ?", "%"+id_kecamatan+"%", "%"+kelurahan+"%").
		Find(&dest).
		Error
	return
}
