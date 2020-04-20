## DATABASE ACCESS

docker exec -it [container_name] psql -U [postgres_user]

psql -U hppoc -v dbname=$DATABASE -v user=$USER -v passwd=$PASSWD -f /scripts/createpg.sql;
psql postgres -v dbname=$DATABASE -v user=$USER -v passwd=$PASSWD -f ./updatepg.sql ;;


CREATE UNIQUE INDEX CONCURRENTLY payments_trxnref
ON payments (trxnref);

ALTER TABLE payments 
ADD CONSTRAINT unique_equip_id 
UNIQUE USING INDEX equipment_equip_id;