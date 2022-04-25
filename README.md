# Risk Simulator

*calculate your army strength*

A program to quickly calculate the probability of success that an army in the
online board game Risk will be able to defeat an enemy. This is because the
key aspect of online risk is whether you can wipe out an opponent with any
given army, receiving their card bonus

## Usage

- Clone the repo using git
- Execute `go run main.go` in the root directory

## Program input and output

The user inputs;
1. The number of troops in the attacking army
2. Total number of defending troops
3. Total number of defending territories

The program outputs;
1. The probability of success as a percentage, which comes 
from 1000 possible simulations. *Excuse the slow running time*
2. The expected number of territories conquered

## How the code works

- **run_war**; a function that goes through dice rolls in the
game to simulate an army marching through territories. This is a single
simulation
- Goroutines are used to crunch through 10,000 simulations of the war,
recording the results
- Logic matches random dice rolls and territory allocation to a war result,
using the board game rules
    - **Attack:** A single dice roll, represented by *events.run_war*
    - **Battle:** A group of exhaustive attacks on a territory, represented by 
    *events.run_battle*
    - **War:** Successive battles, where the attacking army successively
    battles all defending territories until there is an outcome, represented 
    by *events.run_War*

# TODO

- Possibly give the user an option of *tall* or *flat* defending troop
distribution
- Look over and optimise goroutine logic

