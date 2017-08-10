// Code generated by "stringer -type=RelocTypeGeneric,RelocTypeX86_64,RelocTypeARM,RelocTypeARM64 -output reloctype_string.go"; DO NOT EDIT.

package macho

import "fmt"

const _RelocTypeGeneric_name = "GENERIC_RELOC_VANILLAGENERIC_RELOC_PAIRGENERIC_RELOC_SECTDIFFGENERIC_RELOC_PB_LA_PTRGENERIC_RELOC_LOCAL_SECTDIFFGENERIC_RELOC_TLV"

var _RelocTypeGeneric_index = [...]uint8{0, 21, 39, 61, 84, 112, 129}

func (i RelocTypeGeneric) String() string {
	if i < 0 || i >= RelocTypeGeneric(len(_RelocTypeGeneric_index)-1) {
		return fmt.Sprintf("RelocTypeGeneric(%d)", i)
	}
	return _RelocTypeGeneric_name[_RelocTypeGeneric_index[i]:_RelocTypeGeneric_index[i+1]]
}

const _RelocTypeX86_64_name = "X86_64_RELOC_UNSIGNEDX86_64_RELOC_SIGNEDX86_64_RELOC_BRANCHX86_64_RELOC_GOT_LOADX86_64_RELOC_GOTX86_64_RELOC_SUBTRACTORX86_64_RELOC_SIGNED_1X86_64_RELOC_SIGNED_2X86_64_RELOC_SIGNED_4X86_64_RELOC_TLV"

var _RelocTypeX86_64_index = [...]uint8{0, 21, 40, 59, 80, 96, 119, 140, 161, 182, 198}

func (i RelocTypeX86_64) String() string {
	if i < 0 || i >= RelocTypeX86_64(len(_RelocTypeX86_64_index)-1) {
		return fmt.Sprintf("RelocTypeX86_64(%d)", i)
	}
	return _RelocTypeX86_64_name[_RelocTypeX86_64_index[i]:_RelocTypeX86_64_index[i+1]]
}

const _RelocTypeARM_name = "ARM_RELOC_VANILLAARM_RELOC_PAIRARM_RELOC_SECTDIFFARM_RELOC_LOCAL_SECTDIFFARM_RELOC_PB_LA_PTRARM_RELOC_BR24ARM_THUMB_RELOC_BR22ARM_THUMB_32BIT_BRANCHARM_RELOC_HALFARM_RELOC_HALF_SECTDIFF"

var _RelocTypeARM_index = [...]uint8{0, 17, 31, 49, 73, 92, 106, 126, 148, 162, 185}

func (i RelocTypeARM) String() string {
	if i < 0 || i >= RelocTypeARM(len(_RelocTypeARM_index)-1) {
		return fmt.Sprintf("RelocTypeARM(%d)", i)
	}
	return _RelocTypeARM_name[_RelocTypeARM_index[i]:_RelocTypeARM_index[i+1]]
}

const _RelocTypeARM64_name = "ARM64_RELOC_UNSIGNEDARM64_RELOC_SUBTRACTORARM64_RELOC_BRANCH26ARM64_RELOC_PAGE21ARM64_RELOC_PAGEOFF12ARM64_RELOC_GOT_LOAD_PAGE21ARM64_RELOC_GOT_LOAD_PAGEOFF12ARM64_RELOC_POINTER_TO_GOTARM64_RELOC_TLVP_LOAD_PAGE21ARM64_RELOC_TLVP_LOAD_PAGEOFF12ARM64_RELOC_ADDEND"

var _RelocTypeARM64_index = [...]uint16{0, 20, 42, 62, 80, 101, 128, 158, 184, 212, 243, 261}

func (i RelocTypeARM64) String() string {
	if i < 0 || i >= RelocTypeARM64(len(_RelocTypeARM64_index)-1) {
		return fmt.Sprintf("RelocTypeARM64(%d)", i)
	}
	return _RelocTypeARM64_name[_RelocTypeARM64_index[i]:_RelocTypeARM64_index[i+1]]
}