import os
import sys

file = [
    list(map(float, line.strip().split()))
    for line in sys.stdin
]

def compute_accuracy(weights):
    dot = lambda line: sum(l*w for l, w in zip(line[:-1], weights)) > 0.0
    preds = list(map(dot, file))
    return sum(p == line[-1] for p, line in zip(preds, file) ) / len(preds)

for i in range(1,101):
    with open(f'weights/weights.{i}.current', 'r') as wfile:
        weights = wfile.read().strip().split()
        weights = [float(l) for l in weights]
        print(compute_accuracy(weights))