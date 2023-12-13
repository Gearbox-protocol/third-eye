alter table pools add _version integer;
update table pools set _version=1;