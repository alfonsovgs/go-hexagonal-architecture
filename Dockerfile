FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/codelytv/golang-hex-examples
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/codelytv-mooc-api cmd/api/main.go

# Build image with the binary
FROM scratch
COPY --from=build /go/bin/codelytv-mooc-api /go/bin/codelytv-mooc-api
ENTRYPOINT [ "/go/bin/codelytv-mooc-api" ]
