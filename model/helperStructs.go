package model

type SignUp struct {
	IdDepartemen string `json:"id_departemen" binding:"required"`
	Username string `json:"username" binding:"required"`
	Departemen string `json:"departemen" binding:"required"`
	Fakultas string `json:"fakultas" binding:"required"`
	Email string `json:"email" binding:"required"`
	Kontak int64 `json:"kontak,string" binding:"required"`
	Password string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AddRuangan struct {
	Kapasitas string `json:"kapasitas" binding:"required"`
	NamaRuangan string `json:"nama_ruangan" binding:"required"`
	Fasilitas string `json:"fasilitas" binding:"required"`
}

type AddOrder struct {
	IdRuangan string `json:"id_ruangan" binding:"required"`
	IdDepartemen string `json:"id_departemen" binding:"required"`
	Ruangan string `json:"ruangan" binding:"required"`
	Departemen string `json:"departemen" binding:"required"`
	PenanggungJawab string `json:"penanggung_jawab" binding:"required"`
	Telepon string `json:"telepon" binding:"required"`
	Keterangan string `json:"keterangan" binding:"required"`
	Email string `json:"email" binding:"required"`
	StatusProses *bool `json:"status_proses"`
	StatusSurat StatusSurat
	TimeStamp TimeStamp
}

type SearchRuangan struct {
	Kapasitas string `json:"kapasitas" binding:"required"`
	TimeStamp TimeStamp
}

type JadwalDepartemen struct {
	IdDepartemen string `json:"id_departemen"`
	Departemen string `json:"departemen"`
	JadwalRuangan []JadwalRuangan
}

type JadwalRuangan struct {
	Ruangan string `json:"ruangan"`
	IdRuangan string `json:"id_ruangan"`
	JadwalDepartemen []JadwalDepartemen
}

type DetailJadwalRuangan struct {
	IdPemesanan string `json:"id_pemesanan"`
	PenanggungJawab string `json:"penanggung_jawab"`
	Telepon string `json:"telepon"`
	Keterangan string `json:"keterangan"`
	Email string `json:"email"`
	StatusPeminjaman *bool `json:"status_peminjaman"`
	StatusSurat string `json:"status_surat"`
	TimestampStart int64 `json:"timestamp_start,string"`
	TimestampEnd int64 `json:"timestamp_end,string"`
}

//type Cobo struct{
//	Nama string `json:"nama" binding:"required"`
//	Objek ObjectBanyak
//}
//
//type ObjectBanyak struct {
//	Objek1 string `json:"objek_1" binding:"required"`
//}

type Hasil struct {
	Count int `gorm:"column:COUNT(id_departemen)" json:"count" `
	Departemen string `gorm:"column:id_departemen" json:"departemen" `
}

type StatusSurat struct {
	StatusPeminjaman *bool `json:"status_peminjaman" binding:"exists"`
	StatusSurat string `json:"status_surat" binding:"required"`
}

type Jadwal struct {
	Keterangan string `gorm:"column:keterangan" json:"keterangan"`
	TimestampStart int64 `gorm:"column:timestamp_start" json:"timestamp_start,string"`
	TimestampEnd int64 `gorm:"column:timestamp_end" json:"timestamp_end,string"`
}


type TimeStamp struct {
	TimestampStart int64 `json:"timestamp_start,string" binding:"required"`
	TimestampEnd int64 `json:"timestamp_end,string" binding:"required"`
}