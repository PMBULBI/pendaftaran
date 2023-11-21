package prodi

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetAllProdi(ctx context.Context, Mariaconn *gorm.DB) (data []pmbulbi.ProgramStudi, err error) {

	val, err := GetMAllProdi(Mariaconn, ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range val {
		data = append(data, pmbulbi.ProgramStudi{
			ID:               v.ID,
			Fakultas:         v.Fakultas,
			KodeProgramStudi: v.KodeProgramStudi,
			ProgramStudi:     v.ProgramStudi,
		})
	}

	return data, nil
}

func GetOneProdi(ctx context.Context, Mariaconn *gorm.DB, id string) (data pmbulbi.ProgramStudi, err error) {

	prodi, err := GetMOneProdi(Mariaconn, ctx, id)
	if err != nil {
		return pmbulbi.ProgramStudi{}, err
	}

	return prodi, nil
}

func InsertProdi(ctx context.Context, Mariaconn *gorm.DB, val pmbulbi.ProgramStudi) (err error) {
	err = InsProdi(Mariaconn, ctx, val)
	if err != nil {
		return err
	}
	return
}

func UpdateProdi(ctx context.Context, Mariaconn *gorm.DB, id string, val pmbulbi.ProgramStudi) (err error) {
	err = UpdProdi(Mariaconn, ctx, id, val)
	if err != nil {
		return err
	}
	return
}

func DeleteProdi(ctx context.Context, Mariaconn *gorm.DB, id string) (err error) {
	err = DelProdi(Mariaconn, ctx, id)
	if err != nil {
		return err
	}
	return
}
