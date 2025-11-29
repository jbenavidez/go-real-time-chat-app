FRONT_END=frontEndApp


## up_build: build project and run contianer
up_build: build_frontend 

	@echo "Starting Docker images..."
	docker-compose up 
	@echo "Docker images started!"


build_frontend:
	@echo "Building front end  binary..."
	cd ../front-end && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONT_END} ./cmd/web
	@echo "Done!"


