FROM --platform=$BUILDPLATFORM rvolosatovs/protoc:4.0.0 as proto
WORKDIR /build
COPY pkg/proto pkg/proto
RUN mkdir -p pkg/pb
RUN protoc --proto_path=pkg/proto --go_out=pkg/pb \
    --go_opt=paths=source_relative --go-grpc_out=pkg/pb \
    --go-grpc_opt=paths=source_relative pkg/proto/*.proto


FROM --platform=$BUILDPLATFORM golang:1.21-alpine as builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=proto /build/pkg/pb pkg/pb
RUN env GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o profanity-filter-grpc-plugin-server-go_$TARGETOS-$TARGETARCH


FROM alpine:3.19.0
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY --from=builder /build/profanity-filter-grpc-plugin-server-go_$TARGETOS-$TARGETARCH profanity-filter-grpc-plugin-server-go
# Plugin arch gRPC server port
EXPOSE 6565
# Prometheus /metrics web server port
EXPOSE 8080
CMD [ "/app/profanity-filter-grpc-plugin-server-go" ]