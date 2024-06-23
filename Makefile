.PHONY: build run swagger local

# Build the application
build:
    docker-compose build

# Run the application with live reloading
run:
    docker-compose up

# Generate Swagger documentation
swagger:
    docker-compose run --rm app sh -c 'swag init -g ./main.go'

# Start the application with temporary Swagger generation and live reloading
local: swagger
    docker-compose run --rm app sh -c 'air -c .air.local.toml'
