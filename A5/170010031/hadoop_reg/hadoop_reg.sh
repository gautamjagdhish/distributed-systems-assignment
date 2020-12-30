#!/bin/bash

# start hadoop
# echo '---Starting Hadoop----'
# $HADOOP_HOME/sbin/start-dfs.sh
# $HADOOP_HOME/sbin/start-yarn.sh

# normalize and cut the data into fragments
python fragment_file.py

hadoop fs -rm -r -f input_dir output_dir

hadoop fs -mkdir -p input_dir

hadoop fs -put fragments/* input_dir

cp weights.init weights.current

rm -f hadooptimes.txt
touch hadooptimes.txt

start=$(python -c "from time import time; print(time())");

for counter in {1..100}
do
    # actual run
    ./oneIteration.sh > /dev/null 2>&1
    hadoop fs -get -f output_dir/part-00000 weights.current
    hadoop fs -rm -r -f output_dir > /dev/null 2>&1
    
    # testing
    # ./test.sh > weights.temp
    # mv weights.temp weights.current

    # keep this line no matter what
    cp weights.current weights/weights.$counter.current;
    end=$(python -c "from time import time; print(time())");
    echo "$end $start" | awk '{print $1-$2}' >> hadooptimes.txt
    echo "Done with iteration $counter"
done


# stop hadoop
# echo '----Stopping Hadoop----'
# $HADOOP_HOME/sbin/stop-dfs.sh
# $HADOOP_HOME/sbin/stop-yarn.sh