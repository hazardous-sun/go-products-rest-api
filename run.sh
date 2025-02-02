#!/bin/bash

# Setting environment variables (These are used within the docker-compose.yml, not directly by podman)
export DB_HOST="go_db" # Although this is set, it is not used correctly. See explanation below.
export DB_PORT="5432" # Although this is set, it is not used correctly. See explanation below.
export DB_USER="postgres"
export DB_PASSWORD="1234"
export DB_NAME="postgres"

# Cleanup (Use podman-compose for cleanup)
podman-compose down

# Run service (Use podman-compose to manage everything)
podman-compose up -d