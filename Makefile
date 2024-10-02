# Define the image name
IMAGE_NAME = url-shortener

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Run the Docker container
run:
	docker run -p 8080:8080 $(IMAGE_NAME)

# Remove the Docker image
clean:
	docker rmi $(IMAGE_NAME)

# Stop the running container
stop:
	docker ps -q --filter "ancestor=$(IMAGE_NAME)" | xargs -r docker stop

# Remove all stopped containers (optional)
prune:
	docker container prune -f

# Run tests
test:
	go test ./... -v

.PHONY: build run clean stop prune
