all: agent.go module
	go build .

clean:
	rm -f go.mod
	rm -f go.sum
	rm -f agent

module:
	go mod init challenge/agent
	go mod tidy

install: agent
	install -m 0755 agent /usr/local/bin/
