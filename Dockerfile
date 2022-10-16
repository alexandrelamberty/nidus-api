FROM golang:1.19-alpine
WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o /api cmd/api/main.go
EXPOSE 3333
CMD [ "/api" ]