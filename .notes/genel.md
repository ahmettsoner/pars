https://repology.org/
https://keepachangelog.com/en/1.1.0/
https://semver.org/
https://www.conventionalcommits.org/en/v1.0.0/

Available Pars Capabilities

-   Code Generation
-   Project Organization
-   Build Manager

Code Generation (ayrı sayfa)
1- workspace oluştur (yada mevcut birini kullan)
2- opsiyonel olarak grup modeli oluştur
3- proje modellerini oluştur
4- resource modelerini oluştur
5- template modellerini oluştur
6- generate et

Project Organization (ayrı sayfa)
1- workspace oluştur (yada mevcut birini kullan)
2- project model oluştur.
3- projeler arasında ki ilişkileri tanımla
4- projede kullanılan hazır paketleri tanımla
5- proje oluştur veya güncelle

Build Manager (ayrı sayfa)
1- workspace oluştur (yada mevcut birini kullan)
2- project model oluştur.
3- proje oluştur veya güncelle
4- Projeyi build et

---

type: Template
kind: File
name:
metadata:
tags:
specifications:
name:
output:
set:
layers:
template:
content: |
Code template ten farklı olarak doğrudan, "template.code" bulunuyor, benzer şekidle bunu da tamamen dökümante eder misin tüm alanları ve açıklamaları ile bilrlikte

---

https://docs.gomplate.ca/syntax/
https://github.com/hairyhenderson/gomplate
https://github.com/phcollignon/Go-Template
https://github.com/obsproject/obs-studio/wiki/install-instructions
https://gohugo.io/installation/
https://ui.shadcn.com/

https://github.com/obsproject/obs-studio?tab=readme-ov-file
https://hub.docker.com/r/dockurr/macos
https://hub.docker.com/r/almalinux/9-base/tags
https://hub.docker.com/_/centos/tags
https://hub.docker.com/_/ubuntu/tags
https://hub.docker.com/_/debian/tags
https://hub.docker.com/u/opensuse

---

https://github.com/TheZoraiz/ascii-image-converter #https://chatgpt.com/c/d78ff220-d854-4c82-ac82-d59630fc4028

---

https://github.com/parsdevkit/pack-test
https://asiffer.github.io/posts/launchpad/

https://mentors.debian.net/

---

multilanguage sample
https://github.com/squidfunk/mkdocs-material/discussions/2346

installation sample
https://docs.flutter.dev/get-started/install

releases sample
https://docs.flutter.dev/release/archive

desteklenebilecek platformlar listesi
https://docs.sentry.io/platforms/
--
https://github.com/chocolatey/choco/releases
https://github.com/dotnet/core/releases
https://github.com/elastic/elasticsearch/releases
https://github.com/squidfunk/mkdocs-material/releases
https://github.com/xamarin/xamarin-macios/releases
https://github.com/pytorch/pytorch/releases
https://github.com/google/benchmark/releases
https://github.com/zulip/zulip/releases
https://github.com/tensorflow/tensorflow/releases
https://wiki.ubuntu.com/Releases
https://github.com/microsoft/azure-pipelines-agent/releases

---

V 1.0.0
Documentation
mkdocs serve hatalar giderilmeli
Getting started
Concept gözden geçirilmeli
Guides
Commands gözden geçirilmeli
Schemas gözden geçirilmeli
Components index kaldırılmalı, Templates gözden geçirilmeli
Advanced gözden geçirilmeli
Setup tamamlanmalı
Schema tanımları kontrol edilmeli, bazı specification tanımları doğru indentation'da değil (http://localhost:8000/guides/schemas/object/template/file-template-object-model.html#specificationstemplatecontent)
Project, resource, group gibi eksik komutlar eklenmeli(https://docs.parsdevkit.net/v1.0.1/guides/commands/index.html )
Active & Current workspace indicators should be updated (https://docs.parsdevkit.net/v1.0.1/guides/commands/environment/list.html#usage)
Autocomplate tip url hatalı (https://docs.parsdevkit.net/guides/commands/workspace/index.html#-switch)
Workspace kalmış (https://docs.parsdevkit.net/guides/commands/environment/list.html#flags)
Hello world eklenmeli document a
Workflow
Initialize workspace
Set belirle
Create project
_ With set
_ Set layer
Create template
_ With set
_ Set layer
Create resource

-   With set 
    \* Set layer
    google analitics tanımlanmalı (https://squidfunk.github.io/mkdocs-material/setup/setting-up-site-analytics/#google-analytics)
    feedback hazıarlnamlı (https://squidfunk.github.io/mkdocs-material/setup/setting-up-site-analytics/#was-this-page-helpful)
    git-revision-date-localized, git-committers, git-authors aktive edilmeli
    pars uygulama adı, pars-cli olarak değiştirilebilr?
    arm32 kaldırılabişlir, vm bile oluşturulmuyır

Commands

-   Sonraki versiyon'a uygulanacak özellikler pasifize edilmeli
    Functions
-   $data := slice (struct "Info" (struct "Name" "A burrda
    -   slice kullanımı değiştirilmeli
    -   struct için object space'i altında metodlar geliştirilebilir? veya örnek dğeiştirilebilir

Github
organization profile tamamlanmalı
Workflow
sadece src dizininde değişiklik varsa release çıkmalı, docs ta değişiklik varsa döküman yayınlanmalı
docs versioning iptal edilmeli

Package Managers
Windows Choco, MSI, MSIX, Scoop, Winget, cab yapılandırma
Linux APT, YUM, DNF, Pacman, Zypper, Snap, Flatpak, LinuxBrew yapılandırma
MacOS Homebrew, DMG, PacMan, MacPorts, Fink yapılandırma
Platorms: pip, npm, npx, dotnet-tools etc
Markets: Windows Store, Mac App Store
asdf (https://asdf-vm.com/)?
sfx, zip, 7z, tar.gz

webi: https://webinstall.dev/
pkgx: https://pkgx.dev/

Source code:
Build manual, build with make

Manual:
Install scripts

Readme/Guides
badge'ler tanımlanmalı
https://github.com/getsentry/sentry-dotnet/blob/main/README.md

---

Git Repository hazıarlanacak?
Documentation
https://www.mkdocs.org/user-guide/configuration/
https://www.mkdocs.org/user-guide/deploying-your-docs/
döküman page, share butonlar düzenlenmeli, tamamlanmalı (X, f)
docs.parsdevkit.net hazırlanmalı
github linkleri tanımlanmalı
Pars Getting Started
Concept index tamamlanmalı \* layers eklenmeli
içerikler tamamlanmalı
Concept > Task kaldırılmalı
Overview tamamlanmalı
Workflow şimdilk kaldırılmalı
Extensions
Platforms tamamlanmalı
Platforms'a datatype karşılaştırma eklenmeli
Modules, Plugins ve Packages şimdilk kaldırılmalı

    Docs/docs yerine docs/content olarak dönüştürülmesi

    Extensions'ta Language section hazırlanmalı

    Documentation 3 ayrı yapıda planlanmalı, uygulama dökümanı, proje dökümanı ve resmi döküman sayfası şeklinde.
    	uygulama dökümanı - user manuel, uygulama ile kullanım hakkında kullanıcılara sunulmalı kısa ve basitçe
    	proje dökümanı, build, test, run gibi proje yapılandırması ve geliştirme ortamı hazırlanması ile temel kullanım bilgileri (project_root/docs)
    	Resmi döküman sayfası, kapsamlı (docs.parsdevkit.net)

General
Docs, tests ve src dizin yapısına geçilmeli
Pars project temp env tanımlanabilir
Path env tanımlanabilir

Context

-   Workspace eklenecek

Models \* Model yaml tanımları PascalCase olmalı, template te doğrudan kullanılabilmeli
testler çalıştırılmalı
project ve group için submit ve new komutları ayrılmalı

Syntax
_ Basics başlığı dökümanda Syntax olrak değiştirilebilir
_ Dokümanda, loop conditions gibi temel yapılar tanımlanmalıdır
Datatypes
_ Dökümanda DataTypes'lar düzenlenecek (functions ve schema)
_ Yeni Veritipi yapısının uygulanması
_ Time eklenecek
_ Primitive, BuiltIn veya Value olarak değiştirilebilir
_ Value definations tanımlanabilir datatype gibi, bool (true,false,0,1), text (a-z), number (0-9), none/null, Object (struct), Enum
_ Enum eklenebilir? değerlendirilecek

Functions
_ Döküman control edilecek (math.Add)
_ function'larda error return kaldırılmalı, single return olmalı

Commands \* New komutu, yeniden isimlendirilmeli

Models \* Template'te section layer altına alınmalı, section yoksa doğrudan resoruce kullanlır, varsa section

Providers \* dotnet, npx gibi provider'lar host machine öncelikli kullanlmalı

Utils
GetCodeBaseLocation() Environment'e taşınması
GetExecutionLocation() paketleme ve test yapıları için uygun hale getirilebilmeli

Web Site
_ docs.parsdevkit.net yönlendirme hatası
_ www.parsdevkit.net yönlendirme sayfası eklenmeli

Github
organizasyona dönüştürülmeli (https://github.com/settings/organizations)
pages hazıarlanmalı
pages web site entegre edilmeli
Pages, custom domain (https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/managing-a-custom-domain-for-your-github-pages-site)
github issue tempalte'ler tamamlanmalı
Licence
Apache olmalı (https://chatgpt.com/c/1d791ede-b536-42ee-9899-ada7976caa81)

Releases \* indirmelerde sıkıştırmaların versiyon içermesi güzel olur, farklı versiyonları aynı isimde saklamamak için (https://github.com/parsdevkit/pars/releases/download/v1.0.29/pars-ubuntu-386.deb.tar.gz)
ama bu fikir installation sürecinde kullanıcıya versiyon tanımlama zorunluluğu getirebilir unzip'te

Workflow \* CI-CD hazıarlanacak

V 1.1.0
Documentation
docs home page infografic hazırlanmalı (https://www.youtube.com/watch?v=h7YXjpPoegY&ab_channel=WebsiteLearners)
social ve community linkler eklenmeli
Slogan belirlenerek eklenmeli (Anasayfa)
Index sayfaları Grid ile zenginleştirilebilir
Guides
Index tanımlanmalı
Contributing tamamlanmalı
Extras tamamlanmalı
Hello world eklenmeli document a
Context Döküman hazırlanacak
dökümanda şunlar incelenmeli düzenlenmeli
function return'de errorlar kaldırılmalı,
datatype number, string vb camelcase, interface tanımları pars value-type olarak düzenlenmeli
https://vscode.dev/github/parsdevkit/pars/ desteklenmeli
İnstall, uninstall, update işlemleri açıklanmalı
Pars docs dizininde developer guide olmalı github readme bulunmalı sadece
Models
Section ayrılacak
_ Kod düzenleme
_ Resource'ta kullanlıyor
_ Section modeli dökümante edilecek
Task yapısı hazırlanacak
_ Event'lar tanımlanacak

Package
_ Package sign sağlanmalı
_ Package sum hash sağlanmalı

Notification \* Hata, bilgilendirme ve yönlendirme mesajlarının iyileştirilmesi

Error management
_ Hata yönetiminin iyileştirilmesi
Developer Docs
_ https://github.com/getsentry/sentry-dotnet gibi badge'ler eklenebilir

Moduler alt yapi
Alt yapı iyileştirilmesi
"/group/group.yaml.templ" gibi tempalte'ler kullanldığı modülde olmalıu (cmd gibi)

Functions
https://www.postgresql.org/docs/current/functions-string.html incelenmeli function'lar genişletilmeli
_ Time functions
_ düzenlenecek eksikler giderilecek
_ dökümante edilecek
_ UUID eklenecek (https://docs.gomplate.ca/functions/uuid/)
_ Random eklenecek (https://docs.gomplate.ca/functions/random/)
_ Path eklenecek (https://docs.gomplate.ca/functions/path/, https://docs.gomplate.ca/functions/filepath/, https://docs.gomplate.ca/functions/file/)
_ URL (https://docs.gomplate.ca/functions/conv/)
_ Crypto (https://docs.gomplate.ca/functions/crypto/)
_ String
_ Replace
_ Time functions (get date, get time, format date vs)
_ test function'lar eklenebilir (isTrue, isLessThan, isLessOrEqualThan) gibi
_ HTTP function eklenebilri, json, xml gibi değerleri remote serverdan almak isteyebilir (http.Get <url>)
_ Hex(2lik sistem, 3 lük sistem gibi deönüşümler)
_ CSV (expression ile field ve değere erişim sağlanabilir)
_ TOML (expression ile field ve değere erişim sağlanabilir)
_ XML (expression ile field ve değere erişim sağlanabilir)
_ YAML (expression ile field ve değere erişim sağlanabilir)
$ gomplate -i '{{ conv.Dict "name" "Frank" "age" 42 | data.ToYAML }}'
age: 42
name: Frank \* JSON (expression ile field ve değere erişim sağlanabilir [https://github.com/elgs/jsonql])
$ gomplate -i '{{ dict 1 2 3 | toJSON }}'
{"1":2,"3":""}

Test
_ Template Functions testlerin hazırlanması
_ testlerin tamamlanması

http://localhost:8000/guides/schemas/object/resource/object-resource-object-model.html#specificationsattributes, Defination>Datatype: array bilgisi ilgili object ile tanımlanabilir []Attribute gibi? böylece parametre, attribute gibi yapılar ayrı bir dökümana da taşınabilr?

Value Types

-   Language reference sağlanmalı dökümanda
-   Complex type'lar tanımlanabilir
    -   UUID
    -   IP
    -   MAC
    -   ID
    -   Monetary
    -   Range
    -   JSON
    -   XML

Template Dty-Run (https://blog.gopheracademy.com/advent-2017/using-go-templates/ > Verifying Templates)
Repository \* Project Workspace Object repodan doldurulmalı, bu yapı olabildiğince ayrıştırılmalı, gerekirse service katmanı hazırlanmalı

Projects

-   Application
-   Data (https://docs.gomplate.ca/) da olduğu gibi farklı kaynaklardan verileri alarak, birleştirerek data üzerinde manupulation yapabilir ve yeni veri üretebilir. ilerde sadece file generate değil ftp, remote, git gibi output seçenekleri olacak zaten
    Task yapısı ile daha kapsamlı işlemler için kullanlıabilir, örneğin belirli zamanlarda gcp'den bir dosya al, s3'ten al birleştir işle ve sonucu mail olarak gönder gibi (ETL gibi)
    Schedule yapısı ile periyodik işler gerçekleştirebilir
    Subscribe yapısı ile belirli farklı platformlarda yapılan değişikliklere subscribe olarak ona göre aksiyon alınmasını sağlayabilir, örneğin s3'e resim eklendiğinde küçült gibi

Commands
_ Command'lar kontrol edilmeli, build test gibi komutlar için autocomplete gibi fonksiyonlar tamamlanmalı
_ İlgili command'lar için döküman tamamlanmalı
_ Submit --file flag local file alabileceği gibi S3, GCP, OneDrive, Git, Http'den dosya alabilmeli
File flag için supported mime types listesi hazırlanmalı ve açıklanmalı application/yaml
Farklı formatlar da desteklenebilir ilerde (Toml, Json, CSV vs)
_ Submit --content flag oluşturularak, --file ile dosya almak yerine desteklenen kaynaklardan doğrudan body'de alabilmeli. Stdin desteklemeli datasource olarak
_ Remote File import (--file için github gibi)
_ Apply, model new, remove apply ve generate yapıları tamamlanmalı

Context
_ Environment (Host Machine, locale) eklenebilir
_ Time eklenebilr (anlık tarih, gün ay yıl vs)

Providers \* dotnet, npx gibi provider'lar host machine öncelikli kullanlmalı, binary kullanımı provider management alt yapısı ile sağlanmalı

Template
_ Template'ler Pars Template modelleri işleyecek şekilde özelleştirilmiş olmalı, context yapısı ve fucntion'lar ilgili Template yapısına uygun olmalı
_ built-in function ve action'lar gözden geçirilerek pars'ta ki alternatifleri değerlendirilmeli, dökümante edilmeli - if-else-else if, range-else, break, contuniue, with-else, index, slice, len, print, printf, println
_ Map ve array kullanım örnekleri
_ range, index, len vs desteklenenler
_ Pipe yapısı dökümante edilmeli "{{"output" | printf "%q"}}"
_ Değişken tanımlama örneklenmeli dökümante edilmeli, değişken argumentler örnejkenmeli, scope önemi vurgulanmalı
$variable := pipeline
$variable = pipeline
range $index, $element := pipeline \* Inner template yapısı hazırlanmalı, test edilmeli ve dökümante edilmeli "block", "define", "template" keywordları
Nested ve Associated template yapıları test edilerek dökümante edilmeli (https://pkg.go.dev/text/template#hdr-Associated_templates)
define, template, block incelenerek dökümante edilmeli (https://developer.hashicorp.com/nomad/tutorials/templates/go-template-syntax)

    	Nested template örneklenmeli
    		{{define "T1"}}ONE{{end}}
    		{{define "T2"}}TWO{{end}}
    		{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
    		{{template "T3"}}
    * Parantez "(", ")" ile pipe veya action gruplama yöntemi dökümante edilmeli  "(find .Resource.Attributes "ID").Name" gibi
    Datasource (Context)
    	Note: The initial context (.) is always available as the variable $, so the initial context is always available, even when shadowed with range or with blocks ({{ with "foo" }}now context is {{ . }} but original context {{ $ }} {{ end }})

    Complex yapılar ve işlemler dökümante edilmeli (https://github.com/phcollignon/Go-Template?tab=readme-ov-file#03---add-conditions-and-golang-template-functions)
    {{ range .Number}}
    	{{if eq  .n1 .n2 }}
    		{{- .n1}} = {{.n2}}
    	{{else}}
    		{{- if lt .n1 .n2 }}
    			{{- .n1}} < {{.n2}}
    		{{else}}
    			{{- .n1}} > {{.n2}}
    		{{end}}
    	{{end}}
    {{end}}

Github
_ Labels düzenlenmeli (issue'larda bilgi ve işaretleme amaçlı) (https://github.com/chocolatey/choco/labels)
_ Milestones belirlenmeli (Yapılacak işler listesi, hangi versiyonda hangi issue'ler yapılacak) (https://github.com/dotnet/core/milestones)
_ Discussions aktive edilmeli
_ Actions
_ Workflows
_ Static Code Analyze
_ Linter
_ Github code spac \* https://github.com/parsdevkit/pars/community tamamlanmalı

Workflow
_ Bump Version scriptleri hazırlanmalı
./bump-vesion.sh
./bump-vesion.sh 2
_ Dev Number scriptleri hazırlanmalı
./dev-number.sh

    * Release Number scriptleri hazırlanmalı
    	./release-number.sh v1.3.0
    	./release-number.sh v1.3.0 --del

Release
_ MacOS hazırlanmalı
_ Docker image da release çıkmalı
_ Download için sumcheck kontrolü desteği sağlanmalı
_ Release Body planlanmalı?

Downloads \* Download için reverse proxy olmalı, download.parsdevkit.net ten github release indirmeleri sağlanmalı

Web site \* Downloads ve documents ayrı repolarda ayrı subdomain de olmalı.
dl.pardevkit.net
dl-eu, dl-as, dl-us gibi edge pointler olabilir

    * Dl işleyişinde directory ler os arch version olmalı

    * Latest dizini olmalı

Test
Temel fonksiyonları içeren smoke test, farklı os ve platformlarda yapılacak test.

    Workspace init
    Workspace list
    Group submit
    Project submit
    Resource submit
    Template submit
    Generation check

---

Hesaplar ve domainler hazırlanacak

-   Gmail
-   Domain
-   Github
-   Slack
-   Medium
-   Discord
-   Facebook (Group)
-   Instagram
-   Twitter
-   Tiktok
-   StackeOverFlow
-   LinkedIn
-   Mastodon

-   producthunt.com
-   wiki
-   sourceforge
-   groups.google.com
-   Docker
-   Azure DevOps
-   Bitbucket
-   Jira
-   Conlfliunce
-   Notion
-   Trello
-   Evernote
-   Youtube Channel
-   Telegram
-   Whatsapp
-   Pinterest
-   fosstodon
-   Threads
-   twitch.tv
-   libsyn
-   spotify
-   NPM
-   Yarn
-   Composer
-   Nuget
-   Pip
-   Maven
-   quora
-   Reddit
-   Tumblr
-   Blogspot
-   tutorialspoint
-   w3schools
-   udemy
-   vercel
-   netlify
-   heroku
-   supabase
-   aws
-   digitaloceanfirebase
-   replit
-   postman
-   chocolatey
-   dev.to
-   Hashnode
-   yandex
-   readthedocs
-   official blog
-   bülten

    Şu kullanım boş resource içinde sağlanmalı, bir package bilgisini resource olmadan da sağlayabilmeli?
    {{- $entityContext := getContextByBase . "resource::Product" "layer::Core:Data:Entity" "section::Entity" -}}
    using {{ $entityContext.Section.Package }};

    ApplicationDbContextConcrete template'te type için inner type'ların türü resource ise context'i getirilebilmeli (entityContext := getContextByBase)

Project, task, resource ve template paketi ile, uçtan uca proje template i tanımlama. Örneğin şirketin web api standardı oluşturulabilir.

Document da hello world samples olmalı, farklı platformlar ve desteklenen project type lar için basit projeler sağlabmalı. Mesela angular proje oluşturma, build ve execute sağlanmalı

Context model klasör yapısı düzenlenmeli

template section bölümü, layer altına sections olarak taşınması

Complexity calculator eklenebilir resmi siteye:
4D concept
X line components (tables count)
Y line layers (depends on architecture)
Z line variations (depends on sections)
Calculate formula = X*Y*Z

Type, Category Resource için set, resource, template, section gibi tanımlar olabilir namespace ve isim bulmak için

DataTypeDefinition big short long yerine Int64, Int32, Int16, Int8 gibi tnaımlar kullanılmalı

DataTypeDefinition'lar dökümanda açıklanmalı

Section lar özelleştirilebilir olmalıdır, örneği endpoint section, repository section gibi. Bu yapı hem standardizasyon sağlar hemde template geliştirici ile resource geliştirici arasında contract olur. Options ile yapılan işlemler section üzerinden sağlanabilir

Pars Google reklam (code generation project management)

Paket yönetimi, workflow ve registry yapısı düzene konulmalı

Generators
https://github.com/phcollignon/Go-Template?tab=readme-ov-file

Eklenecek fonksiyonlar

FUNCTION IsEven(number)
IF number MOD 2 == 0 THEN
RETURN True
ELSE
RETURN False
END IF
END FUNCTION

FUNCTION Factorial(n)
IF n == 0 THEN
RETURN 1
ELSE
RETURN n \* Factorial(n - 1)
END IF
END FUNCTION

FUNCTION StringLength(str)
length = 0
FOR EACH character IN str
length = length + 1
END FOR
RETURN length
END FUNCTION

FUNCTION ReverseString(str)
reversedStr = ""
FOR i FROM length(str) - 1 TO 0
reversedStr = reversedStr + str[i]
END FOR
RETURN reversedStr
END FUNCTION

---

Object

-   Number
    -- Bit
    -- Int
    --- UInt8 (UByte)
    --- Int16 (Short)
    --- UInt16 (UShort)
    --- Int32 (Int)
    --- UInt32 (UInt)
    --- Int64 (Long)
    --- UInt64 (ULong)
    --- Int128
    --- UInt128
    -- Float
    --- Float16
    --- Float32
    --- Float64 (Double)
    --- Float126
    -- Decimal
-   Text695
    -- String (Text)
    -- Char
-   Boolean
-   Map (Dictionary)
-   Time
    Null

---

İçerikler
_ Flag ile proje oluşturma.
_ Hazır paketler ile mimari projeler hazıarlnması
_ Best practices uygun framework içeren çözümler sağlanmalı, farklı mimari yaklaşımlar için ayrı ayrı
_ Bir resource için bir layerda birden fazla dosya nasıl oluşturulur
_ Bir resource için birden fazla layerda birer tane dosya nasıl oluşturulur
_ Bir layerda birden fazla resource için birer tane dosya nasıl oluşturulur
_ Switch case için farklı dil ve versiyonlarda nasıl verimli ve efektif kod kullanılır, template nasıl olmalı şeklinde içerik üretilebilir
_ string plus ve stringbuilder arasında ki fark ve stringbuilder template içinde kullanımı örneği olabilir

Test
Komple silme işlemi tamamlanmıyor (Group ve Workspace hata veriyor)
\tests\scenario\clean-microservice\BasicProjectStructureTestSuite_test.go

How to

-   Code template, options ile bir attribute nasıl bulunur
-   Template'e json olarak data nasıl yazdırılır

---

Documentation
Setup
Set Path environment

     Contributing
     * Clonning
    	Fork parsdevkit/pars
    	git clon <yourprofile>/pars
    	git remote add parsdevkit https://github.com/parsdevkit/pars.git
    	git pull parsdevkit main
    	$env:PARS_PROJECT_ROOT = (Get-Location).Path tanımlanmalı
    	IDE/Terminal kapatıp aç
    	go mod tidy
    	create branch (see branching strategies, branching rules)

    	> PARS_PROJECT for custom execution location

     * Environment
    	Install Go

     * Local Testing
    	cd tests
    	go test -count=1 .\...
    	go test -count=1 .\unit\...

     Release

     Issues

     Project

https://squidfunk.github.io/mkdocs-material/publishing-your-site/
https://github.com/softprops/action-gh-release (release optionslar)

---

Git (https://chatgpt.com/c/b0cf2388-cfdd-43e9-884b-96516614bb4d)
main (stabil, latest, no push, locked)
dev (actively development, with PR)
feature (features with name)
release (releases with tag)
hotfix (production development, with tag)
fix/main/issue-<issue-id>
hotfix/main/critical-bug-<description>
bugfix/main/feature-<description>

release channels
release os, archs:
https://gist.github.com/zfarbp/121a76d5a3fde562c3955a606a9d6fcc
https://chatgpt.com/c/2e7de4f6-500b-4ac7-8389-bc2d58b042ce
https://chatgpt.com/c/171c93a1-af68-40a5-ac1b-09e876580a19

    OS				Architectures
    Linux 			386, amd64, arm, arm64 (Ubuntu, Centos, Debian, Fedora, openSuse)
    Darwin 			386, amd64, arm, arm64
    Windows 		386, amd64
    Solaris 		amd64
    NetBSD			amd64, arm64
    FreeBSD			386, amd64, arm
    OpenBSD			386, amd64, arm
    OpenSuse
    Plan9			386, amd64


    Manuel installation,
    	https://multipass.run/install
    Package Managers
    Shell Scripts
    Installers

Branching strategies must be documented

Shared branches (main, dev, test, release)
Locked branches (read-only) (main, test, release)
Public branch (dev)
work branches (fix/,feat/,developer,issue/)
Shared branch'ler da rebase (Squash) uygulanmaz
work branch'lerde yapılan işlemlere ait commitler public branch'e rebase ile birleştirilir ve commit mesajı yazılır

Main: production branch, should be stabil, don't merge everyone, only granted person. Release tagging on this branch. On merging from dev, it need to pass all test

Dev: development branch, base of feature branches and require pull request to merge

Feature: feature specific feature

Release, merge dev to main, run tests, need approval to complete release workflow

Release süreci, main üzerinde tag kullanarak tetikleyecek şekilde, olmalı. Release süreci build işlemlerili tamamladıktan sonra deploymetint ve publish için onay gerekmeli

Release branch kullanılabilir, bu branch sadece küçük düzeltmeler ve release hazırlığı için kullanılır, dev den merge edilir. Gerekli hallerde hot-fix branch burdan türetilir

Feature brachte geliştirme yapılır, dev e pr geçilir , pr için unit ve integration testler koşar başarılı olması zorunludur, static kod analizi başarılı olmalı, pr belirli kullanıcılar tarafından onaylanmalı ve pr rule ları geçmeli, pr açıklaması change log a uygun olmalı ve issue id içermeli. Bunlardan sonra dev e eşitlenir. Release branch dev den base alır, gerekli hazırlıklar tamamlanır, push işleminde build oluşturulur, system testler uygulanır, artifact ler hazırlanır, change log dökümanı uygun formatta hazırlanır. Change log'da kapatılan issuelar, geliştirilen feature lar ve katkı sağlayanlar yer alır, main e merge edilir , farklı makine ve ortamlarda sistem testleri uygulanır, herşey başarılı ise geçer, tag verilir ve release başlar . Manuel onaydn sonra yayın gönderilir .

Reviewer rotation uygulanmalı

Branch ler uygun rule lar tanımlanmalk

Release checklist hazırlanması
Change log template hazırlanması
Code review guide hazırlanması
Release guide hazırlanması

---

Change Log Titles

## General

## What's Changed

## Maintenance

## Deprecations

## Bug Fixes

Highlights
Breaking Changes
Deprecated Feature
Features
Bug Fixes
Improvements
Documentation
New Contributors
Contributors

Tracking Changes
https://github.com/dotnet/core/compare/v9.0.0-preview.5...v9.0.0-preview.6

---

MkDocs
mkdocs gh-deploy ile gh actions workflow hazırlandı
mike deploy --push --update-aliases v1.3.1 latest ile yeni versiyon çıkıyor

---

Release

-   sprint sonunda (2w) minor versiyon'da release çıkılmalı
    eğer önemli bir geliştirme yapılmışsa (bug-fix,hot-fix) patch çıkabilir arada

    Channels: Dev, Release, Preview, Stabil

    Dev'den çıkılan yayınlar dev channel, Manuel geliştiricinin ihtiyaç duyması halinde (v1.0.0-dev.1)
    Dev -> Release/v* preview channel automatically end of sprint (v1.0.0-preview.1)
    Hotfix, bugfix from/to Release/v*
    Release -> main
    new version tagging manually to stabil release

git merge dev
git push origin main
git tag v1.0.0
git tag v1.0.2 -f
git tag v1.0.2 -d
git tag $(git tag -l) -d
git push origin v1.0.0
git push parsdevkit v1.0.1
git push parsdevkit v1.0.1 --delete

path env ekle/ <local_repository_path>/dist
go build -o ../dist/

Project Types
_ Mail template yönetimi uygulanabilir, invoce project oluşturup, resource olarak contact bilgileri ve template olarak mail body template
_ Application Packing (.msi, .rpm, .deb vs)
_ Infrastructure, provisioning yapılandırma
_ Workflow (iş akışı yönetimi uygulanabilir, ms app flow gibi)
_ Cloud?
Features
_ Language paketleri ile, versiyon uyumlu generation çözümleri hazırlanmalı, bunlar da custom function olarak sunulmalı. Örneğin namespace, function sign, switch case, property, annotations gibi
_ code.Imports, code.FunctionSign, code.ClassSign(context, base, implemens) gibi tanımlanabilir fonksiyonlar
_ Language, Runtime gibi yapıları ayrıca geliştirilmeli ve bağımsız tanımlanabilmeli. Platform da bunlara referans vermeli
_ Library ve language çözümleri için standardizasyon section lar aracılığıyla genişletilebilir, entity section, dto section, endpoint section, repository section, command section , command handler section gibi yapılar hazırlanmalı, burda tanımlı standartlar template te kullanılabilir olmalı
_ Template function olarak platform ve dil için kolaylaştırıcı çözümler üretilebilir, örneğin switch case, yada method gibi, mesela method sign function method nesnesi alır ve print eder, dil versiyona göre farklı çıktılar sağlayabilir
_ Section layer tarafından çağırılır, class ile tanır. Bu hazır paket dışında ilgili class ile kullanıcının kendi özel section u nu tanımlamasını sağlar. Özellikle endpoint tanımlama gibi uygulamaya özel yapılar section class ile çoğaltılabilir.
_ Resource ta section ilişkilendirme, hem name hemde selector (label, class) ile sağlanabilir
_ Pars project [project _name] analyze, check
_ Pars project scann (harici proje tanıma çözümleme
_ Apply operasyonu sonra ki bir versiyona bırakılabilir, şimdilik commit , submit , store, register, install gibi bir komut new yerine kullanılabilir
_ Provider gereksinimleri için local environment kullanılabilir, internal providers yerine host providers kullanılabilir, bu hem opsiyon olarak bulunur hemde ilk versiyonda provider management gerekmez
_ Section yapısı mevcut haliyle kullanılmaya devam edecek bunun dışında referans yoluyla section modellerde yine bir section başlığında tanımlanarak bağlanabilecek
_ Database, meyhaneci kurt için repository ve servis katmanı ayrı yapılandırılabilir servis katmanında referansların yönetimi sağlanabilir örneğin resource ya da project'in words space'i yamuldu üzerinden değil de doğrudan servis aşamasına servis katmanında modele tanımlanabilir böylece tabanından alınmış nesne üzerinden sağlanır bu durumda da yamul üzerinde doğrudan bir tanımlama yapılmamış olur referans adı veya selektör dışında çözülmüş olabilir

Packages
_ Library paketleri hazırlanmalı, örneğin orm paketi ile c# ta entity framework one to many relations yazdıran fonksiyon hazırlanmalı. Entity hazırlanması için fonksiyon olmalı gibi.
_ Mapper çözümü sağlanmalı, resource lar veya section lar için ilgili frameworke uygun Mapping kodları hazırlanmalı \* Put, patch gibi http verbler için kullanıma dair içerikler hazırlanabilir, resource, section ve template nasıl hazırlanır gibi
related projects

-   https://www.youtube.com/watch?v=fth5uiRSCfI

---

packing
cd src
set GOOS=windows; set GOARCH=amd64; go build -o .\bins\pars-windows-x64.exe .\pars.go
mv ..\bins\pars-windows-x64.exe ..\packing\msi\Source\
cd ..\packing\msi\Setup\

---

pars project list --fields=name gibi sadece gösterilecek bilgiler seçilebilir
filter ve fields gibi yapılar ile belirli filtreyle listelenen kayıtlar delete comutu ile hızlı silinebilir
`pars project remove $(pars project list --field=Name --filter=create-date:04.01.2024)` gibi

---

Marketing

-   Potansiyel contributer'lara ulaşmak
    Github contribution yapan kullanıcılara ulaşmak pars'ı tanıtmak
    Trending Repository'leri takip etmek
-   Use Case, Kısa trick short'lar yayınlanması
    Youtube, instagram gibi mecrelarda
-   Use Case yazı-makale paylaşılması
    Medium, linkedin, blog'da

Contributing Rules

-   Run unittests
-   Run lints

---

Branching rules
PR rules
Issue
rules/templates
labels

---

Pars (pars cli)
Pars Dashboard (Management UI)
Pars Server
Pars Cron (Scheduling)
Pars Worker

Redis
MongoDB
Postgresql

---

repository'lerde belirlenen klasör dışında değişiklikler PR aşamasında reddedilmeli (pars için src dışında (.github, .gitignore vs korumak için))

VERSION dosyasında
iteration start-end date bilgileri yer alabilir?

Content managment
https://www.youtube.com/watch?v=8_9mEzSjGY8
https://www.youtube.com/watch?v=HnRrEofRMUg
https://www.youtube.com/watch?v=Rxw9INjRP38
https://www.youtube.com/watch?v=R2onqpgDwvc
https://www.youtube.com/watch?v=8BedriWb7Lw&t=876s
https://www.youtube.com/watch?v=MpS_kpxhHP4
https://www.youtube.com/watch?v=lqXWVzWkkyc

---

Basic Dotnet WebAPI Project

NodeJS Microservice Project

Dotnet Console Application

Angular Frontend Project

Microservices Architecture with Multiple Projects

Cross-Platform Mobile Application (Xamarin)

Monolithic Dotnet WebApp with Multiple Layers

Go Lang Backend Service

NodeJS Project with Custom NPM Packages

Dotnet Library for Shared Functionality

Complex Configuration with External Dependencies

Project with Specific Workspace Configuration

Dotnet Project with NuGet Package Management

Advanced Dotnet Project with Versioning

NodeJS Project with Environment Labels

---

Versioning

-   Dead Version: çözülemeyen hatalardan dolayı Publish phase geçememiş, zamanında çıkamamış ve bir sonra ki sürüm ile birleştirilmiş versiyon
-   Deferred Version: Geliştirme sürecinin plansız uzamasından kaynaklı olarak Development phase'dan çıkamamış bir sonra ki iterasyon'a sarkmış version
-   On-Time Version: Zamanında ve başarılı olarak çıkmış version
-   Consolidated Version: Dead version'ları içeren birleşik version
-   Merged Version: Bir biri ile ilişkili feature'lar sebebiyle bir sonra ki geliştirme'yi beklemesi gereken version. planlı bekletme. Development fazından release faza geçer fakat aynı release</version>'da toplanır
-   Unified Version: Geliştirme sürecinin planlı uzamasından kaynaklı olarak Development phase'dan çıkamamış bir sonra ki iterasyon'a sarkmış version

---

1. On-Time Version (Zamanında Çıkan Versiyon)
   Tanım: Planlanan tarihte, zamanında ve başarıyla çıkan sürüm. Geliştirme, test ve yayınlama süreçleri sorunsuz ilerler.
   Senaryo: Planlanan iterasyon sonunda eksiksiz test edilen ve belirlenen tarihte yayınlanan versiyon.
   Alternatif İsimler: Scheduled Version, Planned Version.

2. Delayed Version (Gecikmiş Versiyon)
   Tanım: Geliştirme veya test süreçlerinde yaşanan aksaklıklardan dolayı planlanan tarihten daha sonra çıkan sürüm.
   Senaryo: Özellikler veya geliştirmeler planlandığı gibi tamamlanamamış ve bir sonraki iterasyona sarkmış bir sürüm.
   Alternatif İsimler: Postponed Version, Extended Version.

3. Deferred Version (Ertelenmiş Versiyon)
   Tanım: Plansız uzama nedeniyle geliştirme aşamasında kalmış ve planlanan iterasyon içinde tamamlanamayan sürüm. Bir sonraki iterasyonda tamamlanması beklenir.
   Senaryo: Geliştirme aşamasında bazı özelliklerin yetişmemesi sonucu, bir sonraki iterasyona ertelenen versiyon.
   Alternatif İsimler: Carried-Over Version, Rolled-Over Version.

4. Dead Version (Ölü Versiyon)
   Tanım: Yayınlama aşamasında kritik sorunlar nedeniyle başarılı bir şekilde yayınlanamayan ve tamamen iptal edilen veya bir sonraki versiyona entegre edilen sürüm.
   Senaryo: Yayın öncesi aşamada ciddi hatalar bulunan, bu nedenle bir sonraki versiyon ile birleştirilen ve tek başına yayınlanmayan sürüm.
   Alternatif İsimler: Abandoned Version, Cancelled Version.

5. Consolidated Version (Birleştirilmiş Versiyon)
   Tanım: Dead versiyonları ve bir sonraki versiyonu bir araya getirerek çıkartılan birleştirilmiş sürüm. Hatalar nedeniyle önceki sürümlerle birleşip yayınlanan versiyon.
   Senaryo: Yayınlanamayan bir dead version’ın hataları çözülüp, bir sonraki sürümle birleştirilip yayınlanması.
   Alternatif İsimler: Combined Version, Aggregated Version.

6. Merged Version (Birleştirilen Versiyon)
   Tanım: Birbiriyle ilişkili özellikler (features) nedeniyle planlı bir şekilde, bir sonraki geliştirme sürecini bekleyerek birleştirilen sürüm. Planlı bir bekletme ve birleştirme süreci ile çıkartılır.
   Senaryo: İlgili özelliklerin birbirini tamamlama gerekliliği nedeniyle, geliştirmeleri tamamlanmış olsa da aynı sürüm içinde planlı bir şekilde birleştirilip yayınlanan versiyon.
   Alternatif İsimler: Planned Merge Version, Synchronized Version.

7. Hotfix Consolidated Version (Acil Düzeltme İle Birleştirilen Versiyon)
   Tanım: Yayınlanan bir sürümde ciddi bir hata tespit edilip, hızlıca düzeltilen ve bu düzeltmenin bir sonraki sürüme dahil edildiği durum.
   Senaryo: Yayınlanan bir versiyonda kritik bir hata bulunduğunda, hızlıca hotfix uygulanır ve bu hotfix bir sonraki versiyon ile birlikte yayınlanır.
   Alternatif İsimler: Hotfix Roll-Up Version, Urgent Consolidated Version.

8. Deferred Consolidated Version (Ertelenmiş ve Birleştirilmiş Versiyon)
   Tanım: Geliştirme süreçlerinde yaşanan plansız gecikmeler sonucu zamanında çıkamayan ve bir sonraki versiyon ile birleştirilmiş sürüm.
   Senaryo: Plansız gecikmeler nedeniyle bir sonraki iterasyona sarkan ve orada başka bir versiyonla birleştirilip çıkarılan sürüm.
   Alternatif İsimler: Deferred Combined Version, Rolled-Over Consolidated Version.

9. Skipped Version (Atlanan Versiyon)
   Tanım: Geliştirme sürecinde bazı özellikler planlandığı gibi ilerlememiş ve tamamen iptal edilerek sonraki sürüme geçilmiş versiyon. Yayınlanmamış ama bir sonraki versiyona da dahil edilmemiş sürüm.
   Senaryo: Planlanan sürümde bazı özellikler yeterince iyi çalışmadığı için bu versiyon tamamen iptal edilir ve geliştirmeye bir sonraki versiyonla devam edilir.
   Alternatif İsimler: Dropped Version, Discarded Version.

Branches

-   dev
-   feature/<feature-name>
-   release/<version>
-   fix/<hotfix-name>
-   bug-fix/<bugfix-name>
-   main
    -   tag

Phases

-   Planning
    -   2w iteration
-   Development
    -   2w iteration
    -   otomatik olarak 2 haftada bir hazırlanır. iterasyon başında VERSION dosyası sıfırlanır (dev-number ve working-version)
    -   sumsha > CHECKSUMS dosyasına uygulanabilir phase sonunda
-   Test
    -   2w iteration
    -   dev phase tamamlandıktan sonra release phase'e geçilir, bu esnada dev branch'ten release/<version> branch oluşturulur varsa dev'den merge edilir
    -   ilk hafta alpha sürümü ile testler yapılır
    -   ikinci hafta beta sürümü ile testler yapılır
-   Staging
-   Release
    -   2w iteration
    -   preview sürümü ile testler yapılır
-   Production
    -   iteration independent. sonsuz

Release Channels

-   Dev
    -   dev branch'ten hazırlanır.
    -   ihtiyaç halinde manuel tetiklenir
    -   development phase süresince alınabilir
    -   dev number VERSION dosyasında saklanır
    -   debug symbols içerir
    -   stage tanımlanmaz veya "none" atanır
-   Alpha
    -   release branch'ten hazırlanır
    -   ihtiyaç halinde manuel tetiklenir
    -   test phase süresince alınabilir
    -   alpha number VERSION dosyasında saklanır
    -   debug symbols içerir
    -   stage "test" tanımlanır
-   Beta
    -   release branch'ten hazırlanır
    -   ihtiyaç halinde manuel tetiklenir
    -   test phase süresince alınabilir
    -   beta number VERSION dosyasında saklanır
    -   debug symbols içermez
    -   stage "test" tanımlanır
-   Preview
    -   release branch'ten hazırlanır
    -   ihtiyaç halinde manuel tetiklenir
    -   release phase süresince alınabilir
    -   preview number VERSION dosyasında saklanır
    -   debug symbols içermez
    -   stage "final" tanımlanır
-   Stabil
    -   main branch'ten git tag ile hazırlanır
    -   manuel tetiklenir
    -   publish phase süresince alınabilir
    -   patch number VERSION dosyasında saklanır
    -   debug symbols içermez
    -   stage "final" tanımlanır

Iterasyon 1: Planning
Iterasyon 2: Analyse
Iterasyon 3: Design
Iterasyon 4: Implementation (dev) dev
Iterasyon 5: Test (test) alpha, beta
Iterasyon 6: Production (release, main) preview, rc, stabil

---

pars info
şu bilgiler eklenebilir
Source: Official Repository
Method yada Installer: APT (Debian/Ubuntu)
Package Version: 1.0.0
Installation Date: 2024-09-11

---

Repositories
Pars-Samples
Pars-Packages

---

bintray
JFrog Artifactory
OBS
Launchpad

RPM: Redhad
APT: Lİnux
Ports: BSDs

Installers
msi, sfx, pip, npm, npx, yarn

Tools
dotnet tools (https://learn.microsoft.com/tr-tr/dotnet/core/tools/global-tools)
nuget
go-get

Markets
MacOS App Market
Windows App Market
VS Code Market

---

Template Functions

-   string/array için addprefix
-   dirname: file/path'ten full directory alma
-   filename: file'den dosya adı alma
-   extention: file'den uzantı alma
-   date functions, pars, format vs
-   file functions (read/write)
-   shell functions (bash, powershell)

---

Versions

Pro Version

Enterprise version

-   Collaboration

SaaS

Build Service

-   Build Remotely and push to repo

---

Hata yönetimi: Hata mesajları ve kullanıcı bilgilendirme mesajları iyileştirilmeli
Hata mesajları pars00016 gibi unqiue bir kodla belirtilmeli, internet aramalarında hızlı erişim sağlanması için
dökümanda hata mesajları ile ilgili detaylar paylaşılmalı

---

https://github.com/goreleaser/goreleaser
make commands:
provision, build, bundle, archieve, package, upload, install, uninstall, backup, restore etc, artifacts, metadata, release, specs
Provision Project: kurulacak programların, package manager'ların, bağımlılıkların tnaımlandığı proje tipi. bu yapı ile makina kurulumlaının hızlandırılması ve iac olarak yönetilmesi mümkün olur (https://www.youtube.com/shorts/aStfdeJCQFk)

---

parsdevkit@gmail.com
info@parsdevkit.net
parsgrid.io@yandex.com

<!-- enc -->

cGFyc2RldmtpdC5uZXQNCmZ0cDogcGFyc2RldmtpdCBOTHVsaX5HcjFmY3kNCmdtYWlsIHBhcnNkZXZraXQgcXRnTU5MVHh3YU03DQptaWNyb3NvZnQgcGFyc2RldmtpdCBxdGdNTkxUeHdhTTcNCmdpdGh1YjogcGFyc2RldmtpdCBxdGdNTkxUeHdhTTcNCmZpcnN0IFBBVDogZ2hwX3hhUlVoVVpHM25zT1BNZ0gybm96MFMxNjBqNFhRQTJxRTR3VQ0KbnBtanM6IHBhcnNkZXZraXQgcXRnTU5MVHh3YU03DQpmaXJzdCBQQVQ6IG5wbV9nZG1QVzN0SVNGU1JpYUQ3ekFHdDY5QmljOHNjbGgwNkxEazUNCmRlYmlhbjogKGdtYWlsKSBxdGdNTkxUeHdhTTcNCnVidW50dTogcGFyc2RldmtpdCBxdGdNTkxUeHdhTTcNCmxhdW5jaHBhZDogKHVidW50dSBhY2NvdW50KQ0Kc25hcGNyYWZ0LmlvICh1YnVudHUgYWNjb3VudCkNCnNvdXJjZWZvcmdlOiBwYXJzZGV2a2l0IHF0Z01OTFR4d2FNNyBodHRwczovL3BhcnNkZXZraXQuc291cmNlZm9yZ2UuaW8NCmFtYXpvbjogKGdtYWlsKXBhcnNkZXZraXQgcXRnTU5MVHh3YU03DQpzbGFjazogKGdtYWlsKQ0KcmVkZGl0OiAoZ21haWwpIHF0Z01OTFR4d2FNNyBodHRwczovL3d3dy5yZWRkaXQuY29tL3UvcGFycy1kZXYta2l0DQp5b3V0dWJlIChnbWFpbCkNCmxpbmtlZGluIChnbWFpbCkNCmRvY2tlciAoZ21haWwsZ2l0aHViKQ0KUEFUIDogZGNrcl9wYXRfZlkxMkhsOUFQbHFXWERjV1MyTUhnenNGd2YwDQptYXN0b2Rvbi5zb2NpYWw6IChnbWFpbCkgcXRnTU5MVHh3YU03IGh0dHBzOi8vbWFzdG9kb24uc29jaWFsL0BwYXJzZGV2a2l0DQptYWtlLmNvbSAoZ21haWwpDQp6YXBpZXIuY29tIChnbWFpbCkNCmluc3RhZ3JhbQ0KZmFjZWJvb2sNCnR3aXR0ZXINCnRlbGVncmFtDQpjb21tdW5pdHkuY2hvY29sYXRleS5vcmcgcGFyc2RldmtpdCBxdGdNTkxUeCF3YU03DQpwYXRyZW9uOiAoZ21haWwpIGh0dHBzOi8vd3d3LnBhdHJlb24uY29tL1BhcnNEZXZLaXQvY3JlYXRvcnMNCm9wZW5jb2xsZWN0aXZlIChnbWFpbCkgcXRnTU5MVHghd2FNNw0Ka28tZmkuY29tIChnbWFpbCkgaHR0cHM6Ly9rby1maS5jb20vcGFyc2RldmtpdA0KDQptZW50b3JzLmRlYmlhbi5uZXQgcGFyc2RldmtpdEBnbWFpbCBxdGdNTkxUeHdhTTcNCg0KZ2l0aHViIGFobWV0dHNvbmVyDQpkZXYtbWFjaGluZS1wYXQ6IGdocF9DTnRBTXdUSnZuNlk0OGg1NGRVQkFUemkwUDNQSVE0RjQ2Znc=

---

docker build --tag parsdevkit/pars:latest .

docker pull parsdevkit/pars:latest

docker run --name pars --detach --volume D:\pars-docker\data:/var/lib/pars/data --volume D:\pars-docker\workspaces:/home parsdevkit/pars:latest

docker exec -it pars /bin/sh

## docker rm pars --force

cd src
go build -o ../dist/pars.exe .\pars.go

---

bundle for package

Yazılımınızın Kendi Terminalini Açması (çift tıkladığında platforma göre default terminal'i açması)

msi multiple app in installed apps to remove or modify window

---

Issue Creation Rules: (task, feature, bug)

-   add task into description if necessary
-   follow templates
-   set labels
-   set project

Backlog Refinement

-   set priority, iteration, milestone if required
-   assign to developer if needed

Sprint Planning

-   team leader decide to assign or developer can pick anyone
-   team can set issue rate (work height)

Contributing

-   create feat/fix branch in your forked repo
-   delete feature branch in your forked repo (optional)
    git push origin --delete <branch_name>
-   pr message
    açık ve net şekilde yapılan işlemi tanımlayın

Reviewer

-   for complete issue use footer note in PR
    Fixes #<issue_number>
    Closes #<issue_number>
    Resolves #<issue_number>

-   Squash and Merge: Ensure the conventional commit standartds applied in new pr commit

---

Branching

main, test, ve release için Read-Only
dev allows only PR from forked repo

issue types
feature
bug
task
Improvement
Question

teams
core
infra (github rules)

---

-   graphql code generator : c#, vue.js, go gibi ypaılar için hazır template'ler
-   sql code generator

input:
_ json
_ xml
_ database table
_ sql command \* graphql
tools:

-   vscode gibi ideler'de
    -   seçili class, json, graphql object, proto nesnesi kullanarak kod üretmek
    -   json, class, table, sql vs oluşturma

api: https://github.com/nette/php-generator olduğu gibi pars eski
