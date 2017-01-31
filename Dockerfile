FROM dockerhub.cisco.com/iox-docker-dev-local/base-ap-arm-latest
 
RUN ln -s -T -f /lib/libpthread-2.20.so /lib/libpthread.so.0

RUN opkg update && opkg install python
RUN opkg install python-dev
RUN opkg install iox-toolchain
RUN opkg install python-pip
RUN pip install protobuf
RUN pip install grpcio

RUN git clone ssh://dkourkou@cheetah-build:29418/cheetah-objmodel
COPY cheetah-objmodel/grpc/ /opt/grpc/

CMD python /opt/grpc/python/src/tutorial/quickstart.py

