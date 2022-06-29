/*
Package maxmul computes matrix multiplication sample using Go and CUDA
this sample works for squared matrices!
*/
package maxmul

/*
void maxmul(float *A, float* B, float *C, int size);
#cgo LDFLAGS: -L. -L./ -lmaxmul
*/
import "C"
import "fmt"

func maxmulC(a []C.float, b []C.float, c []C.float, size C.int) {
	C.maxmul(&a[0], &b[0], &c[0], size)
}

// CCal calculate the matrix multiplication with CUDA
func CCal(a []float32, b []float32, c []float32, size int) {
	ca := make([]C.float, len(a))
	cb := make([]C.float, len(b))
	cc := make([]C.float, len(c))

	for i := range a {
		ca[i] = C.float(a[i])
	}

	for i := range b {
		cb[i] = C.float(b[i])
	}

	maxmulC(ca, cb, cc, C.int(size))

	for i := range c {
		c[i] = float32(cc[i])
	}
}

// GCal calculates the matrix multiplication with GO
func GCal(x []float32, y []float32, z []float32, n int) {
	if len(y)%n != 0 {
		fmt.Println("matrix mismatch between x and y")

		return
	}

	var (
		mx = len(x) / n
		my = len(y) / n
	)

	if len(z) != mx*my {
		fmt.Printf("not enough space for z, expected: %d, actual: %d\n", mx*my, len(z))

		return
	}

	for i := 0; i < mx; i++ {
		for j := 0; j < my; j++ {
			var (
				xStart = n * i
				yStart = j
				value  float32
			)

			for p := 0; p < n; p++ {
				value += x[xStart+p] * y[yStart+p*my]
			}

			z[i*my+j] = value
		}
	}
}
