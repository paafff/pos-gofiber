## Struktur Folder Project

Berikut adalah struktur folder dari project ini:

![image](https://github.com/user-attachments/assets/901d9747-48ae-4d9a-bb10-486cfc1354c8)


app/
├── cmd/
│   └── web/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── domain/
│   │   ├── models/
│   │   │   └── user.go
│   │   └── services/
│   │       └── user_service.go
│   ├── handlers/
│   │   └── user_handler.go
│   ├── infrastructure/
│   │   ├── database/
│   │   │   └── database.go
│   │   ├── router/
│   │   │   └── router.go
│   │   └── logger/
│   │       └── logger.go
│   ├── repositories/
│   │   └── user_repository.go
│   ├── routes/
│   │   └── user_routes.go
│   └── auth/
│       ├── auth_handler.go
│       ├── auth_service.go
│       └── auth_middleware.go
├── pkg/
│   └── utils/
│       └── utils.go
├── config.json
├── Makefile
├── go.mod
├── go.sum
└── README.md

- **cmd/**: Berisi entry point dari aplikasi.
  - **web/**: Berisi file [`main.go`](command:_github.copilot.openSymbolFromReferences?%5B%22main.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22c%3A%5C%5CUsers%5C%5Cdanangr%5C%5CDesktop%5C%5Cdevp%5C%5Cgithub%5C%5Cgo-fiber%5C%5CREADME.md%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fc%253A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22path%22%3A%22%2Fc%3A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A3%2C%22character%22%3A12%7D%7D%5D%5D 'Go to definition') yang merupakan entry point dari aplikasi web.
- **internal/**: Berisi kode aplikasi yang tidak diekspos ke luar.
  - **config/**: Berisi konfigurasi aplikasi.
  - **domain/**: Berisi logika bisnis dan model aplikasi.
    - **models/**: Berisi definisi model seperti [`user.go`](command:_github.copilot.openSymbolFromReferences?%5B%22user.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22c%3A%5C%5CUsers%5C%5Cdanangr%5C%5CDesktop%5C%5Cdevp%5C%5Cgithub%5C%5Cgo-fiber%5C%5CREADME.md%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fc%253A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22path%22%3A%22%2Fc%3A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A9%2C%22character%22%3A16%7D%7D%5D%5D 'Go to definition').
    - **services/**: Berisi layanan yang mengimplementasikan logika bisnis seperti [`user_service.go`](command:_github.copilot.openSymbolFromReferences?%5B%22user_service.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22c%3A%5C%5CUsers%5C%5Cdanangr%5C%5CDesktop%5C%5Cdevp%5C%5Cgithub%5C%5Cgo-fiber%5C%5CREADME.md%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fc%253A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22path%22%3A%22%2Fc%3A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A11%2C%22character%22%3A16%7D%7D%5D%5D 'Go to definition').
  - **handlers/**: Berisi handler untuk menangani request HTTP seperti [`user_handler.go`](command:_github.copilot.openSymbolFromReferences?%5B%22user_handler.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22c%3A%5C%5CUsers%5C%5Cdanangr%5C%5CDesktop%5C%5Cdevp%5C%5Cgithub%5C%5Cgo-fiber%5C%5CREADME.md%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fc%253A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22path%22%3A%22%2Fc%3A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A13%2C%22character%22%3A12%7D%7D%5D%5D 'Go to definition').
  - **infrastructure/**: Berisi kode infrastruktur seperti database, router, dan logger.
    - **database/**: Berisi kode untuk koneksi database seperti [`database.go`](command:_github.copilot.openSymbolFromReferences?%5B%22database.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22c%3A%5C%5CUsers%5C%5Cdanangr%5C%5CDesktop%5C%5Cdevp%5C%5Cgithub%5C%5Cgo-fiber%5C%5CREADME.md%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fc%253A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22path%22%3A%22%2Fc%3A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A15%2C%22character%22%3A12%7D%7D%5D%5D 'Go to definition').
    - **router/**: Berisi kode untuk routing seperti [`router.go`](command:_github.copilot.openSymbolFromReferences?%5B%22router.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22c%3A%5C%5CUsers%5C%5Cdanangr%5C%5CDesktop%5C%5Cdevp%5C%5Cgithub%5C%5Cgo-fiber%5C%5CREADME.md%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fc%253A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22path%22%3A%22%2Fc%3A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A17%2C%22character%22%3A12%7D%7D%5D%5D 'Go to definition').
    - **logger/**: Berisi kode untuk logging seperti [`logger.go`](command:_github.copilot.openSymbolFromReferences?%5B%22logger.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22c%3A%5C%5CUsers%5C%5Cdanangr%5C%5CDesktop%5C%5Cdevp%5C%5Cgithub%5C%5Cgo-fiber%5C%5CREADME.md%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fc%253A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22path%22%3A%22%2Fc%3A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A19%2C%22character%22%3A12%7D%7D%5D%5D 'Go to definition').
  - **repositories/**: Berisi kode untuk akses data seperti [`user_repository.go`](command:_github.copilot.openSymbolFromReferences?%5B%22user_repository.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22c%3A%5C%5CUsers%5C%5Cdanangr%5C%5CDesktop%5C%5Cdevp%5C%5Cgithub%5C%5Cgo-fiber%5C%5CREADME.md%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fc%253A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22path%22%3A%22%2Fc%3A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A22%2C%22character%22%3A12%7D%7D%5D%5D 'Go to definition').
  - **routes/**: Berisi definisi rute aplikasi seperti [`user_routes.go`](command:_github.copilot.openSymbolFromReferences?%5B%22user_routes.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22c%3A%5C%5CUsers%5C%5Cdanangr%5C%5CDesktop%5C%5Cdevp%5C%5Cgithub%5C%5Cgo-fiber%5C%5CREADME.md%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fc%253A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22path%22%3A%22%2Fc%3A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A24%2C%22character%22%3A12%7D%7D%5D%5D 'Go to definition').
- **pkg/**: Berisi kode yang dapat digunakan kembali di luar aplikasi.
  - **utils/**: Berisi utilitas umum seperti [`utils.go`](command:_github.copilot.openSymbolFromReferences?%5B%22utils.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22c%3A%5C%5CUsers%5C%5Cdanangr%5C%5CDesktop%5C%5Cdevp%5C%5Cgithub%5C%5Cgo-fiber%5C%5CREADME.md%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fc%253A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22path%22%3A%22%2Fc%3A%2FUsers%2Fdanangr%2FDesktop%2Fdevp%2Fgithub%2Fgo-fiber%2FREADME.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A26%2C%22character%22%3A8%7D%7D%5D%5D 'Go to definition').
- **go.mod**: Berisi informasi modul Go.
- **go.sum**: Berisi checksum dari dependensi Go.
- **README.md**: Berisi dokumentasi project.
