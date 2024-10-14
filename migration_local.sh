#!/bin/bash
source local.env

MIGRATION_DIR=$(echo "$MIGRATION_DIR" | tr -d '[:space:]')

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v