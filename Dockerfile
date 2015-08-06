
# sti-labs-golang1
FROM openshift/base-centos7

# Put the maintainer name in the image metadata
MAINTAINER Alexander Sokolovsky <amsokol@gmail.com>

# Rename the builder environment variable to inform users about application you provide them
ENV LABSGOLANG_VERSION 1.0

# Set labels used in OpenShift to describe the builder image
LABEL io.k8s.description="Labs builder images for golang web projects" \
      io.k8s.display-name="Labs Golang 1.0.0" \
      io.openshift.expose-services="8080:http" \
      io.openshift.tags="builder,golang"

# Install required packages here:
RUN yum install -y git golang mc && yum clean all -y

# TODO (optional): Copy the builder files into /opt/openshift
# COPY ./<builder_folder>/ /opt/openshift/

# Copy the S2I scripts to /usr/local/sti, since openshift/base-centos7 image sets io.openshift.s2i.scripts-url label that way, or update that label
COPY ./.sti/bin/ /usr/local/sti

# Drop the root user and make the content of /opt/openshift owned by user 1001
RUN chown -R 1001:1001 /opt/app-root

# This default user is created in the openshift/base-centos7 image
USER 1001

# Set the default port for applications built using this image
EXPOSE 8080

# Set the default CMD for the image
CMD ["usage"]
