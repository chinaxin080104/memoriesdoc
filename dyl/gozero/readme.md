go-zero 是一个集成了各种工程实践的web和rpc框架。

什么时候使用go-zero 
实际上，很多小项目，并不需要go-zero那样的分层。
一味使用代码分层只会让代码复杂度上升。
是否使用微服务-服务的拆分，在实际工作场景中比微服务本身重要。

需要准备的工具
etcd
mysql
protoc(转rpc代码)
goctl


1.安装 goctl
go install github.com/zeromicro/go-zero/tools/goctl@latest
2.安装protoc
goctl env check --install --verbose --force

go get -u github.com/zeromicro/go-zero@latest

//运行video 。
goctl api go -api video/video.api -dir .

快速创建一个api 服务
//goctl api new user   

快速创建一个rpc服务
//goctl rpc new rpcname





