* extract the two libs into your arm toolchain
``` sh
cd ~/poncat2/cross/arm-mv5sft-linux-gnueabi/sys-root/usr/
tar zxvf ncurseslib_arm.tgz
tar zxvf readlinelib_arm.tgz
```

* example build command
```
CGO_ENABLED=1 CGO_LDFLAGS="-lreadline -lncurses" CC=arm-mv5sft-linux-gnueabi-gcc GOARCH=arm GOARM=5 GOOS=linux  go build example.go 
```
