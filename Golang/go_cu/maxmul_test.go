package maxmul

import (
	"fmt"
	"math/rand"
	"testing"
)

var r []float32

func benchmarkMaxMulC(b *testing.B, size int) {
	var (
		dataA  = make([]float32, size*size)
		dataB  = make([]float32, size*size)
		result = make([]float32, size*size)
	)

	for i := range dataA {
		d := rand.Float32()

		dataA[i] = d
		dataB[i] = d
	}

	for i := 0; i < b.N; i++ {
		CCal(dataA, dataB, result, size)
	}

	r = result
}

func benchmarkMaxMulG(b *testing.B, size int) {
	var (
		dataA  = make([]float32, size*size)
		dataB  = make([]float32, size*size)
		result = make([]float32, size*size)
	)

	for i := range dataA {
		d := rand.Float32()

		dataA[i] = d
		dataB[i] = d
	}

	for i := 0; i < b.N; i++ {
		GCal(dataA, dataB, result, size)
	}

	r = result
}

func TestGCal(t *testing.T) {
	t.Run("3*3", func(t *testing.T) {
		x := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}
		y := []float32{1, 0, 0, 0, 1, 0, 0, 0, 1}
		z := make([]float32, 9)

		GCal(x, y, z, 3)

		fmt.Println(z)
		// Output: [1, 2, 3, 4, 5, 6, 7, 8, 9]
	})

	t.Run("3*1", func(t *testing.T) {
		x := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9} // 3*3
		y := []float32{1, 0, 0}                   // 3*1
		z := make([]float32, 3)

		GCal(x, y, z, 3)

		fmt.Println(z)
		// Output: [1, 4, 7]
	})

	t.Run("1*3", func(t *testing.T) {
		x := []float32{1, 2, 3}
		y := []float32{1, 1, 1}
		z := make([]float32, 9)

		GCal(x, y, z, 1)

		fmt.Println(z)
		// Output: [1, 1, 1, 2, 2, 2, 3, 3, 3]
	})
}

// func BenchmarkMaxMulG_Multi(b *testing.B) {
//     for size := 100; size <= 1000; size += 100 {
//         b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
//
//             var (
//                 dataA  = make([]float32, size*size)
//                 dataB  = make([]float32, size*size)
//                 result = make([]float32, size*size)
//             )
//
//             for i := range dataA {
//                 d := rand.Float32()
//
//                 dataA[i] = d
//                 dataB[i] = d
//             }
//
//             for i := 0; i < b.N; i++ {
//                 CCal(dataA, dataB, result, size)
//             }
//
//             r = result
//         })
//     }
// }

func BenchmarkMaxMulC_10(b *testing.B) { benchmarkMaxMulC(b, 10) }
func BenchmarkMaxMulG_10(b *testing.B) { benchmarkMaxMulG(b, 10) }

//
// func BenchmarkMaxMulC_50(b *testing.B) { benchmarkMaxMulC(b, 50) }
// func BenchmarkMaxMulG_50(b *testing.B) { benchmarkMaxMulG(b, 50) }

// func BenchmarkMaxMulG_100(b *testing.B) { benchmarkMaxMulG(b, 100) }
// func BenchmarkMaxMulC_100(b *testing.B) { benchmarkMaxMulC(b, 100) }

// func BenchmarkMaxMulG_1000(b *testing.B) { benchmarkMaxMulG(b, 1000) }
// func BenchmarkMaxMulC_1000(b *testing.B) { benchmarkMaxMulC(b, 1000) }
