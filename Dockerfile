FROM golang:1.19
LABEL authors="brandalik"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY static ./static
RUN CGO_ENABLED=0 GOOS=linux go build -o /todos

CMD ["/todos"]
