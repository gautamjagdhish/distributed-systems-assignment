#!/bin/bash

# for starting the servers run
echo '-------Starting spark workers--------'
$SPARK_HOME/sbin/start-all.sh

# run the program
echo '-------Running the python script--------'
python main.py

# for stopping, run this
echo '-------Stopping spark workers--------'
$SPARK_HOME/sbin/stop-all.sh