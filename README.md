
命令终端示例

## 编译
./mybuild.sh

（注：如果有依赖的包，手动拷贝到vendor对应目录，后续完善）

## 使用  
config.yaml：配置文件  

## 临时记录
工具形式如何？以易阅读的方式，见字知意  
如：
a.exe misc test1
a.exe misc test2
or
a.exe misc --mode test1
a.exe misc --mode test2
(这种形式也可以)

优化子命令接口 完成

## 工程说明
db：将不同的三级命令整合到一起处理（非官方那样使用cmd结构体），但添加了flag，并使用官方的帮助信息显示功能。  
test：与上类似，自行提示信息  

## 知识点
根命令、二级命令、三级命令，参数，标志。  

## 使用指南
### 常用
mybuild.sh为编译脚本，改target名称为真实程序名。version.h为版本号，代码和编译脚本已实现好，实际中改版本号即可。  
工程名为 cmdtool，需要将包引用的路径的 cmdtool 改为真实工程名。
共用模块在pkg中。  
常量、底层函数在common中。   
main.go不需要改。

### 根命令
cmd/rootCmd.go 为根命令入口，可在该文件 rootCmd  定义处实现 Run 函数，此为一级命令。  
在 rootCmd.go 中引入二级命令目录时，将包改名，再在该文件 Execute 函数中调用 AddCommand添加即可。  

### 二级命令
cmd/db 等为二级命令目录，里面的文件不限制，包名统一为cmd，对外的接口统一为 RegisterCmd，针对不同命令，则需修改 name shortDescription 等描述，其中 name 即为子命令名称。  
一般而言，二级命令的 cmd.go 为命令入口文件，busy.go为业务实现文件，可根据实际添加其它文件，该目录可独立其它子模块，也可引用其子模块。   
固定文件名和对外接口，主要是为了方便代码复用。  

## 官方子命令嵌套

官方示例的第三级命令可认为是单独的子命令，必须按套路编码，有点麻烦，先舍弃不用。  

```
先定义根命令：
var rootCmd = &cobra.Command {

}

再定义二级子命令：
var createCmd = &cobra.Command{
}
在二级子命令包的init函数中添加到根命令：
rootCmd.AddCommand(createCmd)

再定义三级子命令
var fileCmd = &cobra.Command{
}
在三级子命令的init函数添加到二级子命令：
createCmd.AddCommand(fileCmd)

可再定义三级子命令，方法类似
var dirCmd = &cobra.Command{
}
var fileCmd = &cobra.dirCmd{
}

最终在三级子命令中实现具体业务。执行类似：
./foo create file
./foo create dir

在三级命令后，可再加参数，如
./foo create file fff

还可加标志

./foo create file fff -p /tmp

等。
如此一样可以将功能相近的子命令整合到一起使用。  

```
