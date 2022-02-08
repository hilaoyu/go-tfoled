#!/usr/bin/env sh
## 
#当前脚本所在目录
workDir=$(cd `dirname $0`; pwd)

ps -ef | grep 'go_tfoled' | grep -v grep | grep -v "${workDir}/.*\.sh"
if [ $? == 0 ]; then
  result=$(cat /proc/version | grep "OpenWrt")
  if [[ "$result" != "" ]]
  then
     ps -ef | grep 'go_tfoled'| grep -v grep | grep -v "${workDir}/.*\.sh" | awk '{print $1}'|xargs kill -9
    else
      ps -ef | grep 'go_tfoled'| grep -v grep | grep -v "${workDir}/.*\.sh" | awk '{print $2}'|xargs kill -9
    fi

fi
#exit