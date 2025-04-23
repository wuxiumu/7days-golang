windows环境搭建
amd64位系统，安装go语言环境 msi文件安装
arm64位系统，安装go语言环境

linus环境搭建
amd64位系统，安装go语言环境 tar.gz文件安装
arm64位系统，安装go语言环境  

mac环境搭建
amd64位系统，安装go语言环境  intel芯片的go语言环境
arm64位系统，安装go语言环境  m1芯片的go语言环境


goland安装
goland是jetbrains出品的go语言IDE，可以方便的编写go语言代码，安装goland后，可以直接在goland中编写go语言代码，并编译运行。  
下载地址：https://www.jetbrains.com/go/
vscode安装
vscode是微软推出的开源编辑器，可以方便的编写go语言代码，安装vscode后，可以直接在vscode中编写go语言代码，并编译运行。  
下载地址：https://code.visualstudio.com/docs/languages/go
安装go语言环境
下载go语言安装包，根据系统类型选择安装包，下载后解压到指定目录，然后配置环境变量。  
配置环境变量：  
在windows系统中，配置环境变量的方法是：  
右键“我的电脑”->“属性”->“高级系统设置”->“环境变量”->“系统变量”->“Path”->“编辑”->“新建”->将go语言安装目录添加到Path中。  
在mac系统中，配置环境变量的方法是：  
打开终端，输入以下命令：  
```
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
source ~/.bash_profile
```
在linux系统中，配置环境变量的方法是：  
打开终端，输入以下命令：  
```
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```
安装成功后，在终端输入go version命令，如果能看到go版本号，则表示安装成功。

