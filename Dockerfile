# BUILDER
FROM golang:latest AS builder
WORKDIR /go/src/app
COPY . .

# Static build required so that we can safely copy the binary over.
# `-tags timetzdata` embeds zone info from the "time/tzdata" package.
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata

# RUNNING
FROM scratch
COPY --from=builder /go/bin/fire /fire
ENTRYPOINT ["/fire"]