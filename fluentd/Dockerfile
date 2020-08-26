# fluentd/Dockerfile

FROM fluent/fluentd:v1.11-1
USER root
RUN ["gem", "install", "fluent-plugin-elasticsearch"]
USER fluent