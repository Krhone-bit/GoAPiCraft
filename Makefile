# install packages
install:
	go mod tity

# run with air
run:
	air

# generate docs files
swagger:
	swag init

# generate air.toml
air:
	air init