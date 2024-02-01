package snapshot

import (
	"fmt"
	"strconv"
	"strings"
)

// Image is the name of a remote snapshot image.
type Image struct {
	Prefix   string `json:"prefix"`
	Template int    `json:"template"`
	VM       int    `json:"vm"`
	Disk     int    `json:"disk"`
}

// ImageSpec describes the full location of an RBD image
type ImageSpec struct {
	Pool      string
	Namespace string
	Image     Image
}

// String implements the Stringer interface for i.
func (i Image) String() string {
	imageName := fmt.Sprintf("%v-%v", i.Prefix, i.Template)

	// a persistent image will have VM and Disk set to -1, so don't include that in the string
	if i.VM != -1 {
		imageName += fmt.Sprintf("-%v", i.VM)
	}
	if i.Disk != -1 {
		imageName += fmt.Sprintf("-%v", i.Disk)
	}

	return imageName
}

// ParseImage parses s to return an Image. The format is as follows:
//
//	PREFIX-TEMPLATE-VM-DISK
//
// where
//
//	PREFIX is the name of the remote
//	TEMPLATE is the template ID
//	VM is the VM ID
//	DISK is the disk number
//
// to support "persistent" disks, we also need to accept the format:
//
//	PREFIX-TEMPLATE
func ParseImage(s string) (Image, error) {
	fields := strings.Split(s, "-")

	var err error
	var vm, disk int

	vm = -1
	disk = -1

	switch len(fields) {
	case 4:
		vm, err = strconv.Atoi(fields[2])
		if err != nil {
			return Image{}, fmt.Errorf("cannot determine vm from %q in image %q", fields[2], s)
		}

		disk, err = strconv.Atoi(fields[3])
		if err != nil {
			return Image{}, fmt.Errorf("cannot determine disk from %q in image %q", fields[3], s)
		}
	case 2:
		break
	default:
		return Image{}, fmt.Errorf("cannot parse image %q", s)
	}

	template, err := strconv.Atoi(fields[1])
	if err != nil {
		return Image{}, fmt.Errorf("cannot determine template from %q in image %q", fields[1], s)
	}

	image := Image{
		Prefix:   fields[0],
		Template: template,
		VM:       vm,
		Disk:     disk,
	}

	return image, nil
}

// ParseImageSpec parses s to return an ImageSpec.
//
// valid formats for 's' are:
//
//	<image-name>
//	<pool-name>/<image-name>
//	<pool-name>/<namespace>/<image-name>
func ParseImageSpec(s string) (ImageSpec, error) {
	var iSpec ImageSpec
	var err error
	var image int

	pool, namespace := -1, -1

	img := strings.Split(s, "/")

	switch len(img) {
	case 1:
		image = 0
	case 2:
		pool = 0
		image = 1
	case 3:
		pool = 0
		namespace = 1
		image = 2
	default:
		return ImageSpec{}, fmt.Errorf("malformed image spec %v", s)
	}

	if pool != -1 {
		iSpec.Pool = img[pool]
	}

	if namespace != -1 {
		iSpec.Namespace = img[namespace]
	}

	iSpec.Image, err = ParseImage(img[image])
	if err != nil {
		return ImageSpec{}, err
	}

	return iSpec, nil
}
