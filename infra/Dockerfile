FROM golang:1.21.1-alpine3.18 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o golearn

FROM alpine:3.18 as app
WORKDIR /app
COPY --from=build /app/golearn /app/golearn
COPY --from=build /app/static /app/static
COPY --from=build /app/index.html /app/index.html

ENTRYPOINT ["/app/golearn"]
