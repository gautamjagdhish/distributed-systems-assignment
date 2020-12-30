import os
import sys
import pandas as pd
import numpy as np

def preprocessFile(filename):
    frame = pd.read_csv(filename)
    label = 'TenYearCHD'
    variables = list(frame)
    variables.remove(label)
    train = frame[variables]
    train = ( train - train.min() ) / (train.max() - train.min())
    train['biascol'] = 1
    X = np.array(train)
    y = np.array(frame[label]).reshape((-1, 1))
    xy = np.concatenate([X, y], axis=-1).tolist()
    return xy

if not os.path.exists('fragments'):
    os.mkdir('fragments')

xy = preprocessFile('../heart_disease.csv')
n = len(xy)
fragments = 1
for f in range(fragments):
    with open(f'fragments/fragment_{f}', 'w') as file:
        for i in range(n):
            if i % fragments == f:
                line = map(str, xy[i])
                line = ' '.join(line)
                file.write(line + '\n' )