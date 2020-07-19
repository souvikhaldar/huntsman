#!/bin/sh
./goinstaller.py --target ./download --os popular --arch amd64
scp download/linux_amd64 souvik@192.168.0.104:~
	
