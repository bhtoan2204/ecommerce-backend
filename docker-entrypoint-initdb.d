#!/bin/bash
set -e

# Tạo thêm database
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE ecom_user;
    CREATE DATABASE ecom_payment;
EOSQL
