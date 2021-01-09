FROM amazonlinux:2

# Set ENV_VAR for Greengrass RC to be untarred inside Docker Image
ARG GREENGRASS_RELEASE_URL=https://d1onfpft10uf5o.cloudfront.net/greengrass-core/downloads/1.11.0/greengrass-linux-x86-64-1.11.0.tar.gz

# Install Greengrass Core Dependencies
RUN yum update -y && \
    yum install -y shadow-utils tar.x86_64 gzip xz wget iproute java-1.8.0 make && \
    yum install -y openssl-devel python27 python37 && \
    amazon-linux-extras enable python3.8 && \
    yum install -y python3.8 && \
    ln -s /usr/bin/java /usr/local/bin/java8 && \
    wget $GREENGRASS_RELEASE_URL && \
    wget https://nodejs.org/dist/v6.10.2/node-v6.10.2-linux-x64.tar.xz && \
    tar xf node-v6.10.2-linux-x64.tar.xz && \
    cp node-v6.10.2-linux-x64/bin/node /usr/bin/nodejs6.10 && \
    wget https://nodejs.org/dist/v8.10.0/node-v8.10.0-linux-x64.tar.xz && \
    tar xf node-v8.10.0-linux-x64.tar.xz && \
    cp node-v8.10.0-linux-x64/bin/node /usr/bin/nodejs8.10 && \
    wget https://nodejs.org/dist/v12.13.0/node-v12.13.0-linux-x64.tar.xz && \
    tar xf node-v12.13.0-linux-x64.tar.xz && \
    cp node-v12.13.0-linux-x64/bin/node /usr/bin/nodejs12.x && \
    ln -s /usr/bin/nodejs12.x /usr/bin/node && \
    rm -rf node-v6.10.2-linux-x64.tar.xz node-v6.10.2-linux-x64 && \
    rm -rf node-v8.10.0-linux-x64.tar.xz node-v8.10.0-linux-x64 && \
    rm -rf node-v12.13.0-linux-x64 node-v12.13.0-linux-x64.tar.xz && \
    yum remove -y wget && \
    rm -rf /var/cache/yum

# Copy Greengrass Licenses AWS IoT Greengrass Docker Image
COPY greengrass-license-v1.pdf /

# Copy start-up script
COPY "greengrass-entrypoint.sh" /

# Setup Greengrass inside Docker Image
RUN export GREENGRASS_RELEASE=$(basename $GREENGRASS_RELEASE_URL) && \
    tar xzf $GREENGRASS_RELEASE -C / && \
    rm $GREENGRASS_RELEASE && \
    useradd -r ggc_user && \
    groupadd -r ggc_group

# Expose 8883 to pub/sub MQTT messages
EXPOSE 8883
