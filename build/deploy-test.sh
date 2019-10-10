#!/bin/sh

BASE=$(dirname $0)/..

$BASE/build/build.sh

if [ $? -ne 0 ]
then
    echo "build failed"
    exit -1
fi

REMOTE="47.56.198.84"
project=gts

echo "============================="
echo "==== deploy to remote server"
echo "============================="
scp dist/gts root@$REMOTE:/data/gts/gts.new

ssh -tt root@$REMOTE <<EOF
cd /data/${project}
rm -f ${project}
mv ${project}.new ${project}
chmod +x ${project}
ls -al
pm2 reload ${project}
exit
EOF

echo "done"
