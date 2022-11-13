Start docker
docker run -it -p 8080:8080 --env-file ./.env \
    -v /root/server_ssl:/server_ssl \
    -e SERVER_SSL=/server_ssl   \
    id997/theheadhunter-server:0.1.2

Build docker
docker build . --tag theheadhunter-server

tag
docker tag theheadhunter-server id997/theheadhunter-server:0.1.1

Push docker
docker push id997/theheadhunter-server:0.1.1