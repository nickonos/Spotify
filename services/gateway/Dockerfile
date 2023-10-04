FROM golang:1.21.1-alpine

# Add additional packages
COPY /packages/ /packages/

WORKDIR /app/service/

COPY services/gateway/go.* ./

RUN  go mod download

COPY services/gateway/ ./

RUN go build

LABEL org.opencontainers.image.source=https://github.com/nickonos/Spotify/gateway

EXPOSE 8080

CMD [ "./gateway" ]