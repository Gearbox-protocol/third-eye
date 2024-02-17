SUPERUSER="$1"
psql -U $SUPERUSER -d postgres -c 'drop database sample'
psql -U $SUPERUSER -d postgres -c "create database sample with owner $SUPERUSER"
pg_dump "$DCDB" | psql  -U $SUPERUSER -d sample
psql -U $SUPERUSER -d sample < migrations/000016_rankings.up.sql
FORK_BLOCK=`jq .forkBlock.number < <(curl https://anvil.gearbox.foundation/api/forks/$2 )`
psql -U $SUPERUSER -d sample < <(cat db_scripts/local_testing/reset_to_blocknum.sql | sed "s/18246321/$FORK_BLOCK/" ) 