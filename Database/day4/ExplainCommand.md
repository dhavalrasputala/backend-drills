## Environment Setup

- Ran a local PostgreSQL database using Docker:
docker run --name go-postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=appdb -p 5432:5432 -d postgres:15
- Enter postgreSQL using command:-
docker exec -it go-postgres psql -U admin -d appdb

## SQL COMMANDS

1. EXPLAIN COMMAND:-
 
  EXPLAIN SELECT *
  FROM users WHERE AGE>20;
  
## FINAL TABLE:-

appdb=# EXPLAIN SELECT * FROM users WHERE AGE>20;
                        QUERY PLAN
----------------------------------------------------------
 Seq Scan on users  (cost=0.00..13.88 rows=103 width=234)
   Filter: (age > 20)
(2 rows))

## STOPING ENVIRONMENT
- Always close the docker after database query using command:-
1. sudo systemctl stop docker 
2. sudo systemctl stop docker.socket
- Check if inactive using command:-
1. sudo systemctl status docker
