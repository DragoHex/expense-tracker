# Learnings
- [Repository Pattern](Repository-Pattern.md)

## GORM
- Abstracts the data layer.
- a gorm can be initialised with a supported DB object, which can belong to any of the supported SQL engine.
- When a struct and it's object is passed, it easily handles the CRUD operations, without need of explicitly writing SQL quries.
- gorm tags can be used with the struct's definition to give directives.
- By default gorm prints the queries in the logs, which can be silenced or configured using the following in `gorm.Open`.
`&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}`
