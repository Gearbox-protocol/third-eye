#!/bin/bash

/migrate -path /migrations -database "$DATABASE_URL" "$@" up
