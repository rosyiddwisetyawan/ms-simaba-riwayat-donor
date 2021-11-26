package riwayat

type RiwayatRequest struct {
	KodePendonor string `json:"kode_pendonor"`
	Ktp          string `json:"ktp"`
	JenisDonor   string `json:"jenis_donor"`
	JadwalDonor  string `json:"jadwal_donor"`
}

type RiwayatCreateRequest struct {
	KodePendonor string `json:"kode_pendonor"`
	Ktp          string `json:"ktp"`
	JenisDonor   string `json:"jenis_donor"`
	JadwalDonor  string `json:"jadwal_donor"`
	Role         string `json:"role"`
}

type RiwayatResponse struct {
	KodePendonor string `json:"kode_pendonor" gorm:"kode_pendonor"`
	Ktp          string `json:"ktp" gorm:"ktp"`
	JenisDonor   string `json:"jenis_donor" gorm:"jenis_donor"`
	JadwalDonor  string `json:"jadwal_donor" gorm:"jadwal_donor"`
	Status  	 string `json:"status" gorm:"status"`
}

type RiwayatResponseAll struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []RiwayatResponse `json:"data"`
}

