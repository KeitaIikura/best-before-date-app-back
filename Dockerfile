FROM golang:1.19rc1-stretch
WORKDIR /app
COPY . .

RUN go install github.com/cosmtrek/air@latest
EXPOSE 8080
CMD ["air", "-c", "/app/air.conf"]
