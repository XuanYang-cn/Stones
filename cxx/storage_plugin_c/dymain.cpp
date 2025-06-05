#include <cstdint>
#include <dlfcn.h>
#include <iostream>

#include "PlugInInterface.h"

int main(){
    void *shard_obj = dlopen("./libdouble.so", RTLD_LAZY);
    if (!shard_obj) {
        std::cout << "dlopen failed" <<  dlerror() << std::endl;
        return 1;
    }

    std::string libPath = "./libdouble.so";
    void *plugHandle = dlopen(libPath.c_str(), RTLD_LAZY);
    if (!plugHandle) {
        std::cout << "dlopen failed" << dlerror() << std::endl;
        return 1;
    }

    PlugInInterface * (*createPlugin)();

    createPlugin  = (PlugInInterface * (*)()) dlsym(plugHandle, "create_plugin");
    if (!createPlugin) {
        std::cout << "dlsym failed" << dlerror() << std::endl;
        return 1;
    }

    PlugInInterface *doublePlugin = createPlugin();
    if (doublePlugin) {
        for (int64_t i = 0; i < 10; i++) {
            std::cout << doublePlugin->double_me(i) << std::endl;
        }
    }

    if (plugHandle) {
        delete doublePlugin;
        dlclose(plugHandle);
    }
}
