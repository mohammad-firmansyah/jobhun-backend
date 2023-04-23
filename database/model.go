package database

type Mahasiswa struct {
	ID            string   `json:"id"`
	Nama          string   `json:"nama" validate:"required"`
	Usia          int      `json:"usia" validate:"required"`
	Gender        int      `json:"gender" validate:"required"`
	TglRegistrasi string   `json:"tgl_registrasi" validate:"required"`
	Hobi          []string `json:"hobi" validate:"required"`
	Jurusan       string   `json:"jurusan" validate:"required"`
}

type Hobi struct {
	ID       string `json:"id"`
	NamaHobi string `json:"nama_hobi"`
}

type Jurusan struct {
	ID          string `json:"id"`
	NamaJurusan string `json:"nama_jurusan"`
}

type MahasiswaHobi struct {
	ID          string `json:"id"`
	IdMahasiswa string `json:"id_mahasiswa"`
	IdHobi      string `json:"id_hobi"`
}
