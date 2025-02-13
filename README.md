## Commands

- Run project `make run`
- Build project `make build`
- Run tests `make test`

### Migrations

> 1. Run database instance before running migrations
> 2. Install [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

- Create migration `make migration <description>`
- Commit migration `make migrate-up`
- Revert migration `make migrate-down`
