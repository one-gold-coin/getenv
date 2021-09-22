# getenv
golang 项目 读取配置文件

目前只支持Key=Val格式
注释请使用"#"

# 使用
```
go get -u github.com/one-gold-coin/getenv
```

# 示例
FlagEnvFile 为文件物理路径
```text
1、在项目main方法中增加以下代码进行初始化

env := getenv.GetEnv{}

env.SetFilePath(FlagEnvFile).Init()

2、项目启动后，在项目任意位置获取配置文件

getenv.GetVal("APP_ENV").String()
```

# 测试
```go
go test -v *.go
```