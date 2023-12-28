# Git Kolaborasi

### Fork Repo

- Melakukan fork terhadap repo yang ada
  ![fork](gambar-04-fork.png)

- Membuat fork
  ![createfork](gambar-04-create-fork.png)

### Clone

![clone](gambar-04-clone.png)

### Konfigurasi Upstream

![upstream](gambar-04-upstream.png)

```
// Melakukan upstream ke repo utama
$ git remote add upstream https://github.com/farhanmn/hello-world
```

## Mengirimkan Pull Request

### Membuat Perubahan di Repo Lokal

- Melakukan sinkronisasi repo lokal dengan repo dari upstream author
  ![sync](gambar-04-sync.png)

- Melakukan perubahan data
  ![change](gambar-04-change.png)

```
// membuat branch baru dengan nama add-data
$ git checkout -b add-data

// menampilkan isi dari file README.md
$ cat README.md

// edit isi README.md menggunakan nano
$ nano README.md

// cek perubahan
$ git status
```

- Push
  ![push](gambar-04-push.png)
