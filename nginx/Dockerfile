# Using official images
FROM centos:7
LABEL maintainer="k <19749261+keb0@users.noreply.github.com>"

# Copy nginx repository
COPY ./nginx.repo /etc/yum.repos.d/nginx.repo

# Update yum
RUN yum -y update && \
    yum clean all

# Install nginx
RUN yum install -y nginx && \
    yum clean all

# Copy files
COPY ./nginx/index.html /usr/share/nginx/html/nginx/index.html
COPY ./conf.d/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

# Run nginx
CMD ["nginx", "-g", "daemon off;"]
