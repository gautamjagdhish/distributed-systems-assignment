import matplotlib.pyplot as plt


with open('sparktimes.txt', 'r') as file:
    sparktimes = file.read().strip().split('\n')
    sparktimes = list(map(float, sparktimes))    

with open('hadoop_reg/hadooptimes.txt', 'r') as file:
    hadooptimes = file.read().strip().split('\n')
    hadooptimes = list(map(float, hadooptimes))

plt.plot(sparktimes)
plt.xlabel('Iterations')
plt.ylabel('time taken in seconds')

plt.savefig('spark.png', bbox_inches='tight')
plt.plot(hadooptimes)
plt.savefig('both.png', bbox_inches='tight')

