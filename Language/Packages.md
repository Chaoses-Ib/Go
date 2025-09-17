# Packages
> A package is a directory of `.go` files.

[The Go Programming Language Specification - The Go Programming Language](https://go.dev/ref/spec#Packages)

## Package clause
- Why do we need to write `package xxx` instead of just use the directory name?

  Just stupidity.

  [Why do we have to name packages? : r/golang](https://www.reddit.com/r/golang/comments/13nzqld/why_do_we_have_to_name_packages/)
  > I'm pretty sure you *could* have designed this all without requiring the `package` directive, without too much trouble. It probably just seemed like a good idea at the time and then was done that way. Honestly, a lot of Go's design decisions, especially the very early ones, come down to this. It was just a bunch of nerds sitting together in an office, excitedly talking about a new project after all. No one expected the language to become this big and get this scrutinized.

[Go packages must be named like parent folder? : r/golang](https://www.reddit.com/r/golang/comments/1hsgsv4/go_packages_must_be_named_like_parent_folder/)

## Import declarations
[Relative imports in Go - Stack Overflow](https://stackoverflow.com/questions/38517593/relative-imports-in-go)

[How to import local packages in go? - Stack Overflow](https://stackoverflow.com/questions/35480623/how-to-import-local-packages-in-go)
