# Migrations

Currently, this project's DB migration is using [go-pg/migrations](https://github.com/go-pg/migrations). You can take a
look at its repo to learn more about it.

## [Running Migrations](https://github.com/go-pg/migrations#run-migrations)

List of main commands to run migrations, run it on this directory:

- `go run *.go up`: runs all available migrations
- `go run *.go up [version-target]`: runs available migrations up to the target one
- `go run *.go down`: reverts last migration
- `go run *.go reset`: reverts all migration
- `go run *.go version`: checks current db version

For full commands check out the original library's guide
[here](https://github.com/go-pg/migrations#run-migrations).

## Writing Migrations

For 1 migration, you should create 2 SQL files for the migration-up and migration-down. Follow these filename formats:

- `[dbversion]_[migration_name].up.sql`: for up migration (eg. `1_create_table_items.up.sqll`)
- `[dbversion]_[migration_name].down.sql`: for down migration (eg. `1_create_table.down.sql`)

## Reference

- https://github.com/go-pg/migrations
- https://en.wikipedia.org/wiki/Schema_migration
- https://martinfowler.com/articles/evodb.html
