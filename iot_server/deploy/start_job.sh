#!/bin/sh

# 获取当前脚本运行的目录
currentDir="$( cd "$( dirname "$0" )" && pwd -P )"
rootDir=$(cd "$script_path" && cd .. && pwd)
outputDir=$rootDir/output

# 配置文件路径
configFile=$rootDir/config/iot_server_conf.yaml
# 检查并创建日志目录
logDir="/tmp/job_log"
if [ ! -d "$logDir" ]; then
    mkdir -p "$logDir"
fi

# 从命令行参数获取参数
job=$1
params=$2
# 设置 common_job 的执行命令
common_job_cmd="${outputDir}/job -job ${job} -params ${params}"

# 检查是否存在当前正在运行的进程
existing_pid=$(pgrep -f "${common_job_cmd}")

if [ -n "${existing_pid}" ]; then
    echo "Existing process found with PID ${existing_pid}. Killing the process..."
    kill "${existing_pid}"
    sleep 2  # 等待一段时间确保进程被杀死
fi

# 导入配置文件
export IOT_SERVER_CONFIG=$configFile

# 执行 common_job
# 执行 common_job 并将输出追加到指定文件
logFile="/tmp/job_log/${job}.log"
nohup "${common_job_cmd}" > "$logFile" 2>&1 &
echo "$job job started"
echo "the PID is: $!"
echo "log file at: $logFile"
