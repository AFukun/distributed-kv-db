# Simple K-V Distributed Database

## Build Excutable

```shell
$ go mod tidy
$ make build
```

## Run Example

```shell
# Shell 1 
$ ./bin/server 9000
# Shell 2
$ ./bin/server 9001
# Shell 3
$ ./bin/server 9002

# Shell 0 (client should be running after server initialized)
$ ./bin/leader
```