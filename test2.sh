#!/bin/bash

# Base URL
BASE_URL="http://localhost:3000"

# Token placeholders (replace with actual tokens)
VALID_TOKEN=""
INVALID_TOKEN="invalid_token"

# Helper function for test output
run_test() {
    echo "Test: $1"
    echo "Command: $2"
    eval "$2"
    sleep 10
    echo -e "\n========================================\n"
}


echo "Begin Phase 2"

VALID_TOKEN=$(curl -X POST $BASE_URL/auth/user/login \
    -H 'Content-Type: application/json' \
    -d '{"username":"berryroot","password":"E1@6~0kB>uT>p%7}x(MG9Pu1p[i2WLK:"}' | jq -r '.jwt')

if [ "$VALID_TOKEN" == "null" ] || [ -z "$VALID_TOKEN" ]; then
    echo "Failed to retrieve a valid token for berryroot"
    exit 1
fi

echo "Retrieved JWT for berryroot: $VALID_TOKEN"

run_test "Action Creation (No Description)" \
    "curl -X POST $BASE_URL/action/create -H 'Content-Type: application/json' -d '{\"name\":\"select\",\"description\":\"\"}' -H 'Authorization: Bearer $VALID_TOKEN'"


run_test "Action Creation (Both Present)" \
    "curl -X POST $BASE_URL/action/create -H 'Content-Type: application/json' -d '{\"name\":\"select\",\"description\":\"some description\"}' -H 'Authorization: Bearer $VALID_TOKEN'"

run_test "Action Creation Negative (No Name)" \
    "curl -X POST $BASE_URL/action/create -H 'Content-Type: application/json' -d '{\"description\":\"some description\"}' -H 'Authorization: Bearer $VALID_TOKEN'"

run_test "Action Creation Negative (No Token)" \
    "curl -X POST $BASE_URL/action/create -H 'Content-Type: application/json' -d '{\"name\":\"select\",\"description\":\"some description\"}'"

