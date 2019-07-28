# WSL(Windows Subsystem for Linux)安装与使用

官方介绍：WSL(Windows Subsystem for Linux)是适用于Linux 的 Windows 子系统，是一个为在Windows 10和Windows Server 2019上能够原生运行Linux二进制可执行文件的兼容层。
直白的讲，就是**借助WSL可以在Windows10中直接编译出Linux环境的可执行文件，让Windows下的开发有机会摆脱庞大的虚拟机**，有点mac下开发的感觉，特别适合我这种没有mac的贫困码农。

# 打开子系统功能

在win10 power shell输入如下指令：

```shell
Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux
```

# win10 安装Ubuntu

从win10应用商店安装Ubuntu：
![](file://D:/MyBlogs/post-images/1564114438785.jpg)
安装后win10桌面会提示设置用户名和密码：
![](file://D:/MyBlogs/post-images/1564123550388.png)

# GO环境配置

## GOROOT配置

### 配置共享目录
为了方便的在Windows和WSL之间拷贝文件，要先设置共享目录。
WSL中的`/mnt/c/Users/HideOnBush/ `路径为Windows10在Linux下的挂载分区，为方便和win10共享文件，在Windows中新建UbuntuShare文件夹，在WSL添加软链接：
```shell
ln -s /mnt/c/Users/HideOnBush/UbuntuShare/ win10
```
> HideOnBush是我的Windows中的用户名，需要替换为你实际的Windows用户名哦。
### WSL安装Go
从[go语言中文网](https://studygolang.com/dl)下载Linux go安装包(我使用的是1.12.5)，放到刚才设置的共享目录中，就可以在WSL中操作了。
解压到/usr/local：

```shell
cd win10
sudo tar -xzf go1.12.5.linux-amd64.tar.gz -C /usr/local/
```

## GOPATH配置

将Windows和Linux的GOPATH设置为同一个路径，就可以实现在win10写代码时和在Linux编代码共用一套第三方包。
### 查看win10环境变量
![](file://D:/MyBlogs/post-images/1564115020270.png)
![](file://D:/MyBlogs/post-images/1564115125234.png)
可以看到我的Windows中GOPATH路径为`Users/HideOnBush/go`，接下来将WSL中的GOPATH设置为同一路径，完成后续操作就可以方便的在Windows中编写代码并调用Ubuntu编译代码了。
### 设置wsl GOPATH：

```go
vim ~/.bashrc
```

末尾添加：

```shell
export GOPATH=/mnt/c/Users/HideOnBush/go
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH
```

使修改立即生效：

```shell
source .bashrc
```

使用`go env`查看GOPATH是否配置成功：

![](file://D:/MyBlogs/post-images/1564116197193.png)
可以看到我的GOPATH已经设置为和Windows中的GOPATH一致了。

# 更新软件源

```shell
sudo cp /etc/apt/sources.list /etc/apt/sources.list_origin
sudo vim /etc/apt/sources.list
```

```shell
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic main restricted universe multiverse
deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-updates main restricted universe multiverse
deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-backports main restricted universe multiverse
deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-backports main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-security main restricted universe multiverse
deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-security main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-proposed main restricted universe multiverse
deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-proposed main restricted universe multiverse
```

```shell
sudo apt-get update //更新源
sudo apt-get upgrade //根据软件源的信息更新软件
```

# 在Windows中调用WSL bash
接下来根据你使用的编译器，将编译器使用的bash修改为WSL内的bash，就可以在Windows下编译出Linux下的可执行文件。

## 配合goland 使用

### 修改bash为wsl所在路径

找到wsl命令行所在路径，
![](file://D:/MyBlogs/post-images/1564123762337.jpg)

设置goland "shell path",
![](file://D:/MyBlogs/post-images/1564123785522.jpg)
修改goland换行符为UINX下的LF：
![](file://D:/MyBlogs/post-images/1564123833128.jpg)

terminal测试，
![](file://D:/MyBlogs/post-images/1564123793665.jpg)
> 不能直接使用goland的run图标运行，因为goland默认使用的是Windows中的编译环境，必须要在terminal中编译才是使用的WSL的bash编译。

## 配合vs code使用

安装wsl后在vscode terminal可以看到wsl（windows system linux）的窗口，就可以在这儿方便操作Linux。
![](file://D:/MyBlogs/post-images/1564123717689.jpg)

为了防止和Linux内编写的文件换行符有冲突，推荐修改换行符为Linux样式，如下：
![](file://D:/MyBlogs/post-images/1564123732107.jpg)

至此，Windows10的Linux子系统的开发环境已经基本搭建完成，后续根据你的需要在WSL安装对应工具即可，如果觉得WSL的命令行太难用，可以使用Goland的terminal，或者使用Secure CRT等SSH工具。