#Installation
There multiple ways in which you can install `huntsman` on your machine or a target machine.  
1. Install it using golang compiler using `go install` or `go build`
2. Copy the binary from `/download` directory on $PATH.  
3. Use the `goinstaller.py` script.  
	```
	./goinstaller.py --help        ✔  4453  13:50:39
	Install go program for multiple OS and multiple architectures
	Run goinstaller.py --help for all options
	usage: goinstaller.py [-h]
			      [--os {all,popular,linux,darwin,windows,dragonfly,android,freebsd,netbsd,openbsd,plan9,solaris,aixjs}]
			      [--arch {all,popular,amd64,386,arm,ppc64,arm64,ppc64le,mips,mipsle,mips64,mips64le,s390x}]
			      [--source SOURCE] [--target TARGET]

	optional arguments:
	  -h, --help            show this help message and exit
	  --os {all,popular,linux,darwin,windows,dragonfly,android,freebsd,netbsd,openbsd,plan9,solaris,aixjs}
				The target OS. Eg. all,linux,darwin,windows,etc
	  --arch {all,popular,amd64,386,arm,ppc64,arm64,ppc64le,mips,mipsle,mips64,mips64le,s390x}
				The target's architecture. Eg. all,amd64,386,etc
	  --source SOURCE       The directory where source source is present
	  --target TARGET       The target dir where the binary needs to be stored
	```
	Eg. Compiling for **popular** OSes like Windows, Microsoft and Linux for 64-bit architecture can be done using
	`./goinstaller.py --target ./download --os popular --arch amd64`
