check_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger 
swagger:check_install
	MODULE111=off swagger generate spec -o ./swagger.yaml --scan-models