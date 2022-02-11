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
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app

# result stage
FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]