LABEL org.opencontainers.image.source=https://github.com/mj3052/open-dmr

FROM ubuntu:21.04

COPY database database
COPY open-dmr open-dmr

CMD ["./open-dmr"]