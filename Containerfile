FROM golang:1.22.0-bullseye AS build
WORKDIR /usr/src/app
ENV CGO_ENABLED=0

RUN apt-get update && apt-get install -y make && apt-get clean

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make tidy
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" make build

FROM gcr.io/distroless/static AS final
ENV APP_HOME=/home/app
WORKDIR $APP_HOME
COPY --from=build /usr/src/app/bin/api ./api

RUN chown -R nonroot:nonroot $APP_HOME
USER nonroot

CMD ["./api"]