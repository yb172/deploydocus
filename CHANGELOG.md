# Changelog

## [1.7.0](https://github.com/yb172/deploydocus/compare/v1.6.1...v1.7.0) (2025-03-02)


### Features

* add cherrypick instructions to merge back into main ([6ad6dd8](https://github.com/yb172/deploydocus/commit/6ad6dd82d2b1d78bde68f3a1e9b78f5972f833f5))

## [1.6.1](https://github.com/yb172/deploydocus/compare/v1.6.0...v1.6.1) (2025-03-01)


### Bug Fixes

* give execute permission to release/cherrypick.sh ([0535611](https://github.com/yb172/deploydocus/commit/053561117d1979d5f9c787f652ddc382fb02b338))

## [1.6.0](https://github.com/yb172/deploydocus/compare/v1.5.0...v1.6.0) (2025-03-01)


### Features

* meaningless change ([fa86e3a](https://github.com/yb172/deploydocus/commit/fa86e3ac00873c8a41e8ad4ac529fb6f7e835efc))

## [1.5.2](https://github.com/yb172/deploydocus/compare/v1.5.1...v1.5.2) (2025-03-01)


### Bug Fixes

* undo update ([3890042](https://github.com/yb172/deploydocus/commit/389004215838d4b2f86bc88ff167b66ccd58677e))

## [1.5.1](https://github.com/yb172/deploydocus/compare/v1.5.0...v1.5.1) (2025-03-01)


### Bug Fixes

* the fix ([d688659](https://github.com/yb172/deploydocus/commit/d6886596f8a9a0d9f6fa0c0b4a636d839ad3906f))


## [1.5.0](https://github.com/yb172/deploydocus/compare/v1.4.1...v1.5.0) (2025-03-01)


### Features

* add cherrypick command ([bd4f5da](https://github.com/yb172/deploydocus/commit/bd4f5dab5b4b71badf63c189afabee0e12b4a701))
* update cherrypick branch names ([74ba0b4](https://github.com/yb172/deploydocus/commit/74ba0b44337ae8f3e44b243242b4dbd846465808))

## [1.4.1](https://github.com/yb172/deploydocus/compare/v1.4.0...v1.4.1) (2025-03-01)


### Bug Fixes

* fix bug on main page ([3e881ef](https://github.com/yb172/deploydocus/commit/3e881efe2e4b4db88e2900f2887a5f271b19cb49))
* tweak cherrypick config ([f21ef3b](https://github.com/yb172/deploydocus/commit/f21ef3bcac95ca584431971c2dc431199cdc94fd))
* update cherrypick config ([56103d4](https://github.com/yb172/deploydocus/commit/56103d4bde609a85bef249877a8e6aff3f0d85b9))

## [1.4.0](https://github.com/yb172/deploydocus/compare/v1.3.0...v1.4.0) (2025-03-01)


### Features

* center the content ([d114e94](https://github.com/yb172/deploydocus/commit/d114e942bdec9e0b29768eff8de76ac79ed07e1a))

## [1.3.0](https://github.com/yb172/deploydocus/compare/v1.2.0...v1.3.0) (2025-03-01)


### Features

* add cherrypick section to readme ([5d2f7f5](https://github.com/yb172/deploydocus/commit/5d2f7f503d7b52e4bb7f61e026d0c7b56728f5ba))
* implement bug ([124d572](https://github.com/yb172/deploydocus/commit/124d572eb3f8e5ce34b9765bbeadf71ee77a98cc))
* update branch name from 'release' to 'cherrypick' ([d97a06c](https://github.com/yb172/deploydocus/commit/d97a06cbf1587aac58ca6384cef7fb4896e4cfbc))

## [1.2.0](https://github.com/yb172/deploydocus/compare/v1.1.1...v1.2.0) (2025-03-01)


### Features

* add commit linter ([31542a9](https://github.com/yb172/deploydocus/commit/31542a907ecc691b67ddd0a02c0a5a9e89c34785))
* Update README.md with info about prod env ([04cf66b](https://github.com/yb172/deploydocus/commit/04cf66be64866ea2a0390c594f93cdf05860c0a9))

## [1.1.1](https://github.com/yb172/deploydocus/compare/v1.1.0...v1.1.1) (2025-02-26)


### Bug Fixes

* Fix step name in deploy-demo.yml ([f529d9a](https://github.com/yb172/deploydocus/commit/f529d9af6e0aa65f5dff44436008d602c5d8c580))

## [1.1.0](https://github.com/yb172/deploydocus/compare/v1.0.0...v1.1.0) (2025-02-26)


### Features

* deploy-prod.sh is ready to use ([3d5d2f5](https://github.com/yb172/deploydocus/commit/3d5d2f515d18ab1f48bbc78ee76c7c1085545dbf))
* restructure workflows to have only 1 job ([1abbca9](https://github.com/yb172/deploydocus/commit/1abbca9c490b8d7f86aacbc12c4e0d53b955a4e2))


### Bug Fixes

* Add checkout step to prod deploy, cleanup deploy-prod.sh ([e53b227](https://github.com/yb172/deploydocus/commit/e53b22700b2429537a6f1ddcf2900cfa98a5e766))
* Append deployed prod version to the begining of the file ([d69fd4a](https://github.com/yb172/deploydocus/commit/d69fd4a1333a7344078c212fc2492b5fece67ae7))
* fix image name var in deploy staging workflow config ([e6d3ced](https://github.com/yb172/deploydocus/commit/e6d3ced26bc3d17b5fa288ed24d21e04d4b64266))
* Fix version extraction for prod deploy ([45376b3](https://github.com/yb172/deploydocus/commit/45376b347708433185baca31fd96fb118a225375))
* Fix version extraction from file in prod deploy workflow ([e895f66](https://github.com/yb172/deploydocus/commit/e895f665faf1cbfac3232eee5801f4b74d4b5f5d))
* lowercase var name in prod deploy workflow ([0dd6958](https://github.com/yb172/deploydocus/commit/0dd6958c87080627978a9ad3bb65e3d50a05c984))
* one more debug ([57c26fb](https://github.com/yb172/deploydocus/commit/57c26fb6b9b1522bd51e80f90e0fe90bf9aa3027))
* update command to extract version from branch name ([9d27a8a](https://github.com/yb172/deploydocus/commit/9d27a8a9b46990f392fd3a28ea322b65bea122a4))

## 1.0.0 (2025-02-25)


### Features

* return html with nice emoji favicon ([84aec9a](https://github.com/yb172/deploydocus/commit/84aec9a2fc3c13cb8cefbd3c868ac3720c51269f))
