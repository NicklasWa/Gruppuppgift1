FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY data/*.go ./data/
COPY templates/* ./templates/

RUN CGO_ENABLED=1 go build -o /applikationen

EXPOSE 8080

CMD [ "/applikationen" ]
