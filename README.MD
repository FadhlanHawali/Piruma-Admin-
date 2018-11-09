# Piruma (Fakultas)

Piruma(fakultas) adalah sebuah aplikasi peminjaman ruangan yang seolah-olah dimiliki oleh sebuah fakultas atau departemen di UGM. Dalam aplikasi peminjaman
ruangan ini disimulasikan milik DTETI UGM

## Getting Started

Berikut adalah route dalam menggunakan API DTETI. Base URL yang digunakan adalah : https://dteti.au-syd.mybluemix.net

### Sign Up

Membuat akun baru

```
Route : /api/signup
Method : /POST
Body :
{
  "id_departemen":string,
  "username":string,
  "departemen":string,
  "fakultas":string,
  "email":string,
  "kontak":int64,
  "password":string,
  "confirm_password":string
}

Response :
{
  "status":"success"
}
```

### Login

Masuk ke akun yang sudah dibuat
```
Route : /api/login
Method : /POST
Body :
{
  "username":string,
  "password":string
}

Response :
{
  "status":"Success",
  "token":string
}
```
Token akan dipakai diseluruh operasi setelah login.

### Add Ruangan

Menambah list ruangan
```
Route : /api/ruangan/add
Method : /POST
Body :
{
	"id_departemen": string,
	"kapasitas": int64,
	"nama_ruangan": string,
	"fasilitas": string
}

Response :
{
    "status": "success"
}
```

### Detail Ruangan

detail dari ruangan tertentu
```
Route : /api/ruangan/room-1541671063/detail
Method : /GET

```

### List Ruangan

list dari ruangan yang tersedia
```
Route : /api/ruang/list
Method : /GET

```

### Delete Ruangan

menghapus ruangan yang tersedia
```
Route : /api/ruangan/room-1541659527/delete
Method : /GET
```

### Add Order

membuat pesanan peminjaman ruangan
```
Route : /api/addOrder
Method : /POST
Body :
{
 "id_ruangan": string,
 "id_departemen": string,
 "ruangan": string,
 "departemen": string,
 "penanggung_jawab": string,
 "telepon": int64,
 "keterangan": string,
 "email": string,
 "StatusSurat":{
  "status_peminjaman": string,
  "status_surat": string
 },
 "TimeStamp":{
  "timestamp_start": int64,
  "timestap_end" int64
 }
}

Response :
{
  "status": "success"
}
```

## Authors

* **Fadhlan Hawali** - *Back End Developer*
* **M. Reyzaldy Indra** - *Android Developer*
* **Faizal P.** - *Web Developer*
* **Raisul Rusdianawati** - *UI/UX Designer*