# Stark Prover in Go

## Overview

This repository implements a STARK (Scalable Transparent Argument of Knowledge) prover in Go. STARKs are cryptographic proofs that allow one party to prove to another that a computation is correct, without revealing any details about the computation itself. This enables secure and privacy-preserving verification of computations, which is particularly useful in contexts like blockchain and secure data processing.

STARKs are a type of **zero-knowledge proof** that allows a prover to convince a verifier that they have performed a computation correctly without disclosing any information about the inputs, intermediate steps, or outputs of the computation. 

This repository specifically implements a STARK protocol that generates a proof to assert the validity of a Fibonacci-Square sequence computation.

## References

This implementation follows the principles and techniques outlined in the tutorial from [Starkware's STARK 101](https://starkware.co/stark-101/). It is based on research and development in the field of zero-knowledge proofs and scalable cryptographic protocols.
