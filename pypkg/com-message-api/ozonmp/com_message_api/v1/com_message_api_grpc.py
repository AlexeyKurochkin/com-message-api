# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: ozonmp/com_message_api/v1/com_message_api.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import validate.validate_pb2
import google.api.annotations_pb2
import google.protobuf.timestamp_pb2
import ozonmp.com_message_api.v1.com_message_api_pb2


class ComMessageApiServiceBase(abc.ABC):

    @abc.abstractmethod
    async def CreateMessageV1(self, stream: 'grpclib.server.Stream[ozonmp.com_message_api.v1.com_message_api_pb2.CreateMessageV1Request, ozonmp.com_message_api.v1.com_message_api_pb2.CreateMessageV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def DescribeMessageV1(self, stream: 'grpclib.server.Stream[ozonmp.com_message_api.v1.com_message_api_pb2.DescribeMessageV1Request, ozonmp.com_message_api.v1.com_message_api_pb2.DescribeMessageV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def ListMessageV1(self, stream: 'grpclib.server.Stream[ozonmp.com_message_api.v1.com_message_api_pb2.ListMessageV1Request, ozonmp.com_message_api.v1.com_message_api_pb2.ListMessageV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def RemoveMessageV1(self, stream: 'grpclib.server.Stream[ozonmp.com_message_api.v1.com_message_api_pb2.RemoveMessageV1Request, ozonmp.com_message_api.v1.com_message_api_pb2.RemoveMessageV1Response]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/ozonmp.com_message_api.v1.ComMessageApiService/CreateMessageV1': grpclib.const.Handler(
                self.CreateMessageV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.com_message_api.v1.com_message_api_pb2.CreateMessageV1Request,
                ozonmp.com_message_api.v1.com_message_api_pb2.CreateMessageV1Response,
            ),
            '/ozonmp.com_message_api.v1.ComMessageApiService/DescribeMessageV1': grpclib.const.Handler(
                self.DescribeMessageV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.com_message_api.v1.com_message_api_pb2.DescribeMessageV1Request,
                ozonmp.com_message_api.v1.com_message_api_pb2.DescribeMessageV1Response,
            ),
            '/ozonmp.com_message_api.v1.ComMessageApiService/ListMessageV1': grpclib.const.Handler(
                self.ListMessageV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.com_message_api.v1.com_message_api_pb2.ListMessageV1Request,
                ozonmp.com_message_api.v1.com_message_api_pb2.ListMessageV1Response,
            ),
            '/ozonmp.com_message_api.v1.ComMessageApiService/RemoveMessageV1': grpclib.const.Handler(
                self.RemoveMessageV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.com_message_api.v1.com_message_api_pb2.RemoveMessageV1Request,
                ozonmp.com_message_api.v1.com_message_api_pb2.RemoveMessageV1Response,
            ),
        }


class ComMessageApiServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.CreateMessageV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.com_message_api.v1.ComMessageApiService/CreateMessageV1',
            ozonmp.com_message_api.v1.com_message_api_pb2.CreateMessageV1Request,
            ozonmp.com_message_api.v1.com_message_api_pb2.CreateMessageV1Response,
        )
        self.DescribeMessageV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.com_message_api.v1.ComMessageApiService/DescribeMessageV1',
            ozonmp.com_message_api.v1.com_message_api_pb2.DescribeMessageV1Request,
            ozonmp.com_message_api.v1.com_message_api_pb2.DescribeMessageV1Response,
        )
        self.ListMessageV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.com_message_api.v1.ComMessageApiService/ListMessageV1',
            ozonmp.com_message_api.v1.com_message_api_pb2.ListMessageV1Request,
            ozonmp.com_message_api.v1.com_message_api_pb2.ListMessageV1Response,
        )
        self.RemoveMessageV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.com_message_api.v1.ComMessageApiService/RemoveMessageV1',
            ozonmp.com_message_api.v1.com_message_api_pb2.RemoveMessageV1Request,
            ozonmp.com_message_api.v1.com_message_api_pb2.RemoveMessageV1Response,
        )
