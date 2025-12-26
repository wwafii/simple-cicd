FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY static ./static

RUN go build -o /simple-cicd

EXPOSE 3000

CMD ["/simple-cicd"]