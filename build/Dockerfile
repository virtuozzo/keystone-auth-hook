FROM alpine

ADD keystone-auth-hook /
EXPOSE 2000

ADD webhook.pem /etc/ssl/certs/
ADD webhook.key /etc/ssl/certs/

ENTRYPOINT ["/keystone-auth-hook"]
CMD ["--certfile", "/etc/ssl/certs/webhook.pem", "--keyfile", "/etc/ssl/certs/webhook.key"]
