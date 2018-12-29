#!/bin/bash
#
# Imports data previously exported into CSV files into our database.
# See env.sh for connection info
#

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DATA_DIR="../data" # This can't be an absolute path, at least on Windows

if [[ ! -f ${DIR}/env.sh ]] ; then
    echo "./env.sh not found.  Please create it to continue"
    exit 1
fi
. ${DIR}/env.sh


cd ${DIR}
PGPASSWORD=${PASSWORD} psql -h ${HOST} -U ${USER} -v schema=${SCHEMA} -f ${DATA_DIR}/ddl/postgres.sql ${DB}
PGPASSWORD=${PASSWORD} psql -h ${HOST} -U ${USER} -c "\copy ${SCHEMA}.game from '${DATA_DIR}/nes/games.csv' with csv header encoding 'utf-8'" ${DB}
PGPASSWORD=${PASSWORD} psql -h ${HOST} -U ${USER} -c "\copy ${SCHEMA}.publisher from '${DATA_DIR}/nes/publishers.csv' with csv header encoding 'utf-8'" ${DB}
PGPASSWORD=${PASSWORD} psql -h ${HOST} -U ${USER} -c "\copy ${SCHEMA}.game_x_publisher from '${DATA_DIR}/nes/game_x_publisher.csv' with csv header encoding 'utf-8'" ${DB}
PGPASSWORD=${PASSWORD} psql -h ${HOST} -U ${USER} -v schema=${SCHEMA} -f ${DATA_DIR}/ddl/postgres-fix-sequences.sql ${DB}
