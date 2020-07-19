# Huntsman
A swiss army knife for all things security.

# Complete guide
[1. Installation](https://github.com/souvikhaldar/huntsman/blob/master/guide.md#installation)
[2. Reverse shell](https://github.com/souvikhaldar/huntsman/blob/master/guide.md#reverse-shell)

# Things you can do with huntsman
1. Fast concurrent port scanning  
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

2. Location details of IP or `tcpdump`. `huntsman ipinfo --ip=12.121.212.32`  
    [![asciicast](https://asciinema.org/a/342086.svg)](https://asciinema.org/a/342086)
    ### Providing tcpdump log:  
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

3. Run a TCP proxy. `huntsman proxy -s <local-port> -t <target-address> -p <target-port>`  
4. Run a TCP listener. `huntsman listen --port=<port>`  

# Installation 
1. [Install Golang](https://golang.org/doc/install)    
2. `git clone git@github.com:souvikhaldar/huntsman.git`
3. `cd huntsman`  
4. `go install`  

**OR**

Download the binary of your choice from the `download` directory and run it directly. 
Shoutout to [go installer](https://github.com/souvikhaldar/scripts/blob/master/goinstaller.py)

# Static code analysis of Huntsman 
1. Run sonar-qube `docker run -d --name sonarqube -p 9000:9000 sonarqube`  
2. Install `gosec`- `go get github.com/securego/gosec/cmd/gosec`  
3. `./sonarGosec.sh`  



# Path
- [x] Port scanner    
- [x] IP information  
- [x] TCP proxy    
- [x] TCP listener  
- [x] HTTP server  
- [ ] Exploiting DNS  
- [ ] VirusTotal API functionality  
- [ ] SMB and NTLM expotation  
- [ ] Abusing Databases  
- [ ] Packet processing  
- [ ] Fuzzing and shellcode  
- [ ] Cryptography  
- [ ] Windows system analysis  
- [ ] steganography  
- [ ] CNC RAT  

# Inspiration
The inspiration of this tool are primarily the following two sources:
1. Pursuing [Advanced Exercutive Program in Cyber Security and Cyber Defense](https://talentsprint.com/pages/wip/iit-kanpur/v2.5/index.html) at the esteemed [c3i](https://security.cse.iitk.ac.in/) institution of [IIT Kanpur](https://www.iitk.ac.in/) and Professor [Sandeep Shukla](https://www.cse.iitk.ac.in/users/sandeeps/).  
2. The excellent book [Black Hat Go: Go Programming For Hackers and Pentesters](https://www.amazon.in/Black-Hat-Go-Programming-Pentesters-ebook/dp/B073NPY29N) by Tom Steele, Chris Patten and Dan Kottmann. 

