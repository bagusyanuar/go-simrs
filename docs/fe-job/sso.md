# SSO Implementation Guide for Frontend

Dokumentasi ini menjelaskan alur integrasi Single Sign-On (SSO) menggunakan protokol OAuth2 dengan PKCE (Proof Key for Code Exchange).

## 1. Alur Login Manual (Authorization Code Flow)

Digunakan saat user mengisi form login di aplikasi.

### Step 1: Request Authorization Code
**Endpoint:** `POST /sso/authorize`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "yourpassword",
  "client_id": "simrs-vue-app",
  "code_challenge": "BASE64URL_ENCODED_SHA256_HASH", // Hasil hash SHA256 dari code_verifier
  "redirect_uri": "http://localhost:3000/callback"
}
```

**Success Response:**
```json
{
  "success": true,
  "message": "authorize success",
  "data": {
    "code": "8c66c304-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  }
}
```
> **Note:** Request ini akan menanamkan HTTPOnly cookie `sso_session` yang digunakan untuk fitur Silent Login.

---

### Step 2: Exchange Code for Token
**Endpoint:** `POST /sso/token`

**Request Body:**
```json
{
  "grant_type": "authorization_code",
  "client_id": "simrs-vue-app",
  "code": "8c66c304-xxxx-xxxx-xxxx-xxxxxxxxxxxx", // Code dari Step 1
  "code_verifier": "plain-random-string", // String asli sebelum di-hash
  "redirect_uri": "http://localhost:3000/callback"
}
```

**Success Response:**
```json
{
  "success": true,
  "message": "token success",
  "data": {
    "access_token": "eyJhbGciOi...",
    "expires_in": 3600
  }
}
```
> **Note:** Request ini akan menanamkan HTTPOnly cookie `refresh_token`. Field `refresh_token` sengaja dikosongkan di body untuk alasan keamanan.

---

## 2. Silent Login (Cek Sesi SSO)

Digunakan untuk login otomatis tanpa form jika user sudah login di aplikasi lain dalam ekosistem SSO yang sama.

**Endpoint:** `GET /sso/authorize`

**Query Parameters:**
- `client_id`: `simrs-vue-app`
- `code_challenge`: `BASE64URL_ENCODED_SHA256_HASH`
- `redirect_uri`: `http://localhost:3000/callback`

**Behavior:**
- **Success (200 OK):** Jika cookie `sso_session` masih valid, akan mengembalikan `code`. Lanjutkan ke **Step 2** (Exchange Token).
- **Failure (401 Unauthorized):** Jika sesi tidak ada atau expired. Arahkan user ke halaman login.

---

## 3. Refresh Token

Digunakan saat `access_token` expired.

**Endpoint:** `POST /sso/token`

**Request Body:**
```json
{
  "grant_type": "refresh_token",
  "client_id": "simrs-vue-app"
}
```

**Note:** Backend secara otomatis mengambil value refresh token dari HTTPOnly cookie `refresh_token`. FE cukup mengirimkan `grant_type` dan `client_id`.

---

## Ringkasan Cookie (HTTPOnly)

| Nama Cookie | Kegunaan | Lifetime |
| :--- | :--- | :--- |
| `sso_session` | Menjaga sesi login global SSO (Silent Login) | 24 Jam |
| `refresh_token` | Digunakan untuk mendapatkan access token baru | Sesuai Config |

## Keamanan (PKCE)
FE wajib generate:
1.  **Code Verifier**: String random (43-128 karakter).
2.  **Code Challenge**: `BASE64URL(SHA256(code_verifier))`.
