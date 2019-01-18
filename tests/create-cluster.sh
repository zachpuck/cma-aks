#!/bin/bash

CLUSTER_API=${CLUSTER_API:-cluster-manager-api.cnct.io}
CLUSTER_API_PORT=${CLUSTER_API_PORT:-443}
CLUSTER_NAME=${CLUSTER_NAME:-azure-test-$(date +%s)}
K8S_VERSION=${K8S_VERSION:-1.11.5}

# azure specific inputs
LOCATION=${LOCATION:westus}
AZURE_CLIENT_ID=${AZURE_CLIENT_ID}
AZURE_CLIENT_SECRET=${AZURE_CLIENT_SECRET}
AZURE_TENANT_ID=${AZURE_TENANT_ID}
AZURE_SUBSCRIPTION_ID=${AZURE_SUBSCRIPTION_ID}
NODE_SIZE=${NODE_SIZE:Standard_A1}

[[ -n $DEBUG ]] && set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

DATA=$(
  cat <<JSON
{
  "name": "${CLUSTER_NAME}",
  "provider": {
    "name": "azure",
    "k8s_version": "${K8S_VERSION}",
    "azure": {
      "location": "${LOCATION}",
      "credentials": {
        "app_id": "${AZURE_CLIENT_ID}",
        "tenant": "${AZURE_TENANT_ID}",
        "password": "${AZURE_CLIENT_SECRET}",
        "subscription_id": "${AZURE_SUBSCRIPTION_ID}"
      },
      "clusterAccount": {
        "client_id": "${AZURE_CLIENT_ID}",
        "client_secret": "${AZURE_CLIENT_SECRET}"
      },
      "instance_groups": [
        {
          "name": "agentpool1",
          "type": "${NODE_SIZE}",
          "min_quantity": 1
        }
      ]
    },
    "high_availability": true,
    "network_fabric": ""
  },
  "callback": {
    "url": "example.com",
    "request_id": ""
  }

}
JSON
)

main() {
  curl -X POST \
    "https://${CLUSTER_API}:${CLUSTER_API_PORT}/api/v1/cluster" \
    -H 'Cache-Control: no-cache' \
    -H 'Content-Type: application/json' \
    -d "${DATA}" \
    -iks
}

main
