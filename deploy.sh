#!/bin/sh
if [ -z "$1" ]; then
  echo "You must enter in a version number argument."
  exit 1
else
  docker build -t kongaaron/flask-planets:$1 -t kongaaron/flask-planets:latest .
  docker push kongaaron/flask-planets&& docker push kongaaron/flask-planets:$1
fi
