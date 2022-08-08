FROM golang:latest

ENV GOPATH=/

ENV APP_PORT=8000

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o astro ./cmd/main.go

CMD ["./astro"]