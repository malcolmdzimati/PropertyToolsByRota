FROM alpine:latest
WORKDIR /app
COPY app .

RUN chmod +x PropertyToolsByRotaBackend

EXPOSE 8080
CMD ["./PropertyToolsByRotaBackend"]