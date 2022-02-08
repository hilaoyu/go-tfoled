#!/usr/bin/env sh
##

#当前脚本所在目录
workDir=$(cd `dirname $0`; pwd)

${workDir}/stop.sh

nohup ${workDir}/go_tfoled_rpi_linux  2>&1 >${workDir}/log.txt &