SCRIPT=`realpath $0`
SCRIPTPATH=`dirname $SCRIPT`
rm -rf $SCRIPTPATH/../out
mkdir $SCRIPTPATH/../out
go build -o $SCRIPTPATH/../out/hworker $SCRIPTPATH/..