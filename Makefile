build-hello-plugin:
	cd hello && go build && cp hello ../plugins

build-hi-plugin:
	cd hi && go build && cp hi ../plugins

run: build-hello-plugin build-hi-plugin
	go run main.go
	 