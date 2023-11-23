package admin

import (
	"context"
	"github.com/PMBULBI/types/schemas"
)

func (r *EmailPassword) GetByEmailPass(ctx context.Context, email, password string) (data schemas.Admin, err error) {
	err = r.db.WithContext(ctx).
		Where("email = ? AND password = ?", email, password).
		First(&data).
		Error
	return

}

func (r *EmailPassword) GetByPhoneNumPass(ctx context.Context, phoneNum, password string) (data schemas.Admin, err error) {
	err = r.db.WithContext(ctx).
		Where("no_hp = ? AND password = ?", phoneNum, password).
		First(&data).
		Error
	return
}

func (r *EmailPassword) CheckUserExists(ctx context.Context, email, phoneNum string) (data schemas.Admin, err error) {
	err = r.db.WithContext(ctx).
		Where("no_hp = ? AND password = ?", phoneNum, email).
		First(&data).
		Error
	return
}
