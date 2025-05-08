make build.cmake.{VERSION}
make documents.generate.html
make artifacts.prepare
make windows.arm64.binary.build
make windows.arm64.binary.checksum
make windows.arm64.binary.archive
make windows.arm64.binary.artifact

---

make package modulunde deb var bin için kullnaılıyor sdeb source pack için rpm bin, srpm source rpm için olacak şekilde `s` ler hazırlanmalı

lp pass: qtgMNLTxwaM7

gpt scribe api key

<!-- enc -->

c2stc3ZjYWNjdC0teWRRaEhOQmd5SWlEOGE2RF9sdWRDbXJZVFZ5SnhRTlhQcFlxYTVxNUxRbHJmc0xISGluTTlsOWtjVjRfU1QzQmxia0ZKV2lGY3pTWlRRTWlEb2lVS1NuMU1xVzdrR3ZMVVcyU2xwekNsNjR0M0RZWlZOMjZ1MDRobkw2OVdiR3hVWUE=

makecab pars.exe pars.cab
wix build -o pars.msi pars.wxs

make release.version PREFIX=v VERSION=1.0.0 CHANNEL=stabil REVISION=15 DATE=2024.10.16
make release VERSION=v1.0.1 CHANNEL=test REVISION=179 DATE=2024.10.16 MESSAGE=
make release.notes.new MESSAGE=

---

make build.cmake VERSION=v1.0.3-test.120

make build.cmake.linux # linux host için make komutları oluşturur, otomatik version (git tag veya default), build/current klasörüne
make build.cmake.linux VERSION=v1.0.0-beta.3 # linux host için make komutları oluşturur, belirtilen VERSION için, build/current klasörüne
make build.cmake.linux.v1.0.0-beta.3 # linux host için make komutları oluşturur, belirtilen VERSION için, build/VERSION klasörüne
make build.cmake.macos # macos host için make komutları oluşturur, VERSION argümanı veya VERSION klasör yapısını destekler
make build.cmake.windows # windows host için make komutları oluşturur, VERSION argümanı veya VERSION klasör yapısını destekler

make build.binary # VERSION belirtilmişse build/VERSION, belirtilmemişse build/current kullanır, mevcut host makine os'u için host'un arch'i için binary build alır
make build.binary.arm64 # VERSION belirtilmişse build/VERSION, belirtilmemişse build/current kullanır, mevcut host makine os'u için belirtilen arch için binary build alır
make build.binary.darwin.arm64 # VERSION belirtilmişse build/VERSION, belirtilmemişse build/current kullanır, belirtilen os için belirtilen arch için binary build alır
make build.binary.openbsd.all # VERSION belirtilmişse build/VERSION, belirtilmemişse build/current kullanır, belirtilen os için desteklenen tüm arch için binary build alır
make build.binary.all # VERSION belirtilmişse build/VERSION, belirtilmemişse build/current kullanır, desteklenen tüm os ve arch'ler için binary build alır

make build.deb.package.setup # VERSION belirtilmişse build/VERSION, belirtilmemişse build/current kullanır, debian based linux host için deb paket oluşturmada gerekli paketlein kurululmunu sağlar. x86, x86_64, arm ve arm64 mimari destekler. "all" belirtilmişse tüm mimarilere uygun olması için `any` kullanır
make build.deb.package.all.configuration # VERSION belirtilmişse build/VERSION, belirtilmemişse build/current kullanır, deb paket için configuration dosyalarının hazırlanmasını sağlar
make build.deb.package.all.payload # VERSION belirtilmişse build/VERSION, belirtilmemişse build/current kullanır, deb paket için source code'un hazırlanmasını sağlar
make build.deb.package.all.package # VERSION belirtilmişse build/VERSION, belirtilmemişse build/current kullanır, deb paket oluşturulmasını sağlar
make build.deb.package.all

make build.deb.package.all.sign.gpg

make build.deb.package.all.push.ppa

setup.deb
build.deb
sign.deb
sign.deb.binary-package
sign.deb.source-package
push.deb
archive.deb
package.deb
install.deb
uninstall.deb
upgrade.deb

tar -xvf pars_1.0.0-beta.1.tar.xz -C .
cd pars
dpkg-buildpackage -b
sudo dpkg -i pars_1.0.0-beta.1_amd64.deb

---

make changelog.entry.add
TAG?: default from git and channel information
message\*:

make changelog.clear
TAG?: default from git and channel information

    TAG?: default from git and channel information
    OS?: default from host machine os
    ARCH?: default from host machine arch

setup
config
build
publish
add.package
remove.package
add.plugin
remove.plugin
run
payload
test
archive
deploy
publish
installed
uninstall
upgrade
doc
release
backup
restore
clean
rollback
lint

environment.init.linux.build
environment.init.linux.deb-packing
environment.init.linux.snap-packing
environment.init.windows.choco-packing

artifacts.create.from-source
artifacts.create.from-binary
artifacts.create.from-binary
artifacts.create.from-deb-package
artifacts.create.from-snap-package
artifacts.create.from-rpm-package
artifacts.create.from-pkg-package

choco.push.package
snap.push.package
obs.push.binay
obs.push.deb-package
launchpad.ppa.push.binary
launchpad.ppa.push.deb-package
github.push.binary
github.push.deb-package
docker-hub.push.image
aws.ecr.push.image
kubernetes.deploy.image

install.binary
install.remote.deb-package
install.local.deb-package
install.remote.snap-package
install.local.snap-package
install.choco-package
install.rpm-package
install.pkg-package

---

Paket yapısının hazırlanması

    /etc/pars/                 			# Yapılandırma dosyaları
    /var/lib/pars/             			# Kalıcı veriler
    /var/log/pars/             			# Log dosyaları
    /usr/bin/pars              			# Uygulama çalıştırılabilir dosyası
    /usr/lib/pars/             			# Kütüphaneler
    /usr/share/pars/           			# Paylaşılan dosyalar (simgeler, belgeler)
    /usr/share/doc/myapp/               # Genel belgeler
    /usr/share/doc/myapp/README.md      # README dosyası
    /usr/share/doc/myapp/CHANGELOG      # Değişiklik günlüğü
    /usr/share/doc/myapp/LICENSE        # Lisans bilgileri
    /usr/share/doc/myapp/examples/      # Örnek dosyalar
    /usr/share/man/man1/myapp.1.gz      # Man sayfası
    /tmp/                      			# Geçici dosyalar

deb install sonrası " /var/lib/pars/data" erişim sebebiyle pars komutlarının sudo'suz kullanımında hata vermesi
lp:ppa multiple series for package (noble, jammy etc)
amd, amd64, arm, arm64 mimariler ile paket oluşturulması.
paket kurulum testlerinin ypaılması
paket oluşturma için gpg secret phrase istememeli otomasyon kullanımında

---

.
┣ etc/
┃ ┗ pars/
┣ usr/
┃ ┣ bin/
┃ ┃ ┗ pars
┃ ┣ lib/
┃ ┃ ┗ pars/
┃ ┗ share/
┃ ┣ doc/
┃ ┃ ┃ ┗ pars/
┃ ┃ ┃ ┗ user_docs/
┃ ┃ ┃ ┃ ┣ faq.md
┃ ┃ ┃ ┃ ┣ installation.md
┃ ┃ ┃ ┃ ┣ troubleshooting.md
┃ ┃ ┃ ┃ ┗ usage.md
┃ ┗ pars/
┗ var/
┣ cache/
┃ ┃ ┗ pars/
┣ lib/
┃ ┃ ┗ pars/
┗ log/
┃ ┗ pars/

---

C:\
├── Program Files\
│ └── pars\
│ ├── bin\
│ │ └── pars.exe
│ ├── lib\
│ │ └── pars\
│ └── doc\
│ └── user_docs\
│ ├── faq.md
│ ├── installation.md
│ ├── troubleshooting.md
│ └── usage.md
├── ProgramData\
│ └── pars\
│ ├── config\ (Ayarlar ve konfigürasyon dosyaları)
│ │ └── config.yaml
│ ├── cache\ (Genel cache dosyaları)
│ │ └── cache.db
│ ├── lib\ (Ortak bağımlılıklar)
│ └── logs\ (Genel log dosyaları)
│ └── log.txt
└── Users\
 └── <username>\
 ├── AppData\
 │ ├── Local\
 │ │ └── pars\
 │ │ ├── cache\ (Kullanıcıya özel cache)
│ │ ├── config\ (Kullanıcıya özel ayarlar)
│ │ └── logs\ (Kullanıcıya özel loglar)
│ └── Roaming\
 └── Documents\
 └── pars\
 └── projects\ (Kullanıcı projeleri)
**\*\***\*\*\***\*\*** MACOS

/Applications\
└── pars.app\
 ├── Contents\
 │ ├── MacOS\
 │ │ └── pars (Uygulamanın çalıştırılabilir dosyası)
│ ├── Resources\
 │ │ └── user_docs\
 │ │ ├── faq.md
│ │ ├── installation.md
│ │ ├── troubleshooting.md
│ │ └── usage.md
│ └── Info.plist (Uygulama bilgileri ve ayarları)
├── /Library\
│ └── Application Support\
│ └── pars\
│ ├── config\ (Ortak konfigürasyon dosyaları)
│ │ └── config.yaml
│ ├── cache\ (Genel cache dosyaları)
│ │ └── cache.db
│ ├── lib\ (Ortak bağımlılıklar)
│ └── logs\ (Genel log dosyaları)
└── /Users\
 └── <username>\
 ├── Library\
 │ ├── Application Support\
 │ │ └── pars\
 │ │ ├── cache\ (Kullanıcıya özel cache)
│ │ ├── config\ (Kullanıcıya özel ayarlar)
│ │ └── logs\ (Kullanıcıya özel loglar)
└── Documents\
 └── pars\
 └── projects\ (Kullanıcı projeleri)

\***\*\*\*\*** BSD
/usr/local\
└── pars\
 ├── bin\
 │ └── pars (Uygulamanın çalıştırılabilir dosyası)
├── lib\
 │ └── pars\
 └── share\
 └── doc\
 └── pars\
 └── user_docs\
 ├── faq.md
├── installation.md
├── troubleshooting.md
└── usage.md
/var\
└── cache\
 └── pars\
/var\
└── log\
 └── pars\
 └── log.txt
/var\
└── lib\
 └── pars\
/etc\
└── pars\
 └── config.yaml

go build sırasında stage tanımlanır ve buna göre file structure şekillenir
'parsdevkit.net/core/utils.stage=[none, dev, test, prod]'
Project run 'da project root code base inşaa edilir
dev binary'de proje klasörleri executable'ın bulunduğu dizinde inşaa edilir. PARS_PROJECT_ROOT tanımlı ise bu dizini proje dizini kabul eder
test stage'de binary'nin bulunduğu dizinde inşaa edilir
prod stage'de ilgili os için belirlenen dir yapısı kullanılır
PARS_PROJECT_ROOT env var tanımlı ise proje dizini olarak bunu kullanır (birden çok proje olduğunda birincil proje tanımlamak için kullanılır) yoksa her proje kendi base'ini kullanır
Code Base, dev ortamında alınan hızlı build'lerde build alınan code base'in yada env' adresini referans eder. projeden build alındığını belirtir
dev ortamı için build alınmışsa parsdevkit.net/core/utils.stage bilgisi içerir ve kullanılacağı yerde code base env olarak tanımlı ise onu kullanır, yoksa boştur
test ve prod ortamları için build alınmışsa parsdevkit.net/core/utils.stage bilgisi içerir ve kullanılacağı yerde code base olmayacağı anlamına gelir

---

.snap Paketini Hazırlama ve Yükleme
? sudo dnf install rpm-build rpmdevtools

    sudo apt update
    sudo apt install snapd -y
    sudo apt install gnome-keyring -y
    sudo apt install
    sudo snap install snapcraft --classic

    sudo lxd init

    # `src` dizininde çalışıyoruz ve `packages` dizinini ana dizinde tutuyoruz
    mkdir -p ./packages/linux/snap/source
    mkdir -p ./packages/linux/snap/source/amd32/
    mkdir -p ./packages/linux/snap/source/amd64/
    mkdir -p ./packages/linux/snap/source/arm32/
    mkdir -p ./packages/linux/snap/source/arm64/


    cp ./bin/linux/amd32/pars ./packages/linux/snap/source/amd32/
    cp ./bin/linux/amd64/pars ./packages/linux/snap/source/amd64/
    cp ./bin/linux/arm32/pars ./packages/linux/snap/source/arm32/
    cp ./bin/linux/arm64/pars ./packages/linux/snap/source/arm64/

PACKAGE_NAME="pars"
VERSION="1.0.0"
SUMMARY="A brief description of the pars application."
DESCRIPTION="A brief description of the pars application."

cat <<EOF > ./packages/linux/snap/source/snapcraft.yaml
name: $PACKAGE_NAME
version: '$VERSION'
summary: $SUMMARY
description: |
$DESCRIPTION

grade: stable
confinement: strict
base: core22

parts:
$PACKAGE_NAME:
plugin: dump
source: .
source-type: local
stage: - source/amd32/pars - source/amd64/pars - source/arm32/pars - source/arm64/pars
prime: - source/amd32/pars - source/amd64/pars - source/arm32/pars - source/arm64/pars
build-attributes: - no-system-libraries

architectures:

-   build-on: [amd64]
    build-for: [amd64]
-   build-on: [arm64]
    build-for: [arm64]
-   build-on: [i386]
    build-for: [i386]
-   build-on: [armhf]
    build-for: [armhf]

apps:
$PACKAGE_NAME:
    command: bin/$SNAP_ARCH/$PACKAGE_NAME
plugs: [network, home]

# layout:

# /usr/bin/$PACKAGE_NAME:

# bind-file: $SNAP/bin/$SNAP_ARCH/$PACKAGE_NAME

EOF

cd ./packages/linux/snap/source
snapcraft

test
sudo snap install pars_1.0_amd64.snap --dangerous

snapcraft login

---

name: pars
version: '1.0'
summary: A short summary of my app
description: |
A longer description of my app.

grade: stable
confinement: strict
base: core22

architectures:

-   build-on: [amd64]
    run-on: [amd64]
-   build-on: [i386]
    run-on: [i386]

parts:
pars:
plugin: dump
source: .
source-type: local
filesets:
amd64-binaries: - amd64/pars
i386-binaries: - amd32/pars # i386 için amd32 klasörünü tanımlıyoruz
organize:
amd64/pars: usr/bin/pars-amd64
amd32/pars: usr/bin/pars-i386 # i386 binary'sini doğru konuma yerleştiriyoruz
override-build: |
if [ "$(uname -m)" = "x86_64" ]; then
cp amd64/pars $SNAPCRAFT_PART_INSTALL/usr/bin/pars
      elif [ "$(uname -m)" = "i686" ] || [ "$(uname -m)" = "i386" ]; then
cp amd32/pars $SNAPCRAFT_PART_INSTALL/usr/bin/pars
fi

apps:
pars-amd64:
command: pars-amd64
plugs: - network
architecture: amd64 # Bu komut sadece amd64 mimarisinde çalışır

pars-i386:
command: pars-i386
plugs: - network
architecture: i386 # Bu komut sadece i386 mimarisinde çalışır

---

name: pars
version: '1.0'
summary: A short summary of my app
description: |
A longer description of my app.

grade: stable
confinement: strict
base: core22

# architectures:

# - build-on: [amd64]

# build-for: [amd64]

# - build-on: [amd64, arm64]

# build-for: [arm64]

# - build-on: [amd64, armhf]

# build-for: [armhf]

parts:
pars:
plugin: dump
source: .
source-type: local

apps:
pars:
command: pars
plugs: - network

---

PACKAGE_NAME="pars"
VERSION="1.0.0"
SUMMARY="A brief description of the pars application."
DESCRIPTION="A brief description of the pars application."

cat <<EOF > ./packages/linux/snap/source/snapcraft.yaml
name: $PACKAGE_NAME
version: '$VERSION'
summary: $SUMMARY
description: |
$DESCRIPTION

grade: stable
confinement: strict
base: core22

parts:
$PACKAGE_NAME:
plugin: dump
source: .
source-type: local
stage: - source/amd32/pars - source/amd64/pars - source/arm32/pars - source/arm64/pars
prime: - source/amd32/pars - source/amd64/pars - source/arm32/pars - source/arm64/pars
build-attributes: - no-system-libraries

architectures:

-   build-on: [amd64]
    build-for: [amd64]
-   build-on: [arm64]
    build-for: [arm64]
-   build-on: [i386]
    build-for: [i386]
-   build-on: [armhf]
    build-for: [armhf]

apps:
$PACKAGE_NAME:
    command: bin/$SNAP_ARCH/$PACKAGE_NAME
plugs: [network, home]

# layout:

# /usr/bin/$PACKAGE_NAME:

# bind-file: $SNAP/bin/$SNAP_ARCH/$PACKAGE_NAME

EOF

cd ./packages/linux/snap/source
snapcraft

test
sudo snap install pars_1.0_amd64.snap --dangerous

snapcraft login

---

sudo dnf install -y make cmake gcc gcc-c++ golang

---

make build.cmake VERSION=v1.0.3-test.118
make build.cmake VERSION=v1.0.3-test.118
make build.binary
make build.binary OUTPUT=tmp/usr/bin/118/pars

make build.cmake.v1.0.3-test.119
make build.binary VERSION=v1.0.3-test.119
make build.binary VERSION=v1.0.3-test.119 OUTPUT=tmp/usr/bin/119/pars

---

lp pass: qtgMNLTxwaM7

make build.cmake VERSION=v1.0.3-test.120

make vcs.git.move.dev-to-test

make build.deb.package.setup
make build.deb.package.all.configuration
make build.deb.package.all.payload
make build.deb.package.all.package
make build.deb.package.all.sign.gpg
make build.deb.package.all.push.ppa

make build.snap.package.setup
make build.snap.package.all.configuration
make build.snap.package.all.payload
make build.snap.package.all.package
make build.snap.package.all.sign.gpg
make build.snap.package.all.push.ppa

make build.msi.package.setup
make build.msi.package.x86_64.configuration
make build.msi.package.x86_64.payload
make build.msi.package.x86_64.package

make build.rpm.package.setup
make build.rpm.package.all.configuration
make build.rpm.package.all.payload
make build.rpm.package.all.package
make build.rpm.package.all.sign.gpg
make build.rpm.package.all.push.ppa

wix build -o pars.msi D:\AS\W\Pars\pars\dist\v1.0.3-test.120\windows\ins\msi\all\pars\config.wxs

https://chatgpt.com/c/6751989b-0a0c-8004-a99d-64d677483c5f
make build.choco.package.setup
make build.choco.package.all.configuration
make build.choco.package.all.payload
make build.choco.package.all.package
make build.choco.package.all.sign.gpg
make build.choco.package.all.push.ppa

https://chatgpt.com/c/675a8261-ab14-8004-8df9-2d9352112d1e
make build.brew.package.setup
make build.brew.package.all.configuration
make build.brew.package.all.payload
make build.brew.package.all.package

---

C:\Program Files (x86)\oh-my-posh\bin;C:\Users\AHMET SONER\.cargo\bin;C:\Users\AHMET SONER\AppData\Local\Programs\Python\Python312\Scripts\;C:\Users\AHMET SONER\AppData\Local\Programs\Python\Python312\;C:\Users\AHMET SONER\AppData\Local\Microsoft\WindowsApps;C:\Users\AHMET SONER\go\bin;C:\Users\AHMET SONER\AppData\Roaming\npm;C:\Users\AHMET SONER\.dotnet\tools;C:\Users\AHMET SONER\AppData\Local\Yarn\bin;C:\tools\flutter\bin;C:\Dev\flutter\bin;D:\AS\W\Pars\pars\dist\;C:\ProgramData\mingw64\mingw64\bin;C:\Program Files\go-msi\;C:\ProgramData\chocolatey\bin;C:\Program Files\Git\bin;C:\Program Files\CMake\bin;%PATH%;C:\Windows\System32\WindowsPowerShell\v1.0\;C:\Program Files\nodejs

cd "D:\AS\W\Pars\pars\dist\v1.0.3-preview.119\windows\pkg\choco\all\pars"

choco install pars -y -s . --pre

choco uninstall pars
