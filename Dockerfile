FROM alpine:3.12
ENV TZ="Asia/Jakarta"
WORKDIR /opt
COPY .env .
COPY main .
RUN chmod +x main && apk add --no-cache ca-certificates
EXPOSE 5000
ENTRYPOINT ./main
