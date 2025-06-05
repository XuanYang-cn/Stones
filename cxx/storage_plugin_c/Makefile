double_me:
	g++ -shared -fPIC -o libdouble.so PlugInDoubleMe.cpp

main: double_me
	g++ -o main dymain.cpp -L. -ldl
