FRONTEND_BINARY=frontend

up:
	@echo "Starting docker compose..."
	docker compose up -d

down:
	@echo "Stopping docker compose..."
	docker compose down

build:
	@echo "Rebuilding & starting docker compose..."
	docker compose up --build -d

clean:
	@echo "Cleaning containers..."
	docker container prune --force
	@echo "Cleaning images..."
	docker image prune --force
	@echo "Cleaning networks..."
	docker network prune --force
	@echo "Cleaning volumes..."
	docker volume prune --force

start-frontend:
	@echo "Building the frontend binary..."
	cd frontend && go build -o ./${FRONTEND_BINARY} ./cmd/web
	@echo "Starting the frontend..."
	cd frontend && ./cmd/web/${FRONTEND_BINARY} &

stop-frontend:
	@echo "Stopping the frontend..."
	pkill ${FRONTEND_BINARY}
