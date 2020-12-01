FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY cgi /app/cgi
COPY app.properties /app/app.properties
CMD ["./cgi"]