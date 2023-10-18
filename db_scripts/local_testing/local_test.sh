set -e 

MAINNET_IP=$1
PROXY_IP=$2
SUPERUSER=$3

FORK_BLOCK=`jq .forkBlock < <(curl https://anvil.gearbox.foundation/forks/432945bc-3620-11ee-be56-0242ac120002  )`

if [ "$PROXY_IP" = '' ]; then 
    ssh -t debian@$MAINNET_IP "bash /home/debian/db_copy.sh"
    scp debian@$MAINNET_IP:/tmp/db.sql /tmp/db.sql
else 
    ssh -t root@$PROXY_IP 'ssh -t debian@'$MAINNET_IP' "bash /home/debian/db_copy.sh"'
    ssh -t root@$PROXY_IP 'scp debian@'$MAINNET_IP':/tmp/db.sql /tmp/db.sql'
    scp root@$PROXY_IP:/tmp/db.sql /tmp/db.sql
fi

if [ "$SUPERUSER" = "postgres" ]; then
    sudo su postgres
fi


if [ "$SUPERUSER" =  "debian" ]; then
    export TDB="postgres://$SUPERUSER:123Sample@localhost:5432/sample?sslmode=disable"
else 
    export TDB="postgres://$SUPERUSER@localhost:5432/sample?sslmode=disable"
fi

set +e
psql -U $SUPERUSER -d postgres -c 'drop database sample'
psql -U $SUPERUSER -d postgres -c 'create database sample'
psql -U $SUPERUSER -d sample < /tmp/db.sql
set -e 

# psql -U $SUPERUSER -d sample < db_scripts/local_testing/missing_table_from_download_db.sql
psql -U $SUPERUSER -d sample < migrations/000016_rankings.up.sql
migrate -path ./migrations/ -database "$TDB" up


psql -U $SUPERUSER -d sample < <(cat db_scripts/local_testing/reset_to_blocknum.sql | sed "s/18246321/$FORK_BLOCK/" )
set -e
psql -U $SUPERUSER -d postgres -c 'drop database tmp_sample'
set +e
createdb -O $SUPERUSER -T sample tmp_sample

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