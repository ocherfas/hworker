SCRIPT=`realpath $0`
SCRIPTPATH=`dirname $SCRIPT`

mkdir -p /opt/hworker/bin
cp $SCRIPTPATH/../out/hworker /opt/hworker/bin/
cp -n $SCRIPTPATH/../config.yaml /opt/hworker/config.yaml
ln -s /opt/hworker/bin/hworker /usr/bin/hworker