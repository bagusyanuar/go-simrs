# 🚀 Phase 3: Master Units Implementation

Implementasi master data Unit (Poli/Ruangan) yang terhubung dengan Instalasi.

## 📋 Checklist

- [x] **Database & Migration**
  - [x] Buat file migrasi dari `units.schema.dbml`.
  - [x] Jalankan `make migrate-up` untuk tabel `units`.
- [x] **Domain Layer**
  - [x] Definisikan `Unit` entity di `internal/unit/domain/unit.go`.
  - [x] Implementasi GORM Relation (BelongsTo) ke `Installation`.
  - [x] Definisikan `UnitRepository` & `UnitUsecase` interfaces.
- [x] **Repository Layer**
  - [x] Implementasi `UnitRepository` (GORM) di `internal/unit/repository/`.
  - [x] Support Preload `Installation`.
- [x] **Usecase Layer**
  - [x] Implementasi `UnitUsecase` (CRUD).
- [x] **Delivery Layer (HTTP)**
  - [x] Buat `UnitHandler` di `internal/unit/delivery/http/`.
  - [x] Buat DTO (Request dengan `installation_id`).
- [x] **Dependency Injection**
  - [x] Wiring Unit module di `internal/shared/container/`.
- [x] **Bootstrap & Routing**
  - [x] Registrasi routes `/units` di `internal/shared/bootstrap/app.go`.

## 🛠️ Technical Details

- **Relation**: Unit wajib memiliki `installation_id`.
- **Pagination**: Standar pagination dengan default sort by `name`.
- **Search**: Support filter berdasarkan `installation_id` untuk list poli per instalasi.
