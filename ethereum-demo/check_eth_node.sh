#!/bin/sh
tail_result=`tail -n 1 /var/log/geth.log`
FAILPATTERN="Synchronisation failed"
if [ $? -ne 0 ]
then
    echo "tail exec failed!" $?
else
    if [[ $tail_result == *$FAILPATTERN* ]]
    then
        echo "Synchronisation failed, restart geth now!"
        /usr/local/bin/supervisorctl restart geth
        if [ $? -ne 0 ]
        then
        echo "restart failed." $?
        fi
    fi
fi
echo "excute time:" `date`