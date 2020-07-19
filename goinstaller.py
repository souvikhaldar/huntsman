#!/usr/local/bin/python3
import os,argparse
print("Install go program for multiple OS and multiple architectures")
print("Run goinstaller.py --help for all options")
parser = argparse.ArgumentParser()
parser.add_argument("--os",
help="The target OS. Eg. all,linux,darwin,windows,etc",
choices=["all","popular","linux","darwin","windows","dragonfly","android","freebsd",
         "netbsd","openbsd","plan9","solaris","aix""js"],
)

parser.add_argument("--arch",
help="The target's architecture. Eg. all,amd64,386,etc",
choices=["all","popular","amd64","386","arm","ppc64","arm64","ppc64le", "mips", "mipsle",
"mips64", "mips64le", "s390x"],
)
parser.add_argument("--source",
help="The directory where source source is present",
default=".",
)
parser.add_argument("--target",
help="The target dir where the binary needs to be stored",
default=".",
)

args = parser.parse_args()
if args.os is None or args.arch is None:
    print("You need to set --os and --arch")
    os.exit()
print("OS: ",args.os)
print("Arch: ",args.arch)
print("Source: ",args.source)
print("Target: ",args.target)
# source for the envs- https://golang.org/doc/install/source#environment
envs = {}
envs["aix"]=["ppc64"]
envs["android"]= ["386","amd64","arm","arm64"]
envs["darwin"]=["386","amd64","arm", "arm64"]
envs["dragonfly"]=["amd64"]
envs["freebsd"]=["386","amd64","arm", "amd64"]
envs["js"]=["wasm"]
envs["linux"]=["386","amd64", "arm", "arm64", "ppc64", "ppc64le", "mips", "mipsle", "mips64", "mips64le", "s390x"]
envs["netbsd"]=["386","amd64", "arm"]
envs["openbsd"]=["386","amd64", "arm", "arm64"]
envs["plan9"]=["386","amd64","arm"]
envs["solaris"]=["amd64"]
envs["windows"]=["386","amd64"]
os_opt = []
arch_opt = []
if args.os == "all":
    os_opt = ["aix","android","darwin","dragonfly","freebsd","js","linux",
            "netbsd","openbsd","plan9","solaris","windows"]
elif args.os == "popular":
    os_opt = ["linux","windows","darwin"]
else:
     os_opt.append(args.os)

if args.arch == "all":
    arch_opt = ['ppc64','386', 'amd64', 'arm', 'arm64', 'ppc64', 'ppc64le',
            'mips', 'mipsle', 'mips64', 'mips64le', 's390x']
elif args.arch == "popular":
    arch_opt = ['386','amd64']
else:
    arch_opt.append(args.arch)

print("OS options: ",os_opt)
print("Arch options: ",arch_opt)

for goos,arch in envs.items():
    if goos in os_opt:
        for val in arch:
            if val not in arch_opt:
                continue
            if goos == "windows":
                cmd = "cd {};env GOOS={} GOARCH={} go build -o {}/{}_{}.exe".format(
                args.source,goos,val,args.target,goos,val)
            else:
                cmd = "cd {};env GOOS={} GOARCH={} go build -o {}/{}_{}".format(
                args.source,goos,val,args.target,goos,val)
            if os.system(cmd) != 0:
                print("Execution failed")
            else:
                print("Generated: {}_{}".format(goos,val))
