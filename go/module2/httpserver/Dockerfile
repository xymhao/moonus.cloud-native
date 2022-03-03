FROM alpine
ENV MY_SERVICE_PORT = 80
ADD bin/amd64/httpserver /httpserver
ENTRYPOINT /httpserver -port 80

EXPOSE 80