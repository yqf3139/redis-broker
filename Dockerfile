FROM bitnami/minideb:latest
COPY ./redis-broker /redis-broker
COPY ./redis-0.7.1.tgz /redis-chart.tgz
CMD ["/redis-broker", "-logtostderr"]
