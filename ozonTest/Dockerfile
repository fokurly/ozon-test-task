FROM golang:onbuild
RUN go version
ENV GOPATH=/
COPY .. ./
# build go app
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o ozon-test ./cmd/web/main.go
CMD ["./ozon-test"]