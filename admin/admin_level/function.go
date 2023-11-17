package admin_level

import (
	"context"
	"github.com/PMBULBI/pendaftaran"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func GetAllLevelAdm(ctx context.Context, Mariaenv string) (dataLvlAdm []pmbulbi.AdminLevel, err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	val, err := GetAllLevelAdmin(conn, ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range val {
		dataLvlAdm = append(dataLvlAdm, pmbulbi.AdminLevel{
			IDLevel:   v.IDLevel,
			NamaLevel: v.NamaLevel,
		})
	}

	return dataLvlAdm, nil
}

func GetOneLevelAdm(ctx context.Context, Mariaenv, id string) (data pmbulbi.AdminLevel, err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	lvladm, err := GetOneLevelAdmin(conn, ctx, id)
	if err != nil {
		return pmbulbi.AdminLevel{}, err
	}

	return lvladm, nil
}

func InsertLevelAdm(ctx context.Context, Mariaenv string, val pmbulbi.AdminLevel) (err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	data := pmbulbi.AdminLevel{
		IDLevel:   val.IDLevel,
		NamaLevel: val.NamaLevel,
	}
	err = InsertLevelAdmin(conn, ctx, data)
	if err != nil {
		return err
	}
	return
}

func UpdateLevelAdm(ctx context.Context, Mariaenv, id string, val pmbulbi.AdminLevel) (err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	err = UpdateLevelAdmin(conn, ctx, id, val)
	if err != nil {
		return err
	}
	return
}

func DeleteLevelAdm(ctx context.Context, Mariaenv, id string) (err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	err = DeleteLevelAdmin(conn, ctx, id)
	if err != nil {
		return err
	}
	return
}
