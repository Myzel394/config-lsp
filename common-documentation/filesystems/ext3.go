// This file contains documentation for ext3 filesystem options
// Note that this documentation ONLY contains options that are exclusive to ext3.
// Since ext3 is a superset of ext2, the documentation for ext2 options can be found in ext2.go
package commondocumentation

import docvalues "config-lsp/doc-values"

var Ext3DocumentationAssignable = docvalues.KeyEnumAssignmentValue{
	Values: map[docvalues.EnumString]docvalues.Value{
		docvalues.CreateEnumStringWithDoc(
			"journal_dev",
			"When the external journal device's major/minor numbers have changed, these options allow the user to specify the new journal location.  The journal device is identified either through its new major/minor numbers encoded in devnum, or via a path to the device.",
		): docvalues.StringValue{},
		docvalues.CreateEnumStringWithDoc(
			"journal_path",
			"When the external journal device's major/minor numbers have changed, these options allow the user to specify the new journal location.  The journal device is identified either through its new major/minor numbers encoded in devnum, or via a path to the device.",
		): docvalues.StringValue{},
		docvalues.CreateEnumStringWithDoc(
			"data",
			"Specifies the journaling mode for file data.  Metadata is always journaled.  To use modes other than ordered on the root file system, pass the mode to the kernel as boot parameter, e.g. rootflags=data=journal.",
		): docvalues.EnumValue{
			EnforceValues: true,
			Values: []docvalues.EnumString{
				docvalues.CreateEnumStringWithDoc(
					"journal",
					"All data is committed into the journal prior to being written into the main file system.",
				),
				docvalues.CreateEnumStringWithDoc(
					"ordered",
					"This is the default mode.  All data is forced directly out to the main file system prior to its metadata being committed to the journal.",
				),
				docvalues.CreateEnumStringWithDoc(
					"writeback",
					"Data ordering is not preserved – data may be metadata has been committed to the journal.  This written into the main file system after its is rumoured to be the highest-throughput option. It guarantees internal file system integrity, however it can allow old data to appear in files after a crash and journal recovery.",
				),
			},
		},
		docvalues.CreateEnumString(
			"data_err",
		): docvalues.EnumValue{
			EnforceValues: true,
			Values: []docvalues.EnumString{
				docvalues.CreateEnumStringWithDoc(
					"ignore",
					"Just print an error message if an error occurs in a file data buffer in ordered mode.",
				),
				docvalues.CreateEnumStringWithDoc(
					"abort",
					"Abort the journal if an error occurs in a file data buffer in ordered mode.",
				),
			},
		},
		docvalues.CreateEnumStringWithDoc(
			"barrier",
			"This disables / enables the use of write barriers in the jbd code.  barrier=0 disables, barrier=1 enables (default). This also requires an IO stack which can support barriers, and if jbd gets an error on a barrier write, it will disable barriers again with a warning. Write barriers enforce proper on-disk ordering of journal commits, making volatile disk write caches safe to use, at some performance penalty.  If your disks are battery-backed in one way or another, disabling barriers may safely improve performance.",
		): docvalues.EnumValue{
			EnforceValues: true,
			Values: []docvalues.EnumString{
				docvalues.CreateEnumStringWithDoc(
					"0",
					"Disables the use of write barriers in the jbd code.",
				),
				docvalues.CreateEnumStringWithDoc(
					"1",
					"Enables the use of write barriers in the jbd code.",
				),
			},
		},
		docvalues.CreateEnumStringWithDoc(
			"commit",
			"Start a journal commit every nrsec seconds.  The default value is 5 seconds.  Zero means default.",
		): docvalues.PositiveNumberValue(),
		docvalues.CreateEnumStringWithDoc(
			"jqfmt",
			"Apart from the old quota system (as in ext2, jqfmt=vfsold aka version 1 quota) ext3 also supports journaled quotas (version 2 quota). jqfmt=vfsv0 or jqfmt=vfsv1 enables journaled quotas. Journaled quotas have the advantage that even after a crash no quota check is required. When the quota file system feature is enabled, journaled quotas are used automatically, and this mount option is ignored.",
		): docvalues.EnumValue{
			EnforceValues: true,
			Values: []docvalues.EnumString{
				docvalues.CreateEnumString(
					"vfsold",
				),
				docvalues.CreateEnumStringWithDoc(
					"vfsv0",
					"Enable journaled quotas",
				),
				docvalues.CreateEnumStringWithDoc(
					"vfsv1",
					"Enable journaled quotas",
				),
			},
		},
		docvalues.CreateEnumStringWithDoc(
			"usrjquota",
			"For journaled quotas (jqfmt=vfsv0 or jqfmt=vfsv1), the mount options usrjquota=aquota.user and grpjquota=aquota.group are required to tell the quota system which quota database files to use. When the quota file system feature is enabled, journaled quotas are used automatically, and this mount option is ignored.",
		): docvalues.SingleEnumValue("aquota.user"),
		docvalues.CreateEnumStringWithDoc(
			"grpjquota",
			"For journaled quotas (jqfmt=vfsv0 or jqfmt=vfsv1), the mount options usrjquota=aquota.user and grpjquota=aquota.group are required to tell the quota system which quota database files to use. When the quota file system feature is enabled, journaled quotas are used automatically, and this mount option is ignored.",
		): docvalues.SingleEnumValue("aquota.group"),
	},
}

var Ext3DocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"norecovery",
		"Don't load the journal on mounting.  Note that if the file system was not unmounted cleanly, skipping the journal replay will lead to the file system containing inconsistencies that can lead to any number of problems.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noload",
		"Don't load the journal on mounting.  Note that if the file system was not unmounted cleanly, skipping the journal replay will lead to the file system containing inconsistencies that can lead to any number of problems.",
	),
}

