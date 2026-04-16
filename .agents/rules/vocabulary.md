---
trigger:
---

# 📖 SIMRS Vocabulary & Business Definitions

Always strictly adhere to these definitions when discussing features, writing code, or naming variables to ensure ubiquitous language across the project.

## 🏥 Pasien & Rekam Medis

- **Pasien**: Entitas individu yang menerima layanan kesehatan.
- **ORM (Nomor Rekam Medis)**: Identitas unik pasien di tingkat rumah sakit (Unique ID).
- **Rekam Medis (RM)**: Kumpulan dokumen yang berisi riwayat kesehatan, hasil pemeriksaan, dan pengobatan pasien.
- **EMR / CPPT (Catatan Perkembangan Pasien Terintegrasi)**: Rekam medis elektronik yang diisi secara kolaboratif oleh Profesional Pemberi Asuhan (PPA).
- **Diagnosis (ICD-10)**: Standar klasifikasi penyakit internasional.
- **Prosedur/Tindakan (ICD-9-CM)**: Standar klasifikasi tindakan medis internasional.

## 🩺 Pelayanan & Instalasi

- **Pendaftaran / Registrasi**: Proses pencatatan pasien untuk mendapatkan layanan di unit tertentu.
- **Kunjungan (Encounter)**: Satu episode pelayanan (Event) dari mulai daftar sampai pulang.
- **Instalasi / Unit**: Pembagian area pelayanan (e.g., Rawat Jalan, Rawat Inap, IGD, Radiologi).
- **Poli (Poliklinik)**: Unit pelayanan rawat jalan spesifik (e.g., Poli Dalam, Poli Anak).
- **Dokter (DPJP)**: Dokter Penanggung Jawab Pelayanan yang memimpin rencana asuhan pasien.

## 💊 Farmasi & Logistik

- **Resep**: Pesanan obat/alkes dari dokter untuk pasien.
- **E-Resep**: Sistem input pesanan obat secara digital.
- **Formularium**: Daftar obat yang tersedia dan disetujui di rumah sakit.
- **Stok Obat**: Persediaan farmasi yang dikelola per unit/gudang.
- **Dispensing**: Proses penyiapan dan penyerahan obat kepada pasien.

## 🏢 Struktur Organisasi

- **Tenaga Medis (Staff)**: User yang memiliki akses berdasarkan peran (Dokter, Perawat, Admin).

## 💳 Billing & Penjamin

- **Billing**: Akumulasi biaya dari semua tindakan, obat, dan sewa kamar dalam satu kunjungan.
- **Penjamin (Payer)**: Pihak yang membayar biaya layanan (e.g., Umum/Pribadi, BPJS, Asuransi Swasta).
- **SEP (Surat Eligibilitas Peserta)**: Bukti kepesertaan pasien BPJS untuk mendapatkan layanan.
