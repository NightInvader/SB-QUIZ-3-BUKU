# SB-QUIZ-3-BUKU

cara penggunaan : 
1. setelah membuka link, tambahkan /api/signup untuk menambahkan credential pada tabel user menggunakan input json
2. kemudian semua akses terhadap /api lainnya dapat digunakan, 
    GET/api/books --> untuk menampilkan semua buku yang terdapat didalam database
    POST/api/books --> untuk menambahkan buku baru kedalam database
    GET/api/books:id --> untuk menampilkan buku beserta detail menggunakan id sebagai parameter
    DELETE/api/books:id --> untuk menghapus data buku di database dengan menggunakan id sebagai parameter
    GET/api/categories --> untuk menampilkan semua jenis kategori yang terdapat didalam database
    POST/api/categories --> untuk menambahkan kategori baru kedalam tabel categories
    GET/api/categories:id --> untuk menampilkan kategori beserta detailnya berdasarkan id sebagai parameter
    DELETE/api/categories:id --> untuk menghapus data kategori berdasarkan id dari database
    GET/api/categories:id/books --> untuk menampilkan buku yang termasuk dalam kategori yang sama menggunakan id sebagai parameter jenis kategori


