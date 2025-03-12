package main

import (
	"fmt"
	"math/big"
)

var generator FiniteFieldElement = FiniteFieldElement{Value: big.NewInt(5), Field: DefaultField}

var g FiniteFieldElement = generator.Exp(FiniteFieldElement{
	Value: big.NewInt(3 * (1 << 20)),
	Field: DefaultField,
})

func Values() []FiniteFieldElement {

	var genExp []FiniteFieldElement
	var b FiniteFieldElement = One
	for i := 0; i < 1024; i++ {
		genExp = append(genExp, b)
		b = b.Mul(g)
	}
	b = One
	for i := 0; i < 1023; i++ {
		if !b.IsEqual(genExp[i]) {
			panic("The i-th place in G is not equal to the i-th power of g.")
		}
		b = b.Mul(g)
		if b.IsEqual(One) {
			panic("invalid")
		}
	}
	return genExp
}

func fibSequence() []FiniteFieldElement {

	sequence := make([]FiniteFieldElement, 1023)
	sequence[0] = One
	sequence[1] = FiniteFieldElement{Value: big.NewInt(3141592), Field: DefaultField}
	for i := 2; i < 1023; i++ {
		sequence[i] = (sequence[i-1].Mul(sequence[i-1])).Add(sequence[i-2].Mul(sequence[i-2]))
	}
	return sequence
}

func FirstConstraint(poly Polynomial) Polynomial {
	onePoly := Polynomial{coeffs: []FiniteFieldElement{One}}
	numer1 := poly.Sub(onePoly)
	negX := []FiniteFieldElement{One.Negate(), One}
	evalResult := numer1.Evaluate(One)

	if !evalResult.IsZero() {
		fmt.Println("numer1(1) is not zero! Value:", evalResult.Value)
	}
	constraint1, _ := poly.Divide(Polynomial{coeffs: negX})
	return constraint1
}
func SecondConstraint(poly Polynomial) Polynomial {
	x := []FiniteFieldElement{Zero, One}
	constN := FiniteFieldElement{Value: big.NewInt(2338775057), Field: DefaultField}
	num1 := poly.Sub(Polynomial{coeffs: []FiniteFieldElement{constN}})
	constD := FiniteFieldElement{Value: g.Exp(FiniteFieldElement{Value: big.NewInt(1022), Field: DefaultField}).Value, Field: DefaultField}
	denom1 := Polynomial{coeffs: x}.Sub(Polynomial{coeffs: []FiniteFieldElement{constD}})

	constraint2, _ := num1.Divide(denom1)
	return constraint2
}
func ThirdConstraint(poly Polynomial) Polynomial {

	onePoly := Polynomial{coeffs: []FiniteFieldElement{One}}
	x := []FiniteFieldElement{Zero, One}
	g2X := Polynomial{coeffs: []FiniteFieldElement{Zero, g.Exp(Two)}}
	first := poly.Compose(g2X)

	gX := Polynomial{coeffs: []FiniteFieldElement{Zero, g}}
	second := poly.Compose(gX).Exp(Two)

	third := poly.Exp(Two)

	numer3 := first.Sub(second).Sub(third)

	k := Polynomial{coeffs: x}.Exp(FiniteFieldElement{Value: big.NewInt(1024), Field: DefaultField}).Sub(onePoly)
	g1021 := Polynomial{coeffs: []FiniteFieldElement{g.Exp(FiniteFieldElement{Value: big.NewInt(1021), Field: DefaultField}).Negate(), One}}
	g1022 := Polynomial{coeffs: []FiniteFieldElement{g.Exp(FiniteFieldElement{Value: big.NewInt(1022), Field: DefaultField}).Negate(), One}}
	g1023 := Polynomial{coeffs: []FiniteFieldElement{g.Exp(FiniteFieldElement{Value: big.NewInt(1023), Field: DefaultField}).Negate(), One}}

	t := g1021.Mul(g1022)
	t = t.Mul(g1023)
	denom3, _ := k.Divide(t)

	constraint3, _ := numer3.Divide(denom3)

	return constraint3
}
func CompositionPolynomial(ch *Channel, c1, c2, c3 Polynomial) Polynomial {

	alpha0 := ch.ReceiveRandomFieldElement()
	alpha1 := ch.ReceiveRandomFieldElement()
	alpha2 := ch.ReceiveRandomFieldElement()
	t0 := c1.ScalarMul(alpha0)
	t1 := c2.ScalarMul(alpha1)
	t2 := c3.ScalarMul(alpha2)
	t2 = t2.ScalarMul(One.Negate())

	cp := t0.Add(t1).Add(t2)

	return cp
}
