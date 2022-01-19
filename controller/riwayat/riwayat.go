package riwayat

import (
	"fmt"
	"ms-simaba-riwayat-donor/database"
	rwtmdl "ms-simaba-riwayat-donor/models/riwayat"

	_ "github.com/go-sql-driver/mysql"
)

const createQuery = "INSERT INTO %s (kode_pendonor, ktp, jenis_donor, jadwal_donor) VALUES (?,?,?,?)"
const table_riwayat = "riwayat_donor"
const table = "calon_pendonor"

func GetRiwayat(param rwtmdl.RiwayatRequest) (rwtmdl.RiwayatResponseAll, error) {
	var (
		result []rwtmdl.RiwayatResponse
		res    rwtmdl.RiwayatResponseAll
		err    error
	)
	db := database.GetDB().Debug()
	query := db.Table(table)
	if param.Ktp != "" {
		query = query.Select("KUESIONER_IC.TGL, calon_pendonor.* ").Joins("left join KUESIONER_IC on KUESIONER_IC.TRANSAKSI = calon_pendonor.kuesioner_id").Where("calon_pendonor.ktp = ? and status = ? ORDER BY created_at DESC", param.Ktp, param.Status)
	}

	fmt.Println(query)
	err = query.Scan(&result).Error
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}
	res.Data = result
	fmt.Println(result)
	fmt.Println(res)
	return res, err
}

func GetRiwayatDetail(param rwtmdl.RiwayatDetailRequest) (rwtmdl.RiwayatResponseDetailAll, error) {
	var (
		result []rwtmdl.RiwayatResponseDetail
		res    rwtmdl.RiwayatResponseDetailAll
		err    error
	)
	db := database.GetDB().Debug()
	query := db.Table(table)
	if param.Ktp != "" {
		query = query.Select("KUESIONER_IC.TGL, calon_pendonor.* ").Joins("left join KUESIONER_IC on KUESIONER_IC.TRANSAKSI = calon_pendonor.kuesioner_id").Where("calon_pendonor.ktp = ? AND calon_pendonor.kuesioner_id = ?", param.Ktp, param.KuesionerId)
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
	err = db.Exec(fmt.Sprintf(createQuery, table_riwayat), param.KodePendonor, param.Ktp, param.JenisDonor, param.JadwalDonor).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}

func GetRiwayatComplete(param rwtmdl.RiwayatCompleteRequest) (rwtmdl.RiwayatCompleteAll, error) {
	var (
		result []rwtmdl.RiwayatCompleteResponse
		res    rwtmdl.RiwayatCompleteAll
		err    error
	)
	db := database.GetDB().Debug()
	query := db.Table(table_riwayat)
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
