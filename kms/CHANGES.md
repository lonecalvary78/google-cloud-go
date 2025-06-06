# Changes


## [1.22.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.21.2...kms/v1.22.0) (2025-05-21)


### Features

* **kms:** Adding eTag field to AutokeyConfig ([2a9d8ee](https://github.com/googleapis/google-cloud-go/commit/2a9d8eec71a7e6803eb534287c8d2f64903dcddd))


### Documentation

* **kms:** A comment for enum value `DESTROYED` in enum `CryptoKeyVersionState` is changed ([2a9d8ee](https://github.com/googleapis/google-cloud-go/commit/2a9d8eec71a7e6803eb534287c8d2f64903dcddd))
* **kms:** Updating docs for total_size field in KMS List APIs ([2a9d8ee](https://github.com/googleapis/google-cloud-go/commit/2a9d8eec71a7e6803eb534287c8d2f64903dcddd))

## [1.21.2](https://github.com/googleapis/google-cloud-go/compare/kms/v1.21.1...kms/v1.21.2) (2025-04-15)


### Bug Fixes

* **kms:** Update google.golang.org/api to 0.229.0 ([3319672](https://github.com/googleapis/google-cloud-go/commit/3319672f3dba84a7150772ccb5433e02dab7e201))

## [1.21.1](https://github.com/googleapis/google-cloud-go/compare/kms/v1.21.0...kms/v1.21.1) (2025-03-13)


### Bug Fixes

* **kms:** Update golang.org/x/net to 0.37.0 ([1144978](https://github.com/googleapis/google-cloud-go/commit/11449782c7fb4896bf8b8b9cde8e7441c84fb2fd))

## [1.21.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.20.5...kms/v1.21.0) (2025-02-20)


### Features

* **kms:** Add a PublicKeyFormat enum to allow specifying the format the public is going to be exported in ([c08d347](https://github.com/googleapis/google-cloud-go/commit/c08d34776d398a79f6962a26e8e2c75bc4958e2b))
* **kms:** Support PQC asymmetric signing algorithms ML_DSA_65 and SLH_DSA_SHA2_128s ([c08d347](https://github.com/googleapis/google-cloud-go/commit/c08d34776d398a79f6962a26e8e2c75bc4958e2b))

## [1.20.5](https://github.com/googleapis/google-cloud-go/compare/kms/v1.20.4...kms/v1.20.5) (2025-01-08)


### Documentation

* **kms:** Modify enum comment ([2e4feb9](https://github.com/googleapis/google-cloud-go/commit/2e4feb938ce9ab023c8aa6bd1dbdf36fe589213a))

## [1.20.4](https://github.com/googleapis/google-cloud-go/compare/kms/v1.20.3...kms/v1.20.4) (2025-01-02)


### Bug Fixes

* **kms:** Update golang.org/x/net to v0.33.0 ([e9b0b69](https://github.com/googleapis/google-cloud-go/commit/e9b0b69644ea5b276cacff0a707e8a5e87efafc9))

## [1.20.3](https://github.com/googleapis/google-cloud-go/compare/kms/v1.20.2...kms/v1.20.3) (2024-12-18)


### Documentation

* **kms:** Code documentation improvements ([4254053](https://github.com/googleapis/google-cloud-go/commit/42540530e44e5f331e66e0777c4aabf449f5fd90))

## [1.20.2](https://github.com/googleapis/google-cloud-go/compare/kms/v1.20.1...kms/v1.20.2) (2024-12-04)


### Documentation

* **kms:** A comment for enum `CryptoKeyVersionAlgorithm` is changed ([8dedb87](https://github.com/googleapis/google-cloud-go/commit/8dedb878c070cc1e92d62bb9b32358425e3ceffb))

## [1.20.1](https://github.com/googleapis/google-cloud-go/compare/kms/v1.20.0...kms/v1.20.1) (2024-10-23)


### Bug Fixes

* **kms:** Update google.golang.org/api to v0.203.0 ([8bb87d5](https://github.com/googleapis/google-cloud-go/commit/8bb87d56af1cba736e0fe243979723e747e5e11e))
* **kms:** WARNING: On approximately Dec 1, 2024, an update to Protobuf will change service registration function signatures to use an interface instead of a concrete type in generated .pb.go files. This change is expected to affect very few if any users of this client library. For more information, see https://togithub.com/googleapis/google-cloud-go/issues/11020. ([8bb87d5](https://github.com/googleapis/google-cloud-go/commit/8bb87d56af1cba736e0fe243979723e747e5e11e))

## [1.20.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.19.1...kms/v1.20.0) (2024-09-19)


### Features

* **kms:** Adding a state field for AutokeyConfig ([fdb4ea9](https://github.com/googleapis/google-cloud-go/commit/fdb4ea99189657880e5f0e0dce16bef1c3aa0d2f))


### Bug Fixes

* **kms:** Pagination feature is introduced for method `ListKeyHandles` in service `Autokey` ([fdb4ea9](https://github.com/googleapis/google-cloud-go/commit/fdb4ea99189657880e5f0e0dce16bef1c3aa0d2f))


### Documentation

* **kms:** A comment for field `destroy_scheduled_duration` in message `.google.cloud.kms.v1.CryptoKey` is updated for the default duration ([fdb4ea9](https://github.com/googleapis/google-cloud-go/commit/fdb4ea99189657880e5f0e0dce16bef1c3aa0d2f))
* **kms:** Field service_resolvers in message .google.cloud.kms.v1.EkmConnection is Explicitly is marked as to have field behavior of Optional ([fdb4ea9](https://github.com/googleapis/google-cloud-go/commit/fdb4ea99189657880e5f0e0dce16bef1c3aa0d2f))

## [1.19.1](https://github.com/googleapis/google-cloud-go/compare/kms/v1.19.0...kms/v1.19.1) (2024-09-12)


### Bug Fixes

* **kms:** Bump dependencies ([2ddeb15](https://github.com/googleapis/google-cloud-go/commit/2ddeb1544a53188a7592046b98913982f1b0cf04))

## [1.19.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.18.5...kms/v1.19.0) (2024-08-20)


### Features

* **kms:** Add support for Go 1.23 iterators ([84461c0](https://github.com/googleapis/google-cloud-go/commit/84461c0ba464ec2f951987ba60030e37c8a8fc18))

## [1.18.5](https://github.com/googleapis/google-cloud-go/compare/kms/v1.18.4...kms/v1.18.5) (2024-08-08)


### Bug Fixes

* **kms:** Update google.golang.org/api to v0.191.0 ([5b32644](https://github.com/googleapis/google-cloud-go/commit/5b32644eb82eb6bd6021f80b4fad471c60fb9d73))

## [1.18.4](https://github.com/googleapis/google-cloud-go/compare/kms/v1.18.3...kms/v1.18.4) (2024-07-24)


### Bug Fixes

* **kms:** Update dependencies ([257c40b](https://github.com/googleapis/google-cloud-go/commit/257c40bd6d7e59730017cf32bda8823d7a232758))

## [1.18.3](https://github.com/googleapis/google-cloud-go/compare/kms/v1.18.2...kms/v1.18.3) (2024-07-10)


### Bug Fixes

* **kms:** Bump google.golang.org/grpc@v1.64.1 ([8ecc4e9](https://github.com/googleapis/google-cloud-go/commit/8ecc4e9622e5bbe9b90384d5848ab816027226c5))

## [1.18.2](https://github.com/googleapis/google-cloud-go/compare/kms/v1.18.1...kms/v1.18.2) (2024-07-01)


### Bug Fixes

* **kms:** Bump google.golang.org/api@v0.187.0 ([8fa9e39](https://github.com/googleapis/google-cloud-go/commit/8fa9e398e512fd8533fd49060371e61b5725a85b))

## [1.18.1](https://github.com/googleapis/google-cloud-go/compare/kms/v1.18.0...kms/v1.18.1) (2024-06-26)


### Bug Fixes

* **kms:** Enable new auth lib ([b95805f](https://github.com/googleapis/google-cloud-go/commit/b95805f4c87d3e8d10ea23bd7a2d68d7a4157568))

## [1.18.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.17.1...kms/v1.18.0) (2024-06-20)


### Features

* **kms:** Support Key Access Justifications policy configuration ([#10398](https://github.com/googleapis/google-cloud-go/issues/10398)) ([4fa4308](https://github.com/googleapis/google-cloud-go/commit/4fa43082511e153044084c1e6736553de41a9894))

## [1.17.1](https://github.com/googleapis/google-cloud-go/compare/kms/v1.17.0...kms/v1.17.1) (2024-05-22)


### Bug Fixes

* **kms:** Enable cloud.google.com/go/auth ([#10246](https://github.com/googleapis/google-cloud-go/issues/10246)) ([1326df1](https://github.com/googleapis/google-cloud-go/commit/1326df1259942f2fd63b60b54379314ea69b6b80))

## [1.17.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.16.0...kms/v1.17.0) (2024-05-16)


### Features

* **kms:** Add client library for KMS Autokey service, which enables automated KMS key provision and management ([292e812](https://github.com/googleapis/google-cloud-go/commit/292e81231b957ae7ac243b47b8926564cee35920))

## [1.16.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.15.9...kms/v1.16.0) (2024-05-08)


### Features

* **kms:** Introduce Long-Running Operations (LRO) for KMS ([3e25053](https://github.com/googleapis/google-cloud-go/commit/3e250530567ee81ed4f51a3856c5940dbec35289))
* **kms:** Support the ED25519 asymmetric signing algorithm ([3e25053](https://github.com/googleapis/google-cloud-go/commit/3e250530567ee81ed4f51a3856c5940dbec35289))

## [1.15.9](https://github.com/googleapis/google-cloud-go/compare/kms/v1.15.8...kms/v1.15.9) (2024-05-01)


### Bug Fixes

* **kms:** Bump x/net to v0.24.0 ([ba31ed5](https://github.com/googleapis/google-cloud-go/commit/ba31ed5fda2c9664f2e1cf972469295e63deb5b4))


### Documentation

* **kms:** In google.cloud.kms.v1.PublicKey, pem field is always populated ([1d757c6](https://github.com/googleapis/google-cloud-go/commit/1d757c66478963d6cbbef13fee939632c742759c))

## [1.15.8](https://github.com/googleapis/google-cloud-go/compare/kms/v1.15.7...kms/v1.15.8) (2024-03-14)


### Bug Fixes

* **kms:** Update protobuf dep to v1.33.0 ([30b038d](https://github.com/googleapis/google-cloud-go/commit/30b038d8cac0b8cd5dd4761c87f3f298760dd33a))

## [1.15.7](https://github.com/googleapis/google-cloud-go/compare/kms/v1.15.6...kms/v1.15.7) (2024-02-06)


### Documentation

* **kms:** Update comments ([e60a6ba](https://github.com/googleapis/google-cloud-go/commit/e60a6ba01acf2ef2e8d12e23ed5c6e876edeb1b7))

## [1.15.6](https://github.com/googleapis/google-cloud-go/compare/kms/v1.15.5...kms/v1.15.6) (2024-01-30)


### Bug Fixes

* **kms:** Enable universe domain resolution options ([fd1d569](https://github.com/googleapis/google-cloud-go/commit/fd1d56930fa8a747be35a224611f4797b8aeb698))

## [1.15.5](https://github.com/googleapis/google-cloud-go/compare/kms/v1.15.4...kms/v1.15.5) (2023-11-01)


### Bug Fixes

* **kms:** Bump google.golang.org/api to v0.149.0 ([8d2ab9f](https://github.com/googleapis/google-cloud-go/commit/8d2ab9f320a86c1c0fab90513fc05861561d0880))

## [1.15.4](https://github.com/googleapis/google-cloud-go/compare/kms/v1.15.3...kms/v1.15.4) (2023-10-26)


### Bug Fixes

* **kms:** Update grpc-go to v1.59.0 ([81a97b0](https://github.com/googleapis/google-cloud-go/commit/81a97b06cb28b25432e4ece595c55a9857e960b7))

## [1.15.3](https://github.com/googleapis/google-cloud-go/compare/kms/v1.15.2...kms/v1.15.3) (2023-10-12)


### Bug Fixes

* **kms:** Update golang.org/x/net to v0.17.0 ([174da47](https://github.com/googleapis/google-cloud-go/commit/174da47254fefb12921bbfc65b7829a453af6f5d))

## [1.15.2](https://github.com/googleapis/google-cloud-go/compare/kms/v1.15.1...kms/v1.15.2) (2023-09-11)


### Documentation

* **kms:** Minor formatting ([15be57b](https://github.com/googleapis/google-cloud-go/commit/15be57b9264a793494cedc3966034fa20f56d7c5))
* **kms:** Minor formatting ([fbfaf21](https://github.com/googleapis/google-cloud-go/commit/fbfaf21c15ae8a07ab39c6036cf0cee700b5627c))

## [1.15.1](https://github.com/googleapis/google-cloud-go/compare/kms/v1.15.0...kms/v1.15.1) (2023-08-08)


### Documentation

* **kms:** Minor formatting ([#8371](https://github.com/googleapis/google-cloud-go/issues/8371)) ([b4349cc](https://github.com/googleapis/google-cloud-go/commit/b4349cc507870ff8629bbc07de578b63bb889626))

## [1.15.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.14.0...kms/v1.15.0) (2023-07-24)


### Features

* **kms:** Add interoperable symmetric encryption system ([432864c](https://github.com/googleapis/google-cloud-go/commit/432864c7fc0bb551a5017b423bbd5f76c3357dc3))

## [1.14.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.13.0...kms/v1.14.0) (2023-07-18)


### Features

* **kms/inventory:** Add resource_types to SearchAllResources, to allow filtering by resource type ([#8261](https://github.com/googleapis/google-cloud-go/issues/8261)) ([9d55bab](https://github.com/googleapis/google-cloud-go/commit/9d55bab76839af4ef4301544b856b042352a489e))

## [1.13.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.12.1...kms/v1.13.0) (2023-07-10)


### Features

* **kms:** Add interoperable symmetric encryption system ([14b95d3](https://github.com/googleapis/google-cloud-go/commit/14b95d33753d0b391d0b49533e92b551e5dc3072))

## [1.12.1](https://github.com/googleapis/google-cloud-go/compare/kms/v1.12.0...kms/v1.12.1) (2023-06-20)


### Bug Fixes

* **kms:** REST query UpdateMask bug ([df52820](https://github.com/googleapis/google-cloud-go/commit/df52820b0e7721954809a8aa8700b93c5662dc9b))

## [1.12.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.11.0...kms/v1.12.0) (2023-06-13)


### Features

* **kms/inventory:** Promote to GA ([3abdfa1](https://github.com/googleapis/google-cloud-go/commit/3abdfa14dd56cf773c477f289a7f888e20bbbd9a))

## [1.11.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.10.2...kms/v1.11.0) (2023-05-30)


### Features

* **kms:** Update all direct dependencies ([b340d03](https://github.com/googleapis/google-cloud-go/commit/b340d030f2b52a4ce48846ce63984b28583abde6))

## [1.10.2](https://github.com/googleapis/google-cloud-go/compare/kms/v1.10.1...kms/v1.10.2) (2023-05-08)


### Bug Fixes

* **kms:** Update grpc to v1.55.0 ([1147ce0](https://github.com/googleapis/google-cloud-go/commit/1147ce02a990276ca4f8ab7a1ab65c14da4450ef))

## [1.10.1](https://github.com/googleapis/google-cloud-go/compare/kms/v1.10.0...kms/v1.10.1) (2023-04-04)


### Documentation

* **kms:** Publish the API comment changes related to supporting different hash functions/values for ECDSA signing ([#7619](https://github.com/googleapis/google-cloud-go/issues/7619)) ([597ea0f](https://github.com/googleapis/google-cloud-go/commit/597ea0fe09bcea04e884dffe78add850edb2120d))

## [1.10.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.9.0...kms/v1.10.0) (2023-03-15)


### Features

* **kms:** Add support for Coordinated External Keys ([#7517](https://github.com/googleapis/google-cloud-go/issues/7517)) ([64c6a6f](https://github.com/googleapis/google-cloud-go/commit/64c6a6fa30fd8bec40405fdddb73d1078024e985))
* **kms:** Update iam and longrunning deps ([91a1f78](https://github.com/googleapis/google-cloud-go/commit/91a1f784a109da70f63b96414bba8a9b4254cddd))

## [1.9.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.8.0...kms/v1.9.0) (2023-03-01)


### Features

* **kms/inventory:** Start generating apiv1 ([#7504](https://github.com/googleapis/google-cloud-go/issues/7504)) ([0ceff58](https://github.com/googleapis/google-cloud-go/commit/0ceff5837ca7389d52cf344da353ef3c85483055))

## [1.8.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.7.0...kms/v1.8.0) (2023-01-04)


### Features

* **kms:** Add REST client ([06a54a1](https://github.com/googleapis/google-cloud-go/commit/06a54a16a5866cce966547c51e203b9e09a25bc0))

## [1.7.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.6.0...kms/v1.7.0) (2022-12-01)


### Features

* **kms:** add SHA-2 import methods ([7231644](https://github.com/googleapis/google-cloud-go/commit/7231644e71f05abc864924a0065b9ea22a489180))
* **kms:** add support for additional HMAC algorithms ([2a0b1ae](https://github.com/googleapis/google-cloud-go/commit/2a0b1aeb1683222e6aa5c876cb945845c00cef79))

## [1.6.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.5.0...kms/v1.6.0) (2022-11-03)


### Features

* **kms:** rewrite signatures in terms of new location ([3c4b2b3](https://github.com/googleapis/google-cloud-go/commit/3c4b2b34565795537aac1661e6af2442437e34ad))

## [1.5.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.4.0...kms/v1.5.0) (2022-10-25)


### Features

* **kms:** enable generation of Locations mixin ([caf4afa](https://github.com/googleapis/google-cloud-go/commit/caf4afa139ad7b38b6df3e3b17b8357c81e1fd6c))
* **kms:** start generating stubs dir ([de2d180](https://github.com/googleapis/google-cloud-go/commit/de2d18066dc613b72f6f8db93ca60146dabcfdcc))

## [1.4.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.3.0...kms/v1.4.0) (2022-02-23)


### Features

* **kms:** set versionClient to module version ([55f0d92](https://github.com/googleapis/google-cloud-go/commit/55f0d92bf112f14b024b4ab0076c9875a17423c9))

## [1.3.0](https://github.com/googleapis/google-cloud-go/compare/kms/v1.2.0...kms/v1.3.0) (2022-02-14)


### Features

* **kms:** add file for tracking version ([17b36ea](https://github.com/googleapis/google-cloud-go/commit/17b36ead42a96b1a01105122074e65164357519e))

## [1.2.0](https://www.github.com/googleapis/google-cloud-go/compare/kms/v1.1.0...kms/v1.2.0) (2022-02-04)


### Features

* **kms:** add a new EkmService API ([7f48e6b](https://www.github.com/googleapis/google-cloud-go/commit/7f48e6b68e59812208ea87b7861fad60169dc63a))

## [1.1.0](https://www.github.com/googleapis/google-cloud-go/compare/kms/v1.0.0...kms/v1.1.0) (2021-10-18)


### Features

* **kms:** add OAEP+SHA1 to the list of supported algorithms ([8c5c6cf](https://www.github.com/googleapis/google-cloud-go/commit/8c5c6cf9df046b67998a8608d05595bd9e34feb0))
* **kms:** add RPC retry information for MacSign, MacVerify, and GenerateRandomBytes Committer: [@bdhess](https://www.github.com/bdhess) ([1a0720f](https://www.github.com/googleapis/google-cloud-go/commit/1a0720f2f33bb14617f5c6a524946a93209e1266))
* **kms:** add support for Raw PKCS[#1](https://www.github.com/googleapis/google-cloud-go/issues/1) signing keys ([58bea89](https://www.github.com/googleapis/google-cloud-go/commit/58bea89a3d177d5c431ff19310794e3296253353))

## 1.0.0

Stabilize GA surface.

## [0.2.0](https://www.github.com/googleapis/google-cloud-go/compare/kms/v0.1.0...kms/v0.2.0) (2021-08-30)


### Features

* **kms:** add support for Key Reimport ([bf4378b](https://www.github.com/googleapis/google-cloud-go/commit/bf4378b5b859f7b835946891dbfebfee31c4b123))

## v0.1.0

This is the first tag to carve out kms as its own module. See
[Add a module to a multi-module repository](https://github.com/golang/go/wiki/Modules#is-it-possible-to-add-a-module-to-a-multi-module-repository).
