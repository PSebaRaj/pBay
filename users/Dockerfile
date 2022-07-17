FROM golang:1.17-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./

RUN go build -v -o /server

EXPOSE 9000

CMD [ "/server" ]
