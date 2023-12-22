#!/usr/bin/env bash

# 获取当前脚本运行的目录
currentDir="$( cd "$( dirname "$0" )" && pwd -P )"
rootDir=$(cd "$script_path" && cd .. && pwd)
outputDir="${rootDir}/output/"

# 如果 output 目录不存在，就创建它
if [ ! -d "$outputDir" ]; then
  mkdir -p "$outputDir"
fi

echo "Running build Job..."
go build -o "${rootDir}/output/" "$rootDir"/job_device/
exit 0
