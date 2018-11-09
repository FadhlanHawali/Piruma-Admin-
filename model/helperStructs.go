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
	IdDepartemen string `json:"id_departemen" binding:"required"`
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
	StatusSurat StatusSurat
	TimeStamp TimeStamp
}

type SearchRuangan struct {
	Kapasitas string `json:"kapasitas" binding:"required"`
	TimeStamp TimeStamp
}

//type Cobo struct{
//	Nama string `json:"nama" binding:"required"`
//	Objek ObjectBanyak
//}
//
//type ObjectBanyak struct {
//	Objek1 string `json:"objek_1" binding:"required"`
//}

type StatusSurat struct {
	StatusPeminjaman string `json:"status_peminjaman" binding:"required"`
	StatusSurat string `json:"status_surat" binding:"required"`
}

type TimeStamp struct {
	TimestampStart string `json:"timestamp_start" binding:"required"`
	TimestampEnd string `json:"timestap_end" binding:"required"`
}