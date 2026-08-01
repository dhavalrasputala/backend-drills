## Environment Setup

- Ran a local PostgreSQL database using Docker:
docker run --name go-postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=appdb -p 5432:5432 -d postgres:15
- Enter postgreSQL using command:-
docker exec -it go-postgres psql -U admin -d appdb

## SQL COMMANDS

1. CREATE INDEX COMMAND:-
 
  CREATE INDEX idx_age
  ON users(age);
  
## FINAL TABLE:-

appdb=#CREATE INDEX idx_age ON users(age);
CREATE INDEX 0
appdb=# SELECT * FROM users;
id |     name      | age | gender
----+---------------+-----+--------
  1 | BOB           |  23 | M
  2 | EMILY         |  23 | M
  4 | Alice Johnson |  28 | F
  5 | Brian Smith   |  34 | M
  6 | Carla Mendes  |  22 | F
  7 | David Kim     |  41 | M
  8 | Emma Wilson   |  19 | F
  9 | Farhan Ali    |  30 | M
 10 | Grace Lee     |  25 | F
 11 | Hassan Raza   |  37 | M
 12 | Isla Brown    |  45 | F
 13 | Jacob Turner  |  29 | M
 14 | Kavya Nair    |  33 | F
 15 | Liam O'Connor |  26 | M
 16 | Maria Garcia  |  39 | F
 17 | Noah Bennett  |  21 | M
 18 | Olivia Chen   |  31 | F
(17 rows)

## STOPING ENVIRONMENT
- Always close the docker after database query using command:-
1. sudo systemctl stop docker 
2. sudo systemctl stop docker.socket
- Check if inactive using command:-
1. sudo systemctl status docker
