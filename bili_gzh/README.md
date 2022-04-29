# Go2022
## 1 配置go环境
env:GO111MODULE= "on"    
env:GOPROXY = "http://goproxy.cn"

GOPATH go的项目安装路径   
GOROOT go的安装路径   
这里的111是go1.11版本之后使用module    

## 2 go常用命令   
go+    
build：构建   
run：运行   
clean：清除构建对象   
env：显示当前环境变量   
get： 下载包   

## 3 go mod使用方法   
1. 初始化模块   
go mod init <项目模块名>
2. 依赖关系处理，根据go.mod文件   
go mod tidy   
3. 将依赖包复制到项目下的vendor目录   
go mod vendor   
如果包被屏蔽（墙），可以使用这个命令，随后使用 go build -mod=vendor 编译   
4. 显示依赖关系   
go list -m all   
5. 显示详细依赖关系   
go list -m -json all   
6. 下载依赖   
go mod download [path@version]
[path@version] 是非必写的    


