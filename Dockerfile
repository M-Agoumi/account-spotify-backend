# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Mohamed AGOUMI"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# install compileDaemon
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

# allow git status to work without changing files owenership
RUN git config --global --add safe.directory /app

# Expose port 8080 to the outside world
EXPOSE 3000

# Run compileDaemon
ENTRYPOINT CompileDaemon --build="go build -o /build" --command="/build"