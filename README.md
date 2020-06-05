# Huntsman
A swiss army knife for all things security.

# Things you can do with huntsman
1. Fast concurrent port scanning `huntsman scan --target=12.23.121.32`  
2. Location details of IP or `tcpdump`. `huntsman iploc 12.121.212.32`  

# Installation 
1. [Install Golang](https://golang.org/doc/install)    
2. `git clone git@github.com:souvikhaldar/huntsman.git`
3. `cd huntsman`  
4. `go install`

# Example
1. Parsing location info from log file  
```
huntsman iploc --tcpdump --file=server.log
Body:  {"status":"success","country":"United States","countryCode":"US","region":"NJ","regionName":"New Jersey","city":"Newark","zip":"07175","lat":40.7357,"lon":-74.1724,"timezone":"America/New_York","isp":"Facebook, Inc.","org":"Facebook, Inc.","as":"AS32934 Facebook, Inc.","query":"173.252.127.3"}
```
