# go build command
build:
	@echo " >> building binaries"
	@go build -v -o nadc-intro-to-rest app.go 

# go run command
run: build
	./nadc-intro-to-rest

