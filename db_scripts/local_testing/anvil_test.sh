set -e 

PARENT_DIR=$(dirname $0)
REMOTE_DB=$1
SUPERUSER=$2
FORK_URL="$3"
FINAL_DB="$4"

FORK_BLOCK=`jq .forkBlock.number < <(curl  $FORK_URL )`


TMP_DB="tmp_$FINAL_DB"
export TMP_DB_URL="postgres://$SUPERUSER@localhost:5432/$TMP_DB?sslmode=disable"

set +e
psql -U $SUPERUSER -d postgres -c " SELECT  pg_terminate_backend(pid) FROM  pg_stat_activity WHERE  pid <> pg_backend_pid() AND datname = '$TMP_DB';"
psql -U $SUPERUSER -d postgres -c "drop database $TMP_DB"
psql -U $SUPERUSER -d postgres -c "create database $TMP_DB"
pg_dump --no-owner "$REMOTE_DB" | psql  -U $SUPERUSER -d $TMP_DB
set -e 
set +e
pg_dump "$REMOTE_DB" --table public.schema_migrations  | psql  -U $SUPERUSER -d $TMP_DB
set -e

# psql -U $SUPERUSER -d sample < db_scripts/local_testing/missing_table_from_download_db.sql
psql -U $SUPERUSER -d $TMP_DB < $PARENT_DIR/../../migrations/000016_rankings.up.sql
migrate -path $PARENT_DIR/../../migrations/ -database "$TMP_DB_URL" up


psql -U $SUPERUSER -d $TMP_DB < <(cat $PARENT_DIR/reset_to_blocknum.sql | sed "s/18246321/$FORK_BLOCK/" )
set +e
psql -U $SUPERUSER -d postgres -c "drop database $FINAL_DB"
set -e

# reset mf
PWD=`pwd`
LOCAL_DB="host=localhost user=$SUPERUSER  dbname=$TMP_DB"
cd /home/debian/$FINAL_DB-third-eye
go run "scripts/08_merged_pf_version_reset/main.go" "$LOCAL_DB" $FORK_BLOCK
cd $PWD

createdb -O $SUPERUSER -T $TMP_DB $FINAL_DB



# create user debian with encrypted password '123Sample';
# GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO debian;
# ALTER user debian createdb;
# ALTER DATABASE sample OWNER TO debian;
# update schema_migrations set version=27, dirty='f';
#
# SELECT format(
#           'ALTER TABLE public.%I OWNER TO debian',
#           table_name
#        )
# FROM information_schema.tables
# WHERE table_schema = 'public';

# SELECT * FROM information_schema.tables  WHERE table_schema = 'public';