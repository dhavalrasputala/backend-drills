## Environment Setup

- Ran a local PostgreSQL database using Docker:
docker run --name go-postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=appdb -p 5432:5432 -d postgres:15
- Enter postgreSQL using command:-
docker exec -it go-postgres psql -U admin -d appdb

## SQL COMMANDS

1. ALTER TABLE:-
 
  ALTER TABLE users
  ADD COLUMN gender
  VARCHAR(1) DEFAULT 'M';
  
## FINAL TABLE:-
appdb=# SELECT * FROM users;
 id | name  | age | gender
----+-------+-----+--------
  1 | BOB   |  23 | M
  2 | EMILY |  23 | M
(2 rows)

## STOPING ENVIRONMENT
- Always close the docker after database query using command:-
1. sudo systemctl stop docker 
2. sudo systemctl stop docker.socket
- Check if inactive using command:-
1. sudo systemctl status docker
