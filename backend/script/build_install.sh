#!/usr/bin/env bash

cd ..
make build

mkdir -p ~/app/kratos_monolithic_demo

mkdir -p ~/app/kratos_monolithic_demo/admin/service/bin/

mkdir -p ~/app/kratos_monolithic_demo/admin/service/configs/

mv -f ./app/admin/service/bin/server ~/app/kratos_monolithic_demo/admin/service/bin/server

cp -rf ./app/admin/service/configs/*.yaml ~/app/kratos_monolithic_demo/admin/service/configs/

cd ~/app/kratos_monolithic_demo/admin/service/bin/
pm2 start --namespace kratos_monolithic_demo --name admin server -- -conf ../configs/

pm2 save

pm2 restart kratos_monolithic_demo
