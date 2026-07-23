#!/bin/sh

set -e


echo "Checking API health..."

STATUS=$(curl \
    -s \
    -o /tmp/health.json \
    -w "%{http_code}" \
    http://localhost:8080/api/v1/health)


if [ "$STATUS" != "200" ]; then

    echo "API health failed"

    cat /tmp/health.json

    exit 1

fi


echo "API healthy"


echo "Checking response..."

cat /tmp/health.json


echo ""

echo "Docker smoke test passed"