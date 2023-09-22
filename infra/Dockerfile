FROM golang:1.21 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o golearn

FROM alpine:latest as app
RUN apk --no-cache add ca-certificates

WORKDIR /app

RUN addgroup -S app && adduser -S app -G app
USER app:app

COPY --from=build /app/golearn /app/golearn
COPY --from=build /app/static /app/static
COPY --from=build /app/index.html /app/index.html

ENTRYPOINT ["/app/golearn"]
