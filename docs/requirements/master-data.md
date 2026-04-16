# 📋 Master Data Requirements & Roadmap

Dokumentasi urutan implementasi master data untuk SIMRS.

## 🦷 Master Data List

| # | Nama | Keterangan | Relasi |
|---|------|------------|--------|
| 1 | **Instalasi** | Area pelayanan besar (Rawat Jalan, Inap, dll) | - |
| 2 | **Unit / Poli** | Unit spesifik (Poli Anak, Poli Dalam, dll) | → Instalasi |
| 3 | **Dokter (DPJP)** | Tenaga medis penanggung jawab | → Unit / Spesialis |
| 4 | **Jadwal Praktik** | Slot waktu pelayanan dokter | → Dokter, Unit |
| 5 | **Kamar / Bed** | Tempat tidur rawat inap | → Unit, Kelas |
| 6 | **Kelas Rawat** | Kategori perawatan (VIP, Kelas 1, 2, 3) | - |
| 7 | **Metode Pembayaran** | Umum, BPJS, Asuransi | - |
| 8 | **ICD-10** | Referensi Kode Penyakit | - |
| 9 | **ICD-9-CM** | Referensi Kode Tindakan | - |
| 10 | **Tarif** | Harga layanan per kelas | → Tindakan, Kelas |
| 11 | **Obat / Alkes** | Master farmasi & logistik | - |

## 🚀 Implementation Roadmap

- [x] **Phase 1**: Authentication
- [x] **Phase 2**: Master Installations
- [x] **Phase 3**: Master Units
- [ ] **Phase 4**: Master Dokter (DPJP)
- [ ] **Phase 5**: Jadwal Praktik
- [ ] **Phase 6**: Kelas Rawat & Kamar/Bed
- [ ] **Phase 7**: Metode Pembayaran & Penjamin
- [ ] **Phase 8**: Ref. Medis (ICD-10 & ICD-9-CM)
- [ ] **Phase 9**: Master Tarif
- [ ] **Phase 10**: Master Obat & Formularium
