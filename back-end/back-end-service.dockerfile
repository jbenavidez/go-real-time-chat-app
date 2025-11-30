FROM alpine:latest

RUN mkdir /app

COPY backEndApp /app

CMD [ "/app/backEndApp"]