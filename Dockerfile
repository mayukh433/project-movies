FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY server/main.go ./

RUN go build -o /project-movies

EXPOSE 8081

CMD ["/project-movies"]