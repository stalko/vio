The API layer should implement a single HTTP endpoint that, given an IP address, returns information about the IP address' location (e.g. country, city).


go get -u github.com/swaggo/swag/cmd/swag


swag i -g ./cmd/vioapi/main.go -o ./docs
