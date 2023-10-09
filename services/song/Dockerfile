FROM golang:1.21.1-alpine

# Add additional packages
COPY /packages/ /packages/

WORKDIR /app/service/

COPY services/song/go.* ./

RUN  go mod download

COPY services/song/ ./

RUN go build

LABEL org.opencontainers.image.source=https://github.com/nickonos/Spotify/services/song

CMD [ "./song" ]