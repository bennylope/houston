# Daemon launch management for Mac OS

Houston is a port of the popular
[lunchy](https://github.com/eddiezane/lunchy/) program to Go.

## Usage

See the [lunchy](https://github.com/eddiezane/lunchy/) docs for complete
usage guidance.

### Supported commands

  - ls [-l] [pattern]
  - show [pattern]
  - edit [pattern]
  - status [pattern]

### Unsupported commands

  - start [-wF] [pattern]
  - stop [-w] [pattern]
  - restart [pattern]
  - install [-s] [file]
  - uninstall [pattern]

## Why?

Porting a small, useful program to another language is a good way to
develop skills in the new langauage, especially. lunchy is small, easy
to understand, and requires a lot of the basic functionality of any
command line utility. Most importantly, it's commands form a hierarchy
of complexity that allow the developer - me! - to get something useful
finished with short steps, building to a [hopefully] complete port of
the original program including refinements in my own code as I go.
