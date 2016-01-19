Go natsort
=====

[![Build Status](https://travis-ci.org/miraclesu/natsort.svg?branch=master)](https://travis-ci.org/miraclesu/natsort)
[![Coverage Status](https://coveralls.io/repos/miraclesu/natsort/badge.svg?branch=master&service=github&_day=201606)](https://coveralls.io/github/miraclesu/natsort?branch=master)

Natural sorting for Go.

Usage
-------

	versions := []string{"1.6beta10", "1.6", "1.6beta2"}
	// [1.6beta2 1.6beta10 1.6]
	fmt.Printf("%+v\n", natsort.Sort(versions))

	versions := []string{"version-1.9", "version-2.0", "version-1.11", "version-1.10"}
	// [version-1.9 version-1.10 version-1.11 version-2.0]
    fmt.Printf("%+v\n", natsort.Sort(versions))

	versions := []string{"Apple", "Banana", "apple", "banana"}
	// [Apple Banana apple banana]
    fmt.Printf("%+v\n", natsort.Sort(versions))

	versions := []string{"num5.10", "num-3", "num5.3", "num2"}
	// [num-3 num2 num5.3 num5.10]
    fmt.Printf("%+v\n", natsort.Sort(versions))
	// [num5.10 num5.3 num2 num-3]
    fmt.Printf("%+v\n", natsort.RSort(versions))

Inspired by
-------
* [natsort](https://github.com/SethMMorton/natsort)

License
-------

[MIT](LICENSE)
