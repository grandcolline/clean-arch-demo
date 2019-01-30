# Build Stage
FROM golang:latest as build

ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/grandcolline/clean-arch-demo
COPY . .
RUN env CGO_ENABLED=0 go install

# Release Stage
FROM gcr.io/distroless/base

COPY --from=build /go/bin/clean-arch-demo /clean-arch-demo
ENV PORT=8080
EXPOSE 8080

CMD ["/clean-arch-demo"]
