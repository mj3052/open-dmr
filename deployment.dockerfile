FROM ubuntu:21.04

COPY dmr.db dmr.db
COPY open-dmr open-dmr

CMD ["./open-dmr"]