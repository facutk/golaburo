# golaburo

## run locally
```sh
go run main.go
```

### migration guide

https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md


```sh
# migrate create -ext sql -dir db/migrations -seq create_users_table

migrate -database ${POSTGRESQL_URL} -path db/migrations up
# migrate -database ${POSTGRESQL_URL} -path db/migrations down
```

### Database access

For now it's defined as global

https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/
https://www.alexedwards.net/blog/organising-database-access
