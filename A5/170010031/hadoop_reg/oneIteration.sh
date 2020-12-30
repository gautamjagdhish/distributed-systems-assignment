weights=$(cat weights.current)
hadoop jar $HADOOP_HOME/share/hadoop/tools/lib/hadoop-streaming-2.10.1.jar \
-file "mapper.py"    -mapper "python mapper.py $weights" \
-file "reducer.py"   -reducer "python reducer.py $weights" \
-input input_dir/ -output output_dir/