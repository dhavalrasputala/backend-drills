## Day 13: MESSAGE QUEUES

**My Explanation**

- `Message Queues` refers to a queue manatained by the producer(server) to maintain an asynchronous communication with heavy traffic of consumers(client).
- **Componets of `Message Queues`**
1. Message Produer:- message producer is responsible for sending messages to the `message queues` from each client
2. `Message Queue`:- a special FIFO(first in first out) queue mainted by the server to track each message individually.
3. Message Consumer:- responsible for receving the message & working on it for producing the desired results.

## The Flow

Server --> Message Producer --> `Message Queue` --> Worker(message consumer).

## Trade-offs / Why it matters

- Improved Performance:-helps maintain server performance even during huge traffic & handling client request with ease.
- Maintains Asynchronous Communication:-Servers can send or receive messages without waiting for response due to message queues.
- Decoupling:-message queues allows server to decouple applications from each other , hence allowing independent work.
