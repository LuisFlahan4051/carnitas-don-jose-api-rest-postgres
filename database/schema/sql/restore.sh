HOST="localhost"
USER="root"
DATABASE="carnitas_db"

DIRECTORIO="./"

psql -U "$USER" -d "postgres" -c "DROP DATABASE IF EXISTS $DATABASE;"
psql -U "$USER" -d "postgres" -c "CREATE DATABASE $DATABASE;"
