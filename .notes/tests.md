Project, sadece Platform type ile çalışabilir
rollback mekanizmaları
  - veritabanı kayıt rolback
  - Provider işlemi rollback
  - Dosya işlemleri rolback
group/project dosya silme işlemleri `DestroyProject`


Temel bilgiler ile işlem yapılabilmeli `tests/scenario/schemas/simple`

Grouplu proje eklendiğinde `p r Group/Project` sln kalıyor proje siliniyor?


testlerde bazılarında env bozuluyor ve ana db'ye yazıyor

set bağımsız template ve resource'ler işlenebilmeli? 
layer tanımlanmazsa root folder kullanlarak işlem yapılmalı? projede "/" path için layer tanımlı değilse otomatik eklenmeli
---

`ProjectServiceInterface`'te Proje dosya kaldırma func tanımlanmalı

`Project.Specifications.Configuration` da tanımlı references ve layeer Specification'a taşınması