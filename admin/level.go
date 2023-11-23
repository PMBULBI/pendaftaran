package admin

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Level) Fetch(ctx context.Context) (val []pmbulbi.AdminLevel, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Level) GetById(ctx context.Context, id string) (val pmbulbi.AdminLevel, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_level = ? ", id).
		First(&val).
		Error
	return
}

func (r *Level) Insert(ctx context.Context, val pmbulbi.AdminLevel) (err error) {
	err = r.db.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func (r *Level) Update(ctx context.Context, id string, val pmbulbi.AdminLevel) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_level = ?", id).
		Updates(&val).
		Error
	return
}

func (r *Level) Delete(ctx context.Context, id string) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_level = ?", id).
		Delete(&pmbulbi.AdminLevel{}).
		Error
	return
}
