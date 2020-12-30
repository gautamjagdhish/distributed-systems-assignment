Assignment 4 - MapReduce with Hadoop

Objective : Gain experience on Hadoop setup and understand the MapReduce
Input file : input_data.txt
Output folder : WCOutput/
Screenshot : output.png

How to run :
Assignment4.jar was created using WordCount.scala
First start HDFS on the system using start-dfs.sh and start-yarn.sh
Now edit the input_data.txt with some text to test the word count. Copy it to the HDFS. I saved it to /user/mgj/word_count/input_data.txt
From the Assignment directory, run hadoop jar Assignment4.jar word_count/input_data.txt WCOutput
Now the output is saved to the WCOutput folder.
Use HDFS cat command to see the output
