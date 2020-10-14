rm ./files/certs/*.key
rm ./files/certs/*.pem

openssl req -new -nodes -x509 -out ./files/certs/server.pem -keyout ./files/certs/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Rabbit/OU=IT/CN=www.6rabbit.com/emailAddress=rabbit@outlook.com"
openssl req -new -nodes -x509 -out ./files/certs/client.pem -keyout ./files/certs/client.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Rabbit/OU=IT/CN=www.6rabbit.com/emailAddress=rabbit@outlook.com"