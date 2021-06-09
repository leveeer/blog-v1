#!/bin/bash
# 删除mysql中所有表
# 示例：
# Usage: ./script user password dbname
# Usage: ./script user password dbname server-ip
# Usage: ./script user password dbname mysql.nixcraft.in
# ---------------------------------------------------

MUSER="$1"
MPASS="$2"
MDB="$3"

MHOST="localhost"

[ "$4" != "" ] && MHOST="$4"

# 设置命令路径
MYSQL=$(which mysql)
AWK=$(which awk)
GREP=$(which grep)

# help
if [ ! $# -ge 3 ]
then
 echo "Usage: $0 {MySQL-User-Name} {MySQL-User-Password} {MySQL-Database-Name} [host-name]"
 echo "Drops all tables from a MySQL"
 exit 1
fi

# 连接mysql数据库
$MYSQL -u "$MUSER" -p"$MPASS" -h "$MHOST" -e "use $MDB"  &>/dev/null
if [ $? -ne 0 ]
then
 echo "Error - 用户名或密码无效，无法连接mysql数据库"
 exit 2
fi

TABLES=$($MYSQL -u "$MUSER" -p"$MPASS" -h "$MHOST" "$MDB" -e 'show tables' | $AWK '{ print $1}' | $GREP -v '^Tables' )

# make sure tables exits
if [ "$TABLES" == "" ]
then
 echo "Error - 在数据库 $MDB 中未发现相关表"
 exit 3
fi

# let us do it
for t in $TABLES
do
 echo "Truncate $t table from $MDB database..."
 $MYSQL -u "$MUSER" -p"$MPASS" -h "$MHOST" "$MDB" -e "TRUNCATE TABLE  $t"
done