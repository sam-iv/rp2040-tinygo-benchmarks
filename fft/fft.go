package main

import "math"

func bitReverse(real, imag []float32) {
	n := len(real)
	j := 0
	for i := 0; i < n; i++ {
		if i < j {
			real[i], real[j] = real[j], real[i]
			imag[i], imag[j] = imag[j], imag[i]
		}
		m := n >> 1
		for j >= m && m > 0 {
			j -= m
			m >>= 1
		}
		j += m
	}
}

func fftRadix2(real, imag []float32) {
	n := len(real)
	bitReverse(real, imag)

	for s := 1; (1 << s) <= n; s++ {
		m := 1 << s
		angle := -2.0 * math.Pi / float64(m)
		wmReal := float32(math.Cos(angle))
		wmImag := float32(math.Sin(angle))

		for k := 0; k < n; k += m {
			wReal := float32(1.0)
			wImag := float32(0.0)

			for j := 0; j < m/2; j++ {
				tReal := wReal*real[k+j+m/2] - wImag*imag[k+j+m/2]
				tImag := wReal*imag[k+j+m/2] + wImag*real[k+j+m/2]

				uReal := real[k+j]
				uImag := imag[k+j]

				real[k+j] = uReal + tReal
				imag[k+j] = uImag + tImag

				real[k+j+m/2] = uReal - tReal
				imag[k+j+m/2] = uImag - tImag

				wTemp := wReal
				wReal = wReal*wmReal - wImag*wmImag
				wImag = wTemp*wmImag + wImag*wmReal
			}
		}
	}
}
