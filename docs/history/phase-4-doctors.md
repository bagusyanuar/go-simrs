# 🚀 Phase 4: Master Doctors (DPJP) Implementation

Implementasi master data dokter, spesialisasi, dan mapping area praktik (Unit).

## 📋 Checklist

- [x] **Database & Migration**
  - [x] Buat file migrasi dari `doctors.schema.dbml`.
  - [x] Jalankan `make migrate-up` untuk tabel `specialties`, `doctors`, dan `doctor_units`.
- [x] **Module: Specialties**
  - [x] Definisikan `Specialty` entity & interfaces.
  - [x] Implementasi Repository (GORM).
  - [x] Implementasi Usecase (CRUD).
  - [x] Buat `SpecialtyHandler` & DTO.
- [x] **Module: Doctors**
  - [x] Definisikan `Doctor` entity dengan relasi `BelongsTo` (Specialty) & `ManyToMany` (Units).
  - [x] Implementasi `DoctorRepository` dengan support Preload Specialty & Units.
  - [x] Implementasi `DoctorUsecase` (CRUD).
  - [x] Buat `DoctorHandler` & DTO.
- [x] **Dependency Injection**
  - [x] Wiring `Specialty` module.
  - [x] Wiring `Doctor` module di `internal/shared/container/`.
- [x] **Bootstrap & Routing**
  - [x] Registrasi routes `/specialties` & `/doctors` di router.

## 🛠️ Technical Details

- **Relation**: 
    - `Doctor` -> `Specialty` (Many-to-One).
    - `Doctor` <-> `Unit` (Many-to-Many via `doctor_units`).
- **Validation**: 
    - `nik` & `sip` wajib unik.
    - `sip_expiry_date` wajib ada.
- **Features**: 
    - List dokter bisa difilter berdasarkan `specialty_id` atau `unit_id`.
    - Soft delete pada dokter tidak menghapus relasi unit secara permanen (tergantung kebijakan gorm).
