#!/bin/sh
# script for static code analysis using gosec and sonar-qube
# followed from- 
# https://levelup.gitconnected.com/static-code-analysis-for-golang-5f24b555d227
# generate coverage file

# first you need to run the sonar-qube docker
# docker run -d --name sonarqube -p 9000:9000 sonarqube
# then install gosec- go get github.com/securego/gosec/cmd/gosec

go test -short -coverprofile=./cov.out ./...

# generate gosec report in sonarqube format
gosec -fmt=sonarqube -out report.json ./...    

# run sonar-scanner
sonar-scanner

# then visit http://localhost:9000/projects and find the project
