package Admin

import (
	"context"
	"github.com/PMBULBI/pendaftaran"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/carbon/v2"
)

func GetAllPendaftaran(ctx context.Context, Mariaenv string) (dataProvinsi []pmbulbi.ResponsePendaftaran, err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	val, err := GetAllPendaftar(conn, ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range val {
		dataProvinsi = append(dataProvinsi, pmbulbi.ResponsePendaftaran{
			ID:              v.ID,
			NamaMhs:         v.NamaMhs,
			AsalSekolah:     v.AsalSekolah,
			EmailMhs:        v.EmailMhs,
			HpMhs:           v.HpMhs,
			ProvinsiSekolah: v.ProvinsiSekolah.String,
			KotaSekolah:     v.KotaSekolah.String,
			Password:        v.Password,
			StatusMhs:       v.StatusMhs,
			UsernameAdmin:   v.UsernameAdmin.String,
			TglDaftarMhs:    carbon.DateTime{v.TglDaftarMhs},
		})
	}

	return dataProvinsi, nil
}

func GetOnePendaftar(ctx context.Context, Mariaenv, id string) (data pmbulbi.ResponsePendaftaran, err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	user, err := GetSatuPendaftar(conn, ctx, id)
	if err != nil {
		return pmbulbi.ResponsePendaftaran{}, err
	}

	data.ID = user.ID
	data.NamaMhs = user.NamaMhs
	data.AsalSekolah = user.AsalSekolah
	data.EmailMhs = user.EmailMhs
	data.HpMhs = user.HpMhs
	data.ProvinsiSekolah = user.ProvinsiSekolah.String
	data.KotaSekolah = user.KotaSekolah.String
	data.Password = user.Password
	data.StatusMhs = user.StatusMhs
	data.UsernameAdmin = user.UsernameAdmin.String
	data.TglDaftarMhs = carbon.DateTime{Carbon: user.TglDaftarMhs}

	return data, nil
}
