# 🛠️ Frontend Implementation Guide: PKCE SSO

Dokumen ini merangkum tugas dan alur kerja Tim Frontend untuk mengintegrasikan sistem Single Sign-On (SSO) menggunakan flow PKCE (Proof Key for Code Exchange).

## 1. Persiapan Environment
Simpan konfigurasi berikut di file `.env` aplikasi FE:
```env
VITE_SSO_BASE_URL=http://localhost:8081/api/v1/sso
VITE_SSO_CLIENT_ID=simrs-vue-app
VITE_SSO_REDIRECT_URI=http://neurovi-simulation.test:3000/callback
```

## 2. Alur Kerja (Workflow)

### Tahap A: Inisiasi Auth (Aplikasi Target)
Saat user membuka aplikasi dan tidak memiliki token:
1.  **Generate PKCE Pair**:
    *   `code_verifier`: String acak panjang (43-128 karakter, high entropy).
    *   `code_challenge`: `Base64URL(SHA256(code_verifier))`.
2.  **Simpan Verifier**: Simpan `code_verifier` di `sessionStorage`.
3.  **Cek Silent Login**:
    *   Hit `GET {{VITE_SSO_BASE_URL}}/authorize` bawa query params `client_id`, `code_challenge`, & `redirect_uri`.
    *   Jika response `200 OK`, ambil `code` dari data response, lalu lanjut ke **Tahap C**.
    *   Jika response `401 Unauthorized`, arahkan user ke **Halaman Login Portal SSO**.

### Tahap B: Halaman Login (Portal SSO)
Jika Silent Login gagal:
1.  Tampilkan form Email & Password.
2.  Hit `POST {{VITE_SSO_BASE_URL}}/authorize`.
    *   Body: `email`, `password`, `client_id`, `code_challenge`, `redirect_uri`.
3.  Jika sukses, ambil `code` dari response.
4.  **Redirect Manual**: Arahkan browser kembali ke `redirect_uri` dengan query param `code`.
    *   Contoh: `window.location.href = redirect_uri + "?code=" + code`.

### Tahap C: Token Exchange (Aplikasi Target)
Di halaman callback (`/callback`):
1.  **Ambil Code**: Ambil parameter `code` dari URL query string.
2.  **Ambil Verifier**: Ambil `code_verifier` yang tadi disimpan di `sessionStorage`.
3.  **Tukar Token**:
    *   Hit `POST {{VITE_SSO_BASE_URL}}/token`.
    *   Body: `client_id`, `code`, `code_verifier`, `redirect_uri`.
4.  **Simpan Token**: Simpan `access_token` dan `refresh_token` di LocalStorage/Cookie.
5.  **Clean Up**: Hapus `code_verifier` dari `sessionStorage` dan bersihkan URL query param.

## 3. Contoh Snippet PKCE (JavaScript)

```javascript
// Generate Code Verifier
const generateVerifier = () => {
  const array = new Uint32Array(56);
  window.crypto.getRandomValues(array);
  return Array.from(array, dec => ('0' + dec.toString(16)).substr(-2)).join('');
};

// Generate Code Challenge from Verifier
const generateChallenge = async (verifier) => {
  const encoder = new TextEncoder();
  const data = encoder.encode(verifier);
  const hash = await window.crypto.subtle.digest('SHA-256', data);
  return btoa(String.fromCharCode(...new Uint8Array(hash)))
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=+$/, '');
};
```

## 4. Checklist Keamanan
- [ ] Pastikan `code_verifier` tidak pernah dikirim ke endpoint `/authorize`.
- [ ] Pastikan `code_verifier` dihapus setelah token berhasil didapatkan.
- [ ] Gunakan parameter `state` (random string) untuk mencegah CSRF saat redirect.
- [ ] Pastikan `redirect_uri` yang dikirim sama persis dengan yang terdaftar di whitelist backend.

---
**Status Backend**: Ready (Port 8081)
**Database**: PostgreSQL (Tabel `oauth_clients`, `auth_code_sessions`, `sso_sessions`)
