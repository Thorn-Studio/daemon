run: build
	@./bin/vault

build:
	@go build -o ./bin/vault ./vault.go

clean:
	@rm -rf ./bin ./tmp

