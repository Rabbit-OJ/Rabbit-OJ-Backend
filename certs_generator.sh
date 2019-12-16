rm ./certs/*.key
rm ./certs/*.pem

openssl req -new -nodes -x509 -out certs/server.pem -keyout certs/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Rabbit/OU=IT/CN=www.6rabbit.com/emailAddress=rabbit@outlook.com"
openssl req -new -nodes -x509 -out certs/client.pem -keyout certs/client.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Rabbit/OU=IT/CN=www.6rabbit.com/emailAddress=rabbit@outlook.com"