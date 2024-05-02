# TRANS
<<<<<<< HEAD
1. git versi 1.13.8 
=======
1. golang versi 1.13.8 
>>>>>>> 5efb48883ab1d3f50825ff141bdb15bac6885613
2. RUN "go mod vendor" at root folder
3. RUN "go run *.go" at cmd/reconcile-api

## requirements
if run locally, create an .env file
```
#database
DRIVER_NAME=postgres
CONNECTION_STRING=postgres://postgres:postgres@localhost:5432/db_report_trxn?sslmode=disable
MAX_CONNECTION_POLL=10

PORT=8085
