FROM golang:1.23.3

WORKDIR /app

ENV PORT=8080

COPY . .

EXPOSE 8080

CMD [ "go", "run", "main.go" ]
