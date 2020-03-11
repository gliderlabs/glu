FROM gliderlabs/alpine:3.7
ENV GLU_CONTAINER true
ENV GO_VERSION=1.6
ENV GLIBC_VERSION=2.31-r0
ENTRYPOINT ["/bin/cat"]
CMD ["Linux"]

RUN apk --update add curl ca-certificates git mercurial bash && \
    curl -Ls https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.31-r0/glibc-${GLIBC_VERSION}.apk > /tmp/glibc-${GLIBC_VERSION}.apk && \
    apk add --allow-untrusted /tmp/glibc-${GLIBC_VERSION}.apk && \
    curl -Ls https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz > /usr/local/go${GO_VERSION}.linux-amd64.tar.gz && \
    cd /usr/local && \
    tar -zxf go${GO_VERSION}.linux-amd64.tar.gz && \
    ln -s /usr/local/go/bin/go /usr/local/bin/go && \
    rm -rf /tmp/glibc-${GLIBC_VERSION}.apk /usr/local/go${GO_VERSION}.linux-amd64.tar.gz /var/cache/apk/*

COPY ./build /build
RUN cp /build/Linux/glu /bin/glu \
  && tar -cf /Darwin -C /build/Darwin glu \
  && tar -cf /Linux -C /build/Linux glu \
  && rm -rf /build
