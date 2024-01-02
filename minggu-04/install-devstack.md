# Melakukan instalasi DevStack pada Ubuntu

## Langkah-langkah Instalasi Devstack

### Menambah User Stack

1. Melakukan instalasi ubuntu di vmware. Bisa menggunakan VirtualBox atau sejenisnya. Kemudian lakukan update sistem terlebih dahulu.

```
$ sudo apt update && sudo apt upgrade -y
```

2. Membuat user baru dengan nama `stack`. (Langkah ini merupakan langkah opsional)

```
$ sudo useradd -s /bin/bash -d /opt/stack -m stack
```

3. Menambahkan hak akses _execute_ pada direktori `/opt/stack`

```
$ sudo chmod +x /opt/stack
```

4. Mengaktifkan `sudo privileges` terhadap user `stack` tanpa password

```
$ echo "stack ALL=(ALL) NOPASSWD: ALL" | sudo tee /etc/sudoers.d/stack
```

5. Login dengan user `stack`

```
$ sudo -u stack -i
```

6. Instalasi `git` dan `pip`

Hal ini dikarenakan dalam penggunaan Devstack nantinya akan membutuhkan setidaknya 2 plugin ini. git dan python.

```
$ sudo apt install git python3-pip -y
```

### Mengunduh DevStack

1. Kloning repository devstack menggunakan git

```
$ git clone https://opendev.org/openstack/devstack
```

### Membuat file _local.conf_

1. Masuk ke folder devstack dan buat file _local.conf_

```
$ cd devstack
$ cat > local.conf
```

2. Tambah konfigurasi berikut:

```
ADMIN_PASSWORD=secret
DATABASE_PASSWORD=$ADMIN_PASSWORD
RABBIT_PASSWORD=$ADMIN_PASSWORD
SERVICE_PASSWORD=$ADMIN_PASSWORD
```

### instalasi Devstack

Lakukan instalasi devstack dengan perintah berikut:

```
$ ./stack.sh
```

Devstack telah berhasil di-install di _local storage_
