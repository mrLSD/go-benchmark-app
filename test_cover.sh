rm -rf *.coverprofile
X=1
for Package in $(go list ./... | grep -v "vendor") ; do
	echo $Package ;
	go test -v -covermode=count -coverprofile=$X.coverprofile $Package ;
	X=$((X+1)) ;
done
