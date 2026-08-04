## Environment Setup

- Ran a local PostgreSQL database using Docker:
docker run --name go-postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=appdb -p 5432:5432 -d postgres:15
- Enter postgreSQL using command:-
docker exec -it go-postgres psql -U admin -d appdb

## SQL COMMANDS

1. INSERT DATA 
  appdb=# INSERT INTO transactions (user_id, amount) VALUES (5, 80.00); INSERT INTO transactions (user_id, amount) VALUES (6, 60.70); INSERT INTO transactions (user_id, amount) VALUES (7, 132.49); INSERT INTO transactions (user_id, amount) VALUES (8, 35.65);
  INSERT 0 1
  INSERT 0 1
  INSERT 0 1
  INSERT 0 1

2. INNER JOIN (REQUIRES TWO TABLES WITH A COMMON FIELD)
  appdb=# SELECT td.date,u.user_id FROM transactions_date AS td INNER JOIN transactions AS u ON td.user_id = u.user_id;
          date          | user_id
  ------------------------+---------
  2026-08-01 00:15:30+00 |       1
  2026-08-01 00:15:30+00 |       1
  2026-08-01 05:45:00+00 |       2
  2026-08-02 12:50:15+00 |       4
  2026-08-03 22:59:59+00 |       6
  2026-08-03 22:30:00+00 |       7
  2026-08-04 10:15:45+00 |       8
  (7 rows) 

## FINAL TABLE:-

appdb=# SELECT * FROM transactions;
            date          | user_id
  ------------------------+---------
  2026-08-01 00:15:30+00 |       1
  2026-08-01 00:15:30+00 |       1
  2026-08-01 05:45:00+00 |       2
  2026-08-02 12:50:15+00 |       4
  2026-08-03 22:59:59+00 |       6
  2026-08-03 22:30:00+00 |       7
  2026-08-04 10:15:45+00 |       8
  (7 rows) 

## STOPING ENVIRONMENT
- Always close the docker after database query using command:-
1. sudo systemctl stop docker 
2. sudo systemctl stop docker.socket
- Check if inactive using command:-
1. sudo systemctl status docker
