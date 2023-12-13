package jalur

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.JalurPendaftaran, err error) {
	err = r.db.
		WithContext(ctx).
		Where("status = ?", "aktif").
		Order("nama_jalur ASC").
		Find(&val).
		Error
	return
}

func (r *Repository) GetById(ctx context.Context, id string) (val pmbulbi.JalurPendaftaran, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_jalur = ? ", id).
		First(&val).
		Error
	return
}

func (r *Repository) Insert(ctx context.Context, val pmbulbi.JalurPendaftaran) (err error) {
	err = r.db.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func (r *Repository) Update(ctx context.Context, id string, val pmbulbi.JalurPendaftaran) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_jalur = ?", id).
		Updates(&val).
		Error
	return
}

func (r *Repository) Delete(ctx context.Context, id string) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_jalur = ?", id).
		Delete(&pmbulbi.JalurPendaftaran{}).
		Error
	return
}
