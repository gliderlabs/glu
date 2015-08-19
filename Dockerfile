FROM ubuntu-debootstrap
ADD ./build /build
RUN cp /build/Linux/glu /bin/glu \
  && tar -cf /Darwin -C /build/Darwin glu \
  && tar -cf /Linux -C /build/Linux glu \
  && rm -rf /build
ENTRYPOINT ["/bin/cat"]
CMD ["Linux"]
