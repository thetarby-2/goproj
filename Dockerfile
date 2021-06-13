FROM golang:1.16
RUN mkdir /code
WORKDIR /code

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /code/
RUN go build -o ./out/goproj .

EXPOSE 8080

CMD ["./out/goproj"]