#!/bin/bash

HOST=10.202.0.39
PORT=2881
USER=root@test
PASSWORD=''
DB=test


function exec_sql() {
    echo "exec sql: $1"
    mysql -u $USER -h $HOST -P $PORT --password=$PASSWORD -e "$1"
}

for i in {1..32}
do
  echo $i;
  exec_sql "alter table $DB.sbtest$i drop index k_$i;"
  exec_sql "alter table $DB.sbtest$i add index k_$i(k) local;"
done
