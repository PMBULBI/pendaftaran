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
