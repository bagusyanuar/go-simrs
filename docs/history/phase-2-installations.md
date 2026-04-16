# 🚀 Phase 2: Master Installations Implementation

Implementasi master data Instalasi (Unit) rumah sakit untuk manajemen area pelayanan.

## 📋 Checklist

- [x] **Database & Migration**
  - [x] Buat file migrasi dari `installations.schema.dbml`.
  - [x] Jalankan `make migrate-up` untuk tabel `installations`.
- [x] **Domain Layer**
  - [x] Definisikan `Installation` entity di `internal/installation/domain/installation.go`.
  - [x] Definisikan `InstallationRepository` interface.
  - [x] Definisikan `InstallationUsecase` interface.
- [x] **Repository Layer**
  - [x] Implementasi `InstallationRepository` (GORM) di `internal/installation/repository/`.
- [x] **Usecase Layer**
  - [x] Implementasi `InstallationUsecase` (CRUD: List, Detail, Create, Update, Delete).
- [x] **Delivery Layer (HTTP)**
  - [x] Buat `InstallationHandler` di `internal/installation/delivery/http/`.
  - [x] Buat DTO untuk request body dan response.
- [x] **Dependency Injection**
  - [x] Wiring Installation module di `internal/shared/container/`.
- [x] **Bootstrap & Routing**
  - [x] Registrasi routes `/installations` di `internal/shared/bootstrap/app.go`.

## 🛠️ Technical Details

- **Pattern**: Standard CRUD with soft delete.
- **Pagination**: Implementasi server-side pagination (page, limit, total_data, total_page).
- **Prefix Route**: `/api/v1/installations`.
- **Validation**: Strict validation untuk field `code` (unique) dan `name`.
- **Search**: Support filter berdasarkan `name` atau `code` (Next improvement).
