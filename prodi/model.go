package prodi

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.ProgramStudi, err error) {
	err = r.db.
		WithContext(ctx).
		Order("program_studi ASC").
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

func (r *Repository) Delete(ctx context.Context, id string) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("kode_program_studi = ?", id).
		Delete(&pmbulbi.ProgramStudi{}).
		Error
	return
}
