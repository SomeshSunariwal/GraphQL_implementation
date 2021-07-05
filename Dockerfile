FROM golang:1.14.2

WORKDIR $GOPATH/src/github.com/SomeshSunariwal/GraphQL_implementation

COPY . .

RUN go install -v ./...

ENV PORT=8080

# ARG DB_PASSWORD

# ENV DB_PASSWORD $DB_PASSWORD

EXPOSE 8080

CMD ["make", "run"]

# docker build -e PG_PASSWORD=value