import grpc
from concurrent.futures import ThreadPoolExecutor
import greet_pb2
import greet_pb2_grpc

class Greeter(greet_pb2_grpc.GreeterServicer):
    def SayHello(self, request, context):
        return greet_pb2.HelloReply(message=f"Hello, {request.name}!")

def serve():
    server = grpc.server(ThreadPoolExecutor())
    # server = grpc.server(grpc.InsecureServer())
    greet_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()
