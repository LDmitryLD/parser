FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY ./ /app

RUN go build -o main cmd/main.go

FROM migrate/migrate AS migrator 

COPY migrations /migrations

FROM alpine 

COPY --from=builder app/main /main

COPY --from=builder app/public /public

COPY --from=migrator /migrations /migrations

CMD [ "/main" ]