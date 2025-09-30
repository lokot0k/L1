package main

import "fmt"

type Disk struct {
	WWN     string
	Size    int
	IsBlink bool
}

// JBODShelf работает с дисками посылая им команды напрямую
type JBODShelf struct {
	Disks []Disk
}

func (j *JBODShelf) GetDiskByWWN(wwn string) *Disk {
	for _, disk := range j.Disks {
		if disk.WWN == wwn {
			return &disk
		}
	}
	return nil
}

// Включаем подсветку на диске, напрямую обащясь с диском
func (j *JBODShelf) SwitchBlinkOfDisk(wwn string) string {
	disk := j.GetDiskByWWN(wwn)
	if disk == nil {
		return fmt.Sprintln("JBOD: Not found disk with WWN", wwn)
	}
	disk.IsBlink = !disk.IsBlink
	return fmt.Sprintf("JBOD: Switched blink of disk %s\n", wwn)
}

// JBOFShelf же общается только со слотами, а не с дисками.
// В слоты вставлены диски и только так можно посылать команды,
// поэтому нужен маппинг слот-диск
type JBOFShelf struct {
	Slots map[int]Disk // маппинг слот-диск
}

func (j *JBOFShelf) SwitchBlinkOfDisk(slotNumber int) string {
	if disk, ok := j.Slots[slotNumber]; ok {
		disk.IsBlink = !disk.IsBlink
		return fmt.Sprintf("JBOF: switched blink in slot %d\n", slotNumber)
	}
	return fmt.Sprintf("JBOF: not found disk in slot %d\n", slotNumber)
}

// DiskShelfAdapter - унифицированный интерфейс
type DiskShelfAdapter interface {
	SwitchBlinkOfDisk(wwn string) string
}

// адаптер JBOF к JBOD
type JBOFToJBODAdapter struct {
	JbofShelf *JBOFShelf
}

// унифицированный интерфейс для JBOF через адаптер
func (a *JBOFToJBODAdapter) SwitchBlinkOfDisk(wwn string) string {
	diskSlot := -1
	for slot := range a.JbofShelf.Slots {
		if a.JbofShelf.Slots[slot].WWN == wwn {
			diskSlot = slot
			break
		}
	}
	if diskSlot == -1 {
		return fmt.Sprintf("JBOF: Not found disk %s\n", wwn)
	}
	return a.JbofShelf.SwitchBlinkOfDisk(diskSlot)
}

func main() {
	jbod := &JBODShelf{
		Disks: []Disk{
			{WWN: "WD1001", Size: 1000},
			{WWN: "ST2002", Size: 2000},
		},
	}
	jbof := &JBOFShelf{
		Slots: map[int]Disk{
			1: {WWN: "SG3001", Size: 3000},
			2: {WWN: "HV4002", Size: 4000},
		},
	}
	jbofAdapter := &JBOFToJBODAdapter{JbofShelf: jbof}
	shelves := []DiskShelfAdapter{jbod, jbofAdapter}

	wwns := []string{"WD1001", "SG3001", "UNKNOWN"}
	for _, wwn := range wwns {
		fmt.Printf("Current WWN: %s\n", wwn)
		for _, shelf := range shelves {
			out := shelf.SwitchBlinkOfDisk(wwn)
			fmt.Printf("Adapter: %s", out)
		}
	}
}
