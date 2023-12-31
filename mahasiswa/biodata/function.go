package biodata

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func InsertBiodataJalur(ctx context.Context, Mariaconn *gorm.DB, val pmbulbi.BiodataJalur) (err error) {
	err = InsertBioJalur(Mariaconn, ctx, val)
	if err != nil {
		return err
	}
	return
}

func InsertBiodataProdi(ctx context.Context, Mariaconn *gorm.DB, val pmbulbi.BiodataProdi) (err error) {
	err = InsertBioProdi(Mariaconn, ctx, val)
	if err != nil {
		return err
	}
	return
}

func InsertBiodataDiri(ctx context.Context, Mariaconn *gorm.DB, val pmbulbi.BiodataDataDiri) (err error) {
	err = InsertBioDiri(Mariaconn, ctx, val)
	if err != nil {
		return err
	}
	return
}

func InsertBiodataSekolah(ctx context.Context, Mariaconn *gorm.DB, val pmbulbi.BiodataDataSekolah) (err error) {
	err = InsertBioSekolah(Mariaconn, ctx, val)
	if err != nil {
		return err
	}
	return
}

func InsertBiodataOrtu(ctx context.Context, Mariaconn *gorm.DB, val pmbulbi.BiodataDataOrtu) (err error) {
	err = InsertBioOrtu(Mariaconn, ctx, val)
	if err != nil {
		return err
	}
	return
}
