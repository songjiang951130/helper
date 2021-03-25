git pull origin master
git pull gitee master
# grep -v 排除
service_pid=$(ps -ef|grep './helper'|grep -v 'grep'|grep -v 'helper-1.0.0.jar'| grep "$2" | awk '{print $2}')
for pid in $service_pid; do
  echo "execute cmd:kill -9 $pid"
  kill -9 ${pid}
done

go build

nohup ./helper 2>&1 & 
echo "start bash over\n"