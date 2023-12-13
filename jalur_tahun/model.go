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
		Joins("JOIN jalur_pendaftaran japen ON japen.id_jalur = jatah.id_jalur").
		Where("jatah.tahun = ?", tahun).
		Where("japen.status = ?", "aktif").
		Order("japen.nama_jalur ASC").
		Select("jatah.*, japen.*").
		Find(&dest).
		Error
	return
}
