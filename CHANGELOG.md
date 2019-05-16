# Changelog

## v1.8.0 (16.05.2019)

### Enhancements

* File types are now defined via the YAML in the `filetypes/files_type_yaml.go`.
* Now you can add simple pattern search checks via extending
the YAML in the `checks/search_checks_yaml.go`.
* I added a few more checks.
* I added a first tool command `find-files`; it generates a YAML definition 
for file types.
* I changed `-l/--long` flag to `-v/--verbose`.

## v1.7.0 (12.09.2018)

### Enhancements

* New `unmount-volume` command checks if Mesos agent had problems unmounting local persistent volumes.
* New `FindFirstLine` function helps to search in logs.

### Bug fixes

* Bun launches a random subcommand instead of the specified one.

Kudos to Jan and Marvin for sponsoring this release with their company.

## v1.6.0 (10.09.2018)

### Enhancements

* I removed the check command to simplify usage.
* I added a CheckBuilder which dramatically simplified existing checks.
* I added the ReadFromJSON functions to simplify reading from JSON files.
* I got rid of some superficial things like contexts and progress reports.
* I abstracted the filesystem as the next step towards a test infrastructure.
* I documented some public types and functions.

## v1.5.0 (16.08.2018)

### Enhancements

* Alex created a new check, `mesos-actor-mailboxes`, it detects when Mesos actors are backlogged.
* Edgar added alternative names for Mesos actor mailboxes files (`__processess__`).
* I introduced the `-p` flag with which you can specify a path to a diagnostics bundle directory.

Kudos to the first contributors!

## v1.4.0 (04.07.2018)

### Enhancements

* I added the `--long` (`-l`) flag to the `check` command and its subcommands to be able to see detailed report even if the check is OK.

## v1.3.0 (18.06.2018)

### Enhancements

* Now you can specify several paths when describing a bundle file. It's useful for those files which names are different in different DC/OS versions.
* Starting from this version, Bun additionally searches the `health` file using the `3dt-health.json` path. It allows performing checks based on this file with bundles cr

## v1.2.0 (11.06.2018)

### Enhancements

* Introducing proper CLI support and the first command, `bun check`. Please, 
  try `bun --help` for details. 

---

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
