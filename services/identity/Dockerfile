FROM golang:1.21.1-alpine

# Add additional packages
COPY /packages/ /packages/

WORKDIR /app/service/

COPY services/identity/go.* ./

RUN  go mod download

COPY services/identity/ ./

RUN go build

LABEL org.opencontainers.image.source=https://github.com/nickonos/Spotify/services/identity

CMD [ "./identity" ]