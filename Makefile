.DEFAULT_GOAL := build
	# If the only command provided is make then build functionality is operated

.PHONY:fmt vet build test run clean # makes these words 'Phony' so they are not treated as files
	# [./...] means do that recursively so all the files end up being vet, fmt, test
fmt: # formats the code using go formatter 
	go fmt ./...

vet:fmt # checks for unreachable code
	go vet ./...

test:vet # for tests 
	go test ./...

build:vet # doesnt rely on tests 
	go build -o output
run:build
	./output
clean: # removes the binary 
	rm -rf ./output
