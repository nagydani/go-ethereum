# Debian Packaging

Tagged releases and develop branch commits are available as installable Debian packages
for Ubuntu. Packages are built for the all Ubuntu versions which are supported by
Canonical:

- Trusty Tahr (14.04 LTS)
<<<<<<< HEAD
- Wily Werewolf (15.10)
- Xenial Xerus (16.04 LTS)
- Yakkety Yak (16.10)
=======
- Xenial Xerus (16.04 LTS)
- Yakkety Yak (16.10)
- Zesty Zapus (17.04)
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23

Packages of develop branch commits have suffix -unstable and cannot be installed alongside
the stable version. Switching between release streams requires user intervention.

The packages are built and served by launchpad.net. We generate a Debian source package
for each distribution and upload it. Their builder picks up the source package, builds it
and installs the new version into the PPA repository. Launchpad requires a valid signature
by a team member for source package uploads. The signing key is stored in an environment
variable which Travis CI makes available to certain builds.

We want to build go-ethereum with the most recent version of Go, irrespective of the Go
version that is available in the main Ubuntu repository. In order to make this possible,
our PPA depends on the ~gophers/ubuntu/archive PPA. Our source package build-depends on
<<<<<<< HEAD
golang-1.7, which is co-installable alongside the regular golang package. PPA dependencies
can be edited at https://launchpad.net/%7Elp-fjl/+archive/ubuntu/geth-ci-testing/+edit-dependencies
=======
golang-1.9, which is co-installable alongside the regular golang package. PPA dependencies
can be edited at https://launchpad.net/%7Eethereum/+archive/ubuntu/ethereum/+edit-dependencies
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23

## Building Packages Locally (for testing)

You need to run Ubuntu to do test packaging.

<<<<<<< HEAD
Add the gophers PPA and install Go 1.7 and Debian packaging tools:

    $ sudo apt-add-repository ppa:gophers/ubuntu/archive
    $ sudo apt-get update
    $ sudo apt-get install build-essential golang-1.7 devscripts debhelper
=======
Add the gophers PPA and install Go 1.9 and Debian packaging tools:

    $ sudo apt-add-repository ppa:gophers/ubuntu/archive
    $ sudo apt-get update
    $ sudo apt-get install build-essential golang-1.9 devscripts debhelper
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23

Create the source packages:

    $ go run build/ci.go debsrc -workdir dist

Then go into the source package directory for your running distribution and build the package:

<<<<<<< HEAD
    $ cd dist/ethereum-unstable-1.5.0+xenial
=======
    $ cd dist/ethereum-unstable-1.6.0+xenial
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
    $ dpkg-buildpackage

Built packages are placed in the dist/ directory.

    $ cd ..
<<<<<<< HEAD
    $ dpkg-deb -c geth-unstable_1.5.0+xenial_amd64.deb
=======
    $ dpkg-deb -c geth-unstable_1.6.0+xenial_amd64.deb
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
