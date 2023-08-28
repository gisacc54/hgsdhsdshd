##
docker run -d --network=mynetwork --name web3:multistage web3
docker run -p 8080:8080 -d -v ./:/app --network=sis_network web3
docker run -d --network=mynetwork --name mysql-container -v web3-database:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root_password -e MYSQL_DATABASE=mydb mysql:latest