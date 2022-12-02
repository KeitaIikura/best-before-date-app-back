#!/bin/sh
set -e
DB_PASSWORD=${DB_PASSWORD}
DB_USER=${DB_USER}
DB_WRITER_ADDRESS=${DB_HOST}

CONNECT_STR="mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_WRITER_ADDRESS})/${SCHEMA}?parseTime=true"
echo "${CONNECT_STR}"
migrate -database ${CONNECT_STR} -path ${DIR} -verbose goto ${VER}
