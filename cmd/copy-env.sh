#!/bin/sh

if [ -f .env ]; then
    echo ".env file exists. Using it."
else
    echo ".env file does not exist. Using .env.example."
    cp .env.example .env
fi
