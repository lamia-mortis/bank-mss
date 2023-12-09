#!/bin/sh

set -e

echo "run db migration"
${WORK_DIR}/migrate -path ${WORK_DIR}/migrations -database "${AUTH_DB_URL}" -verbose up

echo "start the app"
exec "$@"