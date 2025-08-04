DB=golink
pg-dump:
	pg_dump.exe -U postgres -h localhost -d $(DB) --schema-only -f "C:\laragon\www\golink\db\db.sql"
