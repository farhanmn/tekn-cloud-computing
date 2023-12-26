# Tugas 12 - Learn Kubernetes Basics

## Membuat Cluster Minikube

![start](1-start.png)

```
$ minikube start
```

## Membuka Dashboard

Ada 2 cara untuk membuka dashboard, langsung membuka browser atau dengan mendapatkan URL terlebih dahulu. Kemudian copy URL tersebut di browser.

![dashboard](2-dashboard.png)

```
// cara pertama
$ minikube dashboard

// cara kedua
$ minikube dashboard --url
```

## Melakukan Deployment Aplikasi

1. Menggunakan `kubectl create`

![create](3-create.png)

2. Melihat deployment

![get](4-get.png)

3. Melihat pod

![pods](5-getpods.png)

4. Melihat cluster events

![cluster](6-cluster.png)

5. Melihat konfigurasi `kubectl`

![config](7-config.png)

6. Melihat log aplikasi

![log](8-logs.png)

## Membuat Servis

1. Expose pods ke publik menggunakan `kubectl expose`

![expose](9-expose.png)

2. Melihat servis yang telah dibuat

![service](10-servis.png)

3. Menjalankan servis dan akan membuka browser

![service2](11-run.png)
![browser](12-browser.png)

## Mengaktifkan Addons

1. Melihat addons yang mendukung

![addons](13-addons.png)

2. Mengaktifkan addons metric-server

![metric](14-metric.png)

3. Melihat pods dan servis

![pods](15-pods.png)

4. Cek output dari metric-server

![metric](16-toppods.png)

5. Menonaktifkan addons metric-server

![disable](17-disable.png)

## Pembersihan (Clean up)

1. Membersihkan resource

![cleaning](18-cleanup.png)

2. Menghentikan minikube

![stop](19-stop.png)
