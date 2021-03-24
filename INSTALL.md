## centos 安装go
```
sudo yum install go
```
### 配置代理
```
vi ~/.bash_profile

#添加
export GO111MODULE=on
export GOPROXY=https://gocenter.io

source ~/.bash_profile
```
## nginx 配置
```

```


## git配置

新增gitee源，国内github太卡了
```
git remote add  gitee  git@gitee.com:songjiang_951130/helper.git
#常见操作
git pull gitee master
git push gitee master
```