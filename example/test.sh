#!/bin/sh
GRPS_HOST="localhost:8082"
GRPS_METHOD="ozonmp.com_message_api.v1.ComMessageApiService.CreateMessageV1"
#GRPS_METHOD="ozonmp.com_message_api.v1.ComMessageApiService.DescribeMessageV1"
#GRPS_METHOD="ozonmp.com_message_api.v1.ComMessageApiService.ListMessageV1"
#GRPS_METHOD="ozonmp.com_message_api.v1.ComMessageApiService.RemoveMessageV1"

payload=$(
  cat <<EOF
{
  "from": "me",
  "to": "you",
  "text": "my text",
  "datetime": "2021-11-23T23:03:54.135Z"
}
EOF
)

grpcurl -plaintext -d "${payload}" ${GRPS_HOST} ${GRPS_METHOD}