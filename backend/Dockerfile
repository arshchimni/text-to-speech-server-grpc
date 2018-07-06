FROM alpine
#install speech synth package
RUN apk update && apk add flite
ADD app /app
ENTRYPOINT [ "/app" ]