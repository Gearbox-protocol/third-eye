FROM gcr.io/distroless/base-debian11

ENTRYPOINT ["/3eye"]

COPY 3eye /
COPY migrations /
COPY migrate /
