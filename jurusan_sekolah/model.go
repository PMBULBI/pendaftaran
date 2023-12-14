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

func (r *Repository) GetById(ctx context.Context, id string) (val pmbulbi.JurusanSekolah, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_jurusan = ? ", id).
		First(&val).
		Error
	return
}

func (r *Repository) GetJurusanByIdJnsSekolah(ctx context.Context, id_jenis_sekolah int) (dest []pmbulbi.JurusanSekolah, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_jenis_sekolah = ?", id_jenis_sekolah).
		Find(&dest).
		Error
	return
}

func (r *Repository) Insert(ctx context.Context, val pmbulbi.JurusanSekolah) (err error) {
	err = r.db.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func (r *Repository) Update(ctx context.Context, id string, val pmbulbi.JurusanSekolah) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_jurusan = ?", id).
		Updates(&val).
		Error
	return
}

func (r *Repository) Delete(ctx context.Context, id string) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_jurusan = ?", id).
		Delete(&pmbulbi.JurusanSekolah{}).
		Error
	return
}
