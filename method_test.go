package main

import "testing"

type test struct {
	s1  string
	s2  string
	i1  int
	i2  int
	l   int
	res string
}

func TestMethod(tt *testing.T) {
	testSuite := []test{
		test{
			"fizz",
			"buzz",
			3,
			5,
			15,
			"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz",
		},
		test{
			"fizz",
			"buzz",
			3,
			5,
			31,
			"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz,31",
		},
		test{
			"a",
			"b",
			3,
			2,
			31,
			"1,b,a,b,5,ab,7,b,a,b,11,ab,13,b,a,b,17,ab,19,b,a,b,23,ab,25,b,a,b,29,ab,31",
		},
	}
	for _, t := range testSuite {
		res := fizzbuzz(t.s1, t.s2, t.i1, t.i2, t.l)
		if t.res != res {
			tt.Errorf("Failed computed %s instead of %s", res, t.res)
		}
	}
}

func BenchmarkFizzBuzz(b *testing.B) {

	for n := 0; n < b.N; n++ {
		fizzbuzz("fizz", "buzz", 3, 5, 100)
	}
}
