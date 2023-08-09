# Delete an image or repo from a local docker registry

When using [distribution](https://github.com/distribution/distribution/tree/main/registry)
as the docker image `registry:v2`, you must take the following steps to delete
an image:

1. Start the `registry:v2` container with the environment variable
   `REGISTRY_STORAGE_DELETE_ENABLED=true`
2. Find the checksum of the image. This can be done by running garbage collect
   in dry run mode and looking for the `manifest:` entry. I haven't found a way to
   find this information from the API yet, though I'm sure there is one.
3. Delete the image manifest using the API. You must use the checksum, and this
   will not actually free any disk space:

   ```bash
   curl -X DELETE https://registry.galaxy.sc/v2/i/frigate/manifests/sha256:abcde1234
   ```

4. `exec` into the container and free the disk space by running:

   ```bash
   registry garbage-collect /etc/docker/registry/config.yml
   ```

> NOTE: Add the `--dry-run` flag to the `registry garbage-collect` command to
see what will be deleted.

## Deleting unreferenced blobs

To delete all "orphan" blobs which are not referenced by a tag, exec into the container and run:

```bash
registry garbage-collect --delete-untagged=true /etc/docker/registry/config.yml

    #docker #registry
