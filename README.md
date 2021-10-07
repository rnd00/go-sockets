# go-sockets
---

An attempt to make a connection between client and server through tcp socket.

Important;

- Run server first, and then client after
- Client can be more than just 1

---

## Server

- Server will run on the specified parameter (written in the code)
- If the assigned port is being used for something else, it will exit(1)
- for{} loop inside main thread will keep running until it's being forced to close
- for{} loop was made to receive/accept clients' connections
- each connection will have its own thread.
- functions is written inside `helpers` folder

## Client

- Client will connect to the specified parameter (written in the code)
- Client will use random port for dialing