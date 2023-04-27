include .env
APP_NAME=newsfeed
POSTGRES_NAME=$(APP_NAME)-postgres 

be_dev: postgres_start
	cd backend;	go run .
fe_dev:
	cd frontend; npm run dev
fe_install:
	npm --prefix ./frontend install ./frontend
postgres_run:
	podman run --name $(POSTGRES_NAME) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d -p $(POSTGRES_PORT):$(POSTGRES_PORT) postgres 
postgres_start:
	podman start $(POSTGRES_NAME)
postgres_stop:
	podman stop $(POSTGRES_NAME)
postgres_remove:
	podman rm -f $(POSTGRES_NAME)