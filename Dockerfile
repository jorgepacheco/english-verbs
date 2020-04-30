FROM golang:1.14 AS build-env
#disable crosscompiling
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY . /go/src/english-verbs/server
WORKDIR /go/src/english-verbs/server
RUN go build -o /go/bin/server .


FROM alpine
EXPOSE 80
EXPOSE $PORT
COPY --from=build-env /go/bin/server /app/
COPY --from=build-env /go/src/english-verbs/server/data /app/data
WORKDIR /app
CMD ["./server"]