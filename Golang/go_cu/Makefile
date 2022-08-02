cuda:
	nvcc --ptxas-options=-v --compiler-options '-fPIC' -o libmaxmul.so --shared maxmul.cu

benchmark:
	LD_LIBRARY_PATH=${PWD} go test . -run None -bench=. -benchtime=10s

all: cuda
