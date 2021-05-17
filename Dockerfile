FROM golang:1.14.0

COPY ./code /go/src/docker0

RUN go get -v -u github.com/gorilla/mux

RUN go install docker0

EXPOSE 8080/tcp

CMD ["docker0"]