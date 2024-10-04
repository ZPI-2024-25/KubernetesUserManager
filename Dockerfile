FROM golang:1.22.5

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /client-go-test

EXPOSE 8080

CMD [ "/client-go-test" ]
