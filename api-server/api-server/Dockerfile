FROM golang:latest
LABEL maintainer="Naeem"

WORKDIR /app

# COPY go.mod go.sum .
# RUN go mod download

COPY . .
RUN go mod download

RUN go build -o api-server .

EXPOSE 8080

ENTRYPOINT ["./api-server"]
