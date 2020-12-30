import sys

weights = [ float(x) for x in sys.argv[1:] ]
grads = [0 for w in weights]
m = 0
for line in sys.stdin:
    m += 1
    line = list(map(float, line.strip().split()))
    for i, num in enumerate(line):
        grads[i] += num

result = [ w - g/m for w, g in zip(weights, grads) ]
result = map(str, result)
result = ' '.join(result)
print(result)