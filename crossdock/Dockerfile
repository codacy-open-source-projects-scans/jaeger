# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

FROM scratch
ARG TARGETARCH

COPY crossdock-linux-$TARGETARCH /go/bin/crossdock-linux

EXPOSE 8080
ENTRYPOINT ["/go/bin/crossdock-linux"]
