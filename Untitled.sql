CREATE TABLE `user` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `nama` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) UNIQUE NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `role` VARCHAR(50) NOT NULL
);

CREATE TABLE `quiz` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `judul` VARCHAR(255) NOT NULL,
  `deskripsi` TEXT,
  `waktu_mulai` TIMESTAMP NOT NULL,
  `waktu_selesai` TIMESTAMP NOT NULL
);

CREATE TABLE `pertanyaan` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `pertanyaan` TEXT NOT NULL,
  `opsi_jawaban` TEXT NOT NULL,
  `jawaban_benar` INT NOT NULL,
  `id_quiz` INT
);

CREATE TABLE `jawaban_peserta` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `id_user` INT,
  `id_quiz` INT,
  `id_pertanyaan` INT,
  `jawaban_peserta` INT NOT NULL,
  `skor` INT
);

ALTER TABLE `pertanyaan` ADD FOREIGN KEY (`id_quiz`) REFERENCES `quiz` (`id`);

ALTER TABLE `jawaban_peserta` ADD FOREIGN KEY (`id_user`) REFERENCES `user` (`id`);

ALTER TABLE `jawaban_peserta` ADD FOREIGN KEY (`id_quiz`) REFERENCES `quiz` (`id`);

ALTER TABLE `jawaban_peserta` ADD FOREIGN KEY (`id_pertanyaan`) REFERENCES `pertanyaan` (`id`);
