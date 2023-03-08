FROM golang:1.20 as build
RUN mkdir /app
COPY . /app
WORKDIR /app
ENV TZ = "Europe/Moscow"
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/repair/main.go


FROM alpine:latest
RUN apk --no-cache add ca-certificates
ENV RU UTC
COPY --from=build /app .
EXPOSE 8888
CMD ["./app"]