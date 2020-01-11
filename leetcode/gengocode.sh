#!/bin/bash
#生成代码11
echo $1
mkdir $1
cd $1
touch $1.go $1_test.go
echo "package $1" >> $1.go 
echo "package $1" >> $1_test.go
echo  >> $1_test.go
echo "import \"testing\"" >> $1_test.go
echo  >> $1_test.go
echo "func TestAll(t *testing.T) {" >>  $1_test.go
echo  >> $1_test.go
echo "}" >> $1_test.go