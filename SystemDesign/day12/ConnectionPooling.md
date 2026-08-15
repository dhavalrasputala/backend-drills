## Day 12:Connection Pooling

**My Explanation:**

- `Connection Pooling` refers to pools of connections with database which are kept open instead of making an new connection for every request.this helps in optimizing the performance of the systems by fetching the results faster. 

## The Flow

-  Application Start --> initialize connections --> connection pool --> utilize the connections for every request

## Trade-offs / Why it matters

- Improved Performance:-helps by reducing overhead by reducing the number of database opening & closing opertions performed by every request.
- Resource Efficiency:-`connection pooling` helps optimize the number of idle connections and active connections with the database.
- Concurrency Control:-`connection pooling` helps maintain concurrency by defining the number of simultanous connection to database.
