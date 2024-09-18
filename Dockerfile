FROM golang:1.23-alpine

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./pkg/group/go.mod ./pkg/group/go.sum ./pkg/group/
RUN cd ./pkg/group && go mod download && go mod verify && cd ../..

COPY ./pkg/expense/go.mod ./pkg/expense/go.sum ./pkg/expense/
RUN cd ./pkg/expense && go mod download && go mod verify && cd ../..

COPY ./pkg/router/go.mod ./pkg/router/go.sum ./pkg/router/
RUN cd ./pkg/router && go mod download && go mod verify && cd ../..

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 8080

CMD ["app"]
