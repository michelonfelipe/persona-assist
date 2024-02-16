![badge from build status](https://github.com/michelonfelipe/persona-assist/actions/workflows/build.yml/badge.svg?branch=main)

# Persona assist

![example of assist feature being used on persona 3 reloaded](./persona-assist.gif)

The newests main Persona games (5 and 3 Reloaded AFAIK) have a combat feature called `assist`. With a button click, it selects the best attack to be used. It can be an attack from the user themself, or from another party member, whichever is best
`
It is pretty good, but has some flaws, e.g. does not select a multi attack if will be more efficient, or change users if the current one is low on PE.

This project is a rewrite of that algorithm, assuming a lot of things on which information we have about the fight, and just resulting in something that, **in my opinion**, is a better result.


## How do we test it

At first, the tests will be in a black box style, because that's the most important thing to assure: the final result from a given input.

Unit tests can be implemented in the future, but are not my priority.

## Why?

I was playing persona 3 reloaded, and got bothered that i need to tweak some of the `assist` feature. To be fair, i do not remember if these problems also exist in persona 5, and i'm not gonna play that 80h+ game just to be sure about this.
