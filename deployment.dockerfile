FROM ubuntu:21.04

COPY database database
COPY open-dmr open-dmr

CMD ["./open-dmr"]