dev:
	go run cmd/main.go example $(ARGS)

exec:
	go run cmd/main.go real $(ARGS)

test-all:
	go test ./... -v -run 'TestAllAOC'

test-today:
	go test ./... -v -run 'TestUnitAOC' 