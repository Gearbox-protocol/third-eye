# DB location/configs

Start/stop based on: https://stackoverflow.com/questions/7975556/how-can-i-start-postgresql-server-on-mac-os-x

Depends on how we installed postgres, via homebrew or direct dmg package. 
If homebrew try:
```
/opt/homebrew/opt/postgresql/bin/postgres -D /opt/homebrew/var/postgres
```

If via homebrew these commands are available:
```
brew services info postgres
brew services start postgres
brew services stop postgres
```

If dmg package:
```
/Library/PostgreSQL/13/bin/pg_ctl -D /Library/PostgreSQL/13/data/ stop
```



## Error

>>> there is no unique or exclusion constraint matching the ON CONFLICT
Reason:
If the primaryKey constraint is missing from table schema and the gorm data structure tags have the primarykey then the expected constraint are not found in db.

## Tracking last update on table

```
track_commit = on # update postgresql.conf
```