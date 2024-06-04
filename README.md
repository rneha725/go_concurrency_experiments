Notes:
- the sending and receiving on channel are sync functions, goroutine wait to send and receive.
  - to solve this issue, we also have buffering channels [not covered]
- `select`: it works like a `switch` block and contains `case` and `default`. `select` takes the communication with a channel as `case`s and blocks till there is communication remaining. A communication could be in or out both. When there are multiple choices available to it, this statement chooses one randomly.


### Sequence
1. fan-in.go
2. sequencing.go
3. select.go
4. timeouts.go
5. quit.go
6. lightweight_go_routines.go