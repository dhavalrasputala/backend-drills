## Day 1: Client and Server
- My Explanation

- A client is usually the user's browser or mobile app. Its job is to render the UI and capture user clicks. The server is a program running on a remote machine that holds the business logic and the data. The client doesn't know how to calculate a user's balance; it just asks the server via an HTTP request, and the server replies with the answer.

## The Flow

- User Browser --(HTTP Request)--> Go Server --(Process Data)-->Response sent back

## Trade-offs / Why it matters

- Separating the client and server means we can update our backend logic without forcing the user to download a new app. The client stays lightweight, and the server handles the heavy lifting securely.
