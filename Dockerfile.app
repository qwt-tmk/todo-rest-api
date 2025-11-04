FROM golang:1.25.3

WORKDIR /src

COPY ./app/go.mod ./app/go.mod
COPY ./app/go.sum ./app/go.sum
WORKDIR /src/app

RUN go mod download

COPY ./ ./

WORKDIR /src/app

RUN go install github.com/air-verse/air@latest
RUN go install go.uber.org/mock/mockgen@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["air"]
