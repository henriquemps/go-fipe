FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go install github.com/pilu/fresh

EXPOSE 8080

ENTRYPOINT ["fresh"]