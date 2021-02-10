FROM golang:1.15

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o loginservice cmd/main/main.go
EXPOSE 3993
CMD ["./loginservice"]