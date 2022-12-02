FROM amazon/aws-cli:2.1.2
WORKDIR /tmp
COPY ./build/migrate/migrate.sh /tmp/entrypoint.sh
COPY ./deployments/migrations /tmp/migrations
COPY ./build/migrate/migrate.linux-amd64 /usr/local/bin/migrate
RUN chmod a+x /tmp/entrypoint.sh
RUN yum install -y jq
ENTRYPOINT ["sh","-c","/tmp/entrypoint.sh"]