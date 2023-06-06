# [4.0.0](https://github.ibm.com/coligo/go-sdk/compare/v3.1.0...v4.0.0) (2023-06-06)


### Features

* **api:** Added support for service bindings and service access secret ([#40](https://github.ibm.com/coligo/go-sdk/issues/40)) ([438d269](https://github.ibm.com/coligo/go-sdk/commit/438d2694c95b60dd9797865bb22c4bd8a186531f))
* **api:** SDK update 20230605-120554 ([#38](https://github.ibm.com/coligo/go-sdk/issues/38)) ([a231e21](https://github.ibm.com/coligo/go-sdk/commit/a231e21ee040479743b31c99e4e0dfd0fde5e17c))


### Reverts

* Revert "sdk-update-20230605-120554" (#39) ([c689cb7](https://github.ibm.com/coligo/go-sdk/commit/c689cb731e74f0a31a4f5d728c4ee94154d84cf7)), closes [#39](https://github.ibm.com/coligo/go-sdk/issues/39) [#38](https://github.ibm.com/coligo/go-sdk/issues/38)


### BREAKING CHANGES

* **api:** Removed SoureURL from required fields for build create operations

# [3.1.0](https://github.ibm.com/coligo/go-sdk/compare/v3.0.0...v3.1.0) (2023-03-24)


### Bug Fixes

* **tests:** hardened integration tests ([#34](https://github.ibm.com/coligo/go-sdk/issues/34)) ([06943d4](https://github.ibm.com/coligo/go-sdk/commit/06943d4dddc3f482703e75b24f0fd86b9773aece))


### Features

* **api:** added support to retrieve egress ips ([e9688bc](https://github.ibm.com/coligo/go-sdk/commit/e9688bc8a17dc8396d9fb30921fd37a2a2564aea))

# [3.0.0](https://github.ibm.com/coligo/go-sdk/compare/v2.0.5...v3.0.0) (2023-03-24)


### Features

* **oneOf:** oneOf added for v3 release ([#30](https://github.ibm.com/coligo/go-sdk/issues/30)) ([4f0d815](https://github.ibm.com/coligo/go-sdk/commit/4f0d8151e978e9f0cb22cb05a99481300129a81f))


### BREAKING CHANGES

* **oneOf:** Secret data structs are no long string maps, they are oneOf types with sepcific properties based on the type of secret being created/updated

## [2.0.5](https://github.ibm.com/coligo/go-sdk/compare/v2.0.4...v2.0.5) (2023-03-09)


### Bug Fixes

* **model:** pulled in latest v2 endpoint updates ([38796d2](https://github.ibm.com/coligo/go-sdk/commit/38796d2175b9b01a09776d707c8dc35700e3d74f))

## [2.0.4](https://github.ibm.com/coligo/go-sdk/compare/v2.0.3...v2.0.4) (2023-03-08)


### Bug Fixes

* **build:** regenerated with most recent generator ([629189b](https://github.ibm.com/coligo/go-sdk/commit/629189b74e9fbdb8e5a58f2f0d6896e025140744))
* **examples:** added secret create example ([e630bcd](https://github.ibm.com/coligo/go-sdk/commit/e630bcdaf6ea18e138f0431d03bf21650973b445))

## [2.0.3](https://github.ibm.com/coligo/go-sdk/compare/v2.0.2...v2.0.3) (2022-12-09)


### Bug Fixes

* **documentation:** slightly adjusted the readme ([be0d31d](https://github.ibm.com/coligo/go-sdk/commit/be0d31d57f833e5a0050db5897055ccf4cea86fd))

## [2.0.2](https://github.ibm.com/coligo/go-sdk/compare/v2.0.1...v2.0.2) (2022-12-09)


### Bug Fixes

* **build:** remove package.json from git. added changelog entry for v2.0.0 ([f7d29f4](https://github.ibm.com/coligo/go-sdk/commit/f7d29f48a5f5f02befda1bb5e44906250017a20a))

## [2.0.1](https://github.ibm.com/coligo/go-sdk/compare/v2.0.0...v2.0.1) (2022-12-09)


### Bug Fixes

* **build:** added a travis.yml that should be used for public github ([5ee886d](https://github.ibm.com/coligo/go-sdk/commit/5ee886df82e99c0e06c83ab665ac82f2672acb92))


## [2.0.0] (2022-12-08)

### Features

* **core** added support for [Code Engine API v2](https://cloud.ibm.com/apidocs/codeengine/codeengine-v2.0.0) features which allows to manage projects, apps, jobs, builds, secrets and config maps

### Bug Fixes

* **dependencies** bumped various dependencies
* **dependencies** added support for Golang v1.18
