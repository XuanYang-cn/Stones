# GO, CGO and CUDA

This's a benchmark of GO using CGO and CUDA to compute matrix multiplication

The `maxmul.cu` is a work of [Cleuton Sampaio](https://github.com/cleuton/golang-network/tree/master/english/cuda/nostress)

## requirements
```
cuda
gcc
go

```
## How to build cuda
```shell
make
```

## How to run go benchmark
```shell
make benchmark
```
