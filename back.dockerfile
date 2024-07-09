FROM golang:1.22

WORKDIR /project

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["air", "-c", ".air.toml"]