# ecomm-alpha
## To run the project use following steps
- create .env file in root directory of the project
- add below lines in .env file
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=123456
DB_NAME=ecomm_alpha
SECRET=asd
```
- run command: go run main.go

## To run tests use following steps
- for available tests, run command :  go test .\tests\sellertests\ -v

