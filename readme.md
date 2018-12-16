# Golang Gorm with PostgreSQL

## Things to note
We need to install the PostgreSQL driver to run
```bash
go get -u github.com/lib/pq
```

We also need Gorm
```bash
go get -u github.com/jinzhu/gorm
```

Disable SSL when opening the connection using `db.Open` if the server has no SSL enabled
```
sslmode=disable
```