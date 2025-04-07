Dev feature/\*\* branchler'den PR ile kabul edilecek. (.github/workflows/branch-controls/check-branch-name.yml)

-   PR mesaj formatı grm commit validate commit-id/message gibi bir yapı ile kontrol edilebilir

Changelog dosyası PR mesajı üzerinden takip edilecek, diğer commitlerde bulunan conventional commit standarta uygun mesajlar dahi kabul edilmeyecek?

-   PR commit mesajı kime ait olduğu kullanılarak contributer bilgisi burdan sağlanacak?
-   PR'da bir issue close edilmesi gerekli, bu durum zorunlu olacak. issue olmadan dev'e code geçmeyecek, fix'ler içinde aynı şekilde
-   İleride Major version bilgisi de conventional commit mesajına bağlı olarak yönetilmeli, şu anda minor feature, patch ise main'den sonra fix için kullanıılyor

Tag oluşturma işlemleri kısıtlanacak herkes oluşturmamalı, tag'ler ci-cd sürecinde otomatik hazırlanabilmeli sadece
Süreç içinde tag oluştur - yayın al yaklaşımı benimsenerek tüm release işlemleri böyle ilerlemeli (tag-dev -> buid-dev gibi)

Release çıktıları için `VERSION/pars-linux-amd64.bin.tar.gz` formatında isimlendirme kullanılacak
Release çıktıları sıkıştırılacak. binary, changelog ve user guide bilgileri içerecek ileride?

tüm release'ler için binary, package (deb, rpm), installer (msi) hazırlanmalı `VERSION/pars-windows-x86_64.msi.zip`, `VERSION/pars-linux-amd64.deb.zip` gibi

github release için draft yönetimi yapılmalı? özellikle stabil channel için

stabil versionlar için installer, package hazırlanması tammalanmalı
stabil versionlar için download ve install dökümanları hazırlanarak otomatik güncellenmeli özellikle download page'ler
