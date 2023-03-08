FROM golang:golang:1.20

WORKDIR /go/src/app

ADD ./go_api .

RUN apt update
RUN apt install -y sqlite3

RUN go get -d entgo.io/ent/cmd/ent
RUN go install ariga.io/atlas/cmd/atlas@master

EXPOSE 8080
