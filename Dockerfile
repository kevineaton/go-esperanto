FROM golang

# Please see the readme at github.com/kevineaton/go-esperanto

RUN apt-get update 
RUN apt-get install -y bash git nginx

ADD . $GOPATH/src/github.com/kevineaton/go-esperanto
RUN go get github.com/gin-gonic/gin
RUN go install github.com/kevineaton/go-esperanto

ADD ./phrasebook.txt /go/bin
ADD ./phrasebook.json /go/bin

ENTRYPOINT /go/bin/go-esperanto

EXPOSE 8081