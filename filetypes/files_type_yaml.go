package filetypes

const filesYAML = `
- name: active-buildinfo-full
  contentType: JSON
  paths:
  - opt/mesosphere/active.buildinfo.full.json
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: adminrouter
  contentType: journal
  paths:
  - dcos-adminrouter.service
  description: ""
  dirTypes:
  - master
- name: adminrouter-agent
  contentType: journal
  paths:
  - dcos-adminrouter-agent.service
  description: ""
  dirTypes:
  - agent
  - public agent
- name: backup-master
  contentType: journal
  paths:
  - dcos-backup-master.service
  description: ""
  dirTypes:
  - master
- name: backup-master-socket
  contentType: journal
  paths:
  - dcos-backup-master.socket
  description: ""
  dirTypes:
  - master
- name: binsh-c-cat-etc*-release
  contentType: output
  paths:
  - binsh_-c_cat etc*-release.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: binsh-c-cat-proc`+ "`" +`systemctl-show-dcos-mesos-agent-service-p-main-pid|-cut-d'='-f2`+ "`" +`environ
  contentType: output
  paths:
  - binsh_-c_cat proc`+ "`" +`systemctl show dcos-mesos-slave.service -p MainPID| cut -d'='
    -f2`+ "`" +`environ.output
  description: ""
  dirTypes:
  - agent
  - public agent
- name: binsh-c-cat-proc`+ "`" +`systemctl-show-dcos-mesos-master-service-p-main-pid|-cut-d'='-f2`+ "`" +`environ
  contentType: output
  paths:
  - binsh_-c_cat proc`+ "`" +`systemctl show dcos-mesos-master.service -p MainPID| cut -d'='
    -f2`+ "`" +`environ.output
  description: ""
  dirTypes:
  - master
- name: bouncer
  contentType: journal
  paths:
  - dcos-bouncer.service
  description: ""
  dirTypes:
  - master
- name: ca
  contentType: journal
  paths:
  - dcos-ca.service
  description: ""
  dirTypes:
  - master
- name: checks-api
  contentType: journal
  paths:
  - dcos-checks-api.service
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: checks-api-socket
  contentType: journal
  paths:
  - dcos-checks-api.socket
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: checks-poststart
  contentType: journal
  paths:
  - dcos-checks-poststart.service
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: checks-poststart-timer
  contentType: journal
  paths:
  - dcos-checks-poststart.timer
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: cluster-id
  contentType: other
  paths:
  - var/lib/dcos/cluster-id
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: cluster-linker
  contentType: journal
  paths:
  - dcos-cluster-linker.service
  description: ""
  dirTypes:
  - master
- name: cluster-linker-socket
  contentType: journal
  paths:
  - dcos-cluster-linker.socket
  description: ""
  dirTypes:
  - master
- name: cmdline
  contentType: other
  paths:
  - proc/cmdline
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: cockroach
  contentType: journal
  paths:
  - dcos-cockroach.service
  description: ""
  dirTypes:
  - master
- name: cockroachdb-config-change
  contentType: journal
  paths:
  - dcos-cockroachdb-config-change.service
  description: ""
  dirTypes:
  - master
- name: cockroachdb-config-change-timer
  contentType: journal
  paths:
  - dcos-cockroachdb-config-change.timer
  description: ""
  dirTypes:
  - master
- name: cosmos
  contentType: journal
  paths:
  - dcos-cosmos.service
  description: ""
  dirTypes:
  - master
- name: cpuinfo
  contentType: other
  paths:
  - proc/cpuinfo
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: diagnostics
  contentType: journal
  paths:
  - dcos-diagnostics.service
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: diagnostics-health
  contentType: JSON
  paths:
  - dcos-diagnostics-health.json
  - 3dt-health.json
  description: "contains health of systemd services corresponding to DC/OS components."
  dirTypes:
  - master
  - agent
  - public agent
- name: diagnostics-socket
  contentType: journal
  paths:
  - dcos-diagnostics.socket
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: dmesg-t
  contentType: dmesg
  paths:
  - dmesg_-T.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: docker-gc
  contentType: journal
  paths:
  - dcos-docker-gc.service
  description: ""
  dirTypes:
  - agent
  - public agent
- name: docker-gc-timer
  contentType: journal
  paths:
  - dcos-docker-gc.timer
  description: ""
  dirTypes:
  - agent
  - public agent
- name: docker-ps
  contentType: output
  paths:
  - docker_ps.output
  description: ""
  dirTypes:
  - agent
  - public agent
- name: docker-version
  contentType: output
  paths:
  - docker_--version.output
  description: ""
  dirTypes:
  - agent
  - public agent
- name: exhibitor-log
  contentType: journal
  paths:
  - dcos-exhibitor.service
  description: ""
  dirTypes:
  - master
- name: exhibitor-cluster-list
  contentType: JSON
  paths:
  - 443-exhibitor_exhibitor_v1_cluster_list.json
  description: ""
  dirTypes:
  - master
- name: exhibitor-cluster-log
  contentType: JSON
  paths:
  - 443-exhibitor_exhibitor_v1_cluster_log.json
  description: ""
  dirTypes:
  - master
- name: exhibitor-cluster-state
  contentType: JSON
  paths:
  - 443-exhibitor_exhibitor_v1_cluster_state.json
  description: ""
  dirTypes:
  - master
- name: exhibitor-cluster-status
  contentType: JSON
  paths:
  - 443-exhibitor_exhibitor_v1_cluster_status.json
  description: ""
  dirTypes:
  - master
- name: exhibitor-config-get-state
  contentType: JSON
  paths:
  - 443-exhibitor_exhibitor_v1_config_get-state.json
  description: ""
  dirTypes:
  - master
- name: expanded-config
  contentType: JSON
  paths:
  - opt/mesosphere/etc/expanded.config.json
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: fib-trie
  contentType: other
  paths:
  - proc/net/fib_trie
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: gen-resolvconf
  contentType: journal
  paths:
  - dcos-gen-resolvconf.service
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: gen-resolvconf-timer
  contentType: journal
  paths:
  - dcos-gen-resolvconf.timer
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: history
  contentType: journal
  paths:
  - dcos-history.service
  description: ""
  dirTypes:
  - master
- name: history-service-hour
  contentType: JSON
  paths:
  - 443-dcos-history-service_history_hour.json
  description: ""
  dirTypes:
  - master
- name: history-service-last
  contentType: JSON
  paths:
  - 443-dcos-history-service_history_last.json
  description: ""
  dirTypes:
  - master
- name: history-service-minute
  contentType: JSON
  paths:
  - 443-dcos-history-service_history_minute.json
  description: ""
  dirTypes:
  - master
- name: iam-ldap-sync
  contentType: journal
  paths:
  - dcos-iam-ldap-sync.service
  description: ""
  dirTypes:
  - master
- name: iam-ldap-sync-timer
  contentType: journal
  paths:
  - dcos-iam-ldap-sync.timer
  description: ""
  dirTypes:
  - master
- name: ifconfig-a
  contentType: output
  paths:
  - ifconfig -a.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: ip-addr
  contentType: output
  paths:
  - ip_addr.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: ip-route
  contentType: output
  paths:
  - ip_route.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: ip-vs
  contentType: other
  paths:
  - proc/net/ip_vs
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: ip-vs-conn
  contentType: other
  paths:
  - proc/net/ip_vs_conn
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: iptables-save
  contentType: output
  paths:
  - iptables-save.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: jobs
  contentType: JSON
  paths:
  - 9443-v1_jobs.json
  description: ""
  dirTypes:
  - master
- name: licensing
  contentType: journal
  paths:
  - dcos-licensing.service
  description: ""
  dirTypes:
  - master
- name: licensing-audit-decrypt
  contentType: JSON
  paths:
  - 443-licensing_v1_audit_decrypt_1.json
  description: ""
  dirTypes:
  - master
- name: licensing-socket
  contentType: journal
  paths:
  - dcos-licensing.socket
  description: ""
  dirTypes:
  - master
- name: log-agent
  contentType: journal
  paths:
  - dcos-log-agent.service
  description: ""
  dirTypes:
  - agent
  - public agent
- name: log-agent-socket
  contentType: journal
  paths:
  - dcos-log-agent.socket
  description: ""
  dirTypes:
  - agent
  - public agent
- name: log-master
  contentType: journal
  paths:
  - dcos-log-master.service
  description: ""
  dirTypes:
  - master
- name: log-master-socket
  contentType: journal
  paths:
  - dcos-log-master.socket
  description: ""
  dirTypes:
  - master
- name: logrotate-agent
  contentType: journal
  paths:
  - dcos-logrotate-agent.service
  description: ""
  dirTypes:
  - agent
  - public agent
- name: logrotate-agent-timer
  contentType: journal
  paths:
  - dcos-logrotate-agent.timer
  description: ""
  dirTypes:
  - agent
  - public agent
- name: logrotate-master
  contentType: journal
  paths:
  - dcos-logrotate-master.service
  description: ""
  dirTypes:
  - master
- name: logrotate-master-timer
  contentType: journal
  paths:
  - dcos-logrotate-master.timer
  description: ""
  dirTypes:
  - master
- name: marathon
  contentType: journal
  paths:
  - dcos-marathon.service
  description: ""
  dirTypes:
  - master
- name: marathon-apps
  contentType: JSON
  paths:
  - 8443-v2_apps.json
  description: ""
  dirTypes:
  - master
- name: marathon-deployments
  contentType: JSON
  paths:
  - 8443-v2_deployments.json
  - 8443:v2_deployments.json
  description: "Marathon application deployments"
  dirTypes:
  - master
- name: marathon-groups
  contentType: JSON
  paths:
  - 8443-v2_groups.json
  description: ""
  dirTypes:
  - master
- name: marathon-info
  contentType: JSON
  paths:
  - 8443-v2_info.json
  description: ""
  dirTypes:
  - master
- name: marathon-leader
  contentType: JSON
  paths:
  - 8443-v2_leader.json
  description: ""
  dirTypes:
  - master
- name: marathon-pods
  contentType: JSON
  paths:
  - 8443-v2_pods.json
  description: ""
  dirTypes:
  - master
- name: marathon-queue
  contentType: JSON
  paths:
  - 8443-v2_queue.json
  description: ""
  dirTypes:
  - master
- name: marathon-tasks
  contentType: JSON
  paths:
  - 8443-v2_tasks.json
  description: ""
  dirTypes:
  - master
- name: meminfo
  contentType: other
  paths:
  - proc/meminfo
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: mesos-agent-var-log
  contentType: other
  paths:
  - var/log/mesos/mesos-agent.log
  description: ""
  dirTypes:
  - agent
  - public agent
- name: mesos-agent-log
  contentType: journal
  paths:
  - dcos-mesos-slave.service
  - dcos-mesos-slave-public.service
  description: "Mesos agent jounrald log"
  dirTypes:
  - agent
  - public agent
- name: mesos-agent-containers
  contentType: JSON
  paths:
  - 5051-containers.json
  description: ""
  dirTypes:
  - agent
  - public agent
- name: mesos-agent-flags
  contentType: JSON
  paths:
  - 5051-flags.json
  description: ""
  dirTypes:
  - agent
  - public agent
- name: mesos-agent-metrics-snapshot
  contentType: JSON
  paths:
  - 5051-metrics_snapshot.json
  description: ""
  dirTypes:
  - agent
  - public agent
- name: mesos-agent-overlay
  contentType: JSON
  paths:
  - 5051-overlay-agent_overlay.json
  description: ""
  dirTypes:
  - agent
  - public agent
- name: mesos-agent-processes
  contentType: JSON
  paths:
  - 5051-__processes__.json
  description: ""
  dirTypes:
  - agent
  - public agent
- name: mesos-agent-public
  contentType: journal
  paths:
  - dcos-mesos-slave-public.service
  description: ""
  dirTypes:
  - public agent
- name: mesos-agent-state
  contentType: JSON
  paths:
  - 5051-state.json
  description: ""
  dirTypes:
  - agent
  - public agent
- name: mesos-agent-system-stats
  contentType: JSON
  paths:
  - 5051-system_stats_json.json
  description: ""
  dirTypes:
  - agent
  - public agent
- name: mesos-dns
  contentType: journal
  paths:
  - dcos-mesos-dns.service
  description: ""
  dirTypes:
  - master
- name: mesos-dns-config
  contentType: JSON
  paths:
  - 443-mesos_dns_v1_config.json
  description: ""
  dirTypes:
  - master
- name: mesos-dns-version
  contentType: JSON
  paths:
  - 443-mesos_dns_v1_version.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-var-log
  contentType: other
  paths:
  - var/lib/dcos/mesos/log/mesos-master.log
  description: ""
  dirTypes:
  - master
- name: mesos-master-log
  contentType: journal
  paths:
  - dcos-mesos-master.service
  description: ""
  dirTypes:
  - master
- name: mesos-master-agents
  contentType: JSON
  paths:
  - 5050-master_slaves.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-flags
  contentType: JSON
  paths:
  - 5050-master_flags.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-frameworks
  contentType: JSON
  paths:
  - 5050-master_frameworks.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-maintenance-schedule
  contentType: JSON
  paths:
  - 5050-master_maintenance_schedule.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-maintenance-status
  contentType: JSON
  paths:
  - 5050-master_maintenance_status.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-metrics-snapshot
  contentType: JSON
  paths:
  - 5050-metrics_snapshot.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-overlay-state
  contentType: JSON
  paths:
  - 5050-overlay-master_state.json
  description: ""
  dirTypes:
  - master
- name: mesos-processes
  contentType: JSON
  paths:
  - 5050-__processes__.json
  - 5051-__processes__.json
  - 5050:__processes__.json
  - 5051:__processes__.json
  description: "Contains mailbox contents for all actors in the Mesos process on the host."
  dirTypes:
  - master
  - agent
  - public agent
- name: mesos-master-quota
  contentType: JSON
  paths:
  - 5050-quota.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-registrar-1-registry
  contentType: JSON
  paths:
  - 5050-registrar_1__registry.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-roles
  contentType: JSON
  paths:
  - 5050-master_roles.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-state
  contentType: JSON
  paths:
  - 5050-master_state.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-state-summary
  contentType: JSON
  paths:
  - 5050-master_state-summary.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-system-stats
  contentType: JSON
  paths:
  - 5050-system_stats_json.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-tasks
  contentType: JSON
  paths:
  - 5050-master_tasks.json
  description: ""
  dirTypes:
  - master
- name: mesos-master-version
  contentType: JSON
  paths:
  - 5050-version.json
  description: ""
  dirTypes:
  - master
- name: metronome
  contentType: journal
  paths:
  - dcos-metronome.service
  description: ""
  dirTypes:
  - master
- name: mountinfo
  contentType: other
  paths:
  - proc/self/mountinfo
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: myid
  contentType: other
  paths:
  - var/lib/dcos/exhibitor/zookeeper/snapshot/myid
  description: ""
  dirTypes:
  - master
- name: net
  contentType: journal
  paths:
  - dcos-net.service
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: net-watchdog
  contentType: journal
  paths:
  - dcos-net-watchdog.service
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: optmesospherebincurl-s-http:localhost:62080v1vips
  contentType: output
  paths:
  - optmesospherebincurl_-s_-S_http:localhost:62080v1vips.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: optmesospherebincurl-s-https:localhost:8090-adminv1health
  contentType: output
  paths:
  - optmesospherebincurl_-s_-S_https:localhost:8090_adminv1health.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: optmesospherebincurl-s-https:localhost:8090-statusnodes
  contentType: output
  paths:
  - optmesospherebincurl_-s_-S_https:localhost:8090_statusnodes.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: optmesospherebincurl-s-https:localhost:8090health
  contentType: output
  paths:
  - optmesospherebincurl_-s_-S_https:localhost:8090health.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: optmesospherebindetect-ip
  contentType: output
  paths:
  - optmesospherebindetect_ip.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: pkgpanda-api
  contentType: journal
  paths:
  - dcos-pkgpanda-api.service
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: ps-aux-ww-z
  contentType: output
  paths:
  - ps_aux_ww_Z.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: registry
  contentType: journal
  paths:
  - dcos-registry.service
  description: ""
  dirTypes:
  - master
- name: resolv
  contentType: other
  paths:
  - etc/resolv.conf
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: rexray
  contentType: journal
  paths:
  - dcos-rexray.service
  description: ""
  dirTypes:
  - agent
  - public agent
- name: secrets
  contentType: journal
  paths:
  - dcos-secrets.service
  description: ""
  dirTypes:
  - master
- name: secrets-socket
  contentType: journal
  paths:
  - dcos-secrets.socket
  description: ""
  dirTypes:
  - master
- name: sestatus
  contentType: output
  paths:
  - sestatus.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: signal
  contentType: journal
  paths:
  - dcos-signal.service
  description: ""
  dirTypes:
  - master
- name: signal-timer
  contentType: journal
  paths:
  - dcos-signal.timer
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: summary-errors-report
  contentType: other
  paths:
  - summaryErrorsReport.txt
  description: ""
  dirTypes:
  - root
- name: summary-report
  contentType: other
  paths:
  - summaryReport.txt
  description: ""
  dirTypes:
  - root
- name: systemctl-list-units-dcos*
  contentType: output
  paths:
  - systemctl_list-units_dcos*.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: telegraf
  contentType: journal
  paths:
  - dcos-telegraf.service
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: telegraf-socket
  contentType: journal
  paths:
  - dcos-telegraf.socket
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: timedatectl
  contentType: output
  paths:
  - timedatectl.output
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: user-config
  contentType: other
  paths:
  - opt/mesosphere/etc/user.config.yaml
  description: ""
  dirTypes:
  - master
  - agent
  - public agent
- name: vault
  contentType: journal
  paths:
  - dcos-vault.service
  description: ""
  dirTypes:
  - master
- name: dcos-version
  contentType: JSON
  paths:
  - opt/mesosphere/etc/dcos-version.json
  description: "contains DC/OS version, DC/OS image commit and bootstrap ID."
  dirTypes:
  - master
  - agent
  - public agent
- name: zoo
  contentType: other
  paths:
  - var/lib/dcos/exhibitor/conf/zoo.cfg
  description: ""
  dirTypes:
  - master
`
