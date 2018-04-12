FROM golang:1.10.1-alpine3.7

ENV WORKDIR $GOPATH/src/github.com/sonaak/lydia

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN go get -u github.com/golang/dep/cmd/dep

# Copy the code to the workdir
COPY . $WORKDIR

# copy the entrypoint script
COPY build/entrypoint.sh $WORKDIR/entrypoint.sh

# Expose the volume /go/out
VOLUME /go/out

# setup the work directory
WORKDIR $WORKDIR

ARG VERSION
ENV VERSION ${VERSION}
ARG SHORTHASH
ENV SHORTHASH ${SHORTHASH}

LABEL maintainer="evilwire <knfu@mediamath.com>"
LABEL version=${VERSION}

ENTRYPOINT ["/bin/sh", "entrypoint.sh"]