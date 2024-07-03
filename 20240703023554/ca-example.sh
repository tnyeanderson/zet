#!/bin/bash

# Create the certs using subshell for easy cd-ing
(
	subj='/C=US/ST=State/L=City/O=Corp/OU=IT'

	mkdir -p certs
	cd certs || exit

	# Create a self-signed root CA
	openssl req -newkey rsa:2048 -noenc -keyout ca.key -x509 -out ca.crt -subj "$subj/CN=root"

	# Create an intermediate CA signed by the root CA
	openssl req -newkey rsa:2048 -noenc -keyout intermediate.key -x509 -CA ca.crt -CAkey ca.key -out intermediate.crt -subj "$subj/CN=intermediate"

	# Create an app cert signed by the intermediate CA
	openssl req -newkey rsa:2048 -noenc -keyout app.key -x509 -CA intermediate.crt -CAkey intermediate.key -out app.crt -subj "$subj/CN=app.lcl"

	# Create a pem file with the app and intermediate certs
	cat app.crt intermediate.crt >app.pem
)

# Create a Caddyfile config using non-default ports
cat >Caddyfile <<EOF
{
  http_port 8880
  https_port 4443
}

* {
	tls ./certs/app.pem ./certs/app.key
}
EOF

# Start caddy
caddy run --config ./Caddyfile &>/dev/null &
pid=$!

# Verify the chain
echo | openssl s_client -verifyCAfile certs/ca.crt -connect localhost:4443 -servername app.lcl

# Stop caddy
kill "$pid"
