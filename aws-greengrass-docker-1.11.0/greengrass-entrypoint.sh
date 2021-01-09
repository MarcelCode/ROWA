#!/bin/sh

set -e

# Initial check before starting greengrassd
# Check if default IsolationMode is set to GreengrassContainer or any user lambdas are configured in GreengrassContainer mode
if grep -q "GreengrassContainer" /greengrass/ggc/deployment/group/group.json; then
    echo "Default IsolationMode as GreengrassContainer or User Lambdas with GreengrassContainer mode aren't supported to run inside the GGC Docker Container. For troubleshooting, start a fresh deployment by following this guide: https://docs.aws.amazon.com/greengrass/latest/developerguide/run-gg-in-docker-container.html#docker-no-container. Finally, restart the GGC docker container after bind-mounting an empty deployment folder."
    exit 1;
fi

/greengrass/ggc/core/greengrassd start

daemon_pid=`cat /var/run/greengrassd.pid`
# block docker exit until daemon process dies.
while [ -d /proc/$daemon_pid ]
do
 # Sleep for 1s before checking that greengrass daemon is still alive
 daemon_cmdline=`cat /proc/$daemon_pid/cmdline`
 if [[ $daemon_cmdline != ^/greengrass/ggc/packages/1.11.0/bin/daemon.* ]]; then 
  sleep 1;
 else
  break;
 fi;
done 
