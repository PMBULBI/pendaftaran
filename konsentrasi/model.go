package konsentrasi

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.Konsentrasi, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetById(ctx context.Context, id string) (val pmbulbi.Konsentrasi, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_konsenstrasi = ? ", id).
		First(&val).
		Error
	return
}

func (r *Repository) Insert(ctx context.Context, val pmbulbi.Konsentrasi) (err error) {
	err = r.db.
		WithContext(ctx).
		Create(&val).
		Error
	return
}
