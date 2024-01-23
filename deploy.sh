export $( grep -vE "^(#.*|\s*)$" .env )

sshpass -p ${ssh_pass} ssh pi@10.0.0.251 << EOF
    cd repositories/brain-dump
    git pull
    docker compose up -d --force-recreate
    docker container ls
EOF