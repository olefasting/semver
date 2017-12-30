package semver

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	firstSegRegex = regexp.MustCompile("[0-9]{1,4}$")
)

// Version holds a three segment software version
type Version struct {
	pfx   string
	major uint16
	minor uint16
	patch uint16
}

// GetPrefix will get the prefix currently used
func (v *Version) GetPrefix() string {
	return v.pfx
}

// GetMajor will get the current major version
func (v *Version) GetMajor() uint16 {
	return v.major
}

// GetMinor will get the current minor version
func (v *Version) GetMinor() uint16 {
	return v.minor
}

// GetPatch will get the current patch version
func (v *Version) GetPatch() uint16 {
	return v.patch
}

// SetPrefix will set the prefix used in string versions of this to *val*
func (v *Version) SetPrefix(val string) *Version {
	v.pfx = val
	return v
}

// SetMajor will set the major version to *val*
func (v *Version) SetMajor(val uint16) *Version {
	v.major = val
	return v
}

// SetMinor will set the minor version to *val*
func (v *Version) SetMinor(val uint16) *Version {
	v.minor = val
	return v
}

// SetPatch will set the patch version to *val*
func (v *Version) SetPatch(val uint16) *Version {
	v.patch = val
	return v
}

// Bytes returns a byte array containing this version string
func (v *Version) Bytes() []byte {
	b := []byte(v.String())
	return b
}

// String implements `fmt.Stringer`
func (v *Version) String() string {
	s := fmt.Sprintf("%s%v.%v.%v", v.pfx, v.major, v.minor, v.patch)
	return s
}

// MarshalText implements `encoding.TextMarshaler`
func (v *Version) MarshalText() ([]byte, error) {
	b := v.Bytes()
	return b, nil
}

// UnmarshalText implements `encoding.TextUnmarshaler`
func (v *Version) UnmarshalText(b []byte) error {
	s := string(b)
	segs := strings.Split(s, ".")
	if len(segs) < 3 {
		return fmt.Errorf("error while marshalling: text does not contain three dot separated segments")
	}
	num := firstSegRegex.FindString(segs[0])
	if pe := len(segs[0]) - len(num); pe > 0 {
		v.pfx = segs[0][0:pe]
	} else {
		v.pfx = ""
	}
	u, err := strconv.ParseUint(num, 10, 16)
	if err != nil {
		return err
	}
	v.major = uint16(u)
	u, err = strconv.ParseUint(segs[1], 10, 16)
	if err != nil {
		return err
	}
	v.minor = uint16(u)
	u, err = strconv.ParseUint(segs[2], 10, 16)
	if err != nil {
		return err
	}
	v.patch = uint16(u)
	return nil
}

// MarshalJSON implements `json.Marshaler`
func (v *Version) MarshalJSON() ([]byte, error) {
	b := []byte(strconv.Quote(v.String()))
	return b, nil
}

// UnmarshalJSON implements `json.Unmarshaler`
func (v *Version) UnmarshalJSON(b []byte) error {
	l := len(b)
	if l < 3 {
		return fmt.Errorf("error while unmarshaling: json string is too short")
	}
	return v.UnmarshalText(b[1 : l-1])
}
