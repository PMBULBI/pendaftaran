package agama

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.Agama, err error) {
	err = r.db.
		WithContext(ctx).
		Order("agama ASC").
		Find(&val).
		Error
	return
}

func (r *Repository) GetById(ctx context.Context, id string) (val pmbulbi.Agama, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id = ? ", id).
		First(&val).
		Error
	return
}

func (r *Repository) Insert(ctx context.Context, val pmbulbi.Agama) (err error) {
	err = r.db.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func (r *Repository) Update(ctx context.Context, id string, val pmbulbi.Agama) (err error) {
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
		Delete(&pmbulbi.Agama{}).
		Error
	return
}
