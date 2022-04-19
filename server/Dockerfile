FROM golang:1.17-alpine

WORKDIR /grpcServer
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /server
EXPOSE 4545
CMD [ "/server" ]