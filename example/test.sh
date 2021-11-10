#!/bin/sh
GRPS_HOST="localhost:8082"
#GRPS_METHOD="ozonmp.com_message_api.v1.ComMessageApiService.CreateMessageV1"
GRPS_METHOD="ozonmp.com_message_api.v1.ComMessageApiService.DescribeMessageV1"
#GRPS_METHOD="ozonmp.com_message_api.v1.ComMessageApiService.ListMessageV1"
#GRPS_METHOD="ozonmp.com_message_api.v1.ComMessageApiService.RemoveMessageV1"

payload=$(
  cat <<EOF
{
  "message_id": 1
}
EOF
)

grpcurl -plaintext -d "${payload}" ${GRPS_HOST} ${GRPS_METHOD}