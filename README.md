# Template Clean Architecture Golang

## Deskripsi

Ini dibuat untuk keperluan technical test.

## Arsitektur

![Clean Architecture](architecture.png)

1. Sistem eksternal melakukan request (HTTP, gRPC, Messaging, dll)
2. Delivery membuat berbagai Model dari data request
3. Delivery memanggil Use Case, dan mengeksekusinya menggunakan data Model
4. Use Case membuat Entity untuk kebutuhan logika bisnis
5. Use Case memanggil Repository, dan mengeksekusinya menggunakan data Entity
6. Repository menggunakan data Entity untuk menjalankan operasi database
7. Repository menjalankan operasi database ke database
8. Use Case membuat berbagai Model untuk Gateway atau dari data Entity
9. Use Case memanggil Gateway, dan mengeksekusinya menggunakan data Model
10. Gateway menggunakan data Model untuk menyusun request ke sistem eksternal
11. Gateway menjalankan request ke sistem eksternal (HTTP, gRPC, Messaging, dll)

## Tech Stack

- Golang : [github.com/golang/go](https://github.com/golang/go)
- MySQL (Database) : [github.com/mysql/mysql-server](https://github.com/mysql/mysql-server)
- Redis (Cache) : [github.com/redis/redis](https://github.com/redis/redis)

## Framework & Library

- GoFiber (HTTP Framework) : [github.com/gofiber/fiber](https://github.com/gofiber/fiber)
- GORM (ORM) : [github.com/go-gorm/gorm](https://github.com/go-gorm/gorm)
- Viper (Configuration) : [github.com/spf13/viper](https://github.com/spf13/viper)
- Golang Migrate (Database Migration) : [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- Go Playground Validator (Validation) : [github.com/go-playground/validator](https://github.com/go-playground/validator)
- Logrus (Logger) : [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)

### Aturan Umum untuk Commit

Pada projek menggunakan commit konvensional untuk menangani commit git : [Conventional Commits](https://www.conventionalcommits.org)

- Gunakan `feat: pesan commit` untuk melakukan commit git yang terkait dengan fitur.
- Gunakan `refactor: pesan commit` untuk melakukan commit git yang terkait dengan refactoring kode.
- Gunakan `fix: pesan commit` untuk melakukan commit git yang terkait dengan perbaikan bug.
- Gunakan `test: pesan commit` untuk melakukan commit git yang terkait dengan file test.
- Gunakan `docs: pesan commit` untuk melakukan commit git yang terkait dengan dokumentasi (termasuk file README.md).
- Gunakan `style: pesan commit` untuk melakukan commit git yang terkait dengan gaya kode.

## Konfigurasi

Semua konfigurasi ada di file `config.json`.

## API Spec

Seluruh API Spec ada di folder `api`.

## Migrasi Database

Semua migrasi database ada di folder `db/migrations`.

### Membuat Migrasi

```shell
migrate create -ext sql -dir db/migrations create_table_xxx
```

### Menjalankan Migrasi

```shell
migrate -path db/migrations -database "mysql://root:@tcp(localhost:3301)/db_tech_test?charset=utf8mb4&parseTime=True&loc=Local" up
```

## Menjalankan Aplikasi

### Menjalankan unit test

```bash
go test -v ./test/
```

### Menjalankan web server

```bash
go run cmd/web/main.go
```

### Menjalankan worker

```bash
go run cmd/worker/main.go
```
