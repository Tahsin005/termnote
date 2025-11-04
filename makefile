build:
	@go build -o termnote .

run: build
	@./termnote