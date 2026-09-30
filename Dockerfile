# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY go.mod ./
COPY *.go ./
COPY web ./web
RUN CGO_ENABLED=0 go build -o /out/gen-phones .

FROM alpine:3.20

# PORT is read from .env at build time by build.sh / docker-compose and
# baked in here as the image's default; it can still be overridden at
# runtime with `docker run -e PORT=...`.
ARG PORT=8080
ENV PORT=${PORT}

RUN adduser -D -H app
WORKDIR /app
COPY --from=builder /out/gen-phones ./gen-phones
COPY .env ./.env

USER app
EXPOSE ${PORT}
ENTRYPOINT ["./gen-phones"]
