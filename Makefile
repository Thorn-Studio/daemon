run: build
	@./bin/vault

build:
	@go build -o ./bin/vault ./vault.go

clean:
	@rm -rf ./bin ./tmp ~/.local/bin/vault

localinstall: clean build
	@cp ./bin/vault ~/.local/bin/
	@chmod +x ~/.local/bin/vault

