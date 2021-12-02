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

type RiwayatDetailRequest struct {
	Ktp          string `json:"ktp"`
	KuesionerId   string `json:"kuesioner_id"`
	
}

type RiwayatResponse struct {
	// KodePendonor string `json:"kode_pendonor" gorm:"kode_pendonor"`
	KuesionerId string `json:"kuesioner_id" gorm:"kuesioner_id"`
	Ktp          string `json:"ktp" gorm:"ktp"`
	JenisDonor   string `json:"jenis_donor" gorm:"jenis_donor"`
	JadwalDonor  string `json:"jadwal_donor" gorm:"jadwal_donor"`
	Status  	 string `json:"status" gorm:"status"`
}

type RiwayatResponseDetail struct {
	KodeCalonPendonor 	string `json:"kode_calon_pendonor" gorm:"kode_calon_pendonor"`
	KuesionerId 		string `json:"kuesioner_id" gorm:"kuesioner_id"`
	Ktp          		string `json:"ktp" gorm:"ktp"`
	JenisDonor   		string `json:"jenis_donor" gorm:"jenis_donor"`
	JadwalDonor  		string `json:"jadwal_donor" gorm:"jadwal_donor"`
	Status  	 		string `json:"status" gorm:"status"`
	Nama  	 			string `json:"nama" gorm:"nama"`
	JenisKelamin  	 	string `json:"jenis_kelamin" gorm:"jenis_kelamin"`
	TanggalLahir  	 	string `json:"tanggal_lahir" gorm:"tanggal_lahir"`
	GolonganDarah  	 	string `json:"golongan_darah" gorm:"golongan_darah"`
	Rhesus  	 		string `json:"rhesus" gorm:"rhesus"`
	Pekerjaan  	 		string `json:"pekerjaan" gorm:"pekerjaan"`
	Alamat  	 		string `json:"alamat" gorm:"alamat"`
	Wilayah  	 		string `json:"wilayah" gorm:"wilayah"`
	Kelurahan  	 		string `json:"kelurahan" gorm:"kelurahan"`
	Kecamatan  	 		string `json:"kecamatan" gorm:"kecamatan"`
	NomorTelepon  	 	string `json:"nomor_telepon" gorm:"nomor_telepon"`
	StatusMenikah  	 	string `json:"status_menikah" gorm:"status_menikah"`
	PositifCovid  	 	string `json:"positif_covid" gorm:"positif_covid"`
	Gejala  	 		string `json:"gejala" gorm:"gejala"`
	TanggalSembuh  	 	string `json:"tanggal_sembuh" gorm:"tanggal_sembuh"`
	StatusBerkas  	 	string `json:"status_berkas" gorm:"status_berkas"`
	BeratBadan  	 	string `json:"berat_badan" gorm:"berat_badan"`
}

type RiwayatResponseAll struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []RiwayatResponse `json:"data"`
}

type RiwayatResponseDetailAll struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []RiwayatResponseDetail `json:"data"`
}

