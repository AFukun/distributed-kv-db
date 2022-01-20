build:
	go build -o ./bin/server github.com/AFukun/distributed-kv-db/server
	go build -o ./bin/leader github.com/AFukun/distributed-kv-db/leader

clean:
	rm -rf ./bin/*
