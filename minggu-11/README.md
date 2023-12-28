# Tugas 11 - Application Containerization and Microservice Orchestration

## Setup

![setup](setup.png)

```
// clone repository dari url github
git clone https://github.com/ibnesayeed/linkextractor.git

// move ke direktori linkextractor
cd linkextractor

// move ke branch demo
git checkout demo
```

## Step 0: Basic Link Extractor Script

![step0](step0.png)

```
// move ke branch step0
$ git checkout step0

// show struktur folder
$ tree
```

![step0-1](step0-1.png)

```
// tampilkan isi file linkextractor.py
cat linkextractor.py

// run file linkextractor.py
./linkextractor.py http://example.com/
```

![step0-2](step0-2.png)

```
// melihat hak akses file linkextractor.py
ls -l linkextractor.py

// run file linkextractor.py
python3 linkextractor.py
```

## Step 1: Containerized Link Extractor Script

![step1](step1.png)

```
// move ke branch step1
git checkout step1

// show struktur folder
tree
```

![step1-1](step1-1.png)
![step1-2](step1-2.png)
![step1-3](step1-3.png)

```
// show isi file Dockerfile
cat Dockerfile

// create/build image
docker image build -t linkextractor:step1 .

// show daftar image
docker image ls

// run container docker
docker container run -it --rm linkextractor:step1 http://example.com/
```

![step1-4](step1-4.png)

```
// run container
docker container run -it --rm linkextractor:step1 https://training.play-with-docker.com/
```

## Step 2: Link Extractor Module with Full URI and Anchor Text

![step2](step2.png)

```
// move ke branch step2
git checkout step2

// show struktur folder
tree
```

![step2-1](step2-1.png)

```
// show isi file linkextractor.py
cat linkextractor.py
```

![step2-2](step2-2.png)

```
// create/build image
docker image build -t linkextractor:step2 .
```

![step2-3](step2-3.png)

```
// menampilkan image
docker image ls
```

![step2-4](step2-4.png)

```
// run container docker
docker container run -it --rm linkextractor:step2 https://training.play-with-docker.com/
```

![step2-5](step2-5.png)

```
// run container docker
docker container run -it --rm linkextractor:step1 https://training.play-with-docker.com/
```

## Step 3: Link Extractor API Service

![step3](step3.png)

```
// move ke branch step3
git checkout step3

// menampilkan struktur folder
tree
```

![step3-1](step3-1.png)

```
// show isi file Dockerfile
cat Dockerfile
```

![step3-2](step3-2.png)

```
// show isi file main.py
cat main.py
```

![step3-3](step3-3.png)

```
// create/build docker
docker image build -t linkextractor:step3 .
```

![step3-4](step3-4.png)

```
// run container
docker container run -d -p 5000:5000 --name=linkextractor linkextractor:step3

// show daftar container
docker container ls

// create request HTTP
curl -i http://localhost:5000/api/http://example.com/
```

![step3-5](step3-5.png)

```
// show log container
docker container logs linkextractor

// remove container
docker container rm -f linkextractor
```

## Step 4: Link Extractor API and Web Front End Services

![step4](step4.png)

```
// move ke branch step4
git checkout step4

// menampilkan struktur folder
tree
```

![step4-1](step4-1.png)

```
// show file docker-compose.yml
cat docker-compose.yml
```

![step4-2](step4-2.png)

```
// show file www/index.php
cat www/index.php
```

![step4-3](step4-3.png)

```
// run docker compose
docker-compose up -d --build
```

![step4-4](step4-4.png)

```
// show daftar container
docker container ls

// menghubungkan dengan layanan API
curl -i http://localhost:5000/api/http://example.com/

// modifikasi index file link extractor
sed -i 's/Link Extractor/Super Link Extractor/g' www/index.php

// reset perubahan
git reset --hard

// menghentikan layanan
docker-compose down
```

## Step 5: Redis Service for Caching

![step5](step5.png)

```
// move ke branch step5
$ git checkout step5

// show struktur folder
$ tree

// show file www/Dockerfile
$ cat www/Dockerfile
```

![step5-1](step5-1.png)

```
// show file api/main.py
cat api/main.py
```

![step5-2](step5-2.png)

```
// show file docker-compose.yml
cat docker-compose.yml
```

![step5-3](step5-3.png)

```
// build layanan
docker-compose up -d --build
```

![step5-4](step5-4.png)

```
// client redis
docker-compose exec redis redis-cli monitor

// modifikasi index file link extractor
sed -i 's/Link Extractor/Super Link Extractor/g' www/index.php

// reset perubahan
git reset --hard

// stop service
docker-compose down
```

## Step 6: Swap Python API Service with Ruby

![step6](step6.png)

```
// move ke branch step6
$ git checkout step6

// show struktur folder
$ tree
```

![step6-1](step6-1.png)

```
// show file api/linkextractor.rb
cat api/linkextractor.rb
```

![step6-2](step6-2.png)

```
// show file api/Dockerfile
cat api/Dockerfile
```

![step6-3](step6-3.png)

```
// show file docker-compose.yml
cat docker-compose.yml
```

![step6-4](step6-4.png)

```
// run service
docker-compose up -d --build
```

![step6-5](step6-5.png)

```
// menghubungkan service
curl -i http://localhost:4567/api/http://example.com/
```

![step6-6](step6-6.png)

```
// stop service
docker-compose down

// show log
cat logs/extraction.log
```
