# Stage 1
# Podman sets BUILDPLATFORM, TARGETOS and TARGETARCH automatically
FROM --platform=$BUILDPLATFORM golang:alpine AS BUILD
ARG TARGETOS TARGETARCH
ENV CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH
WORKDIR /src

COPY go.* ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go build -trimpath -ldflags='-s -w' -o /out/bin .

# Stage 2
FROM gcr.io/distroless/static-debian13:nonroot
ARG PORT=8080
ENV PORT=${PORT}
COPY --from=BUILD /out/bin /app/bin
EXPOSE ${PORT}
ENTRYPOINT ["/app/bin"]
