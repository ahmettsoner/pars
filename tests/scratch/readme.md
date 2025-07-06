# Test Senaryoları - Tüm Kombinasyonlar

| ID   | Kategori             | Açıklama                                                                  | Set | Layer | Tag | Label | Trigger  | Beklenen Davranış        |
| ---- | -------------------- | ------------------------------------------------------------------------- | --- | ----- | --- | ----- | -------- | ------------------------ |
| TC01 | Matching Logic       | Sadece Set eşleşmesi                                                      | ✅   | ❌     | ❌   | ❌     | Template | Generate çalışmalı       |
| TC02 | Matching Logic       | Sadece Layer eşleşmesi                                                    | ❌   | ✅     | ❌   | ❌     | Template | Generate çalışmalı       |
| TC03 | Matching Logic       | Sadece Tag eşleşmesi                                                      | ❌   | ❌     | ✅   | ❌     | Template | Generate çalışmalı       |
| TC04 | Matching Logic       | Sadece Label (key only) eşleşmesi                                         | ❌   | ❌     | ❌   | ✅     | Template | Generate çalışmalı       |
| TC05 | Matching Logic       | Set + Layer eşleşmesi                                                     | ✅   | ✅     | ❌   | ❌     | Template | Generate çalışmalı       |
| TC06 | Matching Logic       | Set + Tag eşleşmesi                                                       | ✅   | ❌     | ✅   | ❌     | Template | Generate çalışmalı       |
| TC07 | Matching Logic       | Layer + Label (key+value) eşleşmesi                                       | ❌   | ✅     | ❌   | ✅     | Template | Generate çalışmalı       |
| TC08 | Matching Logic       | Tag + Label eşleşmesi                                                     | ❌   | ❌     | ✅   | ✅     | Template | Generate çalışmalı       |
| TC09 | Matching Logic       | Set + Layer + Tag eşleşmesi                                               | ✅   | ✅     | ✅   | ❌     | Template | Generate çalışmalı       |
| TC10 | Matching Logic       | Set + Layer + Label eşleşmesi                                             | ✅   | ✅     | ❌   | ✅     | Template | Generate çalışmalı       |
| TC11 | Matching Logic       | Layer + Tag + Label eşleşmesi                                             | ❌   | ✅     | ✅   | ✅     | Template | Generate çalışmalı       |
| TC12 | Combination Coverage | Tüm eşleşmeler var                                                        | ✅   | ✅     | ✅   | ✅     | Template | Generate çalışmalı       |
| TC13 | Combination Coverage | Hiçbiri yok                                                               | ❌   | ❌     | ❌   | ❌     | Template | Tüm projelerde çalışmalı |
| TC14 | Negative Case        | Template Set var, ama Project/Resource Set farklı                         | ✅   | ❌     | ❌   | ❌     | Template | Generate yapılmamalı     |
| TC15 | Negative Case        | Template Layer var, Resource layer farklı                                 | ❌   | ✅     | ❌   | ❌     | Template | Generate yapılmamalı     |
| TC16 | Negative Case        | Template Tag var, Resource tag eksik                                      | ❌   | ❌     | ✅   | ❌     | Template | Generate yapılmamalı     |
| TC17 | Negative Case        | Template Label var, Project/Resource eşleşmiyor                           | ❌   | ❌     | ❌   | ✅     | Template | Generate yapılmamalı     |
| TC18 | Conflict Resolution  | Template layer’sız, Resource layer’lı                                     | ❌   | ❌     | ❌   | ❌     | Template | Generate yapılmamalı     |
| TC19 | Conflict Resolution  | Template layer’sız, Resource layer’sız                                    | ❌   | ❌     | ❌   | ❌     | Template | Generate yapılmalı       |
| TC20 | Conflict Resolution  | Template tag’lı, Resource tag’larının tamamı yok                          | ❌   | ❌     | ✅   | ❌     | Template | Generate yapılmamalı     |
| TC21 | Matching Logic       | Label eşleşmesi - sadece key                                              | ❌   | ❌     | ❌   | K=1   | Template | Generate çalışmalı       |
| TC22 | Matching Logic       | Label eşleşmesi - sadece value                                            | ❌   | ❌     | ❌   | V=1   | Template | Generate çalışmalı       |
| TC23 | Matching Logic       | Label eşleşmesi - key+value                                               | ❌   | ❌     | ❌   | K=V   | Template | Generate çalışmalı       |
| TC24 | Trigger Logic        | Resource tetikleyici: eşleşen Template’lere işlem                         | ✅   | ✅     | ✅   | ✅     | Resource | Generate yapılmalı       |
| TC25 | Trigger Logic        | Resource tetikleyici: eşleşmeyen Template varsa işlem atlanmalı           | ✅   | ❌     | ❌   | ❌     | Resource | Generate yapılmamalı     |
| TC26 | Combination Coverage | Birden fazla Project, ortak set                                           | ✅   | ❌     | ❌   | ❌     | Template | Hepsi işlenmeli          |
| TC27 | Combination Coverage | Birden fazla Resource, farklı layer'lara sahip                            | ❌   | ✅     | ❌   | ❌     | Template | Uygun olanlar seçilmeli  |
| TC28 | Combination Coverage | Project tag’leri çoklu, Template subset içeriyor                          | ❌   | ❌     | ✅   | ❌     | Template | Generate yapılmalı       |
| TC29 | Matching Logic       | Template layer’lı, Resource layer yok (default)                           | ❌   | ✅     | ❌   | ❌     | Template | Generate yapılmamalı     |
| TC30 | Matching Logic       | Template layer’sız, Resource layer yok                                    | ❌   | ❌     | ❌   | ❌     | Template | Generate yapılmalı       |
| TC31 | Matching Logic       | Template label’ı K1:V1, Project label’ı sadece K1                         | ❌   | ❌     | ❌   | ✅     | Template | Generate yapılmamalı     |
| TC32 | Matching Logic       | Template tag’ı T1,T2, Project T1,T2,T3                                    | ❌   | ❌     | ✅   | ❌     | Template | Generate yapılmalı       |
| TC33 | Matching Logic       | Template tag’ı T1,T2, Project T1                                          | ❌   | ❌     | ✅   | ❌     | Template | Generate yapılmamalı     |
| TC34 | Conflict Resolution  | Template Set yok, Project Set var                                         | ❌   | ❌     | ❌   | ❌     | Template | Project işlenmemeli      |
| TC35 | Conflict Resolution  | Template Set yok, Project Set yok                                         | ❌   | ❌     | ❌   | ❌     | Template | Project işlenmeli        |
| TC36 | Trigger Logic        | Template tetikleyici, birden fazla Resource’a uygulanmalı                 | ✅   | ✅     | ✅   | ✅     | Template | Her biri işlenmeli       |
| TC37 | Trigger Logic        | Resource tetikleyici, birden fazla Template’e uygulanmalı                 | ✅   | ✅     | ✅   | ✅     | Resource | Her biri işlenmeli       |
| TC38 | Negative Case        | Template Tag subset uymaz                                                 | ❌   | ❌     | ✅   | ❌     | Template | Generate yapılmamalı     |
| TC39 | Matching Logic       | Layer inheritance: Layer olmayan template, tüm layer'lara mı uygulanıyor? | ❌   | ❌     | ❌   | ❌     | Template | Sadece layer’sızlara     |


1. 

| Schema            | Set | Layers | Tags | Labels |
| ----------------- | --- | ------ | ---- | ------ |
| Group             | -   | -      | -    | -      |
| Project   (App1)  | 🗴  | 🗴     | 🗴   | 🗴     |
| Resource  (Res1)  | 🗴  | 🗴     | 🗴   | 🗴     |
| Template  (Tmpl1) | 🗴  | 🗴     | 🗴   | 🗴     |

| Argument | Expected | Notes      |
| -------- | -------- | ---------- |
| Project  | Dotnet   |            |
| Output   | Tmpl1    | Empty file |


```bash
a -f ../tests/scratch/schemas/simple
```
```bash
d -f ../tests/scratch/schemas/simple
```


2. 

| Schema            | Set | Layers | Tags | Labels |
| ----------------- | --- | ------ | ---- | ------ |
| Group             | ˗   | ˗      | ˗    | ˗      |
| Project   (App1)  | 🗴  | 🗴     | 🗴   | 🗴     |
| Resource  (Res1)  | 🗴  | 🗴     | 🗴   | 🗴     |
| Template  (Tmpl1) | 🗴  | 🗴     | 🗴   | 🗴     |

| Argument | Expected | Notes       |
| -------- | -------- | ----------- |
| Project  | Dotnet   |             |
| Output   | Tmpl1    | Sample text |

```bash
a -f ../tests/scratch/schemas/simple2
```
```bash
d -f ../tests/scratch/schemas/simple2
```


3. a. 

| Schema            | Set | Layers | Tags | Labels |
| ----------------- | --- | ------ | ---- | ------ |
| Group             | -   | -      | -    | -      |
| Project   (App1)  | ✔   | 🗴     | 🗴   | 🗴     |
| Resource  (Res1)  | 🗴  | 🗴     | 🗴   | 🗴     |
| Template  (Tmpl1) | 🗴  | 🗴     | 🗴   | 🗴     |


| Argument | Expected | Notes |
| -------- | -------- | ----- |
| Project  | Dotnet   |       |
| Output   | -        |       |


```bash
a -f ../tests/scratch/schemas/simple3-Set-OnlyProject
```
```bash
d -f ../tests/scratch/schemas/simple3-Set-OnlyProject
```

1. b. 

| Schema            | Set | Layers | Tags | Labels |
| ----------------- | --- | ------ | ---- | ------ |
| Group             | -   | -      | -    | -      |
| Project   (App1)  | ✔   | 🗴     | 🗴   | 🗴     |
| Resource  (Res1)  | ✔   | 🗴     | 🗴   | 🗴     |
| Template  (Tmpl1) | ✔   | 🗴     | 🗴   | 🗴     |

| Argument | Expected | Notes       |
| -------- | -------- | ----------- |
| Project  | Dotnet   |             |
| Output   | Tmpl1    | Sample text |


```bash
a -f ../tests/scratch/schemas/simple3-Set
```
```bash
d -f ../tests/scratch/schemas/simple3-Set
```


4. 

| Schema            | Set | Layers | Tags | Labels |
| ----------------- | --- | ------ | ---- | ------ |
| Group             | -   | -      | -    | -      |
| Project   (App1)  | 🗴   | ✔      | 🗴   | 🗴     |
| Resource  (Res1)  | 🗴   | ✔      | 🗴   | 🗴     |
| Template  (Tmpl1) | 🗴   | ✔      | 🗴   | 🗴     |

| Argument | Expected | Notes       |
| -------- | -------- | ----------- |
| Project  | Dotnet   |             |
| Output   | L1/Tmpl1 | Sample text |

```bash
a -f ../tests/scratch/schemas/simple4-Layer
```
```bash
d -f ../tests/scratch/schemas/simple4-Layer
```

5. 

| Schema            | Set | Layers | Tags | Labels |
| ----------------- | --- | ------ | ---- | ------ |
| Group             | -   | -      | -    | -      |
| Project   (App1)  | 🗴   | 🗴     | ✔    | 🗴     |
| Resource  (Res1)  | 🗴   | 🗴     | ✔    | 🗴     |
| Template  (Tmpl1) | 🗴   | 🗴     | ✔    | 🗴     |

| Argument | Expected | Notes       |
| -------- | -------- | ----------- |
| Project  | Dotnet   |             |
| Output   | Tmpl1    | Sample text |

```bash
a -f ../tests/scratch/schemas/simple5-Tag
```
```bash
d -f ../tests/scratch/schemas/simple5-Tag
```

6. 

| Schema            | Set | Layers | Tags | Labels |
| ----------------- | --- | ------ | ---- | ------ |
| Group             | -   | -      | -    | -      |
| Project   (App1)  | 🗴   | 🗴     | 🗴   | ✔      |
| Resource  (Res1)  | 🗴   | 🗴     | 🗴   | ✔      |
| Template  (Tmpl1) | 🗴   | 🗴     | 🗴   | ✔      |

| Argument | Expected | Notes       |
| -------- | -------- | ----------- |
| Project  | Dotnet   |             |
| Output   | Tmpl1    | Sample text |

```bash
a -f ../tests/scratch/schemas/simple6-Label
```
```bash
d -f ../tests/scratch/schemas/simple6-Label
```

6. 

| Schema            | Set | Layers | Tags | Labels |
| ----------------- | --- | ------ | ---- | ------ |
| Group             | -   | -      | -    | -      |
| Project   (App1)  | 🗴   | 🗴     | 🗴   | ✔      |
| Resource  (Res1)  | 🗴   | 🗴     | 🗴   | ✔      |
| Template  (Tmpl1) | 🗴   | 🗴     | 🗴   | ✔      |

| Argument | Expected | Notes       |
| -------- | -------- | ----------- |
| Project  | Dotnet   |             |
| Output   | Tmpl1    | Sample text |

```bash
a -f ../tests/scratch/schemas/simple6-Label-Value
```
```bash
d -f ../tests/scratch/schemas/simple6-Label-Value
```


7. Complex

| Schema            | Set | Layers | Tags | Labels |
| ----------------- | --- | ------ | ---- | ------ |
| Group             | -   | -      | -    | -      |
| Project   (App1)  | 🗴   | 🗴     | 🗴   | ✔      |
| Resource  (Res1)  | 🗴   | 🗴     | 🗴   | ✔      |
| Template  (Tmpl1) | 🗴   | 🗴     | 🗴   | ✔      |

| Argument | Expected | Notes       |
| -------- | -------- | ----------- |
| Project  | Dotnet   |             |
| Output   | Tmpl1    | Sample text |

```bash
a -f ../tests/scratch/schemas/simple6-Label-Value
```
```bash
d -f ../tests/scratch/schemas/simple6-Label-Value
```