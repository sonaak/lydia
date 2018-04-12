FROM alpine:3.7

# the version of the docker file
ARG VERSION
ENV VERSION $VERSION

ARG SHORTHASH
ENV SHORTHASH $SHORTHASH

# Copy the binary
COPY build/${SHORTHASH}/lydia /opt/sonaak/lydia
RUN chmod +x /opt/sonaak/lydia

# Expose the appropriate ports
# 9000 is the actual port to respond to requests
# 9090 is the healthcheck port
EXPOSE 9000 9090

ARG BUILD_TIME
ENV BUILD_TIME $BUILD_TIME

ARG GHASH
ENV GHASH $GHASH

ARG PACKER
ENV PACKER $PACKER

LABEL maintainer="evilwire"
LABEL version=$VERSION

# Set the entry point
ENTRYPOINT ["/opt/sonaak/lydia"]
