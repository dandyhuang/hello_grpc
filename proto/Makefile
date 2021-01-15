PROTO_DIR=$(shell pwd)

PUBLIC_LIB = /data/rtrs/public_lib
THIRD_LIB = ${PUBLIC_LIB}/third_lib
PRIVATE_LIB = ${PUBLIC_LIB}/private_lib
PROTOBUF_DIR = ${THIRD_LIB}/protobuf-3.6.0

LIB_PROTOBUF = -L$(PROTOBUF_DIR)/lib -lprotobuf

INCLUDE=-I. -I${PROTOBUF_DIR}/include

CXX= g++
CPPFLAGS= -g --std=c++11 -Wall -O2 -fPIC
CFLAGS := $(CPPFLAGS)
AR= ar cru

PROTO_SRC = $(wildcard $(PROTO_DIR)/*.proto)
PROTO_OUT = $(patsubst %.proto, %.pb.cc, $(PROTO_SRC))
PROTO_OBJ = $(patsubst %.proto, %.pb.o, $(PROTO_SRC))

LIBA = $(PROTO_DIR)/.lib/libproto.a

all: $(LIBA)

$(PROTO_OUT): $(PROTO_SRC)
	${PROTOBUF_DIR}/bin/protoc $(PROTO_SRC) --proto_path $(PROTO_DIR) --cpp_out=./

%.o: %.cc
	$(CXX) $(CFLAGS) $(INCLUDE) -c $< -o $@ $(LIB_PROTOBUF)

$(LIBA): $(PROTO_OBJ)
	$(shell [ -d ./.lib ] || mkdir .lib )
	rm -f $(LIBA)
	ar q $@ $^

clean:
	rm -f *.o *.h *.cc $(LIBA)
