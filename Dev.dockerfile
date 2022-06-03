ARG GO_VERSION=1.18.1
FROM golang:${GO_VERSION}-alpine as jwt-filter-deps
WORKDIR /jwt-filter
COPY . .
RUN go mod tidy
RUN go build -buildvcs=false -o jwt-filter .


FROM jwt-filter-deps as jwt-filter-runner
WORKDIR /jwt-filter
COPY --from=jwt-filter-deps /jwt-filter/jwt-filter .
RUN chmod 755 ./jwt-filter
ENTRYPOINT ./jwt-filter
