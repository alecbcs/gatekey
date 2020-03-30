FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . . 

RUN go build -o gatekey .

EXPOSE 8080

VOLUME [/config]

CMD ["./gatekey"]