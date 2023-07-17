FROM alpine:3.16
RUN apk --update add bash
RUN apk add curl
RUN mkdir -p /data/code
COPY ./growth_* /data/code/
WORKDIR /data/code/

ENV USER_GROWTH_CONFIG='{"Db":{"Engine":"mysql","Username":"root","Password":"root","Host":"localhost","Port":3306,"Database":"user_growth","Charset":"utf8","ShowSql":true,"MaxIdleConns":2,"MaxOpenConns":10,"ConnMaxLifetime":30},"Cache":{}}'

RUN echo $'#/bin/bash\n\
cd /data/code/\n\
./growth_server > logs_server.txt 2>&1 &\n\
./growth_api > logs_api.txt 2>&1' > /startup.sh && \
chmod +x /startup.sh

EXPOSE 80
EXPOSE 8080
EXPOSE 8081

CMD /startup.sh
