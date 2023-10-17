MAINNET_IP=$1
PROXY_IP=$2
SUERPUSER=$3
BLOCK_NUM=$4

if [ $PROXY_IP = '' ]; then 
    ssh -t debian@$MAINNET_IP "bash /home/debian/db_copy.sh"
    scp debian@$MAINNET_IP:/tmp/db.sql /tmp/db.sql
else 
    ssh -t root@$PROXY_IP 'ssh -t debian@'$MAINNET_IP' "bash /home/debian/db_copy.sh"'
    ssh -t root@$PROXY_IP 'scp debian@'$MAINNET_IP':/tmp/db.sql /tmp/db.sql'
    scp root@$PROXY_IP:/tmp/db.sql /tmp/db.sql
fi




TDB="postgres://harshjain@localhost:5432/sample?sslmode=disable"
psql -U $SUERPUSER -d postgres -c 'drop database sample'
psql -U $SUERPUSER -d postgres -c 'create database sample'
psql -U $SUERPUSER -d sample < /tmp/db.sql
psql -U $SUERPUSER -d sample < db_scripts/local_testing/missing_table_from_download_db.sql
migrate -path ./migrations/ -database "$TDB" up


psql -U harshjain -d sample < <(cat db_scripts/local_testing/reset_to_blocknum.sql | sed "s/18246321/$BLOCK_NUM/" )
psql -U $SUERPUSER -d postgres -c 'drop database tmp_sample'
createdb -O harshjain -T sample tmp_sample

# create user sample with encrypted password '123Sample';
# GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO sample;
# ALTER DATABASE sample OWNER TO sample;
# update schema_migrations set version=27, dirty='f';
#
# SELECT format(
#           'ALTER TABLE public.%I OWNER TO sample',
#           table_name
#        )
# FROM information_schema.tables
# WHERE table_schema = 'public';

# SELECT * FROM information_schema.tables  WHERE table_schema = 'public';