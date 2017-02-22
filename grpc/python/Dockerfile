FROM dockerhub.cisco.com/iox-docker-dev-local/base-ap-arm-latest
 
RUN ln -s -T -f /lib/libpthread-2.20.so /lib/libpthread.so.0

RUN opkg update && opkg install python

COPY share/site-packages /usr/lib/python2.7/site-packages/
COPY devel/cheetah-objmodel/grpc/ /opt/grpc/

CMD python /opt/grpc/python/src/tutorial/quickstart.py

