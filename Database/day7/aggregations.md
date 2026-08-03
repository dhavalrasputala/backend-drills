
## Environment Setup

- Ran a local PostgreSQL database using Docker:
docker run --name go-postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=appdb -p 5432:5432 -d postgres:15
- Enter postgreSQL using command:-
docker exec -it go-postgres psql -U admin -d appdb

## SQL COMMANDS

1.  CREATE TABLE 
  appdb=#  CREATE TABLE transactions(id SERIAL PRIMARY KEY,user_id INT,amount NUMERIC(10,2));
  CREATE TABLE

2. INSERT DATA 
  appdb=# INSERT INTO transactions (user_id, amount) VALUES (1, 50.00); INSERT INTO transactions (user_id, amount) VALUES (1, 20.50); INSERT INTO transactions (user_id, amount) VALUES (2, 100.00); INSERT INTO transactions (user_id, amount) VALUES (4, 15.25);
  INSERT 0 1
  INSERT 0 1
  INSERT 0 1
  INSERT 0 1

3. AGGREGATION QUERY (SUM & GROUP BY):- 
 appdb=# SELECT user_id, SUM(amount) FROM transactions GROUP BY user_id;
  user_id |  sum
  ---------+--------
      4 |  15.25
      2 | 100.00
      1 |  70.50
  (3 rows)

## FINAL TABLE:-

appdb=# SELECT * FROM transactions;
 id | user_id | amount
----+---------+--------
  1 |       1 |  50.00
  2 |       1 |  20.50
  3 |       2 | 100.00
  4 |       4 |  15.25
(4 rows)

## STOPING ENVIRONMENT
- Always close the docker after database query using command:-
1. sudo systemctl stop docker 
2. sudo systemctl stop docker.socket
- Check if inactive using command:-
1. sudo systemctl status docker
