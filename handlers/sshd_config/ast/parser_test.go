package ast

import (
	"config-lsp/utils"
	"testing"
)

func TestSimpleParserExample(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin no
PasswordAuthentication yes
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 2 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 2 options and no comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHDOption)

	if !(firstEntry.Value.Value == "PermitRootLogin no" &&
		firstEntry.LocationRange.Start.Line == 0 &&
		firstEntry.LocationRange.End.Line == 0 &&
		firstEntry.LocationRange.Start.Character == 0 &&
		firstEntry.LocationRange.End.Character == 18 &&
		firstEntry.Key.Value.Value == "PermitRootLogin" &&
		firstEntry.Key.LocationRange.Start.Character == 0 &&
		firstEntry.Key.LocationRange.End.Character == 15 &&
		firstEntry.OptionValue.Value.Value == "no" &&
		firstEntry.OptionValue.LocationRange.Start.Character == 16 &&
		firstEntry.OptionValue.LocationRange.End.Character == 18) {
		t.Errorf("Expected first entry to be PermitRootLogin no, but got: %v", firstEntry)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(1))
	secondEntry := rawSecondEntry.(*SSHDOption)

	if !(secondEntry.Value.Value == "PasswordAuthentication yes" &&
		secondEntry.LocationRange.Start.Line == 1 &&
		secondEntry.LocationRange.End.Line == 1 &&
		secondEntry.LocationRange.Start.Character == 0 &&
		secondEntry.LocationRange.End.Character == 26 &&
		secondEntry.Key.Value.Value == "PasswordAuthentication" &&
		secondEntry.Key.LocationRange.Start.Character == 0 &&
		secondEntry.Key.LocationRange.End.Character == 22 &&
		secondEntry.OptionValue.Value.Value == "yes" &&
		secondEntry.OptionValue.LocationRange.Start.Character == 23 &&
		secondEntry.OptionValue.LocationRange.End.Character == 26) {
		t.Errorf("Expected second entry to be PasswordAuthentication yes, but got: %v", secondEntry)
	}
}

func TestMatchSimpleBlock(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin no

Match Address 192.168.0.1
	PasswordAuthentication yes
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 2 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 1 option and no comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHDOption)
	if !(firstEntry.Value.Value == "PermitRootLogin no") {
		t.Errorf("Expected first entry to be 'PermitRootLogin no', but got: %v", firstEntry.Value)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(2))
	secondEntry := rawSecondEntry.(*SSHDMatchBlock)
	if !(secondEntry.MatchEntry.Value.Value == "Match Address 192.168.0.1") {
		t.Errorf("Expected second entry to be 'Match Address 192.168.0.1', but got: %v", secondEntry.MatchEntry.Value)
	}

	if !(secondEntry.Start.Line == 2 && secondEntry.Start.Character == 0 && secondEntry.End.Line == 3 && secondEntry.End.Character == 27) {
		t.Errorf("Expected second entry's location to be 2:0-3:25, but got: %v", secondEntry.LocationRange)
	}

	if !(secondEntry.MatchValue.Entries[0].Criteria.Type == "Address" && secondEntry.MatchValue.Entries[0].Values.Values[0].Value.Value == "192.168.0.1" && secondEntry.MatchEntry.OptionValue.Start.Character == 6) {
		t.Errorf("Expected second entry to be 'Match Address 192.168.0.1', but got: %v", secondEntry.MatchValue)
	}

	if !(secondEntry.Options.Size() == 1) {
		t.Errorf("Expected 1 option in match block, but got: %v", secondEntry.Options)
	}

	rawThirdEntry, _ := secondEntry.Options.Get(uint32(3))
	thirdEntry := rawThirdEntry.(*SSHDOption)
	if !(thirdEntry.Key.Value.Value == "PasswordAuthentication" && thirdEntry.OptionValue.Value.Value == "yes") {
		t.Errorf("Expected third entry to be 'PasswordAuthentication yes', but got: %v", thirdEntry.Value)
	}
}

func TestMultipleEntriesInMatchBlock(
	t *testing.T,
) {
	input := utils.Dedent(`
Match User lena User root
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	_, matchBlock := p.FindOption(uint32(0))

	if !(matchBlock.MatchEntry.Value.Value == "Match User lena User root") {
		t.Errorf("Expected match block to be 'Match User lena User root', but got: %v", matchBlock.MatchEntry.Value)
	}

	if !(len(matchBlock.MatchValue.Entries) == 2) {
		t.Errorf("Expected 2 entries in match block, but got: %v", matchBlock.MatchValue.Entries)
	}

	if !(matchBlock.MatchValue.Entries[0].Criteria.Type == "User" && matchBlock.MatchValue.Entries[0].Values.Values[0].Value.Value == "lena") {
		t.Errorf("Expected first entry to be 'User lena', but got: %v", matchBlock.MatchValue.Entries[0])
	}

	if !(matchBlock.MatchValue.Entries[1].Criteria.Type == "User" && matchBlock.MatchValue.Entries[1].Values.Values[0].Value.Value == "root") {
		t.Errorf("Expected second entry to be 'User root', but got: %v", matchBlock.MatchValue.Entries[1])
	}
}

func TestIncompleteMatchBlock(
	t *testing.T,
) {
	input := "Match User lena User "

	p := NewSSHConfig()
	errors := p.Parse(input)

	if !(len(errors) == 0) {
		t.Errorf("Expected 0 error, got %v", errors)
	}

	_, matchBlock := p.FindOption(uint32(0))

	if !(matchBlock.MatchEntry.Value.Value == "Match User lena User ") {
		t.Errorf("Expected match block to be 'Match User lena User ', but got: %v", matchBlock.MatchEntry.Value)
	}

	if !(matchBlock.MatchValue.Entries[0].Criteria.Type == "User" && matchBlock.MatchValue.Entries[0].Values.Values[0].Value.Value == "lena") {
		t.Errorf("Expected first entry to be 'User lena', but got: %v", matchBlock.MatchValue.Entries[0])
	}

	if !(matchBlock.MatchValue.Entries[1].Value.Value == "User " && matchBlock.MatchValue.Entries[1].Criteria.Type == "User" && matchBlock.MatchValue.Entries[1].Values == nil) {
		t.Errorf("Expected second entry to be 'User ', but got: %v", matchBlock.MatchValue.Entries[1])
	}
}

func TestMultipleMatchBlocks(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin no

Match User lena
	PasswordAuthentication yes
	AllowUsers root user

Match Address 192.168.0.2
	MaxAuthTries 3
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 3 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 3 options and no comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(2))
	secondEntry := rawSecondEntry.(*SSHDMatchBlock)
	if !(secondEntry.Options.Size() == 2) {
		t.Errorf("Expected 2 options in second match block, but got: %v", secondEntry.Options)
	}

	rawThirdEntry, _ := secondEntry.Options.Get(uint32(3))
	thirdEntry := rawThirdEntry.(*SSHDOption)
	if !(thirdEntry.Key.Value.Value == "PasswordAuthentication" && thirdEntry.OptionValue.Value.Value == "yes" && thirdEntry.LocationRange.Start.Line == 3) {
		t.Errorf("Expected third entry to be 'PasswordAuthentication yes', but got: %v", thirdEntry.Value)
	}

	rawFourthEntry, _ := secondEntry.Options.Get(uint32(4))
	fourthEntry := rawFourthEntry.(*SSHDOption)
	if !(fourthEntry.Key.Value.Value == "AllowUsers" && fourthEntry.OptionValue.Value.Value == "root user" && fourthEntry.LocationRange.Start.Line == 4) {
		t.Errorf("Expected fourth entry to be 'AllowUsers root user', but got: %v", fourthEntry.Value)
	}

	rawFifthEntry, _ := p.Options.Get(uint32(6))
	fifthEntry := rawFifthEntry.(*SSHDMatchBlock)
	if !(fifthEntry.Options.Size() == 1) {
		t.Errorf("Expected 1 option in fifth match block, but got: %v", fifthEntry.Options)
	}

	rawSixthEntry, _ := fifthEntry.Options.Get(uint32(7))
	sixthEntry := rawSixthEntry.(*SSHDOption)
	if !(sixthEntry.Key.Value.Value == "MaxAuthTries" && sixthEntry.OptionValue.Value.Value == "3" && sixthEntry.LocationRange.Start.Line == 7) {
		t.Errorf("Expected sixth entry to be 'MaxAuthTries 3', but got: %v", sixthEntry.Value)
	}

	firstOption, firstMatchBlock := p.FindOption(uint32(3))
	if !(firstOption.Key.Key == "PasswordAuthentication" && firstOption.OptionValue.Value.Value == "yes") {
		t.Errorf("Expected first option to be 'PasswordAuthentication yes' and first match block to be 'Match Address 192.168.0.1', but got: %v, %v", firstOption, firstMatchBlock)
	}

	emptyOption, matchBlock := p.FindOption(uint32(5))
	if !(emptyOption == nil && matchBlock.MatchEntry.Value.Value == "Match User lena" && matchBlock.MatchValue.Entries[0].Values.Values[0].Value.Value == "lena") {
		t.Errorf("Expected empty option and match block to be 'Match User lena', but got: %v, %v", emptyOption, matchBlock)
	}

	matchOption, matchBlock := p.FindOption(uint32(2))
	if !(matchOption.Value.Value == "Match User lena" && matchBlock.MatchEntry.Value.Value == "Match User lena" && matchBlock.MatchValue.Entries[0].Values.Values[0].Value.Value == "lena" && matchBlock.MatchEntry.OptionValue.Start.Character == 6) {
		t.Errorf("Expected match option to be 'Match User lena', but got: %v, %v", matchOption, matchBlock)
	}

	if !(matchOption.Start.Line == 2 && matchOption.End.Line == 2 && matchOption.Start.Character == 0 && matchOption.End.Character == 15) {
		t.Errorf("Expected match option to be at 2:0-14, but got: %v", matchOption.LocationRange)
	}

	if !(matchBlock.Start.Line == 2 &&
		matchBlock.Start.Character == 0 &&
		matchBlock.End.Line == 4 &&
		matchBlock.End.Character == 21) {
		t.Errorf("Expected match block to be at 2:0-4:20, but got: %v", matchBlock.LocationRange)
	}
}

func TestSimpleExampleWithComments(
	t *testing.T,
) {
	input := utils.Dedent(`
# Test
PermitRootLogin no
Port 22
# Second test
AddressFamily any
Sample
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 4 &&
		len(utils.KeysOfMap(p.CommentLines)) == 2) {
		t.Errorf("Expected 3 options and 2 comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawFirstEntry, _ := p.Options.Get(uint32(1))
	firstEntry := rawFirstEntry.(*SSHDOption)
	firstEntryOpt, _ := p.FindOption(uint32(1))
	if !(firstEntry.Value.Value == "PermitRootLogin no" && firstEntry.LocationRange.Start.Line == 1 && firstEntryOpt == firstEntry) {
		t.Errorf("Expected first entry to be 'PermitRootLogin no', but got: %v", firstEntry.Value)
	}

	if len(p.CommentLines) != 2 {
		t.Errorf("Expected 2 comment lines, but got: %v", p.CommentLines)
	}

	if !utils.KeyExists(p.CommentLines, uint32(0)) {
		t.Errorf("Expected comment line 0 to not exist, but it does")
	}

	if !(utils.KeyExists(p.CommentLines, uint32(3))) {
		t.Errorf("Expected comment line 2 to exist, but it does not")
	}

	rawSecondEntry, _ := p.Options.Get(uint32(5))
	secondEntry := rawSecondEntry.(*SSHDOption)

	if !(secondEntry.Value.Value == "Sample") {
		t.Errorf("Expected second entry to be 'Sample', but got: %v", secondEntry.Value)
	}

}

func TestComplexExample(
	t *testing.T,
) {
	// From https://gist.github.com/kjellski/5940875
	input := utils.Dedent(`
# This is the sshd server system-wide configuration file.  See
# sshd_config(5) for more information.

# This sshd was compiled with PATH=/usr/bin:/bin:/usr/sbin:/sbin

# The strategy used for options in the default sshd_config shipped with
# OpenSSH is to specify options with their default value where
# possible, but leave them commented.  Uncommented options change a
# default value.

#Port 22
#AddressFamily any
#ListenAddress 0.0.0.0
#ListenAddress ::

# The default requires explicit activation of protocol 1
#Protocol 2

# HostKey for protocol version 1
#HostKey /etc/ssh/ssh_host_key
# HostKeys for protocol version 2
#HostKey /etc/ssh/ssh_host_rsa_key
#HostKey /etc/ssh/ssh_host_dsa_key
#HostKey /etc/ssh/ssh_host_ecdsa_key

# Lifetime and size of ephemeral version 1 server key
#KeyRegenerationInterval 1h
#ServerKeyBits 1024

# Logging
# obsoletes QuietMode and FascistLogging
#SyslogFacility AUTH
#LogLevel INFO

# Authentication:

#LoginGraceTime 2m
#BC# Root only allowed to login from LAN IP ranges listed at end
PermitRootLogin no
#PermitRootLogin yes
#StrictModes yes
#MaxAuthTries 6
#MaxSessions 10

#RSAAuthentication yes
#PubkeyAuthentication yes
#AuthorizedKeysFile  .ssh/authorized_keys

# For this to work you will also need host keys in /etc/ssh/ssh_known_hosts
#RhostsRSAAuthentication no
# similar for protocol version 2
#HostbasedAuthentication no
# Change to yes if you don't trust ~/.ssh/known_hosts for
# RhostsRSAAuthentication and HostbasedAuthentication
#IgnoreUserKnownHosts no
# Don't read the user's ~/.rhosts and ~/.shosts files
#IgnoreRhosts yes

# To disable tunneled clear text passwords, change to no here!
#BC# Disable password authentication by default (except for LAN IP ranges listed later)
PasswordAuthentication no
PermitEmptyPasswords no
#BC# Have to allow root here because AllowUsers not allowed in Match block.  It will not work though because of PermitRootLogin.
#BC# This is no longer true as of 6.1.  AllowUsers is now allowed in a Match block.
AllowUsers kmk root

# Change to no to disable s/key passwords
#BC# I occasionally use s/key one time passwords generated by a phone app
ChallengeResponseAuthentication yes

# Kerberos options
#KerberosAuthentication no
#KerberosOrLocalPasswd yes
#KerberosTicketCleanup yes
#KerberosGetAFSToken no

# GSSAPI options
#GSSAPIAuthentication no
#GSSAPICleanupCredentials yes

# Set this to 'yes' to enable PAM authentication, account processing, 
# and session processing. If this is enabled, PAM authentication will 
# be allowed through the ChallengeResponseAuthentication and
# PasswordAuthentication.  Depending on your PAM configuration,
# PAM authentication via ChallengeResponseAuthentication may bypass
# the setting of "PermitRootLogin without-password".
# If you just want the PAM account and session checks to run without
# PAM authentication, then enable this but set PasswordAuthentication
# and ChallengeResponseAuthentication to 'no'.
#BC# I would turn this off but I compiled ssh without PAM support so it errors if I set this.
#UsePAM no

#AllowAgentForwarding yes
#AllowTcpForwarding yes
#GatewayPorts no
X11Forwarding yes
#X11DisplayOffset 10
#X11UseLocalhost yes
#PrintMotd yes
#PrintLastLog yes
#TCPKeepAlive yes
#UseLogin no
#UsePrivilegeSeparation yes
#PermitUserEnvironment no
#Compression delayed
#ClientAliveInterval 0
#ClientAliveCountMax 3
#UseDNS yes
#PidFile /var/run/sshd.pid
#MaxStartups 10
#PermitTunnel no
#ChrootDirectory none

# no default banner path
#Banner none

# override default of no subsystems
#Subsystem	sftp	/usr/lib/misc/sftp-server
Subsystem	sftp	internal-sftp

# the following are HPN related configuration options
# tcp receive buffer polling. disable in non autotuning kernels
#TcpRcvBufPoll yes
 
# allow the use of the none cipher
#NoneEnabled no

# disable hpn performance boosts. 
#HPNDisabled no

# buffer size for hpn to non-hpn connections
#HPNBufferSize 2048


# Example of overriding settings on a per-user basis
Match User anoncvs
	X11Forwarding no
	AllowTcpForwarding no
	ForceCommand cvs server

#BC# My internal networks
#BC# Root can log in from here but only with a key and kmk can log in here with a password.
Match Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1
  PermitRootLogin without-password
  PasswordAuthentication yes
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 9 &&
		len(utils.KeysOfMap(p.CommentLines)) == 105) {
		t.Errorf("Expected 9 options and 105 comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawFirstEntry, _ := p.Options.Get(uint32(38))
	firstEntry := rawFirstEntry.(*SSHDOption)
	if !(firstEntry.Key.Value.Value == "PermitRootLogin" && firstEntry.OptionValue.Value.Value == "no") {
		t.Errorf("Expected first entry to be 'PermitRootLogin no', but got: %v", firstEntry.Value)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(60))
	secondEntry := rawSecondEntry.(*SSHDOption)
	if !(secondEntry.Key.Value.Value == "PasswordAuthentication" && secondEntry.OptionValue.Value.Value == "no") {
		t.Errorf("Expected second entry to be 'PasswordAuthentication no', but got: %v", secondEntry.Value)
	}

	rawThirdEntry, _ := p.Options.Get(uint32(118))
	thirdEntry := rawThirdEntry.(*SSHDOption)
	if !(thirdEntry.Key.Value.Value == "Subsystem" && thirdEntry.OptionValue.Value.Value == "sftp\tinternal-sftp") {
		t.Errorf("Expected third entry to be 'Subsystem sftp internal-sftp', but got: %v", thirdEntry.Value)
	}

	rawFourthEntry, _ := p.Options.Get(uint32(135))
	fourthEntry := rawFourthEntry.(*SSHDMatchBlock)
	if !(fourthEntry.MatchEntry.Value.Value == "Match User anoncvs") {
		t.Errorf("Expected fourth entry to be 'Match User anoncvs', but got: %v", fourthEntry.MatchEntry.Value)
	}

	if !(fourthEntry.MatchValue.Entries[0].Criteria.Type == "User" && fourthEntry.MatchValue.Entries[0].Values.Values[0].Value.Value == "anoncvs") {
		t.Errorf("Expected fourth entry to be 'Match User anoncvs', but got: %v", fourthEntry.MatchValue)
	}

	if !(fourthEntry.Options.Size() == 3) {
		t.Errorf("Expected 3 options in fourth match block, but got: %v", fourthEntry.Options)
	}

	rawFifthEntry, _ := fourthEntry.Options.Get(uint32(136))
	fifthEntry := rawFifthEntry.(*SSHDOption)
	if !(fifthEntry.Key.Value.Value == "X11Forwarding" && fifthEntry.OptionValue.Value.Value == "no") {
		t.Errorf("Expected fifth entry to be 'X11Forwarding no', but got: %v", fifthEntry.Value)
	}

	rawSixthEntry, _ := p.Options.Get(uint32(142))
	sixthEntry := rawSixthEntry.(*SSHDMatchBlock)
	if !(sixthEntry.MatchEntry.Value.Value == "Match Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1") {
		t.Errorf("Expected sixth entry to be 'Match Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1', but got: %v", sixthEntry.MatchEntry.Value)
	}

	if !(sixthEntry.MatchEntry.Key.Value.Value == "Match" && sixthEntry.MatchEntry.OptionValue.Value.Value == "Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1") {
		t.Errorf("Expected sixth entry to be 'Match Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1', but got: %v", sixthEntry.MatchEntry.Value)
	}

	if !(sixthEntry.MatchValue.Entries[0].Criteria.Type == "Address" && len(sixthEntry.MatchValue.Entries[0].Values.Values) == 3) {
		t.Errorf("Expected sixth entry to contain 3 values, but got: %v", sixthEntry.MatchValue)
	}

	if !(sixthEntry.Options.Size() == 2) {
		t.Errorf("Expected 2 options in sixth match block, but got: %v", sixthEntry.Options)
	}

	rawSeventhEntry, _ := sixthEntry.Options.Get(uint32(143))
	seventhEntry := rawSeventhEntry.(*SSHDOption)
	if !(seventhEntry.Key.Value.Value == "PermitRootLogin" && seventhEntry.OptionValue.Value.Value == "without-password") {
		t.Errorf("Expected seventh entry to be 'PermitRootLogin without-password', but got: %v", seventhEntry.Value)
	}
}
