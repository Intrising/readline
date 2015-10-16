## marvell poncat arch
* extract the two libs into your arm toolchain
``` sh
cd ~/poncat2/cross/arm-mv5sft-linux-gnueabi/sys-root/usr/
tar zxvf ncurseslib.tgz
tar zxvf readlinelib.tgz
```

* example build command
```
CGO_ENABLED=1 CGO_LDFLAGS="-lreadline -lncurses" CC=arm-mv5sft-linux-gnueabi-gcc GOARCH=arm GOARM=5 GOOS=linux  go build example.go 
```

## openwrt armv7 arch


* extract the two libs into your arm toolchain
``` sh
cd ~/openwrt/OpenWrt-Toolchain-mvebu_gcc-4.8-linaro_musl-1.1.10_eabi.Linux-x86_64/toolchain-arm_cortex-a9+vfpv3_gcc-4.8-linaro_musl-1.1.10_eabi/
tar zxvf ncurseslib.tgz
tar zxvf readlinelib.tgz
```

* example build command
```
CGO_ENABLED=1 CGO_LDFLAGS="-lreadline -lncurses" CC=arm-openwrt-linux-gcc GOARCH=arm GOARM=7 GOOS=linux  go build example.go 
```


