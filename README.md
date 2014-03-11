blend
=====

This package provides support for reading [blend][1] files, which are used by
[Blender][].

[1]: http://www.blender.org/development/architecture/blender-file-format/
[Blender]: http://www.blender.org/

Documentation
-------------

Documentation provided by GoDoc.

- [blend][]: implements parsing of Blender files.
	- [block][]: implements parsing of blend file blocks.

[blend]: http://godoc.org/github.com/mewmew/blend
[block]: http://godoc.org/github.com/mewmew/blend/block

Installation
------------

	go get github.com/mewmew/blend

To parse more complex blend files "block/struct.go" and "block/parse.go" may
have to be regenereted. To regenerate these two files for any given blend file,
use the *blendef* tool.

	go get github.com/mewmew/blend/cmd/blendef
	cd $GOPATH/src/github.com/mewmew/blend/block
	blendef /path/to/complex.blend

A more detailed description is given under the "self-describing format" section.

Examples
--------

* Extract an embedded thumbnail from a blend file.

		go get github.com/mewmew/blend/examples/blendview
		cd $GOPATH/src/github.com/mewmew/blend/testdata
		blendview block.blend

![Extracted thumbnail](https://github.com/mewmew/blend/blob/master/examples/blendview/block.png?raw=true)

* Parse a single block in a blend file.

	http://godoc.org/github.com/mewmew/blend#example_Blend

* Parse all blocks in a blend file and access the data of a pointer.

	http://godoc.org/github.com/mewmew/blend/block#example_Pointer

Self-describing format
----------------------

One unique feature of blend files is that they contain a full definition of
every structure used in it's file blocks. The structure definitions are stored
in the DNA block.

All block structure definitions ("block/struct.go") and the block parsing logic
("block/parse.go") have been generating by parsing the DNA block of
"testdata/block.blend".

The tool which was used to generate these two files is located at:

	github.com/mewmew/blend/cmd/blendef

More complex blend files may contain structures which are not yet defined in the
block package. If so, use *blendef* to regenerate "struct.go" and "parse.go" for
the given blend file.

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
