#!/bin/bash

SERVICE_NAME="serveos-datasync"
ENVIRONMENT="prod"

case "$1" in
    start)
        echo "Starting $SERVICE_NAME service..."
        sudo systemctl start $SERVICE_NAME.service
        ;;
    stop)
        echo "Stopping $SERVICE_NAME service..."
        sudo systemctl stop $SERVICE_NAME.service
        ;;
    restart)
        echo "Restarting $SERVICE_NAME service..."
        sudo systemctl restart $SERVICE_NAME.service
        ;;
    status)
        echo "Getting status of $SERVICE_NAME service..."
        sudo systemctl status $SERVICE_NAME.service
        ;;
    debug)
        echo "Starting $SERVICE_NAME service in debug mode..."
        sudo SYSTEMD_LOG_LEVEL=debug systemctl start $SERVICE_NAME.service
        ;;
    *)
        echo "Usage: $0 {start|stop|restart|status|debug}"
        exit 1
        ;;
esac
