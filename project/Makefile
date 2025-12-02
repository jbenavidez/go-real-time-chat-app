FRONT_END=frontEndApp
BACK_END=backEndApp
BROKER_SERVICE=brokerService


## up_build: build project and run contianer
up_build: build_broker_service build_backend build_frontend

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


build_broker_service:
	@echo "Building broker end  binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_SERVICE} ./cmd/api
	@echo "Done!"