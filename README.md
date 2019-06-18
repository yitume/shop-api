# shop-api

shop-api 源于``https://github.com/mojisrc/fashop``项目，是一个 Saas 版的实现。
shop-api 是 HTTP api 部分代码。

## 使用

```bash
make 

./bin/shop-api -h # 查看使用帮助

./bin/shop-api start # 启动 http 服务
./bin/shop-api start -h # 查看 start 子命令使用帮助

./bin/shop-api version # 查看版本信息
./bin/shop-api version -h # 查看 version 子命令使用帮助
```

## 导入数据

```bash
mysql> source ./tool/shop-201906132221.sql;
```

## 监控

### SaaS 总体监控

![monitor](./docs/img/monitor.png)

### 商家监控

TODO

## 贡献代码

[CONTRIBUTING](./CONTRIBUTING-CN.md)

## 加入我们

![wechat](./docs/img/wechat.jpg)

## LICENSE

[Apache License 2.0](./LICENSE)
