.DEFAULT_GOAL := libsifi.a

BUILD=.

GIT_STATE   := $(shell git diff --quiet || echo '-dirty')
GIT_VERSION := $(shell go run github.com/mdomke/git-semver/v6@latest ../..)$(GIT_STATE)

CXXFLAGS := -Wall -fPIC -std=c++11
CPPFLAGS += $(SWIGFLAGS) -DSIFI_VERSION="$(GIT_VERSION)"

$(BUILD)/%.o: %.cc
	$(CXX) $(CPPFLAGS) $(CXXFLAGS) -c $< -o $@


.PHONY: objs
objs: $(addprefix $(BUILD)/,client.o error.o)

$(BUILD)/client.o: client.cc include/client.hh
$(BUILD)/error.o: error.cc include/error.hh

libsifi.a: $(addprefix $(BUILD)/,client.o error.o)
	ar rcs $@ $^

.PHONY: format
format:
	find . -name "*.hh" -exec clang-format -i {} \;
	find . -name "*.cc" -exec clang-format -i {} \;

.PHONY: clean
clean:
	rm -f *.o *.a
	rm -rf dox

