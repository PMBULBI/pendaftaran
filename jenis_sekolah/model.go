package jenis_sekolah

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.JenisSekolah, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetById(ctx context.Context, id string) (val pmbulbi.JenisSekolah, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_jenis_sekolah = ? ", id).
		First(&val).
		Error
	return
}

func (r *Repository) Insert(ctx context.Context, val pmbulbi.JenisSekolah) (err error) {
	err = r.db.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func (r *Repository) Update(ctx context.Context, id string, val pmbulbi.JenisSekolah) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_jenis_sekolah = ?", id).
		Updates(&val).
		Error
	return
}

func (r *Repository) Delete(ctx context.Context, id string) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_jenis_sekolah = ?", id).
		Delete(&pmbulbi.JenisSekolah{}).
		Error
	return
}
