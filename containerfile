# Stage 1
# FROM golang:1.23-alpine3.20 AS BUILD
FROM golang:alpine AS BUILD

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src

# Cache dependencies
COPY go.mod .
# COPY go.sum .
RUN go mod download

# Copy actual source
COPY . .

RUN CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags='-s -w -extldflags "-static"' -o ./out/bin .

# Stage 2
# FROM alpine:3.20
FROM alpine:latest

ARG PORT=8080

WORKDIR /app

RUN addgroup gopher_group \
  && adduser -S gopher -u 1000 -G gopher_group

RUN apk update \
  && apk -U upgrade \
  && apk add --no-cache ca-certificates \
  && update-ca-certificates --fresh \
  && rm -rf /var/cache/apk/*

COPY --from=BUILD /src/out/bin .

RUN chmod +x ./bin

USER gopher

EXPOSE ${PORT}
# executable for the image
ENTRYPOINT ["/app/bin"]
# default arguments for the executable, if applicable
# CMD [ "run" ]
