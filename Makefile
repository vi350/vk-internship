include .env
# Targets

## build: Build the executable file
.PHONY: build
build:
	go build -o botbinary.out .

# Setting default make target to build.
.DEFAULT_GOAL := build

# End of Makefile
