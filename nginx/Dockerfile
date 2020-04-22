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

# Run nginx
CMD ["nginx", "-g", "daemon off;"]
