# Using official images
FROM mysql:8.0.19
LABEL maintainer="k <19749261+keb0@users.noreply.github.com>"

VOLUME /var/lib/mysql

# Copy files
COPY ./conf.d/my.cnf /etc/mysql/conf.d/my.cnf
COPY ./initdb.d /docker-entrypoint-initdb.d
ENTRYPOINT ["docker-entrypoint.sh"]

EXPOSE 3306

# Run mysql
CMD ["mysqld"]
