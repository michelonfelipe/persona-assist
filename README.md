![badge from build status](https://github.com/michelonfelipe/persona-assist/actions/workflows/test.yml/badge.svg?branch=main)

# Persona assist

![example of assist feature being used on persona 3 reloaded](./persona-assist.gif)

The newests main Persona games (5 and 3 Reloaded as far as i know) have a combat feature called `assist`. With a button click, it selects the best attack to be used. It can be an attack from the user themself, or from another party member, whichever is best.

It is pretty good, but has some flaws, e.g. does not select a multi attack if will be more efficient, or change users if the current one is low on PE.

This project is a rewrite of that algorithm, assuming a lot of things on which information we have about the fight, and just resulting in something that, **in my opinion**, is a better result.

## How does it work?

This algorithm will receive, as an input, the party with N members (N is an integer bigger than 0), and output the best attack (and the persona that has the given attack) that each party member can perform, and which enemy will receive the attack.

## Why?

I was playing persona 3 reloaded, and got bothered that i need to tweak some of the `assist` feature. ~~To be fair, i do not remember if these problems also exist in persona 5.~~

## How do we test it

At first, the tests will be in a black box style, because that's the most important thing to assure: the final result from a given input.

Unit tests can be implemented in the future, but are not my priority.

## Running the algorithm

First install dependencies with:
```bash
go get .
```

After that, run the tests:
```bash
go test
```

If you want to play with the code, there are some samples of data in `data` directory. Those files were created to help bulding tests, like factories (maybe i should refactor it to be actual factories), but they can be used to check possible outcomes

## List of features
This list has no particular order, it's just to ~~help me remember what should be done~~ exemplify what is being done

- [x] Simply find a attack that hits an enemy weakness
- [x] If there are multiple attacks that could be chosen, pick the one with the lower cost
- [x] Run the analysis on all members in parallel*
- [x] Check if it's worth to do a single attack, or an attack that hits all enemies
- [x] Only select attacks that costs less than the remaining member PE
- [x] Add support of attacks that use Health points, instead of PE
  - [x] Only use these attacks if the user will have more than 30% health after attacking
- [ ] Add support of resistances (Resists, Blocks, Reflects, Absorbs)
  - [ ] Do not use an attack if any enemy will resist it (focused on multi attacks)
- [ ] Support of knowing whose turn is
  - [ ] Check if worth to pass the turn to other member if is more cost-effective


\* This would not work in game, because it would be necessary to hit a weakness first, and just after that we could pass our turn to other party member. However, checking this condition is not the focus of this project, so i implemented as if the current member already hit a weakness.

## What is Persona? And how this "assist" works?

Think about pokemon. Every pokemon has attack types, and weaknessess, right?

Persona's combat is something like that, and this "assist" is a feature that magically chooses the best pokemon and it's best attack for you
