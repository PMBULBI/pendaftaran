package jalur_tahun

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.JalurTahun, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetJalurByTahun(ctx context.Context, tahun int) (dest []pmbulbi.JalurTahunJoin, err error) {
	err = r.db.
		WithContext(ctx).
		Table("jalur_tahun jatah").
		Joins("JOIN jalur_pendaftaran japen ON japen.jalur = jatah.jalur").
		Where("jatah.tahun = ?", tahun).
		Select("jatah.*, japen.*").
		Find(&dest).
		Error
	return
}
