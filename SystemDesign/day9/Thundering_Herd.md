## Day 9: The Thundering Herd Problem
- My Explanation

- The thundering herd problem occurs when many process are fired simulutaneously causing traffic burst due to this CPU power also increases.During this type of Surges , only handful ammount of threads or process get the required resources to complete their work causing other process or threads to eventually go to sleep or in deadlock.

## The Flow

-  Cache Expired --> 10,000 Requests Arrive --> All miss cache --> all queryies go into database --> Database overloaded`(thunder herd)`

## Trade-offs / Why it matters

- Rate limiting the request can help solve this issue.
- another good option is `single flight` only one reuqest fetches data other waits.
- this techniques reduces the load on database.
