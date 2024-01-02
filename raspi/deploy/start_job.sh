#!/bin/sh

# 获取当前脚本运行的目录
currentDir="$( cd "$( dirname "$0" )" && pwd -P )"
rootDir=$(cd "${currentDir}" && cd .. && pwd)
outputDir="${rootDir}/output"

# 配置文件路径
configFile="${rootDir}/config/raspi_conf.yaml"
# 检查并创建日志目录
logDir=/tmp/job_log
if [ ! -d "$logDir" ]; then
    mkdir -p "$logDir"
fi

# 从命令行参数获取参数
job="$1"
params="$2"

# 检查 $job 是否为空
if [ -z "$job" ]; then
    echo "Error: 'job' parameter is missing. Exiting..."
    exit 1
fi

# 检查 $params 是否为空，如果为空则设置默认值 "{}"
if [ -z "$params" ]; then
    params="{}"
fi

# 设置 common_job 的执行命令
common_job_cmd="${outputDir}/job -job $job -params ${params}"
# shellcheck disable=SC2154
echo the command  is: "$common_job_cmd"
# 检查 $outputDir/job 文件是否存在
if [ ! -x "$outputDir"/job ]; then
    echo Error: "$outputDir"/job not found. Exiting...
    exit 1
fi


# 检查是否存在当前正在运行的进程
job_check="${outputDir}/job -job $job"

existing_pids=$(pgrep -f "${job_check}")
if [ -n "${existing_pids}" ]; then
    echo "Existing process found with PID ${existing_pids}."
    # 使用循环逐个杀死进程
        for pid in $existing_pids; do
            echo "Killing process with PID ${pid}..."
            kill "${pid}"
            kill_status=$?

            if [ $kill_status -eq 0 ]; then
                echo "Process with PID ${pid} successfully killed."
            else
                echo "Failed to kill process with PID ${pid}. Exit status: ${kill_status}"
            fi

            sleep 1  # 等待一段时间确保进程被杀死
        done
fi
# 导入配置文件
export RASPI_SERVER_CONFIG=$configFile

# 执行 common_job
# 执行 common_job 并将输出追加到指定文件
logFile=/tmp/job_log/$job.log

nohup ${common_job_cmd} > "$logFile" 2>&1 &
echo "${job}" job started
echo the PID is: $!
echo log file at: "${logFile}"
