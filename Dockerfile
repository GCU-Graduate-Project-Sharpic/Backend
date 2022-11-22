FROM golang:alpine

WORKDIR /Backend
COPY . .
RUN go mod download

RUN go build -o /run-server

EXPOSE 8005

CMD [ "/run-server" ]
