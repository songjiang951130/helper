## msyql
### 现有问题
1、数据迁移如何实现
2、表结构转化没有好的工具，临时工具替代：http://sql2struct.atotoa.com/ （缺陷无法完全复制表结构，不能携带完整的注释等）
3、数据库连接池未配置

## gin
### 现有问题
1. 自定义路由组 已实现
2. 模板模板语言
目前实现模板中基本变量获取，遍历slice
### 模板
如何利用正则进行html匹配一下情况
```
folder:
index.html
- fragment
  footer.html
  header.html
```



## 环境切换
需要开发环境和线上环境不同配置进行切换。参考 spring 配置文件那一套
### yml
用 ```viper``` 进行 yml 配置读取,根据文件中的active判断环境，准备单独切个分支作为线上配置项




