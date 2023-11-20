package admin_level

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetAllLevelAdm(ctx context.Context, Mariaconn *gorm.DB) (dataLvlAdm []pmbulbi.AdminLevel, err error) {

	val, err := GetAllLevelAdmin(Mariaconn, ctx)

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

func GetOneLevelAdm(ctx context.Context, Mariaconn *gorm.DB, id string) (data pmbulbi.AdminLevel, err error) {

	lvladm, err := GetOneLevelAdmin(Mariaconn, ctx, id)
	if err != nil {
		return pmbulbi.AdminLevel{}, err
	}

	return lvladm, nil
}

func InsertLevelAdm(ctx context.Context, Mariaconn *gorm.DB, val pmbulbi.AdminLevel) (err error) {

	data := pmbulbi.AdminLevel{
		IDLevel:   val.IDLevel,
		NamaLevel: val.NamaLevel,
	}
	err = InsertLevelAdmin(Mariaconn, ctx, data)
	if err != nil {
		return err
	}
	return
}

func UpdateLevelAdm(ctx context.Context, Mariaconn *gorm.DB, id string, val pmbulbi.AdminLevel) (err error) {

	err = UpdateLevelAdmin(Mariaconn, ctx, id, val)
	if err != nil {
		return err
	}
	return
}

func DeleteLevelAdm(ctx context.Context, Mariaconn *gorm.DB, id string) (err error) {
	err = DeleteLevelAdmin(Mariaconn, ctx, id)
	if err != nil {
		return err
	}
	return
}
