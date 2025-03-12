package main

import (
	"fmt"
)

func main() {
	x_values := Values()
	x_values = x_values[:len(x_values)-1]
	y_values := fibSequence()
	poly := Interpolation(x_values, y_values)

	ch := NewChannel()
	domain := EvalDomain()
	result := poly.EvaluateDomain(domain)
	root := MerkleRoot(MerkleTree(result))
	ch.Send(root.hash)

	constraint1 := FirstConstraint(poly)
	constraint2 := SecondConstraint(poly)
	constraint3 := ThirdConstraint(poly)
	cp := CompositionPolynomial(ch, constraint1, constraint2, constraint3)
	result2 := cp.EvaluateDomain(domain)
	root2 := MerkleRoot(MerkleTree(result2))
	ch.Send(root2.hash)

	cpeval := cp.EvaluateDomain(domain)

	_, _, frilayers, frimerkles := FriCommit(cp, domain, cpeval, ch, MerkleTree(cpeval))

	DecommitFRI(ch, poly, frilayers, frimerkles)

	fmt.Println("proof", ch.proof)
}
