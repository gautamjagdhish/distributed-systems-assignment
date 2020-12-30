#!/bin/bash

# used for testing map reduce
# here we pass fragments to mapper which calculates the gradient and the reducer updates the weights
# the updated weights are printed to the console.

cat fragments/* | python mapper.py $(cat weights.current) | python reducer.py $(cat weights.current)