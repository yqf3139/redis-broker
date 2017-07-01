FROM bitnami/minideb:latest
COPY ./redis-broker /redis-broker
ADD https://kubernetes-charts.storage.googleapis.com/redis-0.7.1.tgz /redis-chart.tgz
CMD ["/redis-broker", "-logtostderr"]
