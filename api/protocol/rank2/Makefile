PROTO_DIR=$(shell pwd)

THIRD_LIB =  ../../third_party
PROTOBUF_DIR = ${THIRD_LIB}/protobuf-3.6.1

INCLUDE=-I. -I${PROTOBUF_DIR}/include
export LD_LIBRARY_PATH=${THIRD_LIB}/protobuf-3.6.1/lib/:$LD_LIBRARY_PATH
CXX= g++
CPPFLAGS= -g --std=c++11 -Wall -O2 -fPIC
CFLAGS := $(CPPFLAGS)
AR= ar cru

PROTO_SRC = $(wildcard $(PROTO_DIR)/*.proto)
PROTO_OUT = $(patsubst %.proto, %.pb.cc, $(PROTO_SRC))

all: $(PROTO_OUT)

$(PROTO_OUT): $(PROTO_SRC)
	${PROTOBUF_DIR}/bin/protoc $(PROTO_SRC) --proto_path $(PROTO_DIR) --cpp_out=./

clean:
	rm -f *.o *.h *.cc
