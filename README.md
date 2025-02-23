# Go Realworld Example

## Install go-migrate

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Running API tests locally

To locally run the provided Postman collection against your backend, execute:

```bash
APIURL=http://localhost:3000/api ./scripts/api-tests.sh
```
