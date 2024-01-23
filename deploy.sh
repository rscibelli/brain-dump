export $( grep -vE "^(#.*|\s*)$" .env )

sshpass -p ${ssh_pass} ssh pi@10.0.0.251 << EOF
    cd repositories/brain-dump
    git pull
    cd /backend
    env GOOS=linux GOARCH=arm GOARM=7 go build
    docker compose up -d --force-restart
    docker container ls
EOF