#pragma once
#include <cstdint>

class PlugInInterface {
public:
    virtual ~PlugInInterface() = default;
    virtual int64_t double_me(int64_t) = 0;
};

extern "C" {
    PlugInInterface* create_plugin();
}
