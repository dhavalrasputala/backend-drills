## Day 11: Replication LAG
- My Explanation

- `Replication Lag` refers to delay between data update occuring on one node and the update propogating or showcasing in other node.In `distributed systems` data replication is one of the most important things for fault tolerance & performance.
- Mainly there are 7 types of Replication Lag
1. Network Induced lag:-Occures when thier is delay in transferring of data acorss nodes due to network latency.
2. Write Latency Lag:-Occurs when the writing speed of the source node is faster then replication speed. 
3. Disk I/O Lag:-Occurs due to time taken by the system to implement the changes or showcase it into the system.
4. Backlog Lag:-Occurs when thier are burst amount of pending write operations to be implemented followers nodes.
5. Processing Lag:-Occurs due to limited processing power of the servers.
6. Clock Skew:-Occurs due to time synchronization issues casued due to source node.
7. Leader-Follower Lag:-Occurs due to unsynchronized leader & followers node.

## The Flow

-  N/A

## Trade-offs / Why it matters

- Causes Data Inconsistency.
- Stale Reads.
- Potential Data Loss.
