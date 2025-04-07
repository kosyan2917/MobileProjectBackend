FROM golang:latest as build

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build main.go
EXPOSE 1337
CMD [ "./main" ]