🌐[Ana Sayfa](./README.md)

---

# Git Release Manager v1.0.0-dev.1 Sürüm Notları - 2024-12-10

## Özet:

-   ✨ Yeni Özellikler & Geliştirmeler: 7 commit

-   🐞 Çözümlenen Problemler: 1 commit

-   📚 Dökümantasyon: 2 commit

-   🔧 Kod Tabanı Bakımı ve Güncellemeleri: 5 commit

ve 2 ❤️ katılımcı

## Değişiklikler

### ✨ Yeni Özellikler & Geliştirmeler:

-   <a id="commit-bf437f1"></a>**cmake** [#bf437f1](https://github.com/ahmettsoner/pars/commit/bf437f142c44e7140917d66242993cfd28aafd13): kurulum/kaldırma scriptleri ile choco paket desteği ekleyin (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

    Sistem ortam yollarını yönetmek ve Program Ekle/Kaldır desteği sağlamak için chocolateyInstall.ps1 ve chocolateyUninstall.ps1 scriptleri eklendi. Bu geliştirme, kullanılabilirliği iyileştirir ve sistem paket yönetimi standartlarına uyumu sağlar.

-   <a id="commit-7c974b7"></a>**cmake** [#7c974b7](https://github.com/ahmettsoner/pars/commit/7c974b7a0e560a6b184e84d4009031494d7ec9c0): sürüm bilgisi üzerinde pre-release tespiti ekleyin (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

    Sürüm dizesinin pre-release belirteçlerini (ör., alpha, beta, rc) inceleyerek CMake'de pre-release versiyonları algılama mantığı uygulanmıştır. Bu özellik, build süreci boyunca pre-release durumunu otomatik olarak belirleyerek versiyon yönetimini geliştirir.

-   <a id="commit-81b604c"></a>**cmake** [#81b604c](https://github.com/ahmettsoner/pars/commit/81b604c76f22b2e28c2a7e5dc11886efcae2c1a2): Go derlemelerine 386 mimari desteği ekleyin (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

    Go build sürecinde 386 mimarı desteği etkinleştirildi.

-   <a id="commit-17628e3"></a>**choco** [#17628e3](https://github.com/ahmettsoner/pars/commit/17628e3fbc9d8cee7f4ed5d88f64751c2ebc5a28): Add/Remove Programs ve PATH yönetimi ile paketi iyileştirin (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

    Chocolatey paket kurulumu sırasında Program Ekle/Kaldır girişi için destek eklendi.

    Kurulum sonrasında otomatik olarak uygulamayı PATH ortam değişkenine ekler.

    Uygulama ayrıntılarını Windows Registry'den kaldırmak ve PATH ortam değişkenini temizlemek için kaldırma süreci geliştirilmiştir.

-   <a id="commit-3b6244b"></a>**installer** [#3b6244b](https://github.com/ahmettsoner/pars/commit/3b6244bfaeb3248b46f283ca1bb27f43ce4e3c87): MSI konfigürasyonunda görünen uygulama adını geliştirin (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-6137359"></a>**cmake** [#6137359](https://github.com/ahmettsoner/pars/commit/61373597d297c05baf8f8f868348495ec888db0f): Chocolatey paketleme desteği ekleyin (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-da170e1"></a>**cmake** [#da170e1](https://github.com/ahmettsoner/pars/commit/da170e1b49fc666b25eee7c11599dbb328872fc9): Chocolatey paketleme desteği ekleyin (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

### 🐞 Çözümüleri Çözülenler:

-   <a id="commit-51acc69"></a>**cmake** [#51acc69](https://github.com/ahmettsoner/pars/commit/51acc69ea1d463b5a8dd8da6f6631611320be30c): MSI yükleyici konfigürasyonundaki sorunları çözün (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

### 📚 Dökümantasyon:

-   <a id="commit-a2bf7fe"></a>[#a2bf7fe](https://github.com/ahmettsoner/pars/commit/a2bf7fe6a401fbbe88c0662af42d6e911be77b26): kapsamlı README.md ekleyin (by [@Ahmet Soner](mailto:ahmettsoner@gmail.com)) (Issue: [#385](https://github.com/ahmettsoner/pars/issues/385))

    Proje genel bakışı, anahtar özellikler, başlama rehberi, yükleme talimatları, katkıda bulunma kuralları, dokümantasyon linkleri, lisans bilgileri ve topluluk kaynakları bölümlerini içeren detaylı bir `README.md` dosyası Pars deposuna eklendi.

-   <a id="commit-a6b7695"></a>[#a6b7695](https://github.com/ahmettsoner/pars/commit/a6b7695f2fd0ad04cdc3b70e24e9c072ec0c6b4c): kapsamlı README.md ekleyin (#385) (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) (Issue: [#385](https://github.com/ahmettsoner/pars/issues/385))

    Proje genel bakışı, anahtar özellikler, başlama rehberi, yükleme talimatları, katkıda bulunma kuralları, dokümantasyon linkleri, lisans bilgileri ve topluluk kaynakları bölümlerini içeren detaylı bir README.md dosyası Pars deposuna eklendi.

### 🔧 Kod Tabanı Bakımı ve Güncellemeleri:

-   <a id="commit-92ac517"></a>**vscode** [#92ac517](https://github.com/ahmettsoner/pars/commit/92ac517401e6f31156d605392a872834980eceaa): source.organizeImports'ı ayarlarda hasır belirtin (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

    İçe aktarma ifadelerinin, proje kuralarına göre açıkça organize edilmesini sağlar.

-   <a id="commit-2010a74"></a>**project** [#2010a74](https://github.com/ahmettsoner/pars/commit/2010a74dda273e2b3db1a3ac1bbe9858cfbcdbb1): launch, settings ve task yapılandırmalarıyla birlikte .vscode klasörü ekleyin (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-7857c8f"></a>[#7857c8f](https://github.com/ahmettsoner/pars/commit/7857c8f6d2d6f09f4503e40620b68d4ead79e9a2): kullanılmayan dosyaları kaldırın (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-02ae3da"></a>[#02ae3da](https://github.com/ahmettsoner/pars/commit/02ae3da3d79d97525806d935c832e2b4ba23a443): çakışmaları çözme (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-80ea61f"></a>[#80ea61f](https://github.com/ahmettsoner/pars/commit/80ea61fac776a2fcddfb087f07b78f79ceded35d): repository for Semantic Release için depoyu başlatın (by [@Ahmet Soner](mailto:ahmettsoner@gmail.com))

## Kullanışlı Linkler

-   📜 [Tüm Değişiklikler](https://github.com/ahmettsoner/pars/blob/main/CHANGELOG.md)

## Katılımcılar:

-   [CI/CD Bot](mailto:ci-bot@parsdevkit.net) (Kod, Konfigürasyonlar, Dokumantasyon, Testler)

-   [Ahmet Soner](mailto:ahmettsoner@gmail.com) (Dokumantasyon, Konfigürasyonlar, Kod, Testler)
