# Third Eye

[![GoDoc][doc-img]][doc] [![Github release][release-img]][release] [![Build Status][ci-img]][ci]  [![Go Report Card][report-card-img]][report-card]

<!-- `[![Coverage Status][cov-img]][cov]` -->

Third-eye is used for getting gearbox contracts-related data from evm compatible chains, calculating interesting parameters, and storing them. Architecture docs are available [here](https://github.com/Gearbox-protocol/third-eye/blob/master/docs).

### Instructions to run
```
migrate -path migrations -database "$DURL" up
go build ./cmd/main.go && ./main
```

It's also possible to run migrate with docker:
```
docker run -v $(pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "${url}" up
```

## Licensing

The primary license for the Gearbox-Contracts is the Business Source License 1.1 (BUSL-1.1), see [LICENSE](https://github.com/Gearbox-protocol/third-eye/blob/master/LICENSE). The files licensed under the BUSL-1.1 have appropriate SPDX headers.

## Disclaimer

This application is provided "as is" and "with all faults." Me as developer makes no representations or
warranties of any kind concerning the safety, suitability, lack of viruses, inaccuracies, typographical
errors, or other harmful components of this software. There are inherent dangers in the use of any software,
and you are solely responsible for determining whether this software product is compatible with your equipment and
other software installed on your equipment. You are also solely responsible for the protection of your equipment
and backup of your data, and THE PROVIDER will not be liable for any damages you may suffer in connection with using,
modifying, or distributing this software product.



[doc-img]: http://img.shields.io/badge/GoDoc-Reference-blue.svg
[doc]: https://github.com/Gearbox-protocol/third-eye/blob/master/docs

[release-img]: https://img.shields.io/github/v/release/Gearbox-protocol/third-eye.svg
[release]: https://github.com/Gearbox-protocol/third-eye/releases

[ci-img]: https://github.com/Gearbox-protocol/third-eye/actions/workflows/go.yml/badge.svg
[ci]: https://github.com/Gearbox-protocol/third-eye/actions/workflows/go.yml

[cov-img]: https://codecov.io/gh/Gearbox-protocol/third-eye/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/Gearbox-protocol/third-eye/branch/master

[report-card-img]: https://goreportcard.com/badge/github.com/Gearbox-protocol/third-eye
[report-card]: https://goreportcard.com/report/github.com/Gearbox-protocol/third-eye



### Important information for contributors
As a contributor to the Gearbox Protocol GitHub repository, your pull requests indicate acceptance of our Gearbox Contribution Agreement. This agreement outlines that you assign the Intellectual Property Rights of your contributions to the Gearbox Foundation. This helps safeguard the Gearbox protocol and ensure the accumulation of its intellectual property. Contributions become part of the repository and may be used for various purposes, including commercial. As recognition for your expertise and work, you receive the opportunity to participate in the protocol's development and the potential to see your work integrated within it. The full Gearbox Contribution Agreement is accessible within the [repository](/ContributionAgreement) for comprehensive understanding. [Let's innovate together!]