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

#ENV GIN_MODE=release

CMD ["./main"]