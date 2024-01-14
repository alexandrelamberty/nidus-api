FROM golang:1.19-alpine
WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o /api cmd/main.go
EXPOSE 3333
CMD [ "/api" ]
# Fiber config Prefork
# https://docs.gofiber.io/api/fiber#config
# https://github.com/gofiber/fiber/issues/1021#issuecomment-730537971
# CMD ["sh", "-c", "/api"]