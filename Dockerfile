# syntax=docker/dockerfile:1

FROM golang:1.22 as builder
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /out/app ./cmd/server

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /out/app /app/app
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/app/app"]