package admin

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *UserRepository) Fetch(ctx context.Context) (val []pmbulbi.Admin, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *UserRepository) GetById(ctx context.Context, id string) (val pmbulbi.Admin, err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_admin = ? ", id).
		First(&val).
		Error
	return
}

func (r *UserRepository) Insert(ctx context.Context, val pmbulbi.Admin) (err error) {
	err = r.db.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func (r *UserRepository) Update(ctx context.Context, id string, val pmbulbi.Admin) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_admin = ?", id).
		Updates(&val).
		Error
	return
}

func (r *UserRepository) Delete(ctx context.Context, id string) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id_admin = ?", id).
		Delete(&pmbulbi.Admin{}).
		Error
	return
}

func (r *UserRepository) GetByEmailPass(ctx context.Context, email, password string) (data pmbulbi.Admin, err error) {
	err = r.db.WithContext(ctx).
		Where("email = ? AND password = ?", email, password).
		First(&data).
		Error
	return

}

func (r *UserRepository) GetByPhoneNumPass(ctx context.Context, phoneNum, password string) (data pmbulbi.Admin, err error) {
	err = r.db.WithContext(ctx).
		Where("no_hp = ? AND password = ?", phoneNum, password).
		First(&data).
		Error
	return
}

func (r *UserRepository) CheckUserExists(ctx context.Context, email, phoneNum string) (data pmbulbi.Admin, err error) {
	err = r.db.WithContext(ctx).
		Where("no_hp = ? AND password = ?", phoneNum, email).
		First(&data).
		Error
	return
}
