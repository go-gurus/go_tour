# build stage
FROM golang:1.17.6-alpine AS build

RUN mkdir -p /app
WORKDIR /app

# app src
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /bin/app main.go

# result stage
FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]