module github.com/nickonos/Spotify/services/identity

go 1.21.1

// import local packages
replace github.com/nickonos/Spotify/packages/logging => ../../packages/logging

replace github.com/nickonos/Spotify/packages/routes => ../../packages/routes

replace github.com/nickonos/Spotify/packages/broker => ../../packages/broker

require github.com/nickonos/Spotify/packages/logging v0.0.0

require github.com/nickonos/Spotify/packages/routes v0.0.0

require github.com/nickonos/Spotify/packages/broker v0.0.0

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/nats-io/nats.go v1.30.2
)

require (
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/nats-io/nats-server/v2 v2.10.1 // indirect
	github.com/nats-io/nkeys v0.4.5 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/rs/zerolog v1.31.0 // indirect
	golang.org/x/crypto v0.13.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
