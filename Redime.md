##
docker run -d --network=mynetwork --name web3:multistage web3
docker run -p 8080:8080 -d -v ./:/app --network=mynetwork web3
docker run -d --network=mynetwork --name mysql2-container  -e MYSQL_ROOT_PASSWORD=root_password -e MYSQL_DATABASE=mydb mysql:latest



91f9ef54b891235ebcfe08dc9ca3811cb6a648dee802b3ab5d6eba030cef4d06
91f9ef54b891235ebcfe08dc9ca3811cb6a648dee802b3ab5d6eba030cef4d06