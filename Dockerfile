FROM alpine:3.12

RUN mkdir "/app"
WORKDIR /app

COPY main /app/app

EXPOSE 8082

ENTRYPOINT ["./app"]

