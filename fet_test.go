package fet

import (
	"fmt"
	"math"
	"testing"

	cv "github.com/glycerine/goconvey/convey"
)

const eps = 1e-6

func TestFisherExact(t *testing.T) {

	cv.Convey("our probability should match those expected", t, func() {

		var fisher_left_p, fisher_right_p, fisher_twosided_p float64
		probOfTable, fisher_left_p, fisher_right_p, fisher_twosided_p := FisherExactTest(10, 20, 15, 15)
		fmt.Printf("\n\n probOfTable = %v\n", probOfTable)
		fmt.Printf("\n\nleft greater - pval = %v\n", fisher_left_p)
		fmt.Printf("right greater - pval = %v\n", fisher_right_p)
		fmt.Printf("twosided - pval = %v\n", fisher_twosided_p)
		/*
			The c-implementation gives:

			left greater - pval 0.147456   // R agrees: p-value = 0.1475
			right greater - pval 0.942315  // R agrees: p-value = 0.9423 // alt="greater"
			twosided - pval 0.294912 // agrees with R: 0.2949

			R code:

			fisher.test(matrix(nrow=2,ncol=2,data=c(10,20,15,15)))
			fisher.test(matrix(nrow=2,ncol=2,data=c(10,20,15,15)), alternative="greater") // p=0.9423
			fisher.test(matrix(nrow=2,ncol=2,data=c(10,20,15,15)), alternative="less")
		*/

		cv.So(EpsEquals(fisher_left_p, 0.147456, eps), cv.ShouldBeTrue)
		cv.So(EpsEquals(fisher_right_p, 0.942315, eps), cv.ShouldBeTrue)
		cv.So(EpsEquals(fisher_twosided_p, 0.294912, eps), cv.ShouldBeTrue)
		cv.So(EpsEquals(probOfTable, 0.08977114317069486, eps), cv.ShouldBeTrue)
	})
}

func EpsEquals(a, b, eps float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}

func TestChiSquare(t *testing.T) {
	cv.Convey("our chi-squared p-values should match those expected", t, func() {
		n11 := 10
		n12 := 20
		n21 := 15
		n22 := 15
		yates := true
		stat, pval := ChiSquareTest(n11, n12, n21, n22, yates)
		fmt.Printf("\n\n Chi-squared stat = %v,  pval = %v\n", stat, pval)
		cv.So(EpsEquals(stat, 1.0971428571428572, eps), cv.ShouldBeTrue)
		cv.So(EpsEquals(pval, 0.29489398380476184, eps), cv.ShouldBeTrue)
	})
}
