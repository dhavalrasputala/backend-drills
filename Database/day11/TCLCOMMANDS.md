## Environment Setup
-make an local Postgresql database using docker:
docker run --name go-postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=appdb -p 5432:5432 -d postgres:15
- Run a local PostgreSQL database using Docker:
sudo systemctl start docker
docker start go-postgres
- Enter postgreSQL using command:-
docker exec -it go-postgres psql -U admin -d appdb

## SQL COMMANDS
1. appdb=# SELECT * FROM transactions;
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
2. appdb=# BEGIN TRANSACTION;UPDATE transactions SET amount=amount+50 WHERE user_id=1;UPDATE transactions SET amount=amount-50 WHERE user_id=2;END TRANSACTION;
  NOTE:-(TRANSACTION is keyword here not table name)
  BEGIN
  UPDATE 2
  UPDATE 1
  COMMIT
3. appdb=# SELECT * FROM transactions;
 id | user_id | amount
  ----+---------+--------
  4 |       4 |  15.25
  5 |       6 |  60.70
  6 |       7 | 132.49
  7 |       8 |  35.65
  1 |       1 | 100.00  (--change--)
  2 |       1 |  70.50
  3 |       2 |  50.00  (--change--)
  (7 rows)
## FINAL TABLE:-

appdb=# SELECT * FROM transactions;
 id | user_id | amount
----+---------+--------
  4 |       4 |  15.25
  5 |       6 |  60.70
  6 |       7 | 132.49
  7 |       8 |  35.65
  1 |       1 | 100.00
  2 |       1 |  70.50
  3 |       2 |  50.00
(7 rows)
## STOPING ENVIRONMENT
- Always close the docker after database query using command:-
1. sudo systemctl stop docker 
2. sudo systemctl stop docker.socket
- Check if inactive using command:-
1. sudo systemctl status docker
