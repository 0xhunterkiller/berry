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
    sleep 2
    echo -e "\n========================================\n"
}


echo "Begin Phase 2"

VALID_TOKEN=$(curl -s -X POST $BASE_URL/auth/user/login \
    -H 'Content-Type: application/json' \
    -d '{"username":"berryroot","password":"E1@6~0kB>uT>p%7}x(MG9Pu1p[i2WLK:"}' | jq -r '.jwt')

if [ "$VALID_TOKEN" == "null" ] || [ -z "$VALID_TOKEN" ]; then
    echo "Failed to retrieve a valid token for berryroot"
    return
fi

echo "Retrieved JWT for berryroot: $VALID_TOKEN"

# Test 1: Create Action (Valid Token, Valid Data)
run_test "Create action (valid token and data)" \
    "curl -X POST $BASE_URL/action/create \
    -H 'Authorization: Bearer $VALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"name\":\"Test Action\", \"description\":\"This is a test action\"}'"

# Test 2: Create Action (Invalid Token)
run_test "Create action (invalid token)" \
    "curl -X POST $BASE_URL/action/create \
    -H 'Authorization: Bearer $INVALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"name\":\"Test Action\", \"description\":\"This is a test action\"}'"

# Test 3: Create Action (Valid Token, Missing Required Field)
run_test "Create action (missing required field)" \
    "curl -X POST $BASE_URL/action/create \
    -H 'Authorization: Bearer $VALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"description\":\"This is a test action\"}'"


# Get VALID ID for delete tests
VALID_ID=$(curl -s -X POST $BASE_URL/action/create \
    -H "Authorization: Bearer $VALID_TOKEN" \
    -H 'Content-Type: application/json' \
    -d '{"name":"Test Action 2", "description":"This is also a test action"}' | jq -r '.id')

    echo $VALID_ID

if [ "$VALID_ID" == "null" ] || [ -z "$VALID_ID" ]; then
    echo "Failed to retrieve a valid token for berryroot"
    return
fi

echo "Retrieved JWT for berryroot: $VALID_TOKEN"
# Test 4: Delete Action (Invalid Token)
run_test "Delete action (invalid token)" \
    "curl -X GET $BASE_URL/action/delete?id=$VALID_ID \
    -H 'Authorization: Bearer $INVALID_TOKEN'"

# Test 5: Delete Action (Valid Token, Missing ID)
run_test "Delete action (missing ID)" \
    "curl -X GET $BASE_URL/action/delete \
    -H 'Authorization: Bearer $VALID_TOKEN'"

# Test 6: Delete Action (Valid Token, Nonexistent ID)
run_test "Delete action (nonexistent ID)" \
    "curl -X GET $BASE_URL/action/delete?id=9999 \
    -H 'Authorization: Bearer $VALID_TOKEN'"

# Test 7: Delete Action (Valid Token, Valid ID)
run_test "Delete action (valid token and ID)" \
    "curl -X GET $BASE_URL/action/delete?id=$VALID_ID \
    -H 'Authorization: Bearer $VALID_TOKEN'"


### Roles

# Test 1: Create Role (Valid Token, Valid Data)
run_test "Create role (valid token and data)" \
    "curl -X POST $BASE_URL/role/create \
    -H 'Authorization: Bearer $VALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"name\":\"Test Role\", \"description\":\"This is a test role\"}'"

# Test 2: Create Role (Invalid Token)
run_test "Create role (invalid token)" \
    "curl -X POST $BASE_URL/role/create \
    -H 'Authorization: Bearer $INVALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"name\":\"Test Role\", \"description\":\"This is a test role\"}'"

# Test 3: Create Role (Valid Token, Missing Required Field)
run_test "Create role (missing required field)" \
    "curl -X POST $BASE_URL/role/create \
    -H 'Authorization: Bearer $VALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"description\":\"This is a test role\"}'"

# Get VALID ID for delete tests
VALID_ID=$(curl -s -X POST $BASE_URL/role/create \
    -H "Authorization: Bearer ${VALID_TOKEN}" \
    -H 'Content-Type: application/json' \
    -d '{"name":"Test Role 2", "description":"This is also a test role"}' | jq -r '.id')

if [ "$VALID_ID" == "null" ] || [ -z "$VALID_ID" ]; then
    echo "Failed to retrieve a valid id"
    return
fi

# Test 4: Delete Role (Invalid Token)
run_test "Delete role (invalid token)" \
    "curl -X GET $BASE_URL/role/delete?id=$VALID_ID \
    -H 'Authorization: Bearer $INVALID_TOKEN'"

# Test 5: Delete Role (Valid Token, Missing ID)
run_test "Delete role (missing ID)" \
    "curl -X GET $BASE_URL/role/delete \
    -H 'Authorization: Bearer $VALID_TOKEN'"

# Test 6: Delete Role (Valid Token, Nonexistent ID)
run_test "Delete role (nonexistent ID)" \
    "curl -X GET $BASE_URL/role/delete?id=9999 \
    -H 'Authorization: Bearer $VALID_TOKEN'"

# Test 7: Delete Role (Valid Token, Valid ID)
run_test "Delete role (valid token and ID)" \
    "curl -X GET $BASE_URL/role/delete?id=$VALID_ID \
    -H 'Authorization: Bearer $VALID_TOKEN'"


### Permissions

# Test 1: Create Permission (Valid Token, Valid Data)
run_test "Create permission (valid token and data)" \
    "curl -X POST $BASE_URL/permission/create \
    -H 'Authorization: Bearer $VALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"name\":\"Test Permission\", \"description\":\"This is a test permission\"}'"

# Test 2: Create Permission (Invalid Token)
run_test "Create permission (invalid token)" \
    "curl -X POST $BASE_URL/permission/create \
    -H 'Authorization: Bearer $INVALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"name\":\"Test Permission\", \"description\":\"This is a test permission\"}'"

# Test 3: Create Permission (Valid Token, Missing Required Field)
run_test "Create permission (missing required field)" \
    "curl -X POST $BASE_URL/permission/create \
    -H 'Authorization: Bearer $VALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"description\":\"This is a test permission\"}'"

# Get VALID ID for delete tests
VALID_ID=$(curl -s -X POST $BASE_URL/permission/create \
    -H "Authorization: Bearer $VALID_TOKEN" \
    -H 'Content-Type: application/json' \
    -d '{"name":"Test Permission 2", "description":"This is also a test permission"}' | jq -r '.id')

if [ "$VALID_ID" == "null" ] || [ -z "$VALID_ID" ]; then
    echo "Failed to retrieve a valid token for berryroot"
    return
fi
# Test 4: Delete Permission (Invalid Token)
run_test "Delete permission (invalid token)" \
    "curl -X GET $BASE_URL/permission/delete?id=$VALID_ID \
    -H 'Authorization: Bearer $INVALID_TOKEN'"

# Test 5: Delete Permission (Valid Token, Missing ID)
run_test "Delete permission (missing ID)" \
    "curl -X GET $BASE_URL/permission/delete \
    -H 'Authorization: Bearer $VALID_TOKEN'"

# Test 6: Delete Permission (Valid Token, Nonexistent ID)
run_test "Delete permission (nonexistent ID)" \
    "curl -X GET $BASE_URL/permission/delete?id=9999 \
    -H 'Authorization: Bearer $VALID_TOKEN'"

# Test 7: Delete Permission (Valid Token, Valid ID)
run_test "Delete permission (valid token and ID)" \
    "curl -X GET $BASE_URL/permission/delete?id=$VALID_ID \
    -H 'Authorization: Bearer $VALID_TOKEN'"

### Resources

# Test 1: Create Resource (Valid Token, Valid Data)
run_test "Create resource (valid token and data)" \
    "curl -X POST $BASE_URL/resource/create \
    -H 'Authorization: Bearer $VALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"name\":\"Test Resource\", \"description\":\"This is a test resource\"}'"

# Test 2: Create Resource (Invalid Token)
run_test "Create resource (invalid token)" \
    "curl -X POST $BASE_URL/resource/create \
    -H 'Authorization: Bearer $INVALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"name\":\"Test Resource\", \"description\":\"This is a test resource\"}'"

# Test 3: Create Resource (Valid Token, Missing Required Field)
run_test "Create resource (missing required field)" \
    "curl -X POST $BASE_URL/resource/create \
    -H 'Authorization: Bearer $VALID_TOKEN' \
    -H 'Content-Type: application/json' \
    -d '{\"description\":\"This is a test resource\"}'"

# Get VALID ID for delete tests
VALID_ID=$(curl -s -X POST $BASE_URL/resource/create \
    -H "Authorization: Bearer $VALID_TOKEN" \
    -H 'Content-Type: application/json' \
    -d '{"name":"Test Resource 2", "description":"This is also a test resource"}' | jq -r '.id')

if [ "$VALID_ID" == "null" ] || [ -z "$VALID_ID" ]; then
    echo "Failed to retrieve a valid token for berryroot"
    return
fi
# Test 4: Delete Resource (Invalid Token)
run_test "Delete resource (invalid token)" \
    "curl -X GET $BASE_URL/resource/delete?id=$VALID_ID \
    -H 'Authorization: Bearer $INVALID_TOKEN'"

# Test 5: Delete Resource (Valid Token, Missing ID)
run_test "Delete resource (missing ID)" \
    "curl -X GET $BASE_URL/resource/delete \
    -H 'Authorization: Bearer $VALID_TOKEN'"

# Test 6: Delete Resource (Valid Token, Nonexistent ID)
run_test "Delete resource (nonexistent ID)" \
    "curl -X GET $BASE_URL/resource/delete?id=9999 \
    -H 'Authorization: Bearer $VALID_TOKEN'"

# Test 7: Delete Resource (Valid Token, Valid ID)
run_test "Delete resource (valid token and ID)" \
    "curl -X GET $BASE_URL/resource/delete?id=$VALID_ID \
    -H 'Authorization: Bearer $VALID_TOKEN'"

echo "All tests successfully completed"