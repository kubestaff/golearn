FROM golang:1.21.3-bullseye AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o golearn

FROM debian:stable-slim as app

WORKDIR /app

COPY --from=build /app/golearn /app/golearn
COPY --from=build /app/static /app/static
COPY --from=build /app/html /app/html
COPY --from=build /app/usersDatabase.json /app/usersDatabase.json
COPY --from=build /app/videosDatabase.json /app/videosDatabase.json

ENTRYPOINT ["/app/golearn"]
