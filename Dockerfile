FROM golang:1.20.0-alpine3.15 as builder

WORKDIR /go/src/github.com/go-interview

# Install tools required to build the project
RUN apk add --no-cache make git

# Copy the entire project and build it
COPY . .

RUN make cli

# Path: internal/Dockerfile
FROM alpine:3.15

# Install tools required to run the project
RUN apk add --no-cache ca-certificates

# Copy the binary from the builder image
COPY --from=builder /go/src/github.com/argoproj/argo-cd/dist/argocd-linux-amd64 /usr/local/bin/argocd

# Run the binary
ENTRYPOINT ["/usr/local/bin/argocd"]
