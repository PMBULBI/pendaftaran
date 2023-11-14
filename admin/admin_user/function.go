package admin_user

import (
	"context"
	"github.com/PMBULBI/pendaftaran"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func GetAllAdm(ctx context.Context, Mariaenv string) (dataLvlAdm []pmbulbi.Admin, err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	val, err := GetAllAdmin(conn, ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range val {
		dataLvlAdm = append(dataLvlAdm, pmbulbi.Admin{
			IDAdmin:       v.IDAdmin,
			NamaAdmin:     v.NamaAdmin,
			UsernameAdmin: v.UsernameAdmin,
			Email:         v.Email,
			NoHp:          v.NoHp,
			Password:      v.Password,
			TglBuatAkun:   v.TglBuatAkun,
			IsAktif:       v.IsAktif,
			Level:         v.Level,
		})
	}

	return dataLvlAdm, nil
}