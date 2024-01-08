FROM golang:alpine as first_stage

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go mod tidy

COPY . /app

EXPOSE 8080

RUN go build -o ozonApp cmd/main.go

FROM alpine

COPY --from=first_stage /app/ozonApp .

CMD ["./ozonApp"]

