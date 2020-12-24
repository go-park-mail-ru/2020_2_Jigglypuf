#!/usr/bin/env bash

chmod 600 ./id_key
ssh -i ./id_key $SSH_USER@$SSH_HOST << EOF
if [ -f ~/.bash_exports ]; then
    . ~/.bash_exports
fi
cd 2020_2_Jigglypuf
git checkout main
git pull origin
sudo docker-compose down
cd build/dockerfiles
./run_containers.sh
sudo docker image prune -f
cd ../../
sudo docker-compose up -d
EOF