# [1.7.0](https://github.com/think-root/telegram-connector/compare/v1.6.1...v1.7.0) (2025-03-04)


### Bug Fixes

* **telegram:** update PingHandler to use correct ChatID for responses ([e1405fb](https://github.com/think-root/telegram-connector/commit/e1405fb90c87922565df6a60ace7b838ca606d67))
* update Dockerfile to build from the correct main.go path ([3654233](https://github.com/think-root/telegram-connector/commit/3654233be521649331fa8a55f6ee61e8cd1ef6b4))


### Features

* add main application entry point ([b01863b](https://github.com/think-root/telegram-connector/commit/b01863b12487c59923f766ec0e7e72fd5eb2a978))
* implement HTTP server with Telegram message endpoint ([76216b2](https://github.com/think-root/telegram-connector/commit/76216b20b73e97ef40140ff145ce82e7efabc879))
* **telegram:** add logging middleware to log incoming Telegram updates ([5d48a50](https://github.com/think-root/telegram-connector/commit/5d48a50842da8cf16616bd063e2e10b7cd10f2d9))
* **telegram:** add SendMessage function to send repository notifications with images ([06176dc](https://github.com/think-root/telegram-connector/commit/06176dc10247134d440bbf0d30b3d25add5fc2b7))
* **telegram:** add StartHandler to respond with GitHub Ukraine channel link ([3dbcfdf](https://github.com/think-root/telegram-connector/commit/3dbcfdf6015fd7ac5879826004957d2bb0df0d66))

## [1.6.1](https://github.com/think-root/telegram-connector/compare/v1.6.0...v1.6.1) (2025-03-02)


### Bug Fixes

* update network name in docker-compose.yml from alchemist_network to think-root-network ([4b84dd1](https://github.com/think-root/telegram-connector/commit/4b84dd1d529bc8752330c9b57d8c069603cb8d1c))

# [1.6.0](https://github.com/Think-Root/chappie_bot/compare/v1.5.0...v1.6.0) (2025-03-02)


### Bug Fixes

* update image file in SendPostHandler function ([ef7561f](https://github.com/Think-Root/chappie_bot/commit/ef7561f8624dae3ff3eed31cfba38db7a15796db))
* update image file in SendPostToXHandler function ([321699c](https://github.com/Think-Root/chappie_bot/commit/321699cd9d1f03d38643f739bc6eec3034838f7e))
* update image file path in SendMessageCron function ([7730177](https://github.com/Think-Root/chappie_bot/commit/7730177e97cc6698e71ece18a28d0cc846b23151))


### Features

* add assets directory to Dockerfile ([d100f8b](https://github.com/Think-Root/chappie_bot/commit/d100f8b957607286108a21a169e208fd4831d669))
* add new banner image ([f8af6ff](https://github.com/Think-Root/chappie_bot/commit/f8af6ffbb5cdd5a5b42a510ba619c4ab2e3b9c6d))

# [1.5.0](https://github.com/Think-Root/chappie_bot/compare/v1.4.0...v1.5.0) (2025-02-21)


### Bug Fixes

* update release workflow to trigger on Gitleaks Security Scan ([324addd](https://github.com/think-root/telegram-connector/commit/324addd300b02b5892140cb43222e59d681cf02c))


### Features

* add APP_VERSION argument to Docker build in docker-compose.yml ([4224316](https://github.com/think-root/telegram-connector/commit/422431617a2644cd245502a3b6bc3d78e4254eba))
* add APP_VERSION argument to Dockerfile for versioning in build ([6041f15](https://github.com/think-root/telegram-connector/commit/6041f15a5070c1afdfe00d0fd2be347d2204a9b6))
* update deployment workflow to trigger on release and include versioning in Docker build ([3029347](https://github.com/think-root/telegram-connector/commit/3029347c3c5f303cc8a1c75a80467b7a499ea104))

# [1.4.0](https://github.com/think-root/telegram-connector/compare/v1.3.0...v1.4.0) (2025-02-21)


### Features

* add success and error messages for sending posts to X ([98d9539](https://github.com/think-root/telegram-connector/commit/98d9539f186f80708e388938a9c327371cf2da99))

# [1.3.0](https://github.com/think-root/telegram-connector/compare/v1.2.0...v1.3.0) (2025-02-21)


### Features

* enhance error logging in CreateXPost function for better debugging ([3176ca1](https://github.com/think-root/telegram-connector/commit/3176ca1306de3c33d55e52b9a0a26bd93ca2a225))

# [1.2.0](https://github.com/think-root/telegram-connector/compare/v1.1.0...v1.2.0) (2025-02-21)


### Features

* add /xsend command to README for sending posts to X (Twitter) ([fda0081](https://github.com/think-root/telegram-connector/commit/fda00814260638298ea7393423fd06616b619aaf))
* implement SendPostToXHandler for sending posts to X (Twitter) ([e68f970](https://github.com/think-root/telegram-connector/commit/e68f9704a6ef6aa013c88a370c885356edb7fe8d))
* register /xsend command handler for sending posts to X (Twitter) ([ef794bf](https://github.com/think-root/telegram-connector/commit/ef794bf462a0e675cbdd00527b7c12331aa78284))

# [1.1.0](https://github.com/think-root/telegram-connector/compare/v1.0.5...v1.1.0) (2025-02-18)


### Bug Fixes

* improve log messages for clarity in SendMessageCron function ([ce6829c](https://github.com/think-root/telegram-connector/commit/ce6829cea540367c7273fe3ee6b8b1dfef27b1c3))


### Features

* add GitHub Actions workflows for deployment, security scanning, and release management ([f1cd803](https://github.com/think-root/telegram-connector/commit/f1cd80379e881ab455644d99ddf9def979cd60d5))
* add Gitleaks configuration to allowlist README.md ([3a8f9f4](https://github.com/think-root/telegram-connector/commit/3a8f9f45cfa5d8bcc0e4be2783c74b228b14bae3))
* add issue templates for bug reports, documentation, feature requests, and questions ([9c7fc22](https://github.com/think-root/telegram-connector/commit/9c7fc2257c94ba6afc3b4c01bb14eab5627fac45))
* add semantic release configuration and dependencies for automated versioning ([dba89d6](https://github.com/think-root/telegram-connector/commit/dba89d6f3662e9702164b3aa42ffcd078906daea))
* update README badges for Go version, changelog, and deploy status ([b8d268d](https://github.com/think-root/telegram-connector/commit/b8d268d13793718c75420b24f41d165595135026))
