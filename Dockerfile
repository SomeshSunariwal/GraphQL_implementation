FROM golang:1.14.2

WORKDIR $GOPATH/src/github.com/SomeshSunariwal/GraphQL_implementation

ENV PORT=8080

# ARG DB_PASSWORD

# ENV DB_PASSWORD $DB_PASSWORD

COPY bin/GraphQL_implementation .

EXPOSE 8080

CMD ["./GraphQL_implementation"]

# docker run -p 8080:8080 --env-file ./env.list someshdokerbox/graphql-boiler-plate