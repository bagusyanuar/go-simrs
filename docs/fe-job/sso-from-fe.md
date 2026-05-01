# BE Job: Seamless SSO (Silent Login) Implementation

## 🎯 Goal

Memungkinkan user yang sudah login di satu aplikasi (misal: Master Data) untuk otomatis login di aplikasi lain (misal: Inventory) tanpa harus memasukkan kredensial lagi dan tanpa "kedip" redirect (Silent Login).

---

## 🛠️ Perubahan Utama: Endpoint `/authorize` (GET)

Backend perlu memodifikasi logika pada endpoint `GET /authorize` untuk mendukung respon format JSON jika diminta oleh Frontend via AJAX.

### 1. Deteksi Request Silent Login

Backend harus mengecek dua hal:

- Apakah ada header `Accept: application/json` atau `X-Requested-With: XMLHttpRequest`?
- Apakah ada parameter `prompt=none` di URL? (Opsional tapi disarankan sesuai spek OAuth2).

### 2. Logika Response

Jika request terdeteksi sebagai Silent Login (AJAX):

#### A. Jika User SUDAH memiliki Global Session yang valid:

Jangan melakukan **302 Redirect**. Berikan respon **200 OK** dengan body JSON.

- **Body JSON**:
  ```json
  {
    "status": "success",
    "data": {
      "code": "AUTH_CODE_GENERATED",
      "state": "STATE_FROM_CLIENT"
    }
  }
  ```

#### B. Jika User BELUM login / Session Habis:

Jangan redirect ke halaman Login. Berikan respon **401 Unauthorized** atau JSON error agar Frontend bisa memicu flow login manual.

- **Body JSON**:
  ```json
  {
    "status": "error",
    "message": "login_required",
    "error": "login_required"
  }
  ```

---

## 🔒 Security Requirements

1. **CORS**: Pastikan endpoint `/authorize` mengizinkan request dari semua subdomain aplikasi (`*.neurovi-simulation.test`) dan mendukung `Allow Credentials`.
2. **Cookie Domain**: Pastikan session cookie (misal `sid`) diset dengan `Domain=.neurovi-simulation.test` (pakai titik di depan) supaya terbaca di semua aplikasi.
3. **HTTPOnly**: Tetap gunakan `HttpOnly` dan `Secure` untuk cookie session.

---

## 📝 Flow Checklist untuk BE

- [ ] Tambahkan pengecekan header `Accept: application/json` di handler `/authorize`.
- [ ] Jika user sudah login (session exist), bypass view login dan langsung generate `code`.
- [ ] Kirim `code` tersebut dalam format JSON ke FE.
- [ ] Pastikan `state` yang dikirim FE dikembalikan secara utuh di dalam JSON.

---

> [!IMPORTANT]
> Dengan perubahan ini, Frontend (SDK) bisa melakukan "jabat tangan" di background tanpa mengganggu kenyamanan user. User akan merasa aplikasi sangat seamless.
