# vi: ft=Dockerfile:

ARG GO_VERSION=1.18

FROM --platform=$BUILDPLATFORM golang:$GO_VERSION AS go

RUN apt-get update && apt-get dist-upgrade -y && apt-get install -y build-essential git

WORKDIR $GOPATH/src/sylr.dev/fix

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# -----------------------------------------------------------------------------

FROM --platform=$BUILDPLATFORM go AS builder

ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

# Switch shell to bash
SHELL ["bash", "-c"]

# Run a git command otherwise git describe in the Makefile could report a dirty git dir
RUN git diff --exit-code || true

RUN make build GOOS=${TARGETOS} GOARCH=${TARGETARCH} GOARM=${TARGETVARIANT/v/} GO_BUILD_TARGET=dist/${TARGETPLATFORM}/fix

# -----------------------------------------------------------------------------

FROM scratch

ARG TARGETPLATFORM

WORKDIR /usr/local/bin

COPY --from=builder "/go/src/sylr.dev/fix/dist/$TARGETPLATFORM/fix" .

CMD ["/usr/local/bin/fix"]