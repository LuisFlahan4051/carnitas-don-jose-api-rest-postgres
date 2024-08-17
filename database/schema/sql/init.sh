HOST="localhost"
USER="root"
DATABASE="carnitas_db"

DIRECTORIO="./"

for archivo in "$DIRECTORIO"/*.sql
do
  if [ -f "$archivo" ]; then
    echo "Ejecutando $archivo..."
    psql -h "$HOST" -U "$USER" -d "$DATABASE" -f "$archivo"
  else
    echo "No se encontraron archivos .sql en $DIRECTORIO"
  fi
done