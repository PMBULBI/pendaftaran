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
func (r *Repository) Update(ctx context.Context, id string, data pmbulbi.Pendaftaran) (err error) {
	err = r.db.
		WithContext(ctx).
		Where("id = ?", id).
		Updates(&data).
		Error
	return
}

func (r *Repository) GetByEmailPass(ctx context.Context, email, password string) (data pmbulbi.Pendaftaran, err error) {
	err = r.db.WithContext(ctx).
		Where("email_mhs = ? AND password = ?", email, password).
		First(&data).
		Error
	return
}
func (r *Repository) GetByPhoneNumPass(ctx context.Context, email, password string) (data pmbulbi.Pendaftaran, err error) {
	err = r.db.WithContext(ctx).
		Where("hp_mhs = ? AND password = ?", email, password).
		First(&data).
		Error
	return
}

func (r *Repository) CheckUserExists(ctx context.Context, email, phoneNum string) (res bool) {
	data := pmbulbi.Pendaftaran{}
	err := r.db.WithContext(ctx).
		Where("email_mhs = ? OR hp_mhs = ?", email, phoneNum).
		First(&data).
		Error
	if err != nil {
		return
	}

	res = data != (pmbulbi.Pendaftaran{})
	return
}
func (r *Repository) GetByEmailOrPhone(ctx context.Context, email, phoneNum string) (res pmbulbi.Pendaftaran, err error) {

	res, err = r.GetByEmail(ctx, email)
	if err != nil {
		res, err = r.GetByPhone(ctx, phoneNum)
	}

	return
}

func (r *Repository) GetByPhone(ctx context.Context, phoneNum string) (res pmbulbi.Pendaftaran, err error) {
	err = r.db.WithContext(ctx).
		Where("hp_mhs = ?", phoneNum).
		First(&res).
		Error

	return
}
func (r *Repository) GetByEmail(ctx context.Context, email string) (res pmbulbi.Pendaftaran, err error) {
	err = r.db.WithContext(ctx).
		Where("email_mhs = ?", email).
		First(&res).
		Error

	return
}
