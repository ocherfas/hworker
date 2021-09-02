mkdir -p /usr/local/hworker/bin
cp ./out/hworker /usr/local/hworker/bin/
cp -n config.yaml /usr/local/hworker/config.yaml
ln -s /usr/local/hworker/bin/hworker /usr/local/bin/hworker