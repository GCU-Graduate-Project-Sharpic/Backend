FROM golang:alpine

WORKDIR /Backend

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .
RUN go build -o /run-server

EXPOSE 8005

CMD [ "/run-server" ]
