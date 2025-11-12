#!/bin/bash
set -e

DATABASES=("auth" "profile" "users" "chats")

# Проверяем количество баз
expected_count=${#DATABASES[@]}
actual_count=$(psql -U $POSTGRES_USER -d template1 -t -c \
  "SELECT COUNT(*) FROM pg_database WHERE datname IN ('auth','profile','users','chats');" | tr -d ' ')

if [ "$actual_count" -ne "$expected_count" ]; then
  echo "ERROR: Expected $expected_count databases, found $actual_count"
  exit 1
fi

# Проверяем подключение к каждой базе
for db in "${DATABASES[@]}"; do
  echo "Checking database: $db"
  if ! psql -U $POSTGRES_USER -d "$db" -c "SELECT 1;" > /dev/null 2>&1; then
    echo "ERROR: Cannot connect to database: $db"
    exit 1
  fi
done

echo "SUCCESS: All $expected_count databases are healthy"
exit 0