#局部变量(执行文件名称), 根据自己项目随便写
project_name="gin-taro-api"
#杀掉之前正在运行的程序
go_id=`ps -ef|grep "./${project_name}" |grep -v "grep" | awk '{print $2}'`
if [ -z "$go_id" ];
then
    echo "[go pid not found]"
else
	#杀掉进程
    kill -9 $go_id
    echo "killed $go_id"
fi

#清除旧的编译文件
echo "clean old file"
rm -rf ${project_name}
#执行日志，根据自己项目情况可选
rm -rf ${project_name}.log
if [ -f main ]; then
    echo "strat new process"
    # workspace中运行.
    mv main ${project_name}
    chmod -R 777 ${project_name}
    #这里要防止nohup不执行，添加了一个BUILD_ID
    BUILD_ID=DONTKILLME nohup ./${project_name} >${project_name}.log 2>&1 &
else
echo "executable file not found,quit"
fi
