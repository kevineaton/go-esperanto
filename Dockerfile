FROM --platform=linux/amd64 golang:1.22 AS base

ENV GOFLAGS -mod=vendor

# Please see the readme at github.com/kevineaton/go-esperanto

ADD . /go/src/github.com/kevineaton/go-esperanto
ADD ./api/phrasebook.txt /go/src/github.com/kevineaton/go-esperanto
ADD ./api/phrasebook.json /go/src/github.com/kevineaton/go-esperanto
WORKDIR /go/src/github.com/kevineaton/go-esperanto

RUN go build -mod=vendor .

FROM --platform=linux/amd64 busybox:glibc
WORKDIR /go/src/github.com/kevineaton/go-esperanto
COPY --from=base /go/src/github.com/kevineaton/go-esperanto/go-esperanto .
COPY --from=base /go/src/github.com/kevineaton/go-esperanto/phrasebook.json .
COPY --from=base /go/src/github.com/kevineaton/go-esperanto/phrasebook.txt .
RUN pwd
RUN ls -la

CMD ["./go-esperanto"]