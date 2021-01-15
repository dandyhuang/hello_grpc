#
# make lib proto
#

PUBLIC_LIB = /data/rtrs/brucexiao/public_lib/third_lib

#CC = $(PUBLIC_LIB)/gcc-4.9.3/bin/g++
CC = gcc
CXX = g++
# proto files
PROTO=$(wildcard *.proto)
PROTOC = $(PUBLIC_LIB)/protobuf-3.6.0/bin/protoc
CXXFLAGS = -O3 -g3 --std=c++11 -pthread -Wall -Werror -Wno-deprecated-declarations

CXXFLAGS += -I$(PUBLIC_LIB)/protobuf-3.6.0/include

all: lib/libproto.a

# library of vipshop recommend engine protobuf
lib/libproto.a: $(PROTO:.proto=.pb.o)
	@mkdir -p $(dir $@)
	$(AR) rcs $@ $^

# condition
$(PROTO:.proto=.pb.o): %.pb.o: %.pb.cc $(PROTO:.proto=.pb.h)
	$(CC) -c $< -o $@ -O3 $(CXXFLAGS)
# generate cpp
cpp: $(PROTO:.proto=.pb.h) $(PROTO:.proto=.pb.cc);
# generate cpp
%.pb.h %.pb.cc: %.proto
	$(PROTOC) $< --proto_path=. --cpp_out=.


# cleanup
clean:
	$(RM) -rf *.o *.h *.cc lib
