Start docker
docker run -d -p 8080:8080 --env-file ./.env \
    -v /root/server_ssl:/server_ssl \
    -e SERVER_SSL=/server_ssl   \
    id997/theheadhunter-server:0.1.4

Build docker
docker build . --tag theheadhunter-server

tag
docker tag theheadhunter-server id997/theheadhunter-server:0.1.4

Push docker
docker push id997/theheadhunter-server:0.1.4

# Backend task



## Tasks
✅ DONE ❌ NOT DONE 👷 WIP
1. ✅ Go through the frontend task and implement its server side REST api.
2. ✅ Use the Go programming language.
3. ✅ Use whatever serialization (or in memory) solution you like to store data.
4. ✅ There is no need for any authentication related functionality.
5. ✅ Implement validation as described in the frontend task.
6. ✅ Please send Github linkith your solution. (We accept half-finished solutions as well.)

### For extra points
1. ✅ Use a postgres database to save data.
2. 👷 Write tests.
3. ✅ Dockerize your app.
4. ✅ Install your server somewhere (and share its URL with us).
5. ✅ Create a README in your repo and add notes bout installation, implementation details or any important info about your solution.