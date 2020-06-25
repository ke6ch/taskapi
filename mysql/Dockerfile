# Using official images
FROM mysql:8.0.19
LABEL maintainer="ke6ch <19749261+ke6ch@users.noreply.github.com>"

VOLUME /var/lib/mysql

# Copy files
COPY ./conf.d/my.cnf /etc/mysql/conf.d/my.cnf
COPY ./initdb.d /docker-entrypoint-initdb.d
ENTRYPOINT ["docker-entrypoint.sh"]

EXPOSE 3306

# Run mysql
CMD ["mysqld"]
