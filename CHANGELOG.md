# Changelog

## 1.3.0 (2025-04-10)

Full Changelog: [v1.2.0...v1.3.0](https://github.com/Find-AI/find-ai-go/compare/v1.2.0...v1.3.0)

### Features

* add SKIP_BREW env var to ./scripts/bootstrap ([#38](https://github.com/Find-AI/find-ai-go/issues/38)) ([ab4283a](https://github.com/Find-AI/find-ai-go/commit/ab4283a6931d57dfc4d6111ce76ebfdb026d1dd5))
* **api:** manual updates ([#16](https://github.com/Find-AI/find-ai-go/issues/16)) ([a0a33f9](https://github.com/Find-AI/find-ai-go/commit/a0a33f9084cbf5a5cbfd9967d3fb2b618311d9a6))
* **api:** manual updates ([#18](https://github.com/Find-AI/find-ai-go/issues/18)) ([1220acc](https://github.com/Find-AI/find-ai-go/commit/1220acc736673781cd6610848be4649b439c9a7c))
* **client:** accept RFC6838 JSON content types ([#39](https://github.com/Find-AI/find-ai-go/issues/39)) ([b860e87](https://github.com/Find-AI/find-ai-go/commit/b860e8725956c005c98b402409e09598201029e6))
* **client:** allow custom baseurls without trailing slash ([#37](https://github.com/Find-AI/find-ai-go/issues/37)) ([36e8731](https://github.com/Find-AI/find-ai-go/commit/36e8731abba3d74dc54eec30a449f535ab1a834b))
* **client:** improve default client options support ([#41](https://github.com/Find-AI/find-ai-go/issues/41)) ([7824b31](https://github.com/Find-AI/find-ai-go/commit/7824b315d9b4cb8b361bfbd6d3c98e629ee06b52))
* **client:** send `X-Stainless-Timeout` header ([#30](https://github.com/Find-AI/find-ai-go/issues/30)) ([09f62b1](https://github.com/Find-AI/find-ai-go/commit/09f62b17b03868fb6081612639121ddede552f89))
* **client:** support custom http clients ([#48](https://github.com/Find-AI/find-ai-go/issues/48)) ([d78e6d4](https://github.com/Find-AI/find-ai-go/commit/d78e6d4d5f0c4f4f31581f3c2c1d887f505eb06c))


### Bug Fixes

* **client:** don't truncate manually specified filenames ([#34](https://github.com/Find-AI/find-ai-go/issues/34)) ([13a0af2](https://github.com/Find-AI/find-ai-go/commit/13a0af2629b69cf12f26b5da1e4a0b03fe3a39c2))
* **client:** return error on bad custom url instead of panic ([#47](https://github.com/Find-AI/find-ai-go/issues/47)) ([5055a9b](https://github.com/Find-AI/find-ai-go/commit/5055a9be0582a9cd5845a36050255f9b731b28bc))
* do not call path.Base on ContentType ([#33](https://github.com/Find-AI/find-ai-go/issues/33)) ([4f861e5](https://github.com/Find-AI/find-ai-go/commit/4f861e5a2d5d61caa27b8ddeee87408c1d5e6d01))
* fix apijson.Port for embedded structs ([#25](https://github.com/Find-AI/find-ai-go/issues/25)) ([d4650a0](https://github.com/Find-AI/find-ai-go/commit/d4650a0b8f21113bf7db6a10e044d01b4ee5de2b))
* fix apijson.Port for embedded structs ([#26](https://github.com/Find-AI/find-ai-go/issues/26)) ([dcd5083](https://github.com/Find-AI/find-ai-go/commit/dcd5083a3f6de5834229a6425d6b0368b2f73082))
* fix early cancel when RequestTimeout is provided for streaming requests ([#32](https://github.com/Find-AI/find-ai-go/issues/32)) ([9b6ce1e](https://github.com/Find-AI/find-ai-go/commit/9b6ce1e57f13f620dc55d8c6a87a1fec9cfd5345))
* fix unicode encoding for json ([#28](https://github.com/Find-AI/find-ai-go/issues/28)) ([6879cfd](https://github.com/Find-AI/find-ai-go/commit/6879cfda506301e1d4a850d7c711492833d7522f))
* **test:** return early after test failure ([#45](https://github.com/Find-AI/find-ai-go/issues/45)) ([cd5ad78](https://github.com/Find-AI/find-ai-go/commit/cd5ad78d2e576fb09b887a42113ff6c30258aa0f))


### Chores

* add request options to client tests ([#44](https://github.com/Find-AI/find-ai-go/issues/44)) ([b440e0a](https://github.com/Find-AI/find-ai-go/commit/b440e0a46f8dbe57cedb4a3fdcbe818b5d412802))
* add UnionUnmarshaler for responses that are interfaces ([#31](https://github.com/Find-AI/find-ai-go/issues/31)) ([0164df8](https://github.com/Find-AI/find-ai-go/commit/0164df8ccbadfaf794aea0df95dfd2b25b319527))
* **docs:** improve security documentation ([#43](https://github.com/Find-AI/find-ai-go/issues/43)) ([fb411a8](https://github.com/Find-AI/find-ai-go/commit/fb411a8d7b4ce7039d4b1b371a7c01ab4f5e03aa))
* fix typos ([#46](https://github.com/Find-AI/find-ai-go/issues/46)) ([8b2f261](https://github.com/Find-AI/find-ai-go/commit/8b2f26134138e21e6852600b11c71f600bed59af))
* **internal:** codegen related update ([#21](https://github.com/Find-AI/find-ai-go/issues/21)) ([d7ae1d9](https://github.com/Find-AI/find-ai-go/commit/d7ae1d982063b84606c5eeb43b880469434e57fb))
* **internal:** codegen related update ([#22](https://github.com/Find-AI/find-ai-go/issues/22)) ([cbc70ee](https://github.com/Find-AI/find-ai-go/commit/cbc70eee2c495e3c8554438cd8c0ec07e0b872dc))
* **internal:** codegen related update ([#23](https://github.com/Find-AI/find-ai-go/issues/23)) ([598b15e](https://github.com/Find-AI/find-ai-go/commit/598b15e75e2ff0fc8dab52b9d3de5a5e9ae85976))
* **internal:** codegen related update ([#24](https://github.com/Find-AI/find-ai-go/issues/24)) ([19dba6d](https://github.com/Find-AI/find-ai-go/commit/19dba6d1a51342a569e376f140591bb2fb6314a2))
* **internal:** expand CI branch coverage ([856fcb0](https://github.com/Find-AI/find-ai-go/commit/856fcb084fadd52bff4da611c1311b4c9fb061ef))
* **internal:** fix devcontainers setup ([#35](https://github.com/Find-AI/find-ai-go/issues/35)) ([b3e01bb](https://github.com/Find-AI/find-ai-go/commit/b3e01bb726db5392a4987a5a6d553f022086a26a))
* **internal:** reduce CI branch coverage ([3445dcb](https://github.com/Find-AI/find-ai-go/commit/3445dcb137cc104db8f477ab133bee8f3addffe3))
* **internal:** remove extra empty newlines ([#42](https://github.com/Find-AI/find-ai-go/issues/42)) ([d7d2799](https://github.com/Find-AI/find-ai-go/commit/d7d2799bea9c736bb4367210918136f3df3fb2c7))
* rebuild project due to codegen change ([#20](https://github.com/Find-AI/find-ai-go/issues/20)) ([f7590c4](https://github.com/Find-AI/find-ai-go/commit/f7590c41008d59ee8ac3c3ac5a411e170f7440bb))
* refactor client tests ([#27](https://github.com/Find-AI/find-ai-go/issues/27)) ([58199e8](https://github.com/Find-AI/find-ai-go/commit/58199e8115b2368f7fcaf770c76994c7a887e29c))


### Documentation

* document raw responses ([#29](https://github.com/Find-AI/find-ai-go/issues/29)) ([a83f329](https://github.com/Find-AI/find-ai-go/commit/a83f3297d2497a9f89e16a9c8db16fef7c0d6a98))
* update URLs from stainlessapi.com to stainless.com ([#36](https://github.com/Find-AI/find-ai-go/issues/36)) ([b0cc4ff](https://github.com/Find-AI/find-ai-go/commit/b0cc4ff92d2b809d4b4c2d4058bd4be2fd36d433))


### Refactors

* tidy up dependencies ([#40](https://github.com/Find-AI/find-ai-go/issues/40)) ([01b9bfb](https://github.com/Find-AI/find-ai-go/commit/01b9bfb7439c10a39bf77ebb3bda31a93f954718))

## 1.2.0 (2024-10-08)

Full Changelog: [v0.1.0-alpha.3...v1.2.0](https://github.com/Find-AI/find-ai-go/compare/v0.1.0-alpha.3...v1.2.0)

### Features

* **api:** OpenAPI spec update via Stainless API ([#13](https://github.com/Find-AI/find-ai-go/issues/13)) ([d5c972d](https://github.com/Find-AI/find-ai-go/commit/d5c972def5a0f19e02ca3fe706a79ef622b6856b))

## 0.1.0-alpha.3 (2024-10-02)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/Find-AI/find-ai-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** OpenAPI spec update via Stainless API ([#9](https://github.com/Find-AI/find-ai-go/issues/9)) ([d41d1a8](https://github.com/Find-AI/find-ai-go/commit/d41d1a890eb84119d69adb5dce51956f8aa473ec))


### Chores

* **internal:** codegen related update ([#11](https://github.com/Find-AI/find-ai-go/issues/11)) ([d0ee6be](https://github.com/Find-AI/find-ai-go/commit/d0ee6be7f1d08ec88d5ac64116dbcb61f8c8716c))

## 0.1.0-alpha.2 (2024-09-26)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/Find-AI/find-ai-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** manual updates updated ([#6](https://github.com/Find-AI/find-ai-go/issues/6)) ([fa7e2ad](https://github.com/Find-AI/find-ai-go/commit/fa7e2ad99741119fa3e867084220299659312bab))

## 0.1.0-alpha.1 (2024-09-25)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/Find-AI/find-ai-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** manual updates ([#3](https://github.com/Find-AI/find-ai-go/issues/3)) ([67520c2](https://github.com/Find-AI/find-ai-go/commit/67520c2a870c58ed08a215876b728b22b287af2e))
* **api:** OpenAPI spec update via Stainless API ([#4](https://github.com/Find-AI/find-ai-go/issues/4)) ([159cb9b](https://github.com/Find-AI/find-ai-go/commit/159cb9b196efc9a5e1e4c409558a2df240ae88aa))


### Chores

* configure new SDK language ([a7305ac](https://github.com/Find-AI/find-ai-go/commit/a7305ac15b009e56db28b2130a0722ecf9211b50))
* go live ([#1](https://github.com/Find-AI/find-ai-go/issues/1)) ([9297399](https://github.com/Find-AI/find-ai-go/commit/92973998cd92fca9c3203bf53f7f288803522e6c))
