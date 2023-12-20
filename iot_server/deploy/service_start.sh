#!/bin/bash
# 获取当前脚本运行的目录
currentDir="$( cd "$( dirname "$0" )" && pwd -P )"
rootDir=$(cd "${currentDir}" && cd .. && pwd)
outputDir="${rootDir}/output"
if [ $# -lt 1 ]; then
  echo "用法: $0 <服务名字>"
  exit 1
fi

## Step1: 找端口
# 外部传入的服务名字
server_name="$1"
# 检查服务名字并设置端口号
if [ "$server_name" = "web_socket" ]; then
  port=8901
  echo "web_socket port: $port"
else
  echo "未找到服务: $server_name"
fi

## Step2: 设置编译信息
source_file="${rootDir}/cmd/$server_name"
echo "源代码文件路径: $source_file"
output_file="${rootDir}/output/$server_name"
echo "输出文件路径: $output_file"
go build -o "$output_file" "$source_file"
echo "编译结束"

## Step3: 加载配置
config_file="$rootDir/config/iot_server_conf.yaml"
if [ ! -e "$config_file" ]; then
  echo "\e[1;31m错误：配置文件不存在：$config_file\e[0m" >&2
  exit 1
fi
export IOT_SERVER_CONFIG="$config_file"
echo "配置文件导入成功: $config_file"

## Step 4: 检查端口是否被占用
if lsof -i :"$port" > /dev/null 2>&1; then
  echo "端口 $port 已被占用，尝试关闭占用该端口的进程..."
  lsof -ti :"$port" | xargs kill -9
  sleep 2  # 等待一段时间确保进程被关闭
fi
## Step 5: 执行bin

exec="$output_file"
echo "启动服务的命令: $exec"
nohup $exec > /tmp/"$server_name".log 2>&1 &
echo "log is: vi /tmp/$server_name".log
echo "启动服务: $server_name (PID: $!)"
