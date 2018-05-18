# Multi-Stage build
# Build stage
FROM golang:alpine AS build-env
ENV SRC_DIR=/go/src/github.com/Mardiniii/serapis/
COPY . $SRC_DIR/
RUN cd $SRC_DIR; go build cmd/evaluator/main.go;

# Final stage
FROM alpine
ENV SRC_DIR=/go/src/github.com/Mardiniii/serapis/
WORKDIR /app
COPY --from=build-env $SRC_DIR/ /app/
ENTRYPOINT ["./main"]
