# Installation
There multiple ways in which you can install `huntsman` on your machine or a target machine.  
1. Install it using golang compiler using `go install` or `go build`
	
	1. [Install Golang](https://golang.org/doc/install)    
	2. `git clone git@github.com:souvikhaldar/huntsman.git`
	3. `cd huntsman`  
	4. `go install`  
2. Download and copy the binary from `/download` directory on $PATH.  
3. Use the `goinstaller.py` script.  
	```
	./goinstaller.py --help 
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

# Transfer to a target
Once you've compiled huntsman for the target OS and arch, you can transfer it 
using `scp` or any tool of choice:  
Eg, transfering linux binary to target machine:  
`scp ./download/linux_amd64 username@address:location`

# Fast concurrent port scanning  
``` 
huntsman portScan --help
Concurrently scan the provided range (by default 0 to 65535) to check if any port is open

Usage:
  huntsman portScan [flags]

Flags:
  -e, --end int32       last port number (default 65535)
  -h, --help            help for portScan
  -s, --start int32     starting port number (default 1)
      --target string   IP/URL address of the machine to be scanned
  -t, --threads int32   the number of goroutines to execute at a time (default 100)
```

Example:  
`huntsman portScan --target abc.com`  

# Location details of IP or tcpdump  

`huntsman ipinfo --ip=12.121.212.32`  

[![asciicast](https://asciinema.org/a/342086.svg)](https://asciinema.org/a/342086)

    ## Providing tcpdump log:  
        Parsing location info from log file of `tcpdump`
        ```
        huntsman iploc --tcp-dump server.log
        Request came from:  fwdproxy-frc-003.fbsv.net
        Details of the requester:  {"status":"success","country":"United States","countryCode":"US","region":"NJ","regionName":"New Jersey","city":"Newark","zip":"07175","lat":40.7357,"lon":-74.1724,"timezone":"America/New_York","isp":"Facebook, Inc.","org":"Facebook, Inc.","as":"AS32934 Facebook, Inc.","query":"173.252.127.3"}. 
        ```
  

        P.S- The log file was populated by:  
        ```
        tcpdump -s 0 -A 'tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x47455420' > server.log
        ```
        
        The above command was run on the server where HTTP service was running on port 80.

# Run a TCP proxy

`huntsman proxy -s <local-port> -t <target-address> -p <target-port>`   

# Run a TCP listener

`huntsman listen --port=<port>`


# Reverse shell
First you need to compile the binary for the target machine using the 
`goinstaller.py` or anything of choice. Then preferably use `scp` to transfer
the binary to the target machine (see `Installation` section) then execute it
using `./<binary-name> reverseshell --port <port-number>`. Now the listener is
running to which you will be sending instructions to execute.   

We will be using [netcat](http://netcat.sourceforge.net/) as the client for 
sending the commands over the network.  
`nc <address-of-target> <port-number>`  
[Youtube link for the video demonstration](https://youtu.be/eE0k0GVZXyc)
