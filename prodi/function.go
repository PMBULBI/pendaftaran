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
