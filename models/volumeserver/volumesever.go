package volumeserver

type Volumeservers struct {
	VolumeserverID *string
	IP             *string
	Path           *string
	Name           *string
	Memory         *int
	Created        *string
	Active         *bool
	Groups         *string
	DiskType       *string
}
