#!/bin/bash

CONFIG_FILE="./sysbench.conf"
TABLES=32
TABLE_SIZE=1000000
THREADS=(32 64 128)
LOAD_TYPES=("oltp_point_select" "oltp_read_only" "oltp_read_write")

# sysbench --config-file=sysbench.conf oltp_point_select --tables=32 --table-size=10000 --threads=128 prepare
function do_sysbench() {
    echo "do sysbench $1, thread=$2"
    sysbench $1 run --config-file=$CONFIG_FILE --tables=$TABLES --table-size=$TABLE_SIZE --threads=$2 --report-interval=10 --mysql-ignore-errors=1062,2013,8028,9007
}

for load in ${LOAD_TYPES[*]}
do
    for thread in ${THREADS[*]}
    do
        do_sysbench $load $thread
    done
done

