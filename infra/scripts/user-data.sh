#!/bin/bash

docker 2>/dev/null && echo "Docker is already installed" || {
	echo "Installing Docker"
	apt update
	apt install -y ca-certificates curl gnupg lsb-release

	mkdir -p /etc/apt/keyrings
	curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

	echo \
		"deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
		$(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

	apt update
	apt install -y docker-ce docker-ce-cli containerd.io
	usermod -aG docker ubuntu
}

service docker start

RUN_SCRIPT_PATH=/home/ubuntu/run.sh

cat <<- EOF > $RUN_SCRIPT_PATH
docker run -d -it -p 25565:25565 \
 	-e EULA=TRUE \
	-v /home/ubuntu/minecraft:/data \
	itzg/minecraft-server
EOF

chmod +x $RUN_SCRIPT_PATH

$RUN_SCRIPT_PATH
