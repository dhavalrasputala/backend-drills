## Environment Setup

- Ran a local PostgreSQL database using Docker:
docker run --name go-postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=appdb -p 5432:5432 -d postgres:15
- Enter postgreSQL using command:-
docker exec -it go-postgres psql -U admin -d appdb

## SQL COMMANDS

1. CREATE TABLE 
  CREATE TABLE users (
      id SERIAL PRIMARY KEY,
      name VARCHAR(100),
      age INT
  );

2. INSERT DATA 
  INSERT INTO users (name, age) VALUES ('Alice', 28);
  INSERT INTO users (name, age) VALUES ('Bob', 22);
  INSERT INTO users (name, age) VALUES ('Charlie', 35);

3. READ DATA
  SELECT * FROM users;

## STOPING ENVIRONMENT
- Always close the docker after database query using command:-
1. sudo systemctl stop docker 
2. sudo systemctl stop docker.socket
- Check if inactive using command:-
1. sudo systemctl status docker
