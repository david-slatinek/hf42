FROM golang:1.20-alpine3.17 AS build

WORKDIR /api

COPY . .

RUN go mod download -x
RUN go build -v main.go


FROM alpine:3.17

WORKDIR /api

COPY --from=build /api/main ./
RUN mkdir env
COPY --from=build /api/env/.env ./env

RUN mkdir logo
COPY --from=build /api/logo ./logo

ENV GIN_MODE=release
ENV BOOK_URL=book-service

CMD ["./main"]