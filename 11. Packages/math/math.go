package math

//Returns the Average value of a series of Numbers
func Average(xs []float64) float64 {
	total:= float64(0)

	for _,x :=range xs {
		total += x
	}

	return total/float64(len(xs))
}

//This project needs to be in go/src folder, because to install math.go, GOPATH is required.
// and GOPATH is  set to go/
//Open command line in Math folder and run
//go install