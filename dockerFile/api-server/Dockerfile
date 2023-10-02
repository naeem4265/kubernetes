# syntax=docker/dockerfile:1

FROM golang:latest
LABEL maintainer="Naeem"
WORKDIR /app
ARG LOG_DIR=/app/logs
RUN mkdir -p ${LOG_DIR}
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log 
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

EXPOSE 8000
VOLUME [${LOG_DIR}]
CMD ["./main"]

