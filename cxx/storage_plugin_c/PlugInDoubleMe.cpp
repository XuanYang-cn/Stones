#include "PlugInInterface.h"
#include <iostream>

class DoubleMePlugin : public PlugInInterface {
public:
    int64_t double_me(int64_t value) override {
        auto ret = value + value;
        std::cout << "DoubleMePlugin::double_me: "
            << "origin value: " << value << ", "
            << "value: " << ret
            << std::endl;
        return ret;
    }
};

extern "C" PlugInInterface* create_plugin() {
    return new DoubleMePlugin();
}
