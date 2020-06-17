#!/bin/bash

# ca_cert.pem
# private/ca_key.pem
# server/server_cert.pem
# server/server_key.pem
# client/client_cert.pem
# client/client_key.pem

mkdir testca certs private
chmod 700 private
echo 01 > serial
touch index.txt

CA_CN=MyTestCA
SERVER_CN=localhost
CLIENT_CN=localhost
SERVER_PASSWORD=MySecretPassword
CLIENT_PASSWORD=MySecretPassword


echo "--->openssl req -x509 -config openssl.cnf -newkey rsa:2048 -days 365 -out ca_cert.pem -outform PEM -subj /CN=$CA_CN/ -nodes"

# ca
echo "general ca cert..."
openssl req -x509 -config openssl.cnf -newkey rsa:2048 -days 365 -out ca_cert.pem -outform PEM -subj /CN=$CA_CN/ -nodes
openssl x509 -in ca_cert.pem -out ca.cer -outform DER

# server cert
echo "general server cert..."
mkdir server
cd server
openssl genrsa -out server_key.pem 2048
openssl req -new -key server_key.pem -out req.pem -outform PEM -subj /CN=$SERVER_CN/O=server/ -nodes
cd ..
openssl ca -config openssl.cnf -in ./server/req.pem -out ./server/server_cert.pem -notext -batch -extensions server_ca_extensions
cd server
openssl pkcs12 -export -out server_cert.p12 -in server_cert.pem -inkey server_key.pem -passout pass:$SERVER_PASSWORD
cd ..

# client cert
echo "general client cert..."
mkdir client
cd client
openssl genrsa -out client_key.pem 2048
openssl req -new -key client_key.pem -out req.pem -outform PEM -subj /CN=$CLIENT_CN/O=client/ -nodes
cd ..
openssl ca -config openssl.cnf -in ./client/req.pem -out ./client/client_cert.pem -notext -batch -extensions client_ca_extensions
cd client
openssl pkcs12 -export -out client_cert.p12 -in client_cert.pem -inkey client_key.pem -passout pass:$CLIENT_PASSWORD
