FROM golang:1.14.4

COPY . /backend
WORKDIR /backend

RUN go build -o backend backend.go

EXPOSE 80

CMD /backend/backend
