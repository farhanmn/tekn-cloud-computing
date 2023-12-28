# Rangkuman Cloud Computing

### Pengertian

_Cloud Computing_ (Komputasi Awan) merupakan sumber daya yang tersedia berdasarkan kebutuhan/ permintaan _(On-demand)_, terutama dalam hal penyimpanan data (cloud storage) dan daya komputasi.

Komputasi awan biasanya bergantung pada pembagian sumber daya dan menggunakan model bayar sesuai pemakaian _(pay-as-you-go)_. Hal itu berfungsi untuk membantu mengurangi pengeluaran modal yang tidak diperlukan.

### Karakteristik _Cloud Computing_

- _On-demand self-service_
  Seorang konsumen memiliki kemampuan untuk secara mandiri menyediakan sumber daya komputasi, seperti waktu server dan penyimpanan jaringan, sesuai dengan kebutuhan mereka secara otomatis, tanpa memerlukan interaksi manusia pada setiap penyedia layanan.

- _Broad network access_
  Kemampuan ini dapat diakses melalui jaringan dan menggunakan mekanisme standar yang mendukung penggunaan oleh platform klien tipis atau tebal yang beragam, seperti ponsel, tablet, laptop, dan workstation.

- _Resource pooling_
  Sumber daya komputasi dari penyedia dikonsolidasikan untuk melayani sejumlah konsumen menggunakan model multi-tenant, di mana sumber daya fisik dan virtual secara dinamis dialokasikan dan dipindahkan sesuai dengan permintaan konsumen.

- _Rapid elasticity_
  Kemampuan ini dapat disediakan dan dilepaskan secara elastis, bahkan secara otomatis dalam beberapa kasus, untuk dengan cepat mengakomodasi lonjakan permintaan baik ke luar maupun ke dalam. Bagi konsumen, ketersediaan sumber daya seringkali tampak tanpa batas dan dapat disesuaikan sesuai kebutuhan, dalam jumlah berapapun, kapanpun.

- _Measured service_
  Sistem cloud secara otomatis mengendalikan dan mengoptimalkan penggunaan sumber daya dengan memanfaatkan kemampuan pengukuran pada berbagai tingkat abstraksi sesuai dengan jenis layanan, seperti penyimpanan, pemrosesan, bandwidth, dan akun pengguna aktif. Penggunaan sumber daya dapat dipantau, dikontrol, dan dilaporkan, memberikan tingkat transparansi yang tinggi bagi penyedia layanan dan konsumen yang menggunakan layanan tersebut.

### Tantangan dan keterbatasan

Salah satu tantangan utama dalam komputasi awan, dibandingkan dengan komputasi on-premise yang lebih tradisional, adalah keamanan dan privasi data. Pengguna cloud menyerahkan data sensitif mereka kepada penyedia layanan pihak ketiga, yang mungkin tidak memiliki langkah-langkah yang memadai untuk melindunginya dari akses, pelanggaran, atau kebocoran yang tidak sah. Tantangan lainnya adalah kurangnya visibilitas dan kontrol bagi pengguna cloud. Mereka mungkin tidak sepenuhnya tahu bagaimana sumber daya cloud mereka dikelola atau dikonfigurasi oleh penyedia layanan, dan kemampuan mereka untuk menyesuaikan layanan sesuai kebutuhan atau preferensi terbatas.

Pemahaman teknologi cloud yang kompleks menjadi esensial, meskipun pemahaman penuh mungkin tidak selalu tercapai, terutama mengingat skala, kompleksitas, dan opasitas yang disengaja dari sistem kontemporer. Metafora cloud sering kali menciptakan konsep sesuatu yang tidak dapat dipahami dengan jelas atau diterangkan.

Migrasi cloud juga merupakan tantangan signifikan. Ini melibatkan pemindahan data, aplikasi, atau beban kerja dari satu lingkungan cloud ke lingkungan lain atau dari on-premise ke cloud. Proses migrasi dapat menjadi rumit, memakan waktu, dan mahal, terutama jika terdapat ketidakcocokan antara platform atau arsitektur cloud yang berbeda. Selain itu, migrasi cloud berpotensi menyebabkan waktu henti, penurunan kinerja, atau kehilangan data jika tidak direncanakan dan dijalankan dengan cermat.

### Model layanan

- Infrastruktur sebagai layanan (IaaS)
  Model layanan cloud computing yang menyediakan akses virtual terhadap infrastruktur IT fundamental, seperti komputasi, penyimpanan, dan jaringan. Dalam layanan ini, penyedia cloud mengelola dan menyediakan sumber daya fisik seperti server, ruang penyimpanan, dan konektivitas jaringan, yang dapat disewakan oleh pengguna untuk membangun dan menjalankan aplikasi mereka. IAAS memberikan fleksibilitas dan skalabilitas, memungkinkan pengguna untuk menyesuaikan sumber daya sesuai kebutuhan mereka tanpa harus mengelola secara langsung infrastruktur fisiknya.

- Platform sebagai layanan (PaaS)
  Model layanan cloud computing yang menyediakan platform pengembangan dan pengelolaan aplikasi. Dalam PaaS, penyedia cloud menyediakan infrastruktur dan layanan tingkat lebih tinggi, seperti database, middleware, dan alat pengembangan, sehingga pengembang dapat fokus pada pembuatan aplikasi tanpa harus peduli tentang detail infrastruktur.

- Perangkat lunak sebagai layanan (SaaS)
  Model layanan cloud yang memberikan akses kepada pengguna akhir ke perangkat lunak dan aplikasi melalui internet. Dalam SaaS, perangkat lunak di-host dan dioperasikan oleh penyedia cloud, dan pengguna mengaksesnya melalui antarmuka web. Contoh SaaS termasuk aplikasi produktivitas seperti Google Workspace atau perangkat lunak manajemen keuangan online.

- "Backend" seluler sebagai layanan (MBaaS)
  Layanan cloud yang menyediakan backend dan fungsionalitas server untuk aplikasi mobile. Ini memungkinkan pengembang fokus pada pengembangan frontend dan fungsionalitas aplikasi, sementara infrastruktur backend, seperti manajemen pengguna, penyimpanan data, dan notifikasi push, diatasi oleh penyedia MBaaS.

- Komputasi tanpa server atau Function-as-a-Service (FaaS)
  Layanan cloud yang memungkinkan pengembang menjalankan fungsi atau potongan kecil kode tanpa harus mengelola infrastruktur server secara langsung. Fungsi ini dijalankan secara otomatis sebagai respons terhadap peristiwa tertentu, dan pengguna hanya membayar berdasarkan penggunaan sebenarnya. FaaS sangat cocok untuk penanganan tugas-tugas spesifik dan skala kebutuhan komputasi secara otomatis.

### Arsitektur

Arsitektur cloud, yaitu arsitektur sistem dari sistem perangkat lunak yang terlibat dalam pengiriman komputasi awan, biasanya melibatkan beberapa komponen awan yang berkomunikasi satu sama lain melalui mekanisme kopling longgar seperti antrian pesan. Ketentuan elastis menyiratkan kecerdasan dalam penggunaan kopling yang ketat atau longgar, seperti yang diterapkan pada mekanisme semacam ini dan lainnya.

Rekayasa awan adalah penerapan disiplin ilmu rekayasa komputasi awan. Ini membawa pendekatan sistematis untuk kekhawatiran tingkat tinggi dari komersialisasi, standarisasi, dan tata kelola dalam memahami, mengembangkan, mengoperasikan, dan memelihara sistem komputasi awan. Ini merupakan metode multidisiplin yang melibatkan kontribusi dari berbagai bidang, seperti sistem, perangkat lunak, web, kinerja, teknik teknologi informasi, keamanan, platform, risiko, dan rekayasa kualitas.

### Keamanan dan Privasi

Komputasi awan menimbulkan berbagai masalah terkait privasi, terutama karena penyedia layanan memiliki kemampuan untuk mengakses data yang disimpan di cloud kapan saja. Pengaksesan ini dapat terjadi secara tidak sengaja atau disengaja, yang berpotensi mengakibatkan perubahan atau penghapusan informasi yang vital. Bukan hal yang jarang, banyak penyedia cloud memiliki kebijakan yang memungkinkan berbagi informasi dengan pihak ketiga tanpa surat perintah, jika diperlukan untuk tujuan hukum dan ketertiban.

Solusi untuk mengatasi masalah privasi ini melibatkan implementasi kebijakan dan undang-undang yang kuat, serta memberikan opsi kepada pengguna akhir terkait penyimpanan data mereka. Pengguna memiliki opsi untuk mengenkripsi data yang diproses atau disimpan di dalam cloud, sebagai langkah preventif untuk mencegah akses yang tidak sah. Di samping itu, sistem manajemen identitas dapat memberikan solusi praktis dengan membedakan antara pengguna yang berwenang dan yang tidak sah. Sistem ini menetapkan jumlah data yang dapat diakses oleh setiap entitas, beroperasi melalui pembuatan dan deskripsi identitas, pencatatan aktivitas, serta penghapusan identitas yang tidak lagi digunakan. Dengan pendekatan ini, berbagai aspek keamanan dan privasi dapat diatasi secara lebih terstruktur dalam konteks komputasi awan.
