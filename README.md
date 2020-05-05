# ROWA

## About
ROWA is an indoor farm designed for offices to enhace the workspace by greening it while offering the employees fresh greens.
This repo represents the software running locally on the farm.
The project is sepretaed into backend and frontend. Both parts reside in individual docker containers and can be run locally.
## Architecture
Current running production architecture:
![Image of Architecture](https://github.com/Emil9999/ROWA/blob/dev/documentation/images/prodsetup.png)

In development architecture:
![Image of Future Architecture](https://github.com/Emil9999/ROWA/blob/dev/documentation/images/architecture.png)

## Requirements
Docker

## Building and running
In project folder:
`./start.sh`
(You might have to chmod first)
c
## Alternatively:

### Building the app 
`docker build --tag backend ./backend/.` to build the backend and
`docker build --tag frontend ./frontend/.` to build the frontend.

### Running it
`docker run --publish 8080:80 --name frontend_x --rm frontend`
`docker run --publish 3000:3000 --name backend_x --rm backend`

Then navigate to `localhost:8080`.
Note that this version does not push anything to aws and thus is completely offline.

