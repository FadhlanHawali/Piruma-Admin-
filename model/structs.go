package model

type Ruangan struct {
	ID uint `gorm:"primary_key"`
	IdRuangan string `json:"id_ruangan"`
	IdDepartemen string `json:"id_departemen"`
	Kapasitas string `json:"kapasitas"`
	NamaRuangan string `json:"nama_ruangan"`
	Fasilitas string `json:"fasilitas"`
}

type Departemen struct {
	ID uint `gorm:"primary_key"`
	IdDepartemen string `json:"id_departemen"`
	Username string `json:"username"`
	Departemen string `json:"departemen"`
	Fakultas string `json:"fakultas"`
	Kontak int64 `json:"kontak"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Orders struct {
	ID uint `gorm:"primary_key"`
	IdPemesanan string `json:"id_pemesanan"`
	IdRuangan string `json:"id_ruangan"`
	IdDepartemen string `json:"id_departemen"`
	Ruangan string `json:"ruangan"`
	Departemen string `json:"departemen"`
	PenanggungJawab string `json:"penanggung_jawab"`
	Telepon string `json:"telepon"`
	Keterangan string `json:"keterangan"`
	Email string `json:"email"`
	StatusPeminjaman string `json:"status_peminjaman"`
	StatusSurat string `json:"status_surat"`
	TimestampStart int64 `json:"timestamp_start,string"`
	TimestampEnd int64 `json:"timestamp_end,string"`
}