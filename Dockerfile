FROM golang:1.13-alpine AS build_base
# Install some dependencies needed to build the project
RUN apk add bash ca-certificates git gcc g++ libc-dev
# We want to populate the module cache based on the go.{mod,sum} files.
WORKDIR /app
# Here we copy the rest of the source code
COPY . .
# And compile the project
RUN go build './server.go'
#In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
# Expose port 8080 to the outside world
EXPOSE 50053
# Finally we copy the statically compiled Go binary.
# Run the executable
ENTRYPOINT ["./server"]
