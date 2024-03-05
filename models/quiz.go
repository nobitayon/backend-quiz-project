package models

import "time"

type Quiz struct {
	Id           int       `gorm:"column:id" json:"id"`
	Judul        string    `gorm:"column:judul" json:"judul"`
	Deskripsi    string    `gorm:"column:deskripsi" json:"deskripsi"`
	WaktuMulai   time.Time `gorm:"column:waktu_mulai" json:"waktuMulai"`
	WaktuSelesai time.Time `gorm:"column:waktu_selesai" json:"waktuSelesai"`
}

func (Quiz) TableName() string {
	return "quiz"
}

type Pertanyaan struct {
	Id           int    `gorm:"column:id" json:"id"`
	Pertanyaan   string `gorm:"column:pertanyaan" json:"pertanyaan"`
	OpsiJawaban  string `gorm:"column:opsi_jawaban" json:"opsiJawaban"`
	JawabanBenar int    `gorm:"column:jawaban_benar" json:"jawabanBenar"`
	IdQuiz       int    `gorm:"column:id_quiz" json:"idQuiz"`
	Quiz         Quiz   `gorm:"foreignKey:IdQuiz"`
}

func (Pertanyaan) TableName() string {
	return "pertanyaan"
}

type JawabanPeserta struct {
	Id             int        `gorm:"column:id" json:"id"`
	IdUser         int        `gorm:"column:id_user" json:"idUser"`
	IdQuiz         int        `gorm:"column:id_quiz" json:"idQuiz"`
	IdPertanyaan   int        `gorm:"column:id_pertanyaan" json:"idPertanyaan"`
	Pertanyaan     Pertanyaan `gorm:"foreignKey:IdPertanyaan"`
	JawabanPeserta int        `gorm:"column:jawaban_peserta" json:"jawabanPeserta"`
	Skor           int        `gorm:"column:skor" json:"skor"`
}

func (JawabanPeserta) TableName() string {
	return "jawaban_peserta"
}
