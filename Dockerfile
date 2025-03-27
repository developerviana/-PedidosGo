FROM golang:1.23


WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o pedidosgo ./cmd

CMD [ "./pedidosgo" ]

EXPOSE 8080
