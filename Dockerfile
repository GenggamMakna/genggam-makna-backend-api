FROM golang:1.23.3-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /genggam-makna-backend-api

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

WORKDIR /

COPY --from=build-stage /genggam-makna-backend-api /genggam-makna-backend-api

EXPOSE 8001

ENTRYPOINT ["/genggam-makna-backend-api"]
