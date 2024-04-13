FROM golang:1.22.2-bullseye AS build
WORKDIR /usr/src/app
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go fmt ./... && \
    go mod tidy -v

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build -ldflags='-s -w' -o=./bin/api ./cmd/api

FROM gcr.io/distroless/static AS final
WORKDIR /bin
COPY --from=build /usr/src/app/bin/api ./api

CMD ["./api"]