FROM golang:1.14.2

WORKDIR $GOPATH/src/github.com/SomeshSunariwal/GraphQL_implementation

RUN go install -v ./...

ENV PORT=8080

# ARG DB_PASSWORD

# ENV DB_PASSWORD $DB_PASSWORD

COPY bin/GraphQL_implementation /GraphQL_implementation

EXPOSE 8080

CMD ["make", "run"]