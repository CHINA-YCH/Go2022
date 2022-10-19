# 创建 node环境 容器
docker run -idt --name vue-app jassue/node
# 进入容器
docker exec -it vue-app bash
# 初始化 vue 模板
npm init @vitejs/app . --template vue-ts
# 安装项目依赖
npm install
# 打包
npm run build
# 退出容器
exit
# 拷贝前端文件到 go 项目静态资源文件夹
docker cp vue-app:/app/dist ~/go/src/jassue-gin/static