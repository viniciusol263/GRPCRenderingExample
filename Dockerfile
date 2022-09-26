FROM golang:1.17
WORKDIR /app/

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download -x
COPY . /app/

RUN go build -o /app/server/server /app/server/server.go
CMD [ "/app/server/server" ]
EXPOSE 8080
EXPOSE 8081