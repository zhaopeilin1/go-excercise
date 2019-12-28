#!bash
echo $1
mkdir $1

touch $1.go
touch $1_test.go

echo "package $1" >> $1.go

echo "package $1" >> $1_test.go
echo >> $1_test.go
echo "import (" >> $1_test.go
echo "\"testing\"" >> $1_test.go
echo ")" >> $1_test.go

mv $1.go $1_test.go $1
