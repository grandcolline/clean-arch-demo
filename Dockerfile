ARG GO_VERSION=latest

# Develop & Build Stage
FROM golang:${GO_VERSION} as develop

ENV GOOS=linux
ENV GO111MODULE=on

# install fresh (for development)
RUN GO111MODULE=off go get github.com/pilu/fresh

WORKDIR $GOPATH/src/github.com/grandcolline/clean-arch-demo

# package download
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# build
COPY . .
RUN env CGO_ENABLED=0 go install

# run (for development)
CMD ["fresh"]

# Production Stage
FROM gcr.io/distroless/static as production

COPY --from=develop /go/bin/clean-arch-demo /clean-arch-demo
CMD ["/clean-arch-demo"]
