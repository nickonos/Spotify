# How to run the application?
## backend
### Prerequisites
- Docker
### Commands
From the root of this project, run the following command
```
docker compose up --build
```
This will start the entire backend and will show you the output logs.
If you don't care about the logs and want to run the application in the background you can use
```
docker compose up -d
```
This will run the application as a daemon. To shut down the process run
```
docker compose down
```

## Frontend
### Prerequisites
- Bun
- Linux/MacOS/WSL
### Command
Go into the "web" directory, from here you can run the
```
bun run dev
```
command to start a dev server. Or to build and run a production server run:
```
bun run build && bun run start
```


