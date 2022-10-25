gen:
	goctl api go -api helloworld.api -dir .
	goctl api plugin -plugin goctl-swagger="swagger -filename helloworld.json" -api helloworld.api -dir .
	goctl api doc --dir .

run:
	go run helloworld.go -f etc/helloworld-api.yaml