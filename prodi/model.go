package prodi

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.ProgramStudi, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetById(ctx context.Context, id string) (val pmbulbi.ProgramStudi, err error) {
	err = r.db.
		WithContext(ctx).
		Where("kode_program_studi = ? ", id).
		First(&val).
		Error
	return
}

func (r *Repository) Insert(ctx context.Context, val pmbulbi.ProgramStudi) (err error) {
	err = r.db.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func (r *Repository) Update(ctx context.Context, id string, val pmbulbi.ProgramStudi) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("kode_program_studi = ?", id).
		Updates(&val).
		Error
	return
}

func DelProdi(conn *gorm.DB, ctx context.Context, id string) (err error) {
	err = conn.
		WithContext(ctx).
		Where("kode_program_studi = ?", id).
		Delete(&pmbulbi.ProgramStudi{}).
		Error
	return
}
