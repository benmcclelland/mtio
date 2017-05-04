package mtio

import "bytes"

//MtGet is structure for MTIOCGET - mag tape get status command
type MtGet struct {
	// Type of magtape device
	Type int64

	// Residual count: (not sure)
	// number of bytes ignored, or
	// number of files not skipped, or
	// number of records not skipped
	ResID int64

	// The following registers are device dependent
	// Status register
	DsReg int64

	// Generic (device independent) status
	GStat int64

	// Error register
	ErReg int64

	// The next two fields are not always used
	// Current file number
	FileNo int32
	// Current block number
	BlkNo int32
}

//MtOp is structure for MTIOCTOP - magnetic tape operation command
type MtOp struct {
	// Operation ID
	op int16

	// Padding to match C structures
	Pad int16

	// Operation count
	count int32
}

//MtPos is structure for MTIOCPOS - mag tape get position command
type MtPos struct {
	//BlkNo Current block number
	BlkNo int64
}

const (
	// MTIO ioctl commands

	//MTIOCTOP do a mag tape op
	MTIOCTOP = 0x40086d01
	//MTIOCGET get tape status
	MTIOCGET = 0x80306d02
	//MTIOCPOS get tape position
	MTIOCPOS = 0x80086d03

	//Magnetic Tape operations [Not all operations supported by all drivers]

	//MTRESET +reset drive in case of problems
	MTRESET = 0
	//MTFSF Forward space over FileMark, position at first record of next file
	MTFSF = 1
	//MTBSF Backward space FileMark (position before FM)
	MTBSF = 2
	//MTFSR Forward space record
	MTFSR = 3
	//MTBSR Backward space record
	MTBSR = 4
	//MTWEOF Write an end-of-file record (mark)
	MTWEOF = 5
	//MTREW Rewind
	MTREW = 6
	//MTOFFL Rewind and put the drive offline (eject?)
	MTOFFL = 7
	//MTNOP No op, set status only (read with MTIOCGET)
	MTNOP = 8
	//MTRETEN Retension tape
	MTRETEN = 9
	//MTBSFM +backward space FileMark, position at FM
	MTBSFM = 10
	//MTFSFM +forward space FileMark, position at FM
	MTFSFM = 11
	//MTEOM Goto end of recorded media (for appending files)
	//MTEOM positions after the last FM, ready for appending another file
	MTEOM = 12
	//MTERASE Erase tape -- be careful!
	MTERASE = 13

	//MTRAS1 Run self test 1 (nondestructive)
	MTRAS1 = 14
	//MTRAS2 Run self test 2 (destructive)
	MTRAS2 = 15
	//MTRAS3 Reserved for self test 3
	MTRAS3 = 16

	//MTSETBLK Set block length (SCSI)
	MTSETBLK = 20
	//MTSETDENSITY Set tape density (SCSI)
	MTSETDENSITY = 21
	//MTSEEK Seek to block (Tandberg, etc.)
	MTSEEK = 22
	//MTTELL Tell block (Tandberg, etc.)
	MTTELL = 23
	//MTSETDRVBUFFER Set the drive buffering according to SCSI-2
	//Ordinary buffered operation with code 1
	MTSETDRVBUFFER = 24

	//MTFSS Space forward over setmarks
	MTFSS = 25
	//MTBSS Space backward over setmarks
	MTBSS = 26
	//MTWSM Write setmarks
	MTWSM = 27

	//MTLOCK Lock the drive door
	MTLOCK = 28
	//MTUNLOCK Unlock the drive door
	MTUNLOCK = 29
	//MTLOAD Execute the SCSI load command
	MTLOAD = 30
	//MTUNLOAD Execute the SCSI unload command
	MTUNLOAD = 31
	//MTCOMPRESSION Control compression with SCSI mode page 15
	MTCOMPRESSION = 32
	//MTSETPART Change the active tape partition
	MTSETPART = 33
	//MTMKPART Format the tape with one or two partitions
	MTMKPART = 34

	//Constants for mt_type. Not all of these are supported, and
	//these are not all of the ones that are supported.

	//MTISUNKNOWN unknown
	MTISUNKNOWN = 0x01
	//MTISQIC02 Generic QIC-02 tape streamer
	MTISQIC02 = 0x02
	//MTISWT5150 Wangtek 5150EQ, QIC-150, QIC-02
	MTISWT5150 = 0x03
	//MTISARCHIVE5945L2 Archive 5945L-2, QIC-24, QIC-02?
	MTISARCHIVE5945L2 = 0x04
	//MTISCMSJ500 CMS Jumbo 500 (QIC-02?)
	MTISCMSJ500 = 0x05
	//MTISTDC3610 Tandberg 6310, QIC-24
	MTISTDC3610 = 0x06
	//MTISARCHIVEVP60I Archive VP60i, QIC-02
	MTISARCHIVEVP60I = 0x07
	//MTISARCHIVE2150L Archive Viper 2150L
	MTISARCHIVE2150L = 0x08
	//MTISARCHIVE2060L Archive Viper 2060L
	MTISARCHIVE2060L = 0x09
	//MTISARCHIVESC499 Archive SC-499 QIC-36 controller
	MTISARCHIVESC499 = 0x0A
	//MTISQIC02ALLFEATURES Generic QIC-02 with all features
	MTISQIC02ALLFEATURES = 0x0F
	//MTISWT5099EEN24 Wangtek 5099-een24, 60MB, QIC-24
	MTISWT5099EEN24 = 0x11
	//MTISTEACMT2ST Teac MT-2ST 155mb drive, Teac DC-1 card (Wangtek type)
	MTISTEACMT2ST = 0x12
	//MTISEVEREXFT40A Everex FT40A (QIC-40)
	MTISEVEREXFT40A = 0x32
	//MTISDDS1 DDS device without partitions
	MTISDDS1 = 0x51
	//MTISDDS2 DDS device with partitions
	MTISDDS2 = 0x52
	//MTISONSTREAMSC OnStream SCSI tape drives (SC-x0) and SCSI emulated (DI, DP, USB)
	MTISONSTREAMSC = 0x61
	//MTISSCSI1 Generic ANSI SCSI-1 tape unit
	MTISSCSI1 = 0x71
	//MTISSCSI2 Generic ANSI SCSI-2 tape unit
	MTISSCSI2 = 0x72

	//MTISFTAPEFLAG QIC-40/80/3010/3020 ftape supported drives
	//20bit vendor ID + 0x800000 (see vendors.h in ftape distribution)
	MTISFTAPEFLAG = 0x800000

	//SCSI-tape specific definitions
	//Bitfield shifts in the status

	//MTSTBLKSIZESHIFT blocksize shift
	MTSTBLKSIZESHIFT = 0
	//MTSTBLKSIZEMASK blocksize mask
	MTSTBLKSIZEMASK = 0xffffff
	//MTSTDENSITYSHIFT density shift
	MTSTDENSITYSHIFT = 24
	//MTSTDENSITYMASK density mask
	MTSTDENSITYMASK = 0xff000000

	//MTSTSOFTERRSHIFT soft error shift
	MTSTSOFTERRSHIFT = 0
	//MTSTSOFTERRMASK soft error mask
	MTSTSOFTERRMASK = 0xffff

	//Bitfields for the MTSETDRVBUFFER ioctl

	//MTSTOPTIONS MTSETDRVBUFFER options
	MTSTOPTIONS = 0xf0000000
	//MTSTBOOLEANS MTSETDRVBUFFER booleans
	MTSTBOOLEANS = 0x10000000
	//MTSTSETBOOLEANS MTSETDRVBUFFER set booleans
	MTSTSETBOOLEANS = 0x30000000
	//MTSTCLEARBOOLEANS MTSETDRVBUFFER clear booleans
	MTSTCLEARBOOLEANS = 0x40000000
	//MTSTWRITETHRESHOLD MTSETDRVBUFFER write threshold
	MTSTWRITETHRESHOLD = 0x20000000
	//MTSTDEFBLKSIZE MTSETDRVBUFFER default blocksize
	MTSTDEFBLKSIZE = 0x50000000
	//MTSTDEFOPTIONS MTSETDRVBUFFER default options
	MTSTDEFOPTIONS = 0x60000000
	//MTSTTIMEOUTS timeouts
	MTSTTIMEOUTS = 0x70000000
	//MTSTSETTIMEOUT set timeout
	MTSTSETTIMEOUT = (MTSTTIMEOUTS | 0x000000)
	//MTSTSETLONGTIMEOUT set long timeout
	MTSTSETLONGTIMEOUT = (MTSTTIMEOUTS | 0x100000)
	//MTSTSETCLN set cln
	MTSTSETCLN = 0x80000000

	//MTSTBUFFERWRITES buffered writes
	MTSTBUFFERWRITES = 0x1
	//MTSTASYNCWRITES async writes
	MTSTASYNCWRITES = 0x2
	//MTSTREADAHEAD read ahead
	MTSTREADAHEAD = 0x4
	//MTSTDEBUGGING debugging
	MTSTDEBUGGING = 0x8
	//MTSTTWOFM write two filemarks
	MTSTTWOFM = 0x10
	//MTSTFASTMTEOM send MTEOM directly to drive
	MTSTFASTMTEOM = 0x20
	//MTSTAUTOLOCK auto lock
	MTSTAUTOLOCK = 0x40
	//MTSTDEFWRITES apply settings to drive defaults
	MTSTDEFWRITES = 0x80
	//MTSTCANBSR correct readahaead backspace position
	MTSTCANBSR = 0x100
	//MTSTNOBLKLIMS dont use READ BLOCK LIMITS
	MTSTNOBLKLIMS = 0x200
	//MTSTCANPARTITIONS enable partitions
	MTSTCANPARTITIONS = 0x400
	//MTSTSCSI2LOGICAL use logical block addresses
	MTSTSCSI2LOGICAL = 0x800
	//MTSTSYSV sysv
	MTSTSYSV = 0x1000
	//MTSTNOWAIT no wait
	MTSTNOWAIT = 0x2000
	//MTSTSILI SILI
	MTSTSILI = 0x4000

	//The mode parameters to be controlled
	//Parameter chosen with bits 20-28

	//MTSTCLEARDEFAULT clear default
	MTSTCLEARDEFAULT = 0xfffff
	//MTSTDEFDENSITY default density
	MTSTDEFDENSITY = (MTSTDEFOPTIONS | 0x100000)
	//MTSTDEFCOMPRESSION default compression
	MTSTDEFCOMPRESSION = (MTSTDEFOPTIONS | 0x200000)
	//MTSTDEFDRVBUFFER default buffering
	MTSTDEFDRVBUFFER = (MTSTDEFOPTIONS | 0x300000)

	//MTSTHPLOADEROFFSET arguments for the special HP changer load command
	MTSTHPLOADEROFFSET = 10000
)

//MtTypeToString converts type to string
func MtTypeToString(t int64) string {
	switch t {
	case MTISUNKNOWN:
		return "Unknown type of tape device"
	case MTISQIC02:
		return "Generic QIC-02 tape streamer"
	case MTISWT5150:
		return "Wangtek 5150, QIC-150"
	case MTISARCHIVE5945L2:
		return "Archive 5945L-2"
	case MTISCMSJ500:
		return "CMS Jumbo 500"
	case MTISTDC3610:
		return "Tandberg TDC 3610, QIC-24"
	case MTISARCHIVEVP60I:
		return "Archive VP60i, QIC-02"
	case MTISARCHIVE2150L:
		return "Archive Viper 2150L"
	case MTISARCHIVE2060L:
		return "Archive Viper 2060L"
	case MTISARCHIVESC499:
		return "Archive SC-499 QIC-36 controller"
	case MTISQIC02ALLFEATURES:
		return "Generic QIC-02 tape, all features"
	case MTISWT5099EEN24:
		return "Wangtek 5099-een24, 60MB"
	case MTISTEACMT2ST:
		return "Teac MT-2ST 155mb data cassette drive"
	case MTISEVEREXFT40A:
		return "Everex FT40A, QIC-40"
	case MTISDDS1:
		return "DDS device without partitions"
	case MTISDDS2:
		return "DDS device with partitions"
	case MTISSCSI1:
		return "Generic SCSI-1 tape"
	case MTISSCSI2:
		return "Generic SCSI-2 tape"
	}
	return "invalid type"
}

//IsEOF GStat reports EOF
func IsEOF(s int64) bool {
	return s&0x80000000 != 0
}

//IsBOT GStat reports BOT
func IsBOT(s int64) bool {
	return s&0x40000000 != 0
}

//IsEOT GStat reports EOT
func IsEOT(s int64) bool {
	return s&0x20000000 != 0
}

//IsSM GStat reports DDS setmark
func IsSM(s int64) bool {
	return s&0x10000000 != 0
}

//IsEOD GStat reports DDS EOD
func IsEOD(s int64) bool {
	return s&0x08000000 != 0
}

//IsWrProt GStat reports write protected
func IsWrProt(s int64) bool {
	return s&0x04000000 != 0
}

//IsOnline GStat reports online
func IsOnline(s int64) bool {
	return s&0x01000000 != 0
}

//IsD6250 GStat reports D_6250
func IsD6250(s int64) bool {
	return s&0x00800000 != 0
}

//IsD1600 GStat reports D_1600
func IsD1600(s int64) bool {
	return s&0x00400000 != 0
}

//IsD800 GStat reports D_800
func IsD800(s int64) bool {
	return s&0x00200000 != 0
}

//IsDrOpen GStat reports Door open
func IsDrOpen(s int64) bool {
	return s&0x00040000 != 0
}

//IsImRepEn GStat reports Immediate report mode
func IsImRepEn(s int64) bool {
	return s&0x00010000 != 0
}

//IsCln GStat reports Cleaning requested
func IsCln(s int64) bool {
	return s&0x00008000 != 0
}

func appendBuf(b *bytes.Buffer, s string) error {
	if b.Len() == 0 {
		_, err := b.WriteString(s)
		return err
	}
	_, err := b.WriteString(" ")
	if err != nil {
		return err
	}
	_, err = b.WriteString(s)
	return err
}

//MtStatusToString convert status to string with enabled options listed
func MtStatusToString(s int64) string {
	var b bytes.Buffer
	if IsEOF(s) {
		err := appendBuf(&b, "EOF")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsBOT(s) {
		err := appendBuf(&b, "BOT")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsEOT(s) {
		err := appendBuf(&b, "EOT")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsSM(s) {
		err := appendBuf(&b, "SM")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsEOD(s) {
		err := appendBuf(&b, "EOD")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsWrProt(s) {
		err := appendBuf(&b, "WR_PROT")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsOnline(s) {
		err := appendBuf(&b, "ONLINE")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsD6250(s) {
		err := appendBuf(&b, "D_6250")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsD1600(s) {
		err := appendBuf(&b, "D_1600")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsD800(s) {
		err := appendBuf(&b, "D_800")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsDrOpen(s) {
		err := appendBuf(&b, "DR_OPEN")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsImRepEn(s) {
		err := appendBuf(&b, "IM_REP_EN")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}
	if IsCln(s) {
		err := appendBuf(&b, "CLN")
		if err != nil {
			return "Internal Error " + err.Error()
		}
	}

	return b.String()
}
