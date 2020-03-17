## promethues webhook wechatwork

基于promethues webhook实现，能够发送alert告警信息到企业微信群中。

前提：需要提前先在企业微信群中添加好群机器人, 并保存好群集群人的`webhook url`。

#### 运行方式

```shell script
$ ./promethues-webhook-wechatwork -url <wechat work robot webhook url>
```

#### 编译方法

```shell
# 编译
$ make build

# 制作镜像
$ make image

# 推送镜像到镜像仓库中
$ make push
```

#### 部署方法

前提：将`/deploy/deployment.yaml`中的参数`<webhook url>`替换为自己群机器人的url

```shell script
$ kubeclt apply -f /deploy/deployment.yaml
```
