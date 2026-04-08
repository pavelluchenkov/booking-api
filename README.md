### Выполнение миграции
1. **docker cp migrations/0002_create_tables_table.up.sql booking-api-postgres-1:/tmp/migration.sql**
2. **docker exec -it booking-api-postgres-1 psql -U booking_user -d booking_db -f /tmp/migration.sql**
3. **Проверка что таблица создана: docker exec -it booking-api-postgres-1 psql -U booking_user -d booking_db -c "\dt"**