# Learnings
- [Repository Pattern](Repository-Pattern.md)

## GORM
- Abstracts the data layer.
- a gorm can be initialised with a supported DB object, which can belong to any of the supported SQL engine.
- When a struct and it's object is passed, it easily handles the CRUD operations, without need of explicitly writing SQL quries.
- gorm tags can be used with the struct's definition to give directives.
- By default gorm prints the queries in the logs, which can be silenced or configured using the following in `gorm.Open`.
`&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}`
- When using `gorm.DB.First(&obj, id)`, be very mindfull. This is sufficient if we are using an int ID as our primary key. But in case it is something else, then it is wise to add where condition like this `gorm.DB.Where("column_name = ?", value).First(&obj)`

## encoding/csv
- For exporting slice of struct objects to a csv file. [CODE](../pkg/cmd/export.go)

## SQLITE
- While using following go packages `gorm.io/driver/sqlite` & `github.com/mattn/go-sqlite3`, both internally uses cgo to interface with the SQLite C library.
They interact with the sqlite installed on the user system. This would require the user to have sqlite installed on their system.
But that may not be the case.
- One option that we have to **static linking**, which compile SQLite directly into your Go binary using `cgo`, which then won't need user to install sqlite separately.
- If we have `CGO_ENABLED=1` during the build then sqlite will be compiled with the binary itself. Although setting `CGO_ENABLED=1` is sufficent, following is an explicit way to ensure static linking works every time.
```shell
CGO_ENABLED=1 go build -tags "sqlite_omit_load_extension" -ldflags "-extldflags '-static'" your_program.go
```
- For building cross platform binaries using cgo we may need to use containers for individual platform becaue of missing clibraries.

## SQLC
sqlc generates fully type-safe idiomatic Go code from SQL. Here’s how it works:
1. We write SQL queries.
2. We run sqlc to generate Go code that presents type-safe interfaces to those queries. configurations can be done in `sqlc.yaml` file.
3. We write application code that calls the methods sqlc generated.

This has support for postgres, mysql & sqlite.
Support for SQLite is not up to the par.
When working with SQLite to get proper types in the generated code,
it's better to override them column wise in the `sqlc.yaml` file.

**NOTE:** sqlc doesn't close the transaction object itself. So developer need to handle the rollback in case of any error, and commit in case of success.
So, after auto generating code using sqlc, it's better to maintain a service layer that handles the transactions.

## Miscellaneous
- If only month and year are to be saved in a `time.Time` instance. This will set the day as `1` and time as midnight.
```go
year, month := 2024, time.November
monthYear := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
```
- In case we have two optional flags, say `A` & `B`. And if we have a scenario where we cannot use B in absence of `A`, then we can do the following.
```go
if cmd.Flags().Changed("B") && !cmd.Flags().Changed("A") {
    return errors.New("flag B requires flag A to be set")
}
```
