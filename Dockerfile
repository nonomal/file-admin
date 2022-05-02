FROM --platform=$BUILDPLATFORM alpine:latest as builder

WORKDIR /app
COPY dist .
COPY .github/actions/docker/file.sh .

ARG TARGETPLATFORM

RUN ash ./file.sh

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/file-admin .

RUN set -ex && \
    chmod 755 /app/file-admin && \
    apk --no-cache add ca-certificates

ENV PATH /app:$PATH

CMD ["file-admin"]
