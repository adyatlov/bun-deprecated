package checks

const searchChecksYAML = `
- name: unmount-volume
  description: Checks if Mesos agents had problems unmounting local persistent volumes. MESOS-8830
  fileTypeName: mesos-agent-log
  searchString: Failed to remove rootfs mount point
- name: disk-space-exhibitor
  description: Check disk space errors in Exhibitor logs
  fileTypeName: exhibitor-log
  searchString: No space left on device
- name: migration-in-progress
  description: Detect marathon-upgrade-in-progress flag on failed cluster after upgrade
  fileTypeName: marathon
  searchString: 'Migration Failed: Migration is already in progress'
`
