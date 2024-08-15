#!/bin/bash

# Clean up any existing X server or DBus daemon
if [ -e /tmp/.X20-lock ]; then
  echo "Removing existing X server lock file"
  rm /tmp/.X20-lock
fi

if [ -e /run/dbus/pid ]; then
  echo "Removing existing DBus daemon PID file"
  rm /run/dbus/pid
fi

echo "Starting dbus-daemon"
dbus-daemon --config-file=/usr/share/dbus-1/system.conf --print-address &
echo "Starting Xvfb"
Xvfb :20 -screen 0 1920x1080x24 &
sleep 1
echo "Xvfb started"

echo "Starting openbox"
openbox &

x11vnc -display :20 -N -forever -bg -o "/tmp/x11vnc.log"
sleep 1
echo "x11vnc started"

# Check if X server is running
if xset q &> /dev/null; then
  echo "X server is running"
else
  echo "X server is not running"
fi

# Execute the commands passed as arguments from
# the docker-compose.dev.yml -> services:goCrawlDeeplDev:command
exec "$@"