#!/bin/bash
docker build -t "bomberman-dom" .
docker run -p "8080:8080" bomberman-dom
