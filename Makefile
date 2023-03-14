create:
	cp build/local/.env.example build/local/.env
	PWD=$$(PWD) entrypoint="" docker-compose --file build/local/docker-compose.yaml build pinger
up:
	PWD=$$(PWD) entrypoint="test -f init.toml || cp init_local.toml init.toml && go run ./cmd/pinger/main.go" docker-compose --file build/local/docker-compose.yaml up pinger
down:
	PWD=$$(PWD) entrypoint="" docker-compose --file build/local/docker-compose.yaml down --remove-orphans
bash:
	docker-compose --file build/local/docker-compose.yaml exec pinger bash
test:
	go test -race -v -timeout 10s ./...