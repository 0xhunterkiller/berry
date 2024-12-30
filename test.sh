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

# 1. Register User (Positive Tests)
run_test "Register user (valid data)" \
    "curl -X POST $BASE_URL/auth/user/register -H 'Content-Type: application/json' -d '{\"username\":\"username1\",\"email\":\"username1@example.com\",\"password\":\"Password#3242\"}'"

run_test "Register user (another valid data)" \
    "curl -X POST $BASE_URL/auth/user/register -H 'Content-Type: application/json' -d '{\"username\":\"username2\",\"email\":\"username2@example.com\",\"password\":\"Password*2321\"}'"

# 2. Register User (Negative Tests)
run_test "Register user (missing fields)" \
    "curl -X POST $BASE_URL/auth/user/register -H 'Content-Type: application/json' -d '{\"username\":\"user3\"}'"

run_test "Register user (invalid email)" \
    "curl -X POST $BASE_URL/auth/user/register -H 'Content-Type: application/json' -d '{\"username\":\"user4\",\"email\":\"invalidemail\",\"password\":\"password789\"}'"

run_test "Register user (weak password)" \
    "curl -X POST $BASE_URL/auth/user/register -H 'Content-Type: application/json' -d '{\"username\":\"user5\",\"email\":\"user5@example.com\",\"password\":\"123\"}'"

# 3. Login User (Positive Tests)
run_test "Login user (valid credentials)" \
    "curl -X POST $BASE_URL/auth/user/login -H 'Content-Type: application/json' -d '{\"username\":\"username1\",\"password\":\"Password#3242\"}'"

run_test "Login user (valid credentials for username2)" \
    "curl -X POST $BASE_URL/auth/user/login -H 'Content-Type: application/json' -d '{\"username\":\"username2\",\"password\":\"Password*2321\"}'"

VALID_TOKEN=$(curl -X POST $BASE_URL/auth/user/login \
    -H 'Content-Type: application/json' \
    -d '{"username":"username2","password":"Password*2321"}' | jq -r '.jwt')

if [ "$VALID_TOKEN" == "null" ] || [ -z "$VALID_TOKEN" ]; then
    echo "Failed to retrieve a valid token for user2"
    exit 1
fi

echo "Retrieved JWT for username2: $VALID_TOKEN"


# 4. Login User (Negative Tests)
run_test "Login user (invalid username)" \
    "curl -X POST $BASE_URL/auth/user/login -H 'Content-Type: application/json' -d '{\"username\":\"invaliduser\",\"password\":\"Password#3242\"}'"

run_test "Login user (invalid password)" \
    "curl -X POST $BASE_URL/auth/user/login -H 'Content-Type: application/json' -d '{\"username\":\"username1\",\"password\":\"wrongpassword\"}'"

# 5. Authentication Check (Positive Test)
run_test "Check authentication (valid token)" \
    "curl -X GET $BASE_URL/user/checkauth -H 'Authorization: Bearer $VALID_TOKEN'"

# 6. Authentication Check (Negative Tests)
run_test "Check authentication (no token)" \
    "curl -X GET $BASE_URL/user/checkauth"

run_test "Check authentication (invalid token)" \
    "curl -X GET $BASE_URL/user/checkauth -H 'Authorization: Bearer $INVALID_TOKEN'"

# 7. Update Email (Positive Test)
run_test "Update email (valid token, valid email)" \
    "curl -X PATCH $BASE_URL/user/email -H 'Authorization: Bearer $VALID_TOKEN' -H 'Content-Type: application/json' -d '{\"email\":\"newemail@example.com\"}'"

# 8. Update Email (Negative Tests)
run_test "Update email (invalid email)" \
    "curl -X PATCH $BASE_URL/user/email -H 'Authorization: Bearer $VALID_TOKEN' -H 'Content-Type: application/json' -d '{\"email\":\"invalidemail\"}'"

run_test "Update email (no token)" \
    "curl -X PATCH $BASE_URL/user/email -H 'Content-Type: application/json' -d '{\"email\":\"newemail@example.com\"}'"

# 9. Update Password (Positive Test)
run_test "Update password (valid token, valid password)" \
    "curl -X PATCH $BASE_URL/user/password -H 'Authorization: Bearer $VALID_TOKEN' -H 'Content-Type: application/json' -d '{\"password\":\"newPassword#3242\"}'"

# 10. Update Password (Negative Tests)
run_test "Update password (weak password)" \
    "curl -X PATCH $BASE_URL/user/password -H 'Authorization: Bearer $VALID_TOKEN' -H 'Content-Type: application/json' -d '{\"password\":\"123\"}'"

run_test "Update password (no token)" \
    "curl -X PATCH $BASE_URL/user/password -H 'Content-Type: application/json' -d '{\"password\":\"newPassword#3242\"}'"

# 11. Deactivate User (Positive Test)
run_test "Deactivate user (valid token)" \
    "curl -X PATCH $BASE_URL/user/deactivate -H 'Authorization: Bearer $VALID_TOKEN'"

# 12. Deactivate User (Negative Test)
run_test "Deactivate user (no token)" \
    "curl -X PATCH $BASE_URL/user/deactivate"

# 13. Activate User (Positive Test)
run_test "Activate user (valid token)" \
    "curl -X PATCH $BASE_URL/user/activate -H 'Authorization: Bearer $VALID_TOKEN'"

# 14. Activate User (Negative Test)
run_test "Activate user (no token)" \
    "curl -X PATCH $BASE_URL/user/activate"

# 15. Delete User (Positive Test)
run_test "Delete user (valid token)" \
    "curl -X DELETE $BASE_URL/user -H 'Authorization: Bearer $VALID_TOKEN'"

# 16. Delete User (Negative Test)
run_test "Delete user (no token)" \
    "curl -X DELETE $BASE_URL/user"

# 17. Security Tests
run_test "Tampered token" \
    "curl -X GET $BASE_URL/user/checkauth -H 'Authorization: Bearer tamperedtoken'"

run_test "Rate-limiting simulation" \
    "for i in {1..5}; do curl -X GET $BASE_URL/user/checkauth -H 'Authorization: Bearer $VALID_TOKEN'; done"

run_test "SQL Injection simulation" \
    "curl -X POST $BASE_URL/auth/user/login -H 'Content-Type: application/json' -d '{\"username\":\"admin' OR '1'='1\",\"password\":\"password\"}'"

run_test "XSS Injection" \
    "curl -X POST $BASE_URL/auth/user/register -H 'Content-Type: application/json' -d '{\"username\":\"<script>alert(1)</script>\",\"email\":\"userxss@example.com\",\"password\":\"Password#3242\"}'"

# Add more tests as needed to reach 100!

echo "All tests completed!"
