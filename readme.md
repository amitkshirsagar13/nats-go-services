### Install go packages
- nats-server: `go install github.com/nats-io/nats-server/v2@latest`
- nats-cli: `go install github.com/nats-io/nats.go@latest`

### Setup go module
- `go mod init cmd/main`
- `go mod tidy`