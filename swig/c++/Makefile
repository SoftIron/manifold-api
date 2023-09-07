.DEFAULT_GOAL := libsifi.a

BUILD=.

CXXFLAGS := -Wall -fPIC -std=c++11
CPPFLAGS += $(SWIGFLAGS)

$(BUILD)/%.o: %.cc
	$(CXX) $(CPPFLAGS) $(CXXFLAGS) -c $< -o $@


.PHONY: objs
objs: $(addprefix $(BUILD)/,client.o error.o)

$(BUILD)/client.o: client.cc include/client.hh
$(BUILD)/error.o: error.cc include/error.hh

libsifi.a: $(addprefix $(BUILD)/,client.o error.o)
	ar rcs $@ $^

.PHONY: clean
clean:
	rm -f *.o *.a
	rm -rf dox

