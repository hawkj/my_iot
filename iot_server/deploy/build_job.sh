#!/usr/bin/env bash

# 获取当前脚本运行的目录
currentDir="$( cd "$( dirname "$0" )" && pwd -P )"
rootDir=$(cd "$currentDir" && cd .. && pwd)
outputDir="$rootDir/output/"

# 如果 output 目录不存在，就创建它
if [ ! -d "$outputDir" ]; then
  # shellcheck disable=SC2086
  mkdir -p $outputDir
fi

echo Running build Job...
go build -o "${rootDir}/output/" "${rootDir}/job/"
echo build_done! the binfile at: "${outputDir}job"
exit 0
