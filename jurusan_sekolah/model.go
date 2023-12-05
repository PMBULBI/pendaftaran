package jurusan_sekolah

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.JurusanSekolah, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetJurusanByIdJnsSekolah(ctx context.Context, id_jenis_sekolah string) (dest []pmbulbi.JurusanSekolah, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_jenis_sekolah LIKE ?", "%"+id_jenis_sekolah+"%").
		Find(&dest).
		Error
	return
}
