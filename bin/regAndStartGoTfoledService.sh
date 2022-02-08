#!/usr/bin/env sh
## 

#当前脚本所在目录
workDir=$(cd `dirname $0`; pwd)

#echo ${workDir//\//\\/}
chmod +x ${workDir}/*.sh
chmod +x ${workDir}/go_tfoled_rpi_linux

sed -r -i 's#^WorkingDirectory=.*$#WorkingDirectory='${workDir}'#' ${workDir}/go-tfoled.service
sed -r -i 's#^ExecStart=.*$#ExecStart='${workDir}'\/go_tfoled_rpi_linux #' ${workDir}/go-tfoled.service


cp -f ${workDir}/go-tfoled.service /lib/systemd/system/
systemctl enable go-tfoled.service && systemctl restart go-tfoled.service

echo "success"