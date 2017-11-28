# unit 配置使用以及 go的环境配置步骤

    1: git clone https://github.com/nginx/unit
    2: 在unit/unit目录下执行 make 以及make install ，输出
    install -d sbin
    install -p build/unitd sbin/
    install -d build
    install -d /Users/a123/Grapes/src/unit
    install -p -m644 ./src/go/unit/* /Users/a123/Grapes/src/unit/
    CGO_CFLAGS="-I/Users/a123/Desktop/unit/unit/build -I/Users/a123/Desktop/unit/unit/src" \
		CGO_LDFLAGS="-L/Users/a123/Desktop/unit/unit/build " \
		GOPATH=/Users/a123/Grapes \
		go install -v unit
    3:开发自己的基于unit包的程序，并在unit配置文件写入程序启动路径
