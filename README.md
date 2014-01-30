# Barbowa - Barberton Open Web APIs
Instructions for acessing the APIs will be available shortly, for now we are documenting how we set this up, so any other communities interested in doing the same can see what we did (which is not much yet eaither, but it's a work in progress).

## Installation
We started with a bare-bones server installation of Ubuntu 12.04 LTS. Then we compiled Go version 1.2 from source, following C2Journal's [Go 1.1 Ubuntu Installation Instructions][c2journal]. We did this instead of just using the package because apt only has version 1 available, and that doesn't work with a lot of new examples and frameworks.

To set the environment, source the .go file (you could probably set these variables in /etc/bash.bashrc or someplace else, but since we are not pulling a standard package this helps document what we are doing and makes it easy to reproduce):
	
	. .go

When you do not define those variables globally then you have problems using `sudo go get PACKAGE_NAME`, so we made a script which we can sudo instead (it sources the .go file):
	
	sudo pkg-inst.sh github.com/dougblack/sleepy

Then all you have to do is run the the rest.go file and connect up to "http://server:8080/contacts":
	
	go run rest.go
	
[c2journal]: http://c2journal.com/2013/07/28/installing-go-1-1-on-ubuntu/

