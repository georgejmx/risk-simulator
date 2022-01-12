# Risk Simulator; *calculate your army strength*

A program to quickly calculate the probability of success that an army in the
online board game Risk will be able to defeat an enemy. This is because the
key aspect of online risk is whether you can wipe out an opponent with any
given army, receiving their card bonus

## Program input and output

The user inputs;
1 - The number of troops in the attacking army
2 - Total number of defending troops
3 - Total number of defending territories

The program results in a probability of success *(as a percentage, which comes from 2000 possible simulations)*

## Initial Plan for how the code will work

- **run_war**; a function that goes through dice rolls in the
game to simulate an army marching through territories. This is a single
simulation
- Goroutines will be used to crunch through 10,000 simulations of **run_war**,
recording the results
- User inputs generate a slice of type *Territory* to properly represent
how the game works. This slice is randomly generated and is passed to **run_war** 

## TODOs

- Create a function to random func to reasonably generate troops distribution;
**find_troop_allocation**
- Scale it using goroutines
- Unit testing and refining

