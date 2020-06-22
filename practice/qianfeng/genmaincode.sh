#!/bin/bash
#生成代码11
echo $1
mkdir $1
cd $1
touch $1.go
echo "package main" >> $1.go 
echo "import \"fmt\"" >> $1.go
echo >> $1.go
echo "func main() {" >> $1.go
echo >> $1.go
echo "}" >> $1.go