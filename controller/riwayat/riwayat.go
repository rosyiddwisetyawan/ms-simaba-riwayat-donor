package riwayat

import (
	"fmt"
	"ms-simaba-riwayat-donor/database"
	rwtmdl "ms-simaba-riwayat-donor/models/riwayat"

	_ "github.com/go-sql-driver/mysql"
)

const createQuery = "INSERT INTO %s (kode_pendonor, ktp, jenis_donor, jadwal_donor) VALUES (?,?,?,?)"
const table = "riwayat_donor"

func GetRiwayat(param rwtmdl.RiwayatRequest) (rwtmdl.RiwayatResponseAll, error) {
	var (
		result []rwtmdl.RiwayatResponse
		res    rwtmdl.RiwayatResponseAll
		err    error
	)
	db := database.GetDB().Debug()
	query := db.Table(table)
	if param.KodePendonor != "" {
		query = query.Where("kode_pendonor = ?", param.KodePendonor)
	}
	if param.Ktp != "" {
		query = query.Where("ktp = ?", param.Ktp)
	}
	err = query.Scan(&result).Error
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}
	res.Data = result
	return res, err
}

func CreateRiwayat(param rwtmdl.RiwayatCreateRequest) (err error) {
	db := database.GetDB().Debug()
	err = db.Exec(fmt.Sprintf(createQuery, table), param.KodePendonor, param.Ktp, param.JenisDonor, param.JadwalDonor).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}
