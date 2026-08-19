#!/usr/bin/env bash
set -eu -o pipefail

OPENFGA_APIURL="${OPENFGA_APIURL:-http://localhost:8082}"
OPENFGA_STOREID="${OPENFGA_STOREID:-01KZW369WB6CPEY6RBR42NWF1B}"
OPENFGA_AUTHORIZATIONMODELID="${OPENFGA_AUTHORIZATIONMODELID:-01KZW3ZKH129S283DVN5JD501G}"

TMPDIR=$(mktemp -d)
cleanup() { rm -rf "$TMPDIR"; }
trap cleanup EXIT

echo "OpenFGA validation script"
echo "Using OPENFGA_APIURL=$OPENFGA_APIURL"
echo "Using OPENFGA_STOREID=$OPENFGA_STOREID"
echo "Using OPENFGA_AUTHORIZATIONMODELID=$OPENFGA_AUTHORIZATIONMODELID"

echo "\n1) Fetching stores..."
if ! curl -sf "$OPENFGA_APIURL/stores" -o "$TMPDIR/stores.json"; then
  echo "ERROR: cannot reach $OPENFGA_APIURL/stores" >&2
  exit 2
fi

if command -v jq >/dev/null 2>&1; then
  echo "Stores:"; jq . "$TMPDIR/stores.json"
  if ! jq -r '.stores[].id' "$TMPDIR/stores.json" | grep -q "$OPENFGA_STOREID"; then
    echo "ERROR: store $OPENFGA_STOREID not found" >&2
    exit 3
  fi
else
  if ! grep -q "$OPENFGA_STOREID" "$TMPDIR/stores.json"; then
    echo "ERROR: store $OPENFGA_STOREID not found (install jq for nicer output)" >&2
    exit 3
  fi
fi

echo "\n2) Fetching authorization models for store..."
if ! curl -sf "$OPENFGA_APIURL/stores/$OPENFGA_STOREID/authorization-models" -o "$TMPDIR/models.txt"; then
  echo "WARNING: failed to fetch authorization models for store $OPENFGA_STOREID" >&2
else
  if grep -q "$OPENFGA_AUTHORIZATIONMODELID" "$TMPDIR/models.txt"; then
    echo "Found authorization model $OPENFGA_AUTHORIZATIONMODELID"
  else
    echo "WARNING: authorization model $OPENFGA_AUTHORIZATIONMODELID not found in store" >&2
  fi
fi

echo "\n3) Running sample permission check (may return allowed=false if tuple is not present)"
CHECK_RESPONSE=$(curl -s -X POST "$OPENFGA_APIURL/stores/$OPENFGA_STOREID/check" \
  -H "Content-Type: application/json" \
  -d '{
    "authorization_model_id":"'"$OPENFGA_AUTHORIZATIONMODELID"'",
    "tuple_key":{"object":"vehicle:vehicle-1","relation":"viewer","user":"user:alice"}
  }')

if command -v jq >/dev/null 2>&1; then
  echo "$CHECK_RESPONSE" | jq .
  ALLOWED=$(echo "$CHECK_RESPONSE" | jq -r '.allowed // empty' 2>/dev/null || echo "")
else
  echo "$CHECK_RESPONSE"
  ALLOWED=$(echo "$CHECK_RESPONSE" | grep -o '"allowed"[[:space:]]*:[[:space:]]*[^,}]\+' | sed 's/.*://; s/[^a-zA-Z]*//g' || echo "")
fi

echo "Permission check result: ${ALLOWED:-unknown}"
if [ "${ALLOWED:-}" = "true" ]; then
  echo "Permission allowed for sample tuple"
else
  echo "Permission not allowed or unknown for sample tuple"
fi

exit 0
