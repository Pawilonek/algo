
run:
	go run main.go

test:
	go test ./...

bench-queue:
	cd queue; go test -bench=.
