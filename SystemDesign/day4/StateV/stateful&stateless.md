## Day 4: Stateful V/S Stateless Server 
- My Explanation

- Server side databases are mainly of two types stateless and stateful,stateless database doesn't store client's imformation or past visit which make them memory efficent , whereas stateful databases store client's past visits and imformation and retrive them each from database each time client visits the server hence making them more memory expensive.
- Both have thier pro's and con's and both are used equally. 

## The Flow

- N/A

## Trade-offs / Why it matters

- Stateless:-Stateless architectures excel in isolated, short-term environments where individual requests don't need to retain historical context from previous actions.Common use cases involve:-Restful APIs,microservices,LLM's.

- Stateful:- Stateful applications are necessary when systems require persistent context, continuous feedback loops, or ongoing tracking of user behaviors.Common use cases involve:-User centric application,IoT systems,AI/ML model training.
