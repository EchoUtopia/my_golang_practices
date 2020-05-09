# make ca:

openssl genrsa -out rootCA.key 2048;

openssl req -x509 -new -nodes -key rootCA.key -days 1024 -out rootCA.pem;


# make server cert
openssl genrsa -out server.key 2048

openssl req -new -key server.key -out server.csr

openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 500


# make client cert


openssl genrsa -out client.key 2048

openssl req -new -key client.key -out client.csr

openssl x509 -req -in client.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out client.crt -days 500