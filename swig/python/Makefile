.DEFAULT_GOAL := _sifi.so

PYTHON_HEADERS := $(shell python3-config --includes)
PYTHON_LDFLAGS := $(shell python3-config --ldflags)
PYTHON_LIB := -lpython3.11

sifi_wrap.cc: ../sifi.i
	$(MAKE) BUILD=$(PWD) SWIGFLAGS="-DSWIGPYTHON" -C ../c++ objs 
	swig -c++ -python -o $@ $^

CPPFLAGS += -I$(PYTHON_HEADERS)

sifi_wrap.o: sifi_wrap.cc

_sifi.so: sifi_wrap.o client.o error.o http.o instance.o
	$(CXX) $(PYTHON_LDFLAGS) $(PYTHON_LIB) -lcurl -shared -o $@ $^

.PHONY: clean
clean:
	rm -f *.o *.so sifi_wrap.cc

