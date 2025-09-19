#!/bin/bash

# Папка проекта (текущая)
PROJECT_DIR="$(dirname "$0")"
cd "$PROJECT_DIR" || exit

echo "Pulling latest changes from GitHub..."
git pull

echo "Building Go binary..."
go build -o server

echo "Restarting systemd service..."
sudo systemctl restart tg-index.service

echo "Deployment complete!"
sudo systemctl status tg-index.service --no-pager
