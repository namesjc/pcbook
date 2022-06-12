gen:
	protoc -I proto --go_out=. --go_opt=module=github.com/namesjc/pcbook  --go-grpc_out=. --go-grpc_opt=module=github.com/namesjc/pcbook  proto/*.proto


clean:
	rm pb/*.go

run:
	go run main.go

test:
	go test -cover -race ./...
