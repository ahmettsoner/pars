1. Simple: Basit project, resource ve template
ws'de Class Library `App1` Dotnet oluşturulması, root'ta içi boş `Tmpl1` dosyası generate edilmesi
```bash
a -f ../tests/scratch/schemas/simple
```
```bash
d -f ../tests/scratch/schemas/simple
```


2. Simple: Basit project, resource ve template (lorem)
ws'de Class Library `App1` Dotnet oluşturulması, root'ta lorem metin içeren `Tmpl1` dosyası generate edilmesi
```bash
a -f ../tests/scratch/schemas/simple2
```
```bash
d -f ../tests/scratch/schemas/simple2
```

3. Simple Set: Basit project, resource ve template (lorem)
ws'de `S1` seti ile Class Library `App1` Dotnet oluşturulması, root'ta lorem metin içeren `Tmpl1` dosyası generate edilmesi
```bash
a -f ../tests/scratch/schemas/simple3-Set
```
```bash
d -f ../tests/scratch/schemas/simple3-Set
```


3. a. Simple Set: Basit project, resource ve template (lorem)
ws'de Sadece projede `S1` seti ile Class Library `App1` Dotnet oluşturulması, root'ta dosya generate edilmemesi
```bash
a -f ../tests/scratch/schemas/simple3-Set-OnlyProject
```
```bash
d -f ../tests/scratch/schemas/simple3-Set-OnlyProject
```

4. Simple Layer: Basit project, resource ve template (lorem)
ws'de `L1` layer ile Class Library `App1` Dotnet oluşturulması, root'ta lorem metin içeren `L1/Tmpl1` dosyası generate edilmesi
```bash
a -f ../tests/scratch/schemas/simple4-Layer
```
```bash
d -f ../tests/scratch/schemas/simple4-Layer
```