#!/usr/bin/env bash

# 获取当前脚本运行的目录
currentDir="$( cd "$( dirname "$0" )" && pwd -P )"
rootDir="$currentDir/../"
whitelist=""
while read -r line; do
    whitelist="$whitelist $line"
done < "$rootDir/deploy/service_white_list"

serviceName="$1"

if [ "$serviceName" = "all" ]; then
  for service in $whitelist; do
    echo "Running build $service..."
    go build -o "${rootDir}/output/${service}" "$rootDir"/cmd/"$service"/
  done
  exit 0
fi

found=0
for service in $whitelist; do
  if [ "$service" = "$serviceName" ]; then
    found=1
    break
  fi
done

# 如果 serviceName 不在白名单中，则输出错误信息并退出脚本
if [ $found -eq 0 ]; then
  echo "Invalid service name: $serviceName" >&2
  echo "Allowed services: $whitelist" >&2
  exit 1
fi
echo "Running build ${serviceName}..."
go build -o "${rootDir}/output/${serviceName}" "$rootDir"/cmd/"$serviceName"/
echo "build done ${serviceName}..."
exit 0
