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
