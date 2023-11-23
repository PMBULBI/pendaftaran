package fakultas

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.Fakultas, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetMOneFakultas(ctx context.Context, id string) (val pmbulbi.Fakultas, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_fakultas = ? ", id).
		First(&val).
		Error
	return
}

func (r *Repository) InsFakultas(ctx context.Context, val pmbulbi.Fakultas) (err error) {
	err = r.db.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func (r *Repository) UpdFakultas(ctx context.Context, id string, val pmbulbi.Fakultas) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_fakultas = ?", id).
		Updates(&val).
		Error
	return
}

func (r *Repository) DelFakultas(ctx context.Context, id string) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_fakultas = ?", id).
		Delete(&pmbulbi.Fakultas{}).
		Error
	return
}
