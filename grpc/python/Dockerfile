FROM devhub-docker.cisco.com/iox-docker/ap3k/base-rootfs
 
RUN ln -s -T -f /lib/libpthread-2.20.so /lib/libpthread.so.0

RUN opkg update && opkg install python
RUN opkg install python-unittest python-json python-misc python-pkgutil
RUN opkg install python-threading python-numbers

COPY share/site-packages /usr/lib/python2.7/site-packages/
COPY devel/cheetah-objmodel/grpc/ /opt/grpc/

CMD python /opt/grpc/python/src/tutorial/quickstart.py
