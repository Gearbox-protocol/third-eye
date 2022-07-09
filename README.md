# Third Eye

[![GoDoc][doc-img]][doc] [![Github release][release-img]][release] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report Card][report-card-img]][report-card]


Third-eye is used for getting gearbox contracts-related data from evm compatible chains, calculating interesting parameters, and storing them. Architecture docs are available [here](https://github.com/Gearbox-protocol/third-eye/blob/master/docs).

### Instructions to run
```
migrate -path migrations -database "$DURL" up
go build ./cmd/main.go && ./main
```

## Testing

```
go test ./tests && go test ./ds
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