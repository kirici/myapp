# Changelog

All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit message semantics.

---
## [1.6.0](https://github.com/kirici/myapp/compare/v1.5.0..v1.6.0) - 2025-02-22

### Features

- **(tool)** extend changelog into release task - ([cf2865c](https://github.com/kirici/myapp/commit/cf2865c125e7c244761e73fa98bbf3096ce5b1c5)) - kirici

---
## [1.5.0](https://github.com/kirici/myapp/compare/v1.4.0..v1.5.0) - 2025-02-21

### Features

- **(ctr)** set profile for k6 service - ([67ab155](https://github.com/kirici/myapp/commit/67ab155a454902f8f18084b03af3c133d637634b)) - kirici
- **(obs)** smoothen curves - ([916dee3](https://github.com/kirici/myapp/commit/916dee380ad59042bada88c201b5020ad3ec3de8)) - kirici
- **(obs)** sum client latency metrics by path - ([3535f1f](https://github.com/kirici/myapp/commit/3535f1f78e32a78c45f1cddafa6100d5c0e314a2)) - kirici
- **(obs)** add axis unit - ([056e9f4](https://github.com/kirici/myapp/commit/056e9f4ac6fac68c0020fcb977325f7307438b2e)) - kirici
- **(tool)** add, adjust commands for test profile - ([cbc0889](https://github.com/kirici/myapp/commit/cbc0889182d15874414c66080892b7d7fbd94c30)) - kirici

### Performance

- **(app)** work speed increased by 3 orders of magnitude - ([6450a8d](https://github.com/kirici/myapp/commit/6450a8d28de55da3703154ebb57c40ec57841d53)) - kirici

### Documentation

- update screenshot - ([4799188](https://github.com/kirici/myapp/commit/47991886557a084f2c9b257544cbec6d29926d19)) - kirici

---
## [1.4.0](https://github.com/kirici/myapp/compare/v1.3.0..v1.4.0) - 2025-02-21

### Features

- **(obs)** enable remote-writing k6 metrics to prometheus - ([b934441](https://github.com/kirici/myapp/commit/b9344419e9437849f21fe1a4a5f09fca45ba549e)) - kirici
- **(obs)** use $__rate_interval for graph rates over hardcoded time durations - ([94f48d7](https://github.com/kirici/myapp/commit/94f48d7ff8c2af1c03d1f658a202735352a28a20)) - kirici
- **(obs)** add k6 client metrics, row separators - ([a821726](https://github.com/kirici/myapp/commit/a8217264fd1df41ba95a40a2ef6cf8d83aaa76cb)) - kirici
- **(tool)** replace vegeta with grafanalabs' k6 - ([a0b2ff9](https://github.com/kirici/myapp/commit/a0b2ff918d9e1cc049dc9761e6c4edc2e32de608)) - kirici

---
## [1.3.0](https://github.com/kirici/myapp/compare/v1.2.0..v1.3.0) - 2024-10-26

### Features

- **(conf)** add export statements for a maximally-compatible .env file - ([076f195](https://github.com/kirici/myapp/commit/076f1951756d9231b3f50b6499a5216f9a872744)) - kirici
- **(obs)** add new histogram metric for request latencies - ([c452769](https://github.com/kirici/myapp/commit/c452769a20e6312840b3e82c51946f9f8a1f1cbd)) - kirici
- **(obs)** add new dashboard with api metrics only - ([0f918af](https://github.com/kirici/myapp/commit/0f918af0d8c17ca1a2451372c8a719f161e684a9)) - kirici
- **(tool)** load sample .env file as default, overwrite afterwards - ([0ec3e05](https://github.com/kirici/myapp/commit/0ec3e05c182eb597c81624911de0bc0dec409fb3)) - kirici

### Documentation

- update usage instructions with regards to env vars - ([96144b9](https://github.com/kirici/myapp/commit/96144b9cf6e46c01c72507a301a571335ef77e40)) - kirici

### Revert

- **(conf)** non-default port for grafana - ([2f474a3](https://github.com/kirici/myapp/commit/2f474a33bf6b71198022b074fe2114945991b359)) - kirici

---
## [1.2.0](https://github.com/kirici/myapp/compare/v1.1.0..v1.2.0) - 2024-10-25

### Features

- **(obs)** register error metric collector - ([82814d7](https://github.com/kirici/myapp/commit/82814d7e3c0353dc26656e5cc5f4c4b7f2831b43)) - kirici
- **(obs)** replace label filtering with new http_errors_total metric - ([90edf63](https://github.com/kirici/myapp/commit/90edf63ae058006c5b25f5cfffcb3a4b99735050)) - kirici
- **(tool)** just commands for containers check cli availability - ([4546b5c](https://github.com/kirici/myapp/commit/4546b5c567128f18576197f511da824e22acb037)) - kirici
- **(tool)** add compose build command - ([7ba7fcb](https://github.com/kirici/myapp/commit/7ba7fcb76f5798bbda94e8b3e4e5baf57320a04d)) - kirici
- **(tool)** posix-compliant output silencing - ([2a7d2a2](https://github.com/kirici/myapp/commit/2a7d2a2e44f1fa0b812161fb6307e2cc104feae9)) - kirici

### Bug Fixes

- **(ctr)** remove redundant and out-of-spec args key - ([fd75ac1](https://github.com/kirici/myapp/commit/fd75ac18c1326d8511dd25754d4e5a5a9eddc291)) - kirici

---
## [1.1.0](https://github.com/kirici/myapp/compare/v1.0.0..v1.1.0) - 2024-10-19

### Features

- **(obs)** convert grafana db fixture into provisioning files - ([6e362ed](https://github.com/kirici/myapp/commit/6e362ed4cb40c518c60b54631ccf636242ee1d4f)) - kirici
- **(obs)** replace total count of http codes with rate - ([98f6e1f](https://github.com/kirici/myapp/commit/98f6e1f8a027873df49411de6ac2f8487261e38a)) - kirici

### Documentation

- add dashboard screenshot - ([3be7487](https://github.com/kirici/myapp/commit/3be74879b43ee654813686e60bac042e83206c9c)) - kirici
- add changelog reference, explanations for optional tools - ([9c495d3](https://github.com/kirici/myapp/commit/9c495d355c96149f92debbe3a9d128aa7ca77e7c)) - kirici

---
## [1.0.0](https://github.com/kirici/myapp/compare/v0.3.0..v1.0.0) - 2024-10-18

### Features

- **(api)** [**breaking**] do not serve /math results over 0.8 - ([b0940ac](https://github.com/kirici/myapp/commit/b0940acc62babd3e7b5b4e6a31fb619aa97e14fc)) - kirici
- **(app)** convert all api responses to json with gin-gonic context - ([b375684](https://github.com/kirici/myapp/commit/b37568456d511e797209df9a712bf532b2f342b2)) - kirici
- **(obs)** add new counter metric for errors - ([0e4151b](https://github.com/kirici/myapp/commit/0e4151b71a924e082749b21dd706f99c223acbf2)) - kirici

---
## [0.3.0](https://github.com/kirici/myapp/compare/v0.2.0..v0.3.0) - 2024-10-18

### Features

- **(app)** add simulated "/work" endpoint - ([64611b3](https://github.com/kirici/myapp/commit/64611b372fdf6473887e9252ec460d9e7fb3dcb8)) - kirici
- **(obs)** bake panel for custom metrics into grafana db fixture - ([98cf451](https://github.com/kirici/myapp/commit/98cf4519d6b56b392652f0739235aafefa9b5ba2)) - kirici

### Documentation

- **(app)** fix metrics description - ([5cb5a21](https://github.com/kirici/myapp/commit/5cb5a21058e757197c23e65a082fa7297ac955f0)) - kirici

---
## [0.2.0](https://github.com/kirici/myapp/compare/v0.1.0..v0.2.0) - 2024-10-16

### Features

- **(app)** add two new endpoints at "/" and "/math" - ([2df0b19](https://github.com/kirici/myapp/commit/2df0b19a8fc4dc3b5440067639e46ea6f99be873)) - kirici

### Bug Fixes

- **(obs)** prometheus ignoring config file, target pointing to local container - ([76769b5](https://github.com/kirici/myapp/commit/76769b5f95ba3c104866fc56f8b62969b041fd89)) - kirici

### Documentation

- add vegeta to requirements - ([a38c05b](https://github.com/kirici/myapp/commit/a38c05bc88a75495fae8893b9d81c356c84f3c7f)) - kirici
- add required steps for launching - ([54619d8](https://github.com/kirici/myapp/commit/54619d877ab47ef37f65eee1cb0048d4b1cdb4ca)) - kirici

---
## [0.1.0] - 2024-10-15

### Features

- **(app)** add minimal server with promhttp - ([81f5f3f](https://github.com/kirici/myapp/commit/81f5f3f54907109f9a43b10a4f9dba3d8c2acc40)) - kirici
- **(db)** add postgresql 16 - ([678a4f3](https://github.com/kirici/myapp/commit/678a4f3df0787312baae32efd444df09c7a23286)) - kirici
- **(obs)** add prometheus - ([9858601](https://github.com/kirici/myapp/commit/9858601be452492ad5c87fb588ec3460df7a7c34)) - kirici
- **(obs)** add grafana, import sqlite db - ([53bacaa](https://github.com/kirici/myapp/commit/53bacaaba67446006da60a5f56565dbb6134e079)) - kirici

### Documentation

- add readme file - ([f8585de](https://github.com/kirici/myapp/commit/f8585de652060e612246cd2cfc4a27d6fa84cbfe)) - kirici

<!-- generated by git-cliff -->
