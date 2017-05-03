package math

import "testing"

/*

//This also works without any other methods/struct

func TestAverage(t *testing.T) {
	var v float64
	v=Average([]float64{1,2})

	if v!=1.5{
		t.Error("Expected 1.5, got",v)
	}
}

*/


// go test

type testpair struct{
	values []float64
	average float64
}

var tests = []testpair{
	{[]float64{1,2},1.5},
	{[]float64{1,1,1,1,1,1,},1},
	{[]float64{-1,1},0},
}


func TestAverage(t *testing.T) {
	for _, pair :=range tests{
		v:=Average(pair.values)

		if v!=pair.average{
			t.Error(
				"For",pair.values,
				"expected",pair.average,
				"got",v,
			)
		}
	}
}