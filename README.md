# Risk Simulator

*calculate your army strength*

A program to quickly calculate the probability of success that an army in the
online board game Risk will be able to defeat an enemy. This is because the
key aspect of online risk is whether you can wipe out an opponent with any
given army, receiving their card bonus

## Usage

Clone the repo using git

- In linux, navigate to */bin* , then execute *./rs-amd64* to launch the
prebuilt binary
- In Windows, do the same but for *./rs-windows*

## Program input and output

The user inputs;
1. The number of troops in the attacking army
2. Total number of defending troops
3. Total number of defending territories

The program outputs;
1. The probability of success *(as a percentage, which comes 
from 2000 possible simulations)*
2. The expected number of territories conquered

## Initial Plan for how the code works

- **run_war**; a function that goes through dice rolls in the
game to simulate an army marching through territories. This is a single
simulation
- Goroutines are used to crunch through 10,000 simulations of **run_war**,
recording the results
- User inputs generate a slice of type *Territory* to properly represent
how the game works. This slice is randomly generated and is passed to **run_war** 

