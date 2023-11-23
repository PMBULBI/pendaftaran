package pendaftar

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.Pendaftaran, err error) {
	err = r.db.
		WithContext(ctx).
		Order("id DESC").
		Find(&val).
		Error
	return
}

func (r *Repository) GetById(ctx context.Context, id string) (dest pmbulbi.Pendaftaran, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id = ? ", id).
		First(&dest).
		Error
	return
}

func (r *Repository) Insert(ctx context.Context, val pmbulbi.Pendaftaran) (err error) {
	err = r.db.
		WithContext(ctx).
		Create(&val).
		Error
	return
}
