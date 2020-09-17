FROM golang:1.14.4

COPY . /backend
WORKDIR /backend

RUN go build -o backend backend.go

FROM busybox
COPY --from=0 /backend/backend /app
EXPOSE 80
CMD ["/app"]
