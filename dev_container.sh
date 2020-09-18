docker build -f ./debugger.Dockerfile -t rabbitoj:latest .
docker run -e Role=Server -e Secret=hzytql -p 8888:8888 -p 8090:8090 -d rabbitoj:latest
