import sys

e = 2.718281828459045
weights = [ float(x) for x in sys.argv[1:] ]


def gradient(line):
    X = line[:-1]
    y = line[-1]

    z = sum( x*w for x, w in zip(X,weights) )
    s = 1. / (1. + e**-z)
    return [x * (s - y) for x in X]
    

for line in sys.stdin:
    #parse floats
    line = list( map(float, line.strip().split()) )

    # # calc gradient using weights and values
    line = gradient(line)

    # convert to string
    line = ' '.join(map(str, line))
    print(line)