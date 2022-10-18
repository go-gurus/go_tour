# build stage
FROM golang:1.17.6-alpine AS build

RUN mkdir -p /app
WORKDIR /app

# build src
COPY go.mod .
COPY go.sum .
RUN go mod download

# app src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o beer-fridge-server ./cmd/beer-fridge-server/main.go

# result stage
FROM scratch
COPY --from=build /app/beer-fridge-server /beer-fridge-server
EXPOSE 8080
EXPOSE 443
EXPOSE 80
ENTRYPOINT ["/beer-fridge-server", "--port", "8080"]