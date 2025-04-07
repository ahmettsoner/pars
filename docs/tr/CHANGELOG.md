# Git Sürüm Yöneticisi v1.0.0-dev.1 Sürüm Notları - 2024-12-10

## Özet:

-   ✨ Yeni Özellikler ve Geliştirmeler: 7 taahhüt

-   🐞 Çözülmüş Sorunlar: 1 işlem

-   📚 Belgeler: 2 taahhüt

-   🔧 Kod Tabanı Bakımı ve Güncellemeleri: 5 taahhüt

ve 2 ❤️ katkıda bulunan

## Değişiklikler

### ✨ Yeni Özellikler ve Geliştirmeler:

-   <a id="commit-bf437f1"></a>**cmake** [#bf437f1](https://github.com/ahmettsoner/pars/commit/bf437f142c44e7140917d66242993cfd28aafd13): yükleme/kaldırma ile choco paketi desteği ekleyin komut dosyaları ([@CI/CD tarafından) Bot](mailto:ci-bot@parsdevkit.net))

Sistem ortamı yollarını yönetmek ve Program Ekle/Kaldır desteği sağlamak için ChocolateyInstall.ps1 ve ChocolateyUninstall.ps1 komut dosyaları eklendi. Bu geliştirme, kullanılabilirliğin iyileştirilmesini ve sistem paketi yönetimi standartlarıyla uyumluluğu sağlar.

-   <a id="commit-7c974b7"></a>**cmake** [#7c974b7](https://github.com/ahmettsoner/pars/commit/7c974b7a0e560a6b184e84d4009031494d7ec9c0): sürüme göre yayın öncesi algılama ekleyin bilgi ([@CI/CD tarafından) Bot](mailto:ci-bot@parsdevkit.net))

Sürüm öncesi tanımlayıcılar (ör. alfa, beta, rc) için sürüm dizesini değerlendirerek CMake'de yayın öncesi sürümleri algılamak için uygulanan mantık. Bu işlevsellik, oluşturma işlemi sırasında yayın öncesi durumu otomatik olarak belirleyerek sürüm yönetimini geliştirir.

-   <a id="commit-81b604c"></a>**cmake** [#81b604c](https://github.com/ahmettsoner/pars/commit/81b604c76f22b2e28c2a7e5dc11886efcae2c1a2): Go derlemeleri için 386 mimari desteği ekleyin ( Yazan [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

Go derleme sürecinde 386 mimari desteği etkinleştirildi.

-   <a id="commit-17628e3"></a>**choco** [#17628e3](https://github.com/ahmettsoner/pars/commit/17628e3fbc9d8cee7f4ed5d88f64751c2ebc5a28): Paketi Program Ekle/Kaldır ve ile geliştirin PATH yönetimi ([@CI/CD tarafından) Bot](mailto:ci-bot@parsdevkit.net))

Chocolatey paketi kurulumu sırasında Program Ekle/Kaldır girişi desteği eklendi.

Uygulamayı kurulum sonrasında PATH ortam değişkenine otomatik olarak ekler.

Uygulama ayrıntılarını Windows Kayıt Defteri'nden kaldırmak ve PATH ortam değişkenini temizlemek için iyileştirilmiş kaldırma işlemi.

-   <a id="commit-3b6244b"></a>**installer** [#3b6244b](https://github.com/ahmettsoner/pars/commit/3b6244bfaeb3248b46f283ca1bb27f43ce4e3c87): MSI yapılandırmasında görüntülenen uygulama adını iyileştirin ( Yazan [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-6137359"></a>**cmake** [#6137359](https://github.com/ahmettsoner/pars/commit/61373597d297c05baf8f8f868348495ec888db0f): Chocolatey paketleme desteği ekleyin ([@ tarafından) CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-da170e1"></a>**cmake** [#da170e1](https://github.com/ahmettsoner/pars/commit/da170e1b49fc666b25eee7c11599dbb328872fc9): Chocolatey paketleme desteği ekleyin ([@ tarafından) CI/CD Bot](mailto:ci-bot@parsdevkit.net))

### 🐞 Çözülmüş Sorunlar:

-   <a id="commit-51acc69"></a>**cmake** [#51acc69](https://github.com/ahmettsoner/pars/commit/51acc69ea1d463b5a8dd8da6f6631611320be30c): MSI yükleyici yapılandırmasındaki sorunları çözün (tarafından) [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

### 📚 Belgeler:

-   <a id="commit-a2bf7fe"></a>[#a2bf7fe](https://github.com/ahmettsoner/pars/commit/a2bf7fe6a401fbbe88c0662af42d6e911be77b26): kapsamlı README.md ekleyin (tarafından: [@Ahmet Soner](mailto:ahmettsoner@gmail.com)) (Sayı: [#385](https://github.com/ahmettsoner/pars/issues/385))

Projeye genel bakış bölümleri, temel özellikler, başlangıç ​​kılavuzu, kurulum talimatları, katkı yönergeleri, belge bağlantıları, lisans bilgileri ve topluluk kaynakları dahil olmak üzere Pars deposu için ayrıntılı bir "README.md" dosyası eklendi.

-   <a id="commit-a6b7695"></a>[#a6b7695](https://github.com/ahmettsoner/pars/commit/a6b7695f2fd0ad04cdc3b70e24e9c072ec0c6b4c): kapsamlı README.md ekleyin (#385) (tarafından: [@ CI/CD Bot](mailto:ci-bot@parsdevkit.net)) (Sayı: [#385](https://github.com/ahmettsoner/pars/issues/385))

Projeye genel bakış bölümleri, temel özellikler, başlangıç ​​kılavuzu, kurulum talimatları, katkı yönergeleri, belge bağlantıları, lisans bilgileri ve topluluk kaynakları dahil olmak üzere Pars deposu için ayrıntılı bir README.md dosyası eklendi.

### 🔧 Kod Tabanı Bakımı ve Güncellemeleri:

-   <a id="commit-92ac517"></a>**vscode** [#92ac517](https://github.com/ahmettsoner/pars/commit/92ac517401e6f31156d605392a872834980eceaa): ayarlarda source.organizeImports'u açık olarak ayarlayın ([@CI/CD tarafından) Bot](mailto:ci-bot@parsdevkit.net))

İçe aktarma bildirimlerinin proje kurallarına göre açıkça düzenlenmesini sağlar.

-   <a id="commit-2010a74"></a>**project** [#2010a74](https://github.com/ahmettsoner/pars/commit/2010a74dda273e2b3db1a3ac1bbe9858cfbcdbb1): başlatma, ayarlarla .vscode klasörünü ekleyin ve görev yapılandırmaları ([@CI/CD tarafından) Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-7857c8f"></a>[#7857c8f](https://github.com/ahmettsoner/pars/commit/7857c8f6d2d6f09f4503e40620b68d4ead79e9a2): kullanılmayan dosyalar kaldırıldı ([@CI/CD Bot] tarafından)( mailto:ci-bot@parsdevkit.net))

-   <a id="commit-02ae3da"></a>[#02ae3da](https://github.com/ahmettsoner/pars/commit/02ae3da3d79d97525806d935c832e2b4ba23a443): çakışmaları düzeltme ([@CI/CD Bot](mailto tarafından) :ci-bot@parsdevkit.net))

-   <a id="commit-80ea61f"></a>[#80ea61f](https://github.com/ahmettsoner/pars/commit/80ea61fac776a2fcddfb087f07b78f79ceded35d): Semantic Release için depoyu başlat (tarafından: [@Ahmet) Soner](mailto:ahmettsoner@gmail.com))

## Faydalı Bağlantılar

-   📜 [Tüm Değişiklikler Günlüğü](https://github.com/ahmettsoner/pars/blob/main/CHANGELOG.md)

## Katkıda Bulunanlar:

-   [CI/CD Botu](mailto:ci-bot@parsdevkit.net) (Kod, Yapılandırmalar, Dokümanlar, Testler)

-   [Ahmet Soner](mailto:ahmettsoner@gmail.com) (Dokümanlar, Yapılandırmalar, Kod, Testler)
