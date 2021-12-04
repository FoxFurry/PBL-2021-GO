swagger:
	swagger generate spec -o ./gen/swagger.yaml --scan-models

serve-swagger:
	swagger serve -F=swagger ./gen/swagger.yaml

swag-and-server: swagger serve-swagger

format:
	go fmt ./...
	go vet ./...