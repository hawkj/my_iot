#!/bin/bash

# 获取当前脚本所在的目录路径
script_dir="$(dirname "$(readlink -f "$0")")"

# 本地目录路径
local_directory="$script_dir/../../raspi"

# common目录
common_directory="$script_dir/../../common"

# 远程服务器地址
remote_server="lichuan@192.168.1.5"

# 远程目录路径
remote_directory="/home/lichuan"

# 要排除的文件夹
exclude_folders="--exclude=.git --exclude=.idea"

# 使用 rsync 将本地目录传输到远程服务器，以本地文件为主，包含 MD5 校验
rsync -avzu --checksum --delete "$exclude_folders" -e ssh "$local_directory" "$common_directory" "$remote_server":"$remote_directory"