bazel run //:gazelle -- update-repos -from_file=server-go/go.mod
bazel run //:gazelle -- update-repos -from_file=client-go/go.mod
bazel run //:gazelle
bazel build //...