FROM dockerhub.cisco.com/iox-docker-dev/ap3k/base-rootfs

COPY devel/dna-agent /usr/bin
RUN chmod 777 /usr/bin/dna-agent

COPY devel/startapp /etc/init.d/
RUN chmod 777 /etc/init.d/startapp

RUN cd /etc/rc5.d/; ln -s ../init.d/startapp S50startapp
CMD /etc/rc5.d/S50startapp start
