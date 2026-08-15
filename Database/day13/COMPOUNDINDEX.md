## Environment Setup
-make an local Postgresql database using docker:
docker run --name go-postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=appdb -p 5432:5432 -d postgres:15
- Run a local PostgreSQL database using Docker:
sudo systemctl start docker
docker start go-postgres
- Enter postgreSQL using command:-
docker exec -it go-postgres psql -U admin -d appdb

## SQL COMMANDS
1. appdb=# CREATE INDEX idx_user_id_amount ON transactions(user_id,amount);
  CREATE INDEX
  appdb=# SELECT * FROM transactions;

|id |user_id |amount|
|:---|:---|:----|
|4 |       4 |  15.25 |
|5 |       6 |  60.70 |
|6 |       7 | 132.49 |
|3 |       2 |  50.00 |
|7 |       8 | -14.35 |
|1 |       1 | 200.00 |
|2 |       1 | 170.50 |
(7 rows)

2. appdb=# SELECT * FROM transactions WHERE user_id=8 AND amount=-14.35; //left-right rule

id | user_id | amount
|:---|:---|:---|
|7 |       8 | -14.35|
  (1 row)
  
## FINAL TABLE

appdb=# SELECT * FROM transactions;

|id |user_id |amount|
|:---|:---|:----|
|4 |       4 |  15.25 |
|5 |       6 |  60.70 |
|6 |       7 | 132.49 |
|3 |       2 |  50.00 |
|7 |       8 | -14.35 |
|1 |       1 | 200.00 |
|2 |       1 | 170.50 |
(7 rows)

## STOPING ENVIRONMENT
- Always close the docker after database query using command:-
1. sudo systemctl stop docker 
2. sudo systemctl stop docker.socket
- Check if inactive using command:-
1. sudo systemctl status docker
