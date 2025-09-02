FROM golang:1.25

COPY . .
RUN go mod download

RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
