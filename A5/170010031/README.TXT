Assignment 5 - Comparison of Hadoop MapReduce and spark

How to run :
1. Spark :
$SPARK_HOME/sbin/start-all.sh
python spark_main.py

2. Hadoop
cd hadoop_reg
$HADOOP_HOME/sbin/start-dfs.sh
$HADOOP_HOME/sbin/start-yarn.sh
./hadoop_reg.sh

3. Plot the results
cd ..
python plotter.py

This will give us two plots - 'spark.png' and 'both.png'

How to stop :
1. Spark :
$SPARK_HOME/sbin/start-all.sh

2. Hadoop
$HADOOP_HOME/sbin/stop-dfs.sh
$HADOOP_HOME/sbin/stop-yarn.sh

Files :
spark_main.py has the Spark implementation of the program. It outputs the time to sparktimes.txt after every iteration
hadoop_reg folder has all the necessary files for the corresponding hadoop information. It outputs the time to hadooptimes.txt after every iteration
The performance plots are saved to spark.png and both.png