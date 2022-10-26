# About

`hoshinovactl` is a small CLI to control [Hoshinova](https://github.com/HoloArchivists/hoshinova) remotely.

## Build

```shell
make
make install
```

## Usage

To see all commands and flags, run `hoshinovactl -h`

Example of listing all tasks:

```shell
λ hoshinovactl task list
VideoID      State          Total Size
------------ -------------- ----------------
-kAwbEDuikQ  waiting
```

