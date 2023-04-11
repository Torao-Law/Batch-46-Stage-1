const xhr = new XMLHttpRequest()

xhr.open("METHOD", "LINK URL", "STATUS")
// 1. GET, POST, PATCH, DELETE
// 2. https://api.npoint.io/19dc1d52f9465c39624a
// 3. true/false

xhr.onload = function() { } // ketika mengecek status request, saat di load
xhr.onerror = function() { } // ketika kondisi error, mau menjalankan apa
xhr.send() // kirim sebuah request ke alamat server
