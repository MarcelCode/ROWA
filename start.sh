#!/bin/bash
echo Building Frontend
docker build --tag frontend ./frontend/.
echo Building Backend
docker build --tag backend ./backend/.
echo Starting backend
docker run --publish 3000:3000 --name backend_x --rm backend 
echo Starting frontend
docker run --publish 8080:80 --name frontend_x --rm frontend