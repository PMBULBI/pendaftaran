package jalur_prodi

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.JalurProdi, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetProdiByJalur(ctx context.Context, jalur string) (dest []pmbulbi.JalurProdi, err error) {
	err = r.db.
		WithContext(ctx).
		Where("jalur ILIKE ?", "%"+jalur+"%").
		Find(&dest).
		Error
	return
}

func (r *Repository) Insert(ctx context.Context, val pmbulbi.JalurProdi) (err error) {
	err = r.db.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func (r *Repository) Update(ctx context.Context, id string, val pmbulbi.JalurProdi) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id = ?", id).
		Updates(&val).
		Error
	return
}

func (r *Repository) Delete(ctx context.Context, id string) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&pmbulbi.JalurProdi{}).
		Error
	return
}