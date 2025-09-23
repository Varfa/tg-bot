#!/bin/bash

# Папка проекта (где лежит deploy.sh)
PROJECT_DIR="$(dirname "$0")"
cd "$PROJECT_DIR" || exit

echo "Pulling latest changes from GitHub..."
git pull

echo "Building Go binary..."
# бинарник будем класть в папку Index
go build -o Index/tg-index

echo "Restarting systemd service..."
sudo systemctl restart tg-index.service

echo "Deployment complete!"
sudo systemctl status tg-index.service --no-pager
