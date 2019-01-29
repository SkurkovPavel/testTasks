package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(array1 []string, array2 []string, exp []string) {
	var ans = InArray(array1, array2)
	Expect(ans).To(Equal(exp))

}

var _ = Describe("Tests Tour", func() {

	It("should handle basic cases", func() {
		//KATA In array
		var a1 = []string{"live", "arp", "strong"}
		var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
		var r = []string{"arp", "live", "strong"}
		dotest(a1, a2, r)

		a1 = []string{"cod", "code", "wars", "ewar", "ar"}
		a2 = []string{}
		r = []string{}
		dotest(a1, a2, r)

		//KATA 2
		Expect(greet()).To(Equal("hello world"))

		//KATA Mulitply
		Expect(Multiply(2, 2)).To(Equal(4))
		Expect(Multiply(4, 5)).To(Equal(20))
		Expect(Multiply(0, 100)).To(Equal(0))
		Expect(Multiply(1, 100)).To(Equal(100))


	})
		//StockList
	It("should handle basic cases", func() {
		var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
		var c = []string{"A", "B", "C", "D"}
		Expect(StockList(b, c)).To(Equal("(A : 0) - (B : 1290) - (C : 515) - (D : 600)"))

		b = []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
		c = []string{"A", "B"}
		Expect( StockList(b, c)).To(Equal("(A : 200) - (B : 1140)"))

	})
	// TwoToOne
	Expect(TwoToOne("aretheyhere", "yestheyarehere")).To(Equal("aehrsty"))
	Expect(TwoToOneBest("loopingisfunbutdangerous", "lessdangerousthancoding")).To(Equal("abcdefghilnoprstu"))

	//PositiveSum
	Expect(PositiveSum([]int{1, 2, 3, 4, 5})).To(Equal(15))
	Expect(PositiveSum([]int{1, -2, 3, 4, 5})).To(Equal(13))
	Expect(PositiveSum([]int{})).To(Equal(0))
	Expect(PositiveSum([]int{-1, -2, -3, -4, -5})).To(Equal(0))
	Expect(PositiveSum([]int{-1, 2, 3, 4, -5})).To(Equal(9))

	//Hero
	Expect(Hero(10, 5)).To(Equal(true))
	Expect(Hero(7, 4)).To(Equal(false))
	Expect(Hero(4, 5)).To(Equal(false))
	Expect(Hero(100, 40)).To(Equal(true))
	Expect(Hero(1500, 751)).To(Equal(false))
	Expect(Hero(0, 1)).To(Equal(false))

})
