## Day 10: DATABASE Replication
- My Explanation

- `Database Replication`refers to creating multiple copyies of databases to distribute the workload & also server data in case of system failure mainly there are 5 types of database replication:-
1. Master-Slave Replication:-Only master database has the rights to use insert,delete,update query slave databases have rights to read querys only.
2. Master-Master Replication:-thier are many master databases in system and all have rights to read&write opertions
3. Snapshot Replication:-Only part of database is copied to slave database by taking snapshot at specific point of time 
4. Transactional Replication:-Synchronizes databases as the change occurs
5. Merge Replication:-allows databases to update independely & synchronizes later

## The Flow

-  Write Queries --> Master Database --> Copies Data into slave database --> Replicate Database <-- Read Queries

## Trade-offs / Why it matters

- Ensure high availabilty of data by copying it multiple times.
- improves load balancing by allowing load balancer to send requests to replica database
- provides fault tolerance by making sure data is available all the time.
