# build
build:
	go build -o bin/main cmd/root.go

# run
run:
	bin/main


# generate wire_gen.go
wire_gen:
	wire ./cmd/...
