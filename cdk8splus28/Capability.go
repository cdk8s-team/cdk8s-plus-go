package cdk8splus28


// Capability - complete list of POSIX capabilities.
type Capability string

const (
	// CAP_AUDIT_CONTROL.
	Capability_AUDIT_CONTROL Capability = "AUDIT_CONTROL"
	// CAP_AUDIT_READ.
	Capability_AUDIT_READ Capability = "AUDIT_READ"
	// CAP_AUDIT_WRITE.
	Capability_AUDIT_WRITE Capability = "AUDIT_WRITE"
	// CAP_BLOCK_SUSPEND.
	Capability_BLOCK_SUSPEND Capability = "BLOCK_SUSPEND"
	// CAP_BPF.
	Capability_BPF Capability = "BPF"
	// CAP_CHECKPOINT_RESTORE.
	Capability_CHECKPOINT_RESTORE Capability = "CHECKPOINT_RESTORE"
	// CAP_CHOWN.
	Capability_CHOWN Capability = "CHOWN"
	// CAP_DAC_OVERRIDE.
	Capability_DAC_OVERRIDE Capability = "DAC_OVERRIDE"
	// CAP_DAC_READ_SEARCH.
	Capability_DAC_READ_SEARCH Capability = "DAC_READ_SEARCH"
	// CAP_FOWNER.
	Capability_FOWNER Capability = "FOWNER"
	// CAP_FSETID.
	Capability_FSETID Capability = "FSETID"
	// CAP_IPC_LOCK.
	Capability_IPC_LOCK Capability = "IPC_LOCK"
	// CAP_IPC_OWNER.
	Capability_IPC_OWNER Capability = "IPC_OWNER"
	// CAP_KILL.
	Capability_KILL Capability = "KILL"
	// CAP_LEASE.
	Capability_LEASE Capability = "LEASE"
	// CAP_LINUX_IMMUTABLE.
	Capability_LINUX_IMMUTABLE Capability = "LINUX_IMMUTABLE"
	// CAP_MAC_ADMIN.
	Capability_MAC_ADMIN Capability = "MAC_ADMIN"
	// CAP_MAC_OVERRIDE.
	Capability_MAC_OVERRIDE Capability = "MAC_OVERRIDE"
	// CAP_MKNOD.
	Capability_MKNOD Capability = "MKNOD"
	// CAP_NET_ADMIN.
	Capability_NET_ADMIN Capability = "NET_ADMIN"
	// CAP_NET_BIND_SERVICE.
	Capability_NET_BIND_SERVICE Capability = "NET_BIND_SERVICE"
	// CAP_NET_BROADCAST.
	Capability_NET_BROADCAST Capability = "NET_BROADCAST"
	// CAP_NET_RAW.
	Capability_NET_RAW Capability = "NET_RAW"
	// CAP_PERFMON.
	Capability_PERFMON Capability = "PERFMON"
	// CAP_SETGID.
	Capability_SETGID Capability = "SETGID"
	// CAP_SETFCAP.
	Capability_SETFCAP Capability = "SETFCAP"
	// CAP_SETPCAP.
	Capability_SETPCAP Capability = "SETPCAP"
	// CAP_SETUID.
	Capability_SETUID Capability = "SETUID"
	// CAP_SYS_ADMIN.
	Capability_SYS_ADMIN Capability = "SYS_ADMIN"
	// CAP_SYS_BOOT.
	Capability_SYS_BOOT Capability = "SYS_BOOT"
	// CAP_SYS_CHROOT.
	Capability_SYS_CHROOT Capability = "SYS_CHROOT"
	// CAP_SYS_MODULE.
	Capability_SYS_MODULE Capability = "SYS_MODULE"
	// CAP_SYS_NICE.
	Capability_SYS_NICE Capability = "SYS_NICE"
	// CAP_SYS_PACCT.
	Capability_SYS_PACCT Capability = "SYS_PACCT"
	// CAP_SYS_PTRACE.
	Capability_SYS_PTRACE Capability = "SYS_PTRACE"
	// CAP_SYS_RAWIO.
	Capability_SYS_RAWIO Capability = "SYS_RAWIO"
	// CAP_SYS_RESOURCE.
	Capability_SYS_RESOURCE Capability = "SYS_RESOURCE"
	// CAP_SYS_TIME.
	Capability_SYS_TIME Capability = "SYS_TIME"
	// CAP_SYS_TTY_CONFIG.
	Capability_SYS_TTY_CONFIG Capability = "SYS_TTY_CONFIG"
	// CAP_SYSLOG.
	Capability_SYSLOG Capability = "SYSLOG"
	// CAP_WAKE_ALARM.
	Capability_WAKE_ALARM Capability = "WAKE_ALARM"
)

