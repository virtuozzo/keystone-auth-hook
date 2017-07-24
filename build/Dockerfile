FROM alpine

ADD keystone-auth-hook /
EXPOSE 2000

ADD webhook.pem /etc/ssl/certs/
ADD webhook.key /etc/ssl/certs/

ENV OS_AUTH_URL http://127.0.0.1:5001/v3
ENV OS_USERNAME admin
ENV OS_PASSWORD password
ENV OS_DOMAIN_ID default

ENTRYPOINT ["/keystone-auth-hook"]
CMD ["--certfile", "/etc/ssl/certs/webhook.pem", "--keyfile", "/etc/ssl/certs/webhook.key"]
