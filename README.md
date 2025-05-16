# berry
a simple rbac server

For local setup you need to install: 
- https://github.com/golang-migrate/migrate
- Refer: https://www.freecodecamp.org/news/database-migration-golang-migrate/

Then run this:

```
migrate -path database/migration/ -database "postgresql://postgres:password@localhost:5432/berrydb?sslmode=disable" -verbose up
```