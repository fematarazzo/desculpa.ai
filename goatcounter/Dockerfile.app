FROM arp242/goatcounter:latest

ENV GC_LISTEN=:8080 \
  GC_WEB_ROOT=/ \
  GC_DB=/home/goatcounter/goatcounter-data/goatcounter.sqlite \
  GC_ALLOW_SIGNUPS=false

VOLUME /home/goatcounter/goatcounter-data

EXPOSE 8080

CMD ["goatcounter"]

