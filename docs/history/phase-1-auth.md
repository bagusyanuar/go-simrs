# 🚀 Phase 1: Authentication Implementation

Implementasi sistem autentikasi dasar menggunakan JWT, mencakup fitur Login dan Refresh Token (tanpa Registrasi karena sistem internal).

## 📋 Checklist

- [x] **Database & Migration**
  - [x] Generate migration file dari `users.schema.dbml`.
  - [x] Jalankan `make migrate-up` untuk inisialisasi tabel `users`.
- [x] **Domain Layer**
  - [x] Definisikan `User` entity di `internal/user/domain/user.go`.
  - [x] Definisikan `UserRepository` interface.
  - [x] Definisikan `AuthUsecase` interface.
- [x] **Repository Layer**
  - [x] Implementasi `UserRepository` menggunakan GORM di `internal/user/repository/`.
- [x] **Usecase Layer**
  - [x] Implementasi `AuthUsecase` (Login & Refresh logic).
- [x] **Delivery Layer (HTTP)**
  - [x] Buat Handler `AuthHandler` di `internal/auth/delivery/http/`.
  - [x] Buat DTO untuk Request Login dan Refresh.
- [x] **Dependency Injection**
  - [x] Wiring `UserRepo`, `AuthUC`, dan `AuthHandler` di `internal/shared/container/container.go`.
- [x] **Bootstrap & Routing**
  - [x] Daftarkan route auth di `internal/shared/bootstrap/app.go`.

## 🛠️ Technical Details

- **Auth Method**: JWT (HS256).
- **Password Hashing**: Bcrypt.
- **Validation**: `go-playground/validator/v10`.
- **API Response**: Unified response wrapper (`pkg/response`).
