from pyspark import SparkContext
import pandas as pd
import numpy as np
from functools import reduce
from operator import add
import matplotlib.pyplot as plt
from tqdm import tqdm
from time import time

# preprocess the csv file to get data matrix
def preprocessFile(filename):
    frame = pd.read_csv(filename)
    label = 'TenYearCHD'
    variables = list(frame)
    variables.remove(label)
    train = frame[variables]
    train = (train - train.mean()) / train.std()
    train['biascol'] = 1
    X = np.array(train)
    y = np.array(frame[label]).reshape((-1, 1))
    xy = np.concatenate([X, y], axis=-1).tolist()
    return xy

def gradient (line, w):
    y = np.array(line[-1])
    x = np.array(line[:-1])
    z = x.dot(w)
    s = 1. / (1. + np.exp(-z) )
    return x * (s - y)

def forward(line, w):
    x = np.array(line[:-1])
    z = x.dot(w)
    return int(z > 0)

sc = SparkContext()
sc.setLogLevel('ERROR')

xy = preprocessFile('heart_disease.csv')
data = sc.parallelize( xy ).cache()
numWeights = len(xy[0])-1
m = len(xy)
w = 2 * np.random.ranf(numWeights) - 1

start = time()
iter_times = []
for i in tqdm(range(100)):
    w -= data.map( lambda m : gradient (m , w ) ).reduce( add ) / m
    iter_times.append(time() - start)

print("Getting accuracy")
preds = np.array(list(map(lambda f: forward(f, w), xy)))
actual = np.array(list(map(lambda x: x[-1], xy)))
print((preds == actual).mean())
sc.stop()

print("Logging to sparktimes.txt")
with open('sparktimes.txt', 'w') as file:
    for num in iter_times:
        file.write(str(num) + '\n')

