FROM golang:1.21.1-alpine

# Add additional packages
COPY /packages/ /packages/

WORKDIR /app/service/

COPY services/authorization/go.* ./

RUN  go mod download

COPY services/authorization/ ./

RUN go build

LABEL org.opencontainers.image.source=https://github.com/nickonos/Spotify/services/authorization

CMD [ "./authorization" ]