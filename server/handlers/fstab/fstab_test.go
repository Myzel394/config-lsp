package fstab

import (
	fstabdocumentation "config-lsp/handlers/fstab/documentation"
	handlers "config-lsp/handlers/fstab/handlers"
	"config-lsp/handlers/fstab/parser"
	"config-lsp/utils"
	"testing"
)

func TestValidBasicExample(t *testing.T) {
	input := utils.Dedent(`
LABEL=test /mnt/test ext4 defaults 0 0
`)
	p := parser.FstabParser{}
	p.Clear()

	errors := p.ParseFromContent(input)

	if len(errors) > 0 {
		t.Fatal("ParseFromContent failed with error", errors)
	}

	// Get hover for first field
	rawEntry, _ := p.Entries.Get(uint32(0))
	entry := rawEntry.(parser.FstabEntry)

	println("Getting hover info")
	{
		hover, err := handlers.GetHoverInfo(&entry, uint32(0))

		if err != nil {
			t.Fatal("getHoverInfo failed with error", err)
		}

		if hover.Contents != handlers.SpecHoverField.Contents {
			t.Fatal("getHoverInfo failed to return correct hover content. Got:", hover.Contents, "but expected:", handlers.SpecHoverField.Contents)
		}

		// Get hover for second field
		hover, err = handlers.GetHoverInfo(&entry, uint32(11))
		if err != nil {
			t.Fatal("getHoverInfo failed with error", err)
		}

		if hover.Contents != handlers.MountPointHoverField.Contents {
			t.Fatal("getHoverInfo failed to return correct hover content. Got:", hover.Contents, "but expected:", handlers.MountPointHoverField.Contents)
		}

		hover, err = handlers.GetHoverInfo(&entry, uint32(20))

		if err != nil {
			t.Fatal("getHoverInfo failed with error", err)
		}

		if hover.Contents != handlers.MountPointHoverField.Contents {
			t.Fatal("getHoverInfo failed to return correct hover content. Got:", hover.Contents, "but expected:", handlers.MountPointHoverField.Contents)
		}
	}

	println("Getting completions")
	{
		completions, err := handlers.GetCompletion(entry.Line, uint32(0))

		if err != nil {
			t.Fatal("getCompletion failed with error", err)
		}

		if len(completions) != 4 {
			t.Fatal("getCompletion failed to return correct number of completions. Got:", len(completions), "but expected:", 4)
		}

		if !(completions[0].Label == "LABEL" ||
			completions[1].Label == "LABEL" ||
			completions[2].Label == "LABEL" ||
			completions[3].Label == "LABEL") {
			t.Fatal("getCompletion failed to return correct label. Got:", completions[0].Label, "but expected:", "LABEL")
		}
	}

	{
		completions, err := handlers.GetCompletion(entry.Line, uint32(21))

		if err != nil {
			t.Fatal("getCompletion failed with error", err)
		}

		expectedLength := len(utils.KeysOfMap(fstabdocumentation.MountOptionsMapField))
		if len(completions) != expectedLength {
			t.Fatal("getCompletion failed to return correct number of completions. Got:", len(completions), "but expected:", expectedLength)
		}
	}

	println("Checking values")
	{
		diagnostics := p.AnalyzeValues()

		if len(diagnostics) > 0 {
			t.Fatal("AnalyzeValues failed with error", diagnostics)
		}
	}
}

func TestInvalidOptionsExample(t *testing.T) {
	input := utils.Dedent(`
LABEL=test /mnt/test btrfs subvol=backup,fat=32 0 0
`)
	p := parser.FstabParser{}
	p.Clear()

	errors := p.ParseFromContent(input)

	if len(errors) > 0 {
		t.Fatal("ParseFromContent returned error", errors)
	}

	// Get hover for first field
	println("Checking values")
	{
		diagnostics := p.AnalyzeValues()

		if len(diagnostics) == 0 {
			t.Fatal("AnalyzeValues should have returned error")
		}
	}
}

// func TestExample1(t *testing.T) {
// 	input := utils.Dedent(`
// /dev/disk/by-uuid/19ae4e13-1d6d-4833-965b-a20197aebf27 /mnt/RetroGames auto nosuid,nodev,nofail,x-gvfs-show 0 0
// /dev/disk/by-uuid/02629e02-a66d-4290-8a67-717ec9db6acc /mnt/SteamGames1 auto nosuid,nodev,nofail,x-gvfs-show 0 0
// /dev/disk/by-uuid/eb7c3d96-162f-41b7-b47f-add2c21b0220 /mnt/SteamGames2 auto nosuid,nodev,nofail,x-gvfs-show 0 0
// /dev/disk/by-uuid/ae977b84-cb99-492c-b50c-0e8b1242789f /mnt/Multimedia auto nosuid,nodev,nofail,x-gvfs-show 0 0
// /dev/disk/by-uuid/e5baf492-8653-415f-b0c4-bd88a81d61bd /mnt/Backup auto nosuid,nodev,nofail,x-gvfs-show 0 0
// /dev/disk/by-uuid/a3eb3cce-a34f-4141-b604-9aa293cb40c5 /mnt/Data auto nosuid,nodev,nofail,x-gvfs-show 0 0
// `)
// 	p := parser.FstabParser{}
// 	p.Clear()
//
// 	errors := p.ParseFromContent(input)
//
// 	if len(errors) > 0 {
// 		t.Fatalf("ParseFromContent failed with error %v", errors)
// 	}
// }

func TestArchExample1(t *testing.T) {
	input := utils.Dedent(`
UUID=0a3407de-014b-458b-b5c1-848e92a327a3 /     ext4   defaults  0      1
UUID=f9fe0b69-a280-415d-a03a-a32752370dee none  swap   defaults  0      0
UUID=b411dc99-f0a0-4c87-9e05-184977be8539 /home ext4   defaults  0      2
`)
	p := parser.FstabParser{}
	p.Clear()

	errors := p.ParseFromContent(input)

	if len(errors) > 0 {
		t.Fatalf("ParseFromContent failed with error %v", errors)
	}
}

func TestArchExample2(t *testing.T) {
	input := utils.Dedent(`
/dev/sda1         /boot        vfat          defaults         0      2
/dev/sda2         /            ext4          defaults         0      1
/dev/sda3         /home        ext4          defaults         0      2
/dev/sda4         none         swap          defaults         0      0
`)
	p := parser.FstabParser{}
	p.Clear()

	errors := p.ParseFromContent(input)

	if len(errors) > 0 {
		t.Fatalf("ParseFromContent failed with error %v", errors)
	}
}

func TestArchExample3(t *testing.T) {
	input := utils.Dedent(`
LABEL=ESP         /boot        vfat          defaults         0      2
LABEL=System      /            ext4          defaults         0      1
LABEL=Data        /home        ext4          defaults         0      2
LABEL=Swap        none         swap          defaults         0      0
`)
	p := parser.FstabParser{}
	p.Clear()

	errors := p.ParseFromContent(input)

	if len(errors) > 0 {
		t.Fatalf("ParseFromContent failed with error %v", errors)
	}
}

func TestArchExample4(t *testing.T) {
	input := utils.Dedent(`
UUID=CBB6-24F2                            /boot vfat   defaults  0      2
UUID=0a3407de-014b-458b-b5c1-848e92a327a3 /     ext4   defaults  0      1
UUID=b411dc99-f0a0-4c87-9e05-184977be8539 /home ext4   defaults  0      2
UUID=f9fe0b69-a280-415d-a03a-a32752370dee none  swap   defaults  0      0
`)
	p := parser.FstabParser{}
	p.Clear()

	errors := p.ParseFromContent(input)

	if len(errors) > 0 {
		t.Fatalf("ParseFromContent failed with error %v", errors)
	}
}

func TestArchExample5(t *testing.T) {
	input := utils.Dedent(`
PARTLABEL=EFI\040system\040partition /boot vfat   defaults  0      2
PARTLABEL=GNU/Linux                  /     ext4   defaults  0      1
PARTLABEL=Home                       /home ext4   defaults  0      2
PARTLABEL=Swap                       none  swap   defaults  0      0
`)
	p := parser.FstabParser{}
	p.Clear()

	errors := p.ParseFromContent(input)

	if len(errors) > 0 {
		t.Fatalf("ParseFromContent failed with error %v", errors)
	}
}

func TestArchExample6(t *testing.T) {
	input := utils.Dedent(`
PARTUUID=d0d0d110-0a71-4ed6-936a-304969ea36af /boot vfat   defaults  0      2
PARTUUID=98a81274-10f7-40db-872a-03df048df366 /     ext4   defaults  0      1
PARTUUID=7280201c-fc5d-40f2-a9b2-466611d3d49e /home ext4   defaults  0      2
PARTUUID=039b6c1c-7553-4455-9537-1befbc9fbc5b none  swap   defaults  0      0
`)
	p := parser.FstabParser{}
	p.Clear()

	errors := p.ParseFromContent(input)

	if len(errors) > 0 {
		t.Fatalf("ParseFromContent failed with error %v", errors)
	}
}

