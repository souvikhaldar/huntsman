# Huntsman
A swiss army knife for all things security.

# Complete guide
[1. Installation](https://github.com/souvikhaldar/huntsman/blob/master/guide.md#installation)  
[2. Reverse shell](https://github.com/souvikhaldar/huntsman/blob/master/guide.md#reverse-shell)

# Things you can do with huntsman
2. 
3. 
4. Run a TCP listener. `huntsman listen --port=<port>`  

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
- [x] Reverse Shell listener  
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

