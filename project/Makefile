FRONT_END=frontEndApp
BACK_END=backEndApp


## up_build: build project and run contianer
up_build: build_frontend build_backend

	@echo "Starting Docker images..."
	docker-compose up 
	@echo "Docker images started!"


build_frontend:
	@echo "Building front end  binary..."
	cd ../front-end && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONT_END} ./cmd/web
	@echo "Done!"


build_backend:
	@echo "Building front end  binary..."
	cd ../back-end && env GOOS=linux CGO_ENABLED=0 go build -o ${BACK_END} ./cmd/api
	@echo "Done!"