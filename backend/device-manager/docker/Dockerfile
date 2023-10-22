FROM scratch
LABEL org.opencontainers.image.source https://github.com/0x6flab/nairobi-traffic-pcb
ARG SVC
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY ntp-$SVC /exe
ENTRYPOINT ["/exe"]
