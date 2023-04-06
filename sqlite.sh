#!/bin/bash

if [ -f server.db ]; then
    echo "The server.db file already exists"
else
	sqlite3 server.db < ./sql/quotes_table.sql
	echo "The sqlite.db file was created"
fi
