import asyncio

from grpclib.client import Channel

from ozonmp.com_message_api.v1.com_message_api_grpc import ComMessageApiServiceStub
from ozonmp.com_message_api.v1.com_message_api_pb2 import DescribeMessageV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = ComMessageApiServiceStub(channel)

        req = DescribeMessageV1Request(message_id=1)
        reply = await client.DescribeMessageV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
