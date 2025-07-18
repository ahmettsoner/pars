Grouplu proje eklendiğinde `p r Group/Project` sln kalıyor proje siliniyor?


testlerde bazılarında env bozuluyor ve ana db'ye yazıyor

---
Dotnet Proje oluşturma
| Test ID   | Senaryo Açıklaması | Test Adımları             | Beklenen Sonuç                                                                                                       |
| --------- | ------------------ | ------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| E2E-PD001 | Basic              | Platform Dotnet           | ws'de `Basic/Basic.csproj` olmalı, `pars p l` de `Basic` projesi görünmeli                                           |
| E2E-PD002 | Dependency         | Platform Dotnet, EF Nuget | ws'de `Dependency/Dependency.csproj` olmalı, csproj'de nuget tanımlı olmalı, `pars p l` de `Basic` projesi görünmeli |
|           |                    |                           |                                                                                                                      |




Proje oluşturma
Basit dotnet proje classlib (/tests/scratch/schemas/ts/project/Basic.yml)
Basit dotnet proje classlib, custom package
Basit dotnet proje classlib, dependency
Basit dotnet proje classlib, reference
