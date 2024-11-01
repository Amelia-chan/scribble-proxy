FROM golang:1.23.2-alpine AS build

WORKDIR /usr/src/app
COPY go.* ./
RUN go mod download && go mod verify

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -installsuffix cgo -ldflags="-w -s" -o /usr/local/bin cmd/amelia.go

FROM gcr.io/distroless/static-debian12

COPY --from=build /usr/local/bin /usr/local/bin

EXPOSE 7631
CMD ["amelia"]