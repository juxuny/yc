# yc
远界云微服务开发框架

## Requirements

* protoc v3.12.4
* protoc-gen-go v1.28
* protoc-gen-go-grpc v.2

## Prepare

### Install `protoc`

1. Go plugins for the protocol compiler
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

2. Install protoc
You can get the compiled binary file at [protoc-3.12.4](https://github.com/protocolbuffers/protobuf/releases)
   
**For OSX**

```shell
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.12.4/protoc-3.12.4-osx-x86_64.zip
mkdir -p $GOPATH/bin/protoc.d/protoc-3.12.4
mv protoc-3.12.4-osx-x86_64.zip $GOPATH/bin/protoc.d
cd $GOPATH/bin/protoc.d/protoc-3.12.4/ && unzip protoc-3.12.4-osx-x86_64.zip 
cd $GOPATH/bin
if [ -f protoc ]; then rm protoc; fi 
ln -s $GOPATH/bin/protoc.d/protoc-3.12.4/bin/protoc protoc
```

**For Linux**
```shell
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.12.4/protoc-3.12.4-linux-x86_64.zip
mkdir -p $GOPATH/bin/protoc.d/protoc-3.12.4
mv protoc-3.12.4-linux-x86_64.zip $GOPATH/bin/protoc.d
cd $GOPATH/bin/protoc.d/protoc-3.12.4/ && unzip protoc-3.12.4-linux-x86_64.zip 
cd $GOPATH/bin
if [ -f protoc ]; then rm protoc; fi 
ln -s $GOPATH/bin/protoc.d/protoc-3.12.4/bin/protoc protoc
```

### Validator

* length.min
* length.max
* min
* max
* pattern
* datetime
* date 
* time
* timestamp.log
* password
* required
* in

#### 1.length.min or length.max

```text
// @v: length.min=10
// @msg: 长度最小值为: 10
List []int

// @v: length.max=10
// @msg: 长度最大值为：10
List []int
```

##### 2. min or max

@msg 后面跟着一个 go template，可以通过双重花括号获取当前的值

```text
// @v: min=10
// @msg: 最小值为10,当前值为{{.Money}}
Money float64

// @v: max=10
// @msg: 最大值为10,当前值为{{.Money}}
Money float64
```

##### 3. pattern

```text
// @v: pattern=^([\\d])$
// @msg: invalid password
Password string
```

##### 4. required

表示某个参数必传，没有 @msg 就获取系统默认错误信息，（指针不能为空）

```text
// @v: required
Pagination *dt.Pagination
```

##### 5. in

表示检查取值范围,通过逗号分割表示集合 

```text
// @v: in=1,3,4
Type int64
```

##### 6. password

密码检查，如果密码规则有特别要求，可以用pattern

