<<<<<<< HEAD

=======
🌐 [Main](./README.md)

---
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c

# Git Release Manager v1.0.0-dev.1 Release Notes - 2024-12-10

## Summary:

<<<<<<< HEAD
- ✨ New Features &amp; Enhancements: 7 commits

- 🐞 Resolved Issues: 1 commit

- 📚 Documentation: 2 commits

- 🔧 Codebase Maintenance and Updates: 5 commits


and 2 ❤️ contributors






=======
-   ✨ New Features &amp; Enhancements: 7 commits

-   🐞 Resolved Issues: 1 commit

-   📚 Documentation: 2 commits

-   🔧 Codebase Maintenance and Updates: 5 commits

and 2 ❤️ contributors

>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
## Changes

### ✨ New Features &amp; Enhancements:

<<<<<<< HEAD
- <a id="commit-bf437f1"></a>**cmake** [#bf437f1](https://github.com/ahmettsoner/pars/commit/bf437f142c44e7140917d66242993cfd28aafd13): add choco package support with install/uninstall scripts (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 


	Added chocolateyInstall.ps1 and chocolateyUninstall.ps1 scripts for managing system environment paths and providing Add/Remove Programs support. This enhancement ensures improving usability and compliance with system package management standards.


- <a id="commit-7c974b7"></a>**cmake** [#7c974b7](https://github.com/ahmettsoner/pars/commit/7c974b7a0e560a6b184e84d4009031494d7ec9c0): add pre-release detection based on version information (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 


	Implemented logic to detect pre-release versions in CMake by evaluating the version string for pre-release identifiers (e.g., alpha, beta, rc). This functionality enhances version management by automatically determining the pre-release status during the build process.


- <a id="commit-81b604c"></a>**cmake** [#81b604c](https://github.com/ahmettsoner/pars/commit/81b604c76f22b2e28c2a7e5dc11886efcae2c1a2): add 386 architecture support for Go builds (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 


	Enabled 386 architecture support in the Go build process.


- <a id="commit-17628e3"></a>**choco** [#17628e3](https://github.com/ahmettsoner/pars/commit/17628e3fbc9d8cee7f4ed5d88f64751c2ebc5a28): enhance package with Add/Remove Programs and PATH management (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 


	Added support for Add/Remove Programs entry during Chocolatey package installation.
	
	Automatically adds the application to the PATH environment variable post-installation.
	
	Improved uninstall process to remove application details from Windows Registry and clean up the PATH environment variable.


- <a id="commit-3b6244b"></a>**installer** [#3b6244b](https://github.com/ahmettsoner/pars/commit/3b6244bfaeb3248b46f283ca1bb27f43ce4e3c87): improve displayed application name in MSI config (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 


- <a id="commit-6137359"></a>**cmake** [#6137359](https://github.com/ahmettsoner/pars/commit/61373597d297c05baf8f8f868348495ec888db0f): add Chocolatey packaging support (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 


- <a id="commit-da170e1"></a>**cmake** [#da170e1](https://github.com/ahmettsoner/pars/commit/da170e1b49fc666b25eee7c11599dbb328872fc9): add Chocolatey packaging support (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 



### 🐞 Resolved Issues:

- <a id="commit-51acc69"></a>**cmake** [#51acc69](https://github.com/ahmettsoner/pars/commit/51acc69ea1d463b5a8dd8da6f6631611320be30c): resolve issues in MSI installer configuration (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 



### 📚 Documentation:

- <a id="commit-a2bf7fe"></a>[#a2bf7fe](https://github.com/ahmettsoner/pars/commit/a2bf7fe6a401fbbe88c0662af42d6e911be77b26): add comprehensive README.md (by [@Ahmet Soner](mailto:ahmettsoner@gmail.com))  (Issue: [#385](https://github.com/ahmettsoner/pars/issues/385))


	Added a detailed `README.md` file for the Pars repository, including sections for project overview, key features, getting started guide, installation instructions, contribution guidelines, documentation links, license information, and community resources.


- <a id="commit-a6b7695"></a>[#a6b7695](https://github.com/ahmettsoner/pars/commit/a6b7695f2fd0ad04cdc3b70e24e9c072ec0c6b4c): add comprehensive README.md (#385) (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))  (Issue: [#385](https://github.com/ahmettsoner/pars/issues/385))


	Added a detailed README.md file for the Pars repository, including sections for project overview, key features, getting started guide, installation instructions, contribution guidelines, documentation links, license information, and community resources.



### 🔧 Codebase Maintenance and Updates:

- <a id="commit-92ac517"></a>**vscode** [#92ac517](https://github.com/ahmettsoner/pars/commit/92ac517401e6f31156d605392a872834980eceaa): set source.organizeImports to explicit in settings (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 


	Ensures that import statements are explicitly organized as per project conventions.


- <a id="commit-2010a74"></a>**project** [#2010a74](https://github.com/ahmettsoner/pars/commit/2010a74dda273e2b3db1a3ac1bbe9858cfbcdbb1): add .vscode folder with launch, settings, and tasks configurations (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 


- <a id="commit-7857c8f"></a>[#7857c8f](https://github.com/ahmettsoner/pars/commit/7857c8f6d2d6f09f4503e40620b68d4ead79e9a2): removed unused files (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 


- <a id="commit-02ae3da"></a>[#02ae3da](https://github.com/ahmettsoner/pars/commit/02ae3da3d79d97525806d935c832e2b4ba23a443): fixing conflicts (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) 


- <a id="commit-80ea61f"></a>[#80ea61f](https://github.com/ahmettsoner/pars/commit/80ea61fac776a2fcddfb087f07b78f79ceded35d): initialize repository for Semantic Release (by [@Ahmet Soner](mailto:ahmettsoner@gmail.com)) 





## Useful Links
- 📜 [Full Changelog](https://github.com/ahmettsoner/pars/blob/main/CHANGELOG.md)



## Contributors:

- [CI/CD Bot](mailto:ci-bot@parsdevkit.net) (Code, Configurations, Docs, Tests)

- [Ahmet Soner](mailto:ahmettsoner@gmail.com) (Docs, Configurations, Code, Tests)


=======
-   <a id="commit-bf437f1"></a>**cmake** [#bf437f1](https://github.com/ahmettsoner/pars/commit/bf437f142c44e7140917d66242993cfd28aafd13): add choco package support with install/uninstall scripts (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

    Added chocolateyInstall.ps1 and chocolateyUninstall.ps1 scripts for managing system environment paths and providing Add/Remove Programs support. This enhancement ensures improving usability and compliance with system package management standards.

-   <a id="commit-7c974b7"></a>**cmake** [#7c974b7](https://github.com/ahmettsoner/pars/commit/7c974b7a0e560a6b184e84d4009031494d7ec9c0): add pre-release detection based on version information (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

    Implemented logic to detect pre-release versions in CMake by evaluating the version string for pre-release identifiers (e.g., alpha, beta, rc). This functionality enhances version management by automatically determining the pre-release status during the build process.

-   <a id="commit-81b604c"></a>**cmake** [#81b604c](https://github.com/ahmettsoner/pars/commit/81b604c76f22b2e28c2a7e5dc11886efcae2c1a2): add 386 architecture support for Go builds (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

    Enabled 386 architecture support in the Go build process.

-   <a id="commit-17628e3"></a>**choco** [#17628e3](https://github.com/ahmettsoner/pars/commit/17628e3fbc9d8cee7f4ed5d88f64751c2ebc5a28): enhance package with Add/Remove Programs and PATH management (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

    Added support for Add/Remove Programs entry during Chocolatey package installation.

    Automatically adds the application to the PATH environment variable post-installation.

    Improved uninstall process to remove application details from Windows Registry and clean up the PATH environment variable.

-   <a id="commit-3b6244b"></a>**installer** [#3b6244b](https://github.com/ahmettsoner/pars/commit/3b6244bfaeb3248b46f283ca1bb27f43ce4e3c87): improve displayed application name in MSI config (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-6137359"></a>**cmake** [#6137359](https://github.com/ahmettsoner/pars/commit/61373597d297c05baf8f8f868348495ec888db0f): add Chocolatey packaging support (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-da170e1"></a>**cmake** [#da170e1](https://github.com/ahmettsoner/pars/commit/da170e1b49fc666b25eee7c11599dbb328872fc9): add Chocolatey packaging support (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

### 🐞 Resolved Issues:

-   <a id="commit-51acc69"></a>**cmake** [#51acc69](https://github.com/ahmettsoner/pars/commit/51acc69ea1d463b5a8dd8da6f6631611320be30c): resolve issues in MSI installer configuration (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

### 📚 Documentation:

-   <a id="commit-a2bf7fe"></a>[#a2bf7fe](https://github.com/ahmettsoner/pars/commit/a2bf7fe6a401fbbe88c0662af42d6e911be77b26): add comprehensive README.md (by [@Ahmet Soner](mailto:ahmettsoner@gmail.com)) (Issue: [#385](https://github.com/ahmettsoner/pars/issues/385))

    Added a detailed `README.md` file for the Pars repository, including sections for project overview, key features, getting started guide, installation instructions, contribution guidelines, documentation links, license information, and community resources.

-   <a id="commit-a6b7695"></a>[#a6b7695](https://github.com/ahmettsoner/pars/commit/a6b7695f2fd0ad04cdc3b70e24e9c072ec0c6b4c): add comprehensive README.md (#385) (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net)) (Issue: [#385](https://github.com/ahmettsoner/pars/issues/385))

    Added a detailed README.md file for the Pars repository, including sections for project overview, key features, getting started guide, installation instructions, contribution guidelines, documentation links, license information, and community resources.

### 🔧 Codebase Maintenance and Updates:

-   <a id="commit-92ac517"></a>**vscode** [#92ac517](https://github.com/ahmettsoner/pars/commit/92ac517401e6f31156d605392a872834980eceaa): set source.organizeImports to explicit in settings (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

    Ensures that import statements are explicitly organized as per project conventions.

-   <a id="commit-2010a74"></a>**project** [#2010a74](https://github.com/ahmettsoner/pars/commit/2010a74dda273e2b3db1a3ac1bbe9858cfbcdbb1): add .vscode folder with launch, settings, and tasks configurations (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-7857c8f"></a>[#7857c8f](https://github.com/ahmettsoner/pars/commit/7857c8f6d2d6f09f4503e40620b68d4ead79e9a2): removed unused files (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-02ae3da"></a>[#02ae3da](https://github.com/ahmettsoner/pars/commit/02ae3da3d79d97525806d935c832e2b4ba23a443): fixing conflicts (by [@CI/CD Bot](mailto:ci-bot@parsdevkit.net))

-   <a id="commit-80ea61f"></a>[#80ea61f](https://github.com/ahmettsoner/pars/commit/80ea61fac776a2fcddfb087f07b78f79ceded35d): initialize repository for Semantic Release (by [@Ahmet Soner](mailto:ahmettsoner@gmail.com))

## Useful Links

-   📜 [Full Changelog](https://github.com/ahmettsoner/pars/blob/main/CHANGELOG.md)

## Contributors:

-   [CI/CD Bot](mailto:ci-bot@parsdevkit.net) (Code, Configurations, Docs, Tests)

-   [Ahmet Soner](mailto:ahmettsoner@gmail.com) (Docs, Configurations, Code, Tests)
>>>>>>> 9b114aa382da2ca860f21271f9780439b8929d5c
