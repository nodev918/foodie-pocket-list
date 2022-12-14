FROM golang:1.16-alpine

WORKDIR /app

COPY . /app

RUN go mod download
RUN go build .

EXPOSE 8888

CMD [ "./foodie-pocket-list"]