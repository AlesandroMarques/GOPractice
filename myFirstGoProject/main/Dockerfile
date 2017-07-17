FROM golang

add . /go/src/myFirstGoProject/main

WORKDIR /go/src/myFirstGoProject/main 

RUN go get .
RUN go build .

ENTRYPOINT ./main