mysql -uroot -proot --local-infile=1 sample -e "LOAD DATA LOCAL INFILE '/docker-entrypoint-initdb.d/data.csv' INTO TABLE alphabet FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n' (name)"

mysql -uroot -proot --local-infile=1 sample -e "LOAD DATA LOCAL INFILE '/docker-entrypoint-initdb.d/task.csv' INTO TABLE task_status_history FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n'"

mysql -uroot -proot --local-infile=1 sample -e "LOAD DATA LOCAL INFILE '/docker-entrypoint-initdb.d/user_company.csv' INTO TABLE user_company FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n' (paying, city, company_name)"
