#dev
FROM golang:1.16.4-alpine as base
WORKDIR /go/todo/base
COPY src ./src
COPY go.mod go.sum ./
RUN apk update && apk --no-cache add git && go get github.com/cosmtrek/air@v1.27.3
RUN go mod tidy
WORKDIR /go/todo/base/src
CMD ["air", "-c", ".air.toml"]
EXPOSE 80

#builder
FROM golang:1.16.4-alpine as builder
COPY src ./go/todo/builder/src
WORKDIR go/todo/builder/src
COPY go.mod go.sum ./
RUN apk update && apk --no-cache add git && go get github.com/cosmtrek/air@v1.27.3
RUN go mod tidy
RUN CGO_ENABLE=0 GOOS=linux go build -o /go/todo/builder/binary


##production
FROM alpine as production
WORKDIR go/todo/production
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/todo/builder/binary /go/todo/production
EXPOSE 80
CMD ["/go/todo/production/binary"]


