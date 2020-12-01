FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY svr /app/svr
COPY app.properties /app/app.properties
CMD ["./svr"]