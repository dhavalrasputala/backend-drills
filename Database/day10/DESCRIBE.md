## Environment Setup
-make an local Postgresql database using docker:
docker run --name go-postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=appdb -p 5432:5432 -d postgres:15
- Run a local PostgreSQL database using Docker:
sudo systemctl start docker
docker go-postgres
- Enter postgreSQL using command:-
docker exec -it go-postgres psql -U admin -d appdb

## SQL COMMANDS
1. LEFT JOIN (REQUIRES TWO TABLES WITH A COMMON FIELD)
  appdb=# SELECT * FROM transactions ORDER BY amount DESC LIMIT 3;
    id | user_id | amount
    ----+---------+--------
    6 |       7 | 132.49
    3 |       2 | 100.00
    5 |       6 |  60.70
  (3 rows)

appdb=#
## FINAL TABLE:-

appdb=# SELECT * FROM transactions;
 id | user_id | amount
----+---------+--------
  1 |       1 |  50.00
  2 |       1 |  20.50
  3 |       2 | 100.00
  4 |       4 |  15.25
  5 |       6 |  60.70
  6 |       7 | 132.49
  7 |       8 |  35.65
(7 rows)

## STOPING ENVIRONMENT
- Always close the docker after database query using command:-
1. sudo systemctl stop docker 
2. sudo systemctl stop docker.socket
- Check if inactive using command:-
1. sudo systemctl status docker
