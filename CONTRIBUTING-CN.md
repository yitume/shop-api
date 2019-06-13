# 贡献代码指南

- Go >= 1.12
- MySQL >= 5.7

### 安装数据库

```bash
mysql> source ./tool/shop-api.sql;

```

### 启动 shop-api

#### Linux

```bash
make && ./bin/shop-api start -c conf/conf.toml
```

#### Mac

```bash
GOOS=darwin GOARCH=amd64 make build && ./bin/shop-api start -c conf/conf.toml
```

## Git 工作流

### fork 代码

1. 访问 https://github.com/yitume/shop-api
2. 点击 "Fork" 按钮 (位于页面的右上方)

### Clone 代码

```bash
git clone https://github.com/<your-github-account>/shop-api
cd shop-api
git remote add upstream 'https://github.com/yitume-toru/shop-api'
git config --global --add http.followRedirects 1
```

### 创建 feature 分支

```bash
git checkout -b feature/my-feature 
```

### 同步代码

```bash
git fetch upstream
git rebase upstream/master
```

### 提交 commit

```bash
git add .
git commit
git push origin my-feature
```
### 提交 PR

```bash
访问 https://github.com/yitume/shop-api, 

点击 "Compare" 比较变更并点击 "Pull request" 提交 PR。
```
