set -e 

PARENT_DIR=$(dirname $0)
REMOTE_DB=$1
SUPERUSER=$2

FORK_BLOCK=`jq .forkBlock.number < <(curl https://anvil.gearbox.foundation/api/forks/432945bc-3620-11ee-be56-0242ac120002  )`

# if [ "$PROXY_IP" = '' ]; then 
#     ssh -t debian@$MAINNET_IP "bash /home/debian/db_copy.sh"
#     scp debian@$MAINNET_IP:/tmp/db.sql /tmp/db.sql
# else 
#     ssh -t root@$PROXY_IP 'ssh -t debian@'$MAINNET_IP' "bash /home/debian/db_copy.sh"'
#     ssh -t root@$PROXY_IP 'scp debian@'$MAINNET_IP':/tmp/db.sql /tmp/db.sql'
#     scp root@$PROXY_IP:/tmp/db.sql /tmp/db.sql
# fi

# if [ "$SUPERUSER" = "postgres" ]; then
#     sudo su postgres
# fi


export SAMPLE_DB="postgres://$SUPERUSER@localhost:5432/sample?sslmode=disable"

set +e
psql -U $SUPERUSER -d postgres -c 'drop database sample'
psql -U $SUPERUSER -d postgres -c 'create database sample'
pg_dump "$REMOTE_DB" | psql  -U $SUPERUSER -d sample
set -e 

# psql -U $SUPERUSER -d sample < db_scripts/local_testing/missing_table_from_download_db.sql
psql -U $SUPERUSER -d sample < $PARENT_DIR/../../migrations/000016_rankings.up.sql
migrate -path $PARENT_DIR/../../migrations/ -database "$SAMPLE_DB" up


psql -U $SUPERUSER -d sample < <(cat $PARENT_DIR/reset_to_blocknum.sql | sed "s/18246321/$FORK_BLOCK/" )
set +e
psql -U $SUPERUSER -d postgres -c 'drop database tmp_sample'
set -e
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