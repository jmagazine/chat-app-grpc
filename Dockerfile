FROM ubuntu

WORKDIR /usr/chat-app-grpc

RUN apt-get update && \
    apt-get install -y golang ca-certificates && \
    update-ca-certificates

COPY * .

EXPOSE 50051
EXPOSE 8082

LABEL maintainer="Joshua Magazine"
LABEL version = "0.10"

CMD ["go", "run", "./src/server/server.go"]

