# Artifact stage
FROM alpine:3.14

COPY goshell /goshell


EXPOSE 8080
EXPOSE 8082

ENTRYPOINT ["/goshell", "start", "--configPath=/etc/goshell/configs/goshell"]
