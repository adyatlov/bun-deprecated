# Changelog

## v1.1.0 (05.06.2018)

### Enhancements

* Now `bun` can read from `.gz` files; no need in `gunzip -r ./` anymore.
* The project got `CHANGELOG.md`.
* Added simple build script.

---

## v1.0.0 (04.06.2018)

### Initial Features

* Verifies that all hosts in the cluster have the same DC/OS version installed.
* Counts nodes of each type, checks if the amount of master nodes is odd.
* Checks if all DC/OS components are healthy.
