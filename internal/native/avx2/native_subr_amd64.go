//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	"github.com/abdullayev13/qsonic/loader"
)

const (
	_entry__f32toa                 = 34640
	_entry__f64toa                 = 368
	_entry__format_significand     = 39440
	_entry__format_integer         = 3616
	_entry__fsm_exec               = 22128
	_entry__advance_ns             = 17856
	_entry__advance_string         = 18592
	_entry__advance_string_default = 41024
	_entry__do_skip_number         = 24624
	_entry__get_by_path            = 29616
	_entry__skip_one_fast          = 26736
	_entry__unescape               = 41824
	_entry__unhex16_is             = 11376
	_entry__html_escape            = 11712
	_entry__i64toa                 = 4048
	_entry__u64toa                 = 4320
	_entry__lspace                 = 64
	_entry__quote                  = 5872
	_entry__skip_array             = 22080
	_entry__skip_number            = 26128
	_entry__skip_object            = 24048
	_entry__skip_one               = 26304
	_entry__unquote                = 8816
	_entry__validate_one           = 26368
	_entry__validate_utf8          = 30960
	_entry__validate_utf8_fast     = 31920
	_entry__value                  = 16320
	_entry__vnumber                = 19680
	_entry__atof_eisel_lemire64    = 13760
	_entry__atof_native            = 15712
	_entry__decimal_to_f64         = 14240
	_entry__left_shift             = 39920
	_entry__right_shift            = 40464
	_entry__vsigned                = 21328
	_entry__vstring                = 18352
	_entry__vunsigned              = 21696
)

const (
	_stack__f32toa                 = 64
	_stack__f64toa                 = 80
	_stack__format_significand     = 24
	_stack__format_integer         = 16
	_stack__fsm_exec               = 136
	_stack__advance_ns             = 8
	_stack__advance_string         = 48
	_stack__advance_string_default = 48
	_stack__do_skip_number         = 40
	_stack__get_by_path            = 304
	_stack__skip_one_fast          = 184
	_stack__unescape               = 64
	_stack__unhex16_is             = 8
	_stack__html_escape            = 72
	_stack__i64toa                 = 16
	_stack__u64toa                 = 8
	_stack__lspace                 = 8
	_stack__quote                  = 72
	_stack__skip_array             = 144
	_stack__skip_number            = 96
	_stack__skip_object            = 144
	_stack__skip_one               = 144
	_stack__unquote                = 112
	_stack__validate_one           = 144
	_stack__validate_utf8          = 48
	_stack__validate_utf8_fast     = 176
	_stack__value                  = 352
	_stack__vnumber                = 264
	_stack__atof_eisel_lemire64    = 40
	_stack__atof_native            = 144
	_stack__decimal_to_f64         = 88
	_stack__left_shift             = 32
	_stack__right_shift            = 16
	_stack__vsigned                = 16
	_stack__vstring                = 104
	_stack__vunsigned              = 24
)

const (
	_size__f32toa                 = 3792
	_size__f64toa                 = 3248
	_size__format_significand     = 480
	_size__format_integer         = 432
	_size__fsm_exec               = 1380
	_size__advance_ns             = 496
	_size__advance_string         = 1040
	_size__advance_string_default = 800
	_size__do_skip_number         = 1300
	_size__get_by_path            = 1344
	_size__skip_one_fast          = 2360
	_size__unescape               = 704
	_size__unhex16_is             = 144
	_size__html_escape            = 2048
	_size__i64toa                 = 272
	_size__u64toa                 = 1408
	_size__lspace                 = 256
	_size__quote                  = 2896
	_size__skip_array             = 48
	_size__skip_number            = 160
	_size__skip_object            = 48
	_size__skip_one               = 48
	_size__unquote                = 2560
	_size__validate_one           = 64
	_size__validate_utf8          = 688
	_size__validate_utf8_fast     = 2672
	_size__value                  = 992
	_size__vnumber                = 1648
	_size__atof_eisel_lemire64    = 416
	_size__atof_native            = 608
	_size__decimal_to_f64         = 1472
	_size__left_shift             = 544
	_size__right_shift            = 496
	_size__vsigned                = 368
	_size__vstring                = 144
	_size__vunsigned              = 368
)

var (
	_pcsp__f32toa = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{3734, 64},
		{3738, 48},
		{3739, 40},
		{3741, 32},
		{3743, 24},
		{3745, 16},
		{3747, 8},
		{3751, 0},
		{3781, 64},
	}
	_pcsp__f64toa = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{3124, 56},
		{3128, 48},
		{3129, 40},
		{3131, 32},
		{3133, 24},
		{3135, 16},
		{3137, 8},
		{3141, 0},
		{3234, 56},
	}
	_pcsp__format_significand = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{468, 24},
		{469, 16},
		{471, 8},
		{473, 0},
	}
	_pcsp__format_integer = [][2]uint32{
		{1, 0},
		{4, 8},
		{412, 16},
		{413, 8},
		{414, 0},
		{423, 16},
		{424, 8},
		{426, 0},
	}
	_pcsp__fsm_exec = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{1017, 88},
		{1021, 48},
		{1022, 40},
		{1024, 32},
		{1026, 24},
		{1028, 16},
		{1030, 8},
		{1031, 0},
		{1380, 88},
	}
	_pcsp__advance_ns = [][2]uint32{
		{1, 0},
		{442, 8},
		{446, 0},
		{467, 8},
		{471, 0},
		{489, 8},
	}
	_pcsp__advance_string = [][2]uint32{
		{14, 0},
		{18, 8},
		{20, 16},
		{22, 24},
		{24, 32},
		{26, 40},
		{396, 48},
		{397, 40},
		{399, 32},
		{401, 24},
		{403, 16},
		{405, 8},
		{409, 0},
		{1031, 48},
	}
	_pcsp__advance_string_default = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{314, 48},
		{315, 40},
		{317, 32},
		{319, 24},
		{321, 16},
		{323, 8},
		{327, 0},
		{786, 48},
	}
	_pcsp__do_skip_number = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{1253, 40},
		{1254, 32},
		{1256, 24},
		{1258, 16},
		{1260, 8},
		{1264, 0},
		{1300, 40},
	}
	_pcsp__get_by_path = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{1293, 120},
		{1297, 48},
		{1298, 40},
		{1300, 32},
		{1302, 24},
		{1304, 16},
		{1306, 8},
		{1307, 0},
		{1344, 120},
	}
	_pcsp__skip_one_fast = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{396, 176},
		{397, 168},
		{399, 160},
		{401, 152},
		{403, 144},
		{405, 136},
		{409, 128},
		{2360, 176},
	}
	_pcsp__unescape = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{246, 56},
		{250, 48},
		{251, 40},
		{253, 32},
		{255, 24},
		{257, 16},
		{259, 8},
		{260, 0},
		{695, 56},
	}
	_pcsp__unhex16_is = [][2]uint32{
		{1, 0},
		{35, 8},
		{36, 0},
		{62, 8},
		{63, 0},
		{97, 8},
		{98, 0},
		{121, 8},
		{123, 0},
	}
	_pcsp__html_escape = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{2017, 72},
		{2021, 48},
		{2022, 40},
		{2024, 32},
		{2026, 24},
		{2028, 16},
		{2030, 8},
		{2035, 0},
	}
	_pcsp__i64toa = [][2]uint32{
		{1, 0},
		{171, 8},
		{172, 0},
		{207, 8},
		{208, 0},
		{222, 8},
		{223, 0},
		{247, 8},
		{248, 0},
		{253, 8},
		{259, 0},
	}
	_pcsp__u64toa = [][2]uint32{
		{13, 0},
		{162, 8},
		{163, 0},
		{175, 8},
		{240, 0},
		{498, 8},
		{499, 0},
		{519, 8},
		{592, 0},
		{852, 8},
		{928, 0},
		{1376, 8},
		{1378, 0},
	}
	_pcsp__lspace = [][2]uint32{
		{1, 0},
		{186, 8},
		{190, 0},
		{199, 8},
		{203, 0},
		{210, 8},
		{214, 0},
		{232, 8},
	}
	_pcsp__quote = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{2828, 72},
		{2832, 48},
		{2833, 40},
		{2835, 32},
		{2837, 24},
		{2839, 16},
		{2841, 8},
		{2845, 0},
		{2876, 72},
	}
	_pcsp__skip_array = [][2]uint32{
		{1, 0},
		{28, 8},
		{34, 0},
	}
	_pcsp__skip_number = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{107, 56},
		{111, 48},
		{112, 40},
		{114, 32},
		{116, 24},
		{118, 16},
		{120, 8},
		{121, 0},
		{145, 56},
	}
	_pcsp__skip_object = [][2]uint32{
		{1, 0},
		{28, 8},
		{34, 0},
	}
	_pcsp__skip_one = [][2]uint32{
		{1, 0},
		{28, 8},
		{34, 0},
	}
	_pcsp__unquote = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{1836, 104},
		{1840, 48},
		{1841, 40},
		{1843, 32},
		{1845, 24},
		{1847, 16},
		{1849, 8},
		{1853, 0},
		{2554, 104},
	}
	_pcsp__validate_one = [][2]uint32{
		{1, 0},
		{33, 8},
		{39, 0},
	}
	_pcsp__validate_utf8 = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{11, 40},
		{639, 48},
		{643, 40},
		{644, 32},
		{646, 24},
		{648, 16},
		{650, 8},
		{651, 0},
		{682, 48},
	}
	_pcsp__validate_utf8_fast = [][2]uint32{
		{1, 0},
		{4, 8},
		{5, 16},
		{1706, 176},
		{1707, 168},
		{1711, 160},
		{2003, 176},
		{2004, 168},
		{2008, 160},
		{2656, 176},
	}
	_pcsp__value = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{974, 88},
		{978, 48},
		{979, 40},
		{981, 32},
		{983, 24},
		{985, 16},
		{987, 8},
		{992, 0},
	}
	_pcsp__vnumber = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{150, 120},
		{154, 48},
		{155, 40},
		{157, 32},
		{159, 24},
		{161, 16},
		{163, 8},
		{164, 0},
		{1638, 120},
	}
	_pcsp__atof_eisel_lemire64 = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{315, 40},
		{316, 32},
		{318, 24},
		{320, 16},
		{322, 8},
		{323, 0},
		{387, 40},
	}
	_pcsp__atof_native = [][2]uint32{
		{1, 0},
		{4, 8},
		{596, 56},
		{600, 8},
		{602, 0},
	}
	_pcsp__decimal_to_f64 = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{1431, 56},
		{1435, 48},
		{1436, 40},
		{1438, 32},
		{1440, 24},
		{1442, 16},
		{1444, 8},
		{1448, 0},
		{1460, 56},
	}
	_pcsp__left_shift = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{418, 32},
		{419, 24},
		{421, 16},
		{423, 8},
		{424, 0},
		{539, 32},
	}
	_pcsp__right_shift = [][2]uint32{
		{1, 0},
		{4, 8},
		{452, 16},
		{453, 8},
		{454, 0},
		{462, 16},
		{463, 8},
		{464, 0},
		{472, 16},
		{473, 8},
		{475, 0},
	}
	_pcsp__vsigned = [][2]uint32{
		{1, 0},
		{4, 8},
		{113, 16},
		{114, 8},
		{115, 0},
		{126, 16},
		{127, 8},
		{128, 0},
		{278, 16},
		{279, 8},
		{280, 0},
		{284, 16},
		{285, 8},
		{286, 0},
		{340, 16},
		{341, 8},
		{342, 0},
		{353, 16},
		{354, 8},
		{356, 0},
	}
	_pcsp__vstring = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{11, 40},
		{105, 56},
		{109, 40},
		{110, 32},
		{112, 24},
		{114, 16},
		{116, 8},
		{118, 0},
	}
	_pcsp__vunsigned = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{74, 24},
		{75, 16},
		{77, 8},
		{78, 0},
		{89, 24},
		{90, 16},
		{92, 8},
		{93, 0},
		{116, 24},
		{117, 16},
		{119, 8},
		{120, 0},
		{281, 24},
		{282, 16},
		{284, 8},
		{285, 0},
		{336, 24},
		{337, 16},
		{339, 8},
		{340, 0},
		{348, 24},
		{349, 16},
		{351, 8},
		{353, 0},
	}
)

var Funcs = []loader.CFunc{
	{"__native_entry__", 0, 67, 0, nil},
	{"_f32toa", _entry__f32toa, _size__f32toa, _stack__f32toa, _pcsp__f32toa},
	{"_f64toa", _entry__f64toa, _size__f64toa, _stack__f64toa, _pcsp__f64toa},
	{"_format_significand", _entry__format_significand, _size__format_significand, _stack__format_significand, _pcsp__format_significand},
	{"_format_integer", _entry__format_integer, _size__format_integer, _stack__format_integer, _pcsp__format_integer},
	{"_fsm_exec", _entry__fsm_exec, _size__fsm_exec, _stack__fsm_exec, _pcsp__fsm_exec},
	{"_advance_ns", _entry__advance_ns, _size__advance_ns, _stack__advance_ns, _pcsp__advance_ns},
	{"_advance_string", _entry__advance_string, _size__advance_string, _stack__advance_string, _pcsp__advance_string},
	{"_advance_string_default", _entry__advance_string_default, _size__advance_string_default, _stack__advance_string_default, _pcsp__advance_string_default},
	{"_do_skip_number", _entry__do_skip_number, _size__do_skip_number, _stack__do_skip_number, _pcsp__do_skip_number},
	{"_get_by_path", _entry__get_by_path, _size__get_by_path, _stack__get_by_path, _pcsp__get_by_path},
	{"_skip_one_fast", _entry__skip_one_fast, _size__skip_one_fast, _stack__skip_one_fast, _pcsp__skip_one_fast},
	{"_unescape", _entry__unescape, _size__unescape, _stack__unescape, _pcsp__unescape},
	{"_unhex16_is", _entry__unhex16_is, _size__unhex16_is, _stack__unhex16_is, _pcsp__unhex16_is},
	{"_html_escape", _entry__html_escape, _size__html_escape, _stack__html_escape, _pcsp__html_escape},
	{"_i64toa", _entry__i64toa, _size__i64toa, _stack__i64toa, _pcsp__i64toa},
	{"_u64toa", _entry__u64toa, _size__u64toa, _stack__u64toa, _pcsp__u64toa},
	{"_lspace", _entry__lspace, _size__lspace, _stack__lspace, _pcsp__lspace},
	{"_quote", _entry__quote, _size__quote, _stack__quote, _pcsp__quote},
	{"_skip_array", _entry__skip_array, _size__skip_array, _stack__skip_array, _pcsp__skip_array},
	{"_skip_number", _entry__skip_number, _size__skip_number, _stack__skip_number, _pcsp__skip_number},
	{"_skip_object", _entry__skip_object, _size__skip_object, _stack__skip_object, _pcsp__skip_object},
	{"_skip_one", _entry__skip_one, _size__skip_one, _stack__skip_one, _pcsp__skip_one},
	{"_unquote", _entry__unquote, _size__unquote, _stack__unquote, _pcsp__unquote},
	{"_validate_one", _entry__validate_one, _size__validate_one, _stack__validate_one, _pcsp__validate_one},
	{"_validate_utf8", _entry__validate_utf8, _size__validate_utf8, _stack__validate_utf8, _pcsp__validate_utf8},
	{"_validate_utf8_fast", _entry__validate_utf8_fast, _size__validate_utf8_fast, _stack__validate_utf8_fast, _pcsp__validate_utf8_fast},
	{"_value", _entry__value, _size__value, _stack__value, _pcsp__value},
	{"_vnumber", _entry__vnumber, _size__vnumber, _stack__vnumber, _pcsp__vnumber},
	{"_atof_eisel_lemire64", _entry__atof_eisel_lemire64, _size__atof_eisel_lemire64, _stack__atof_eisel_lemire64, _pcsp__atof_eisel_lemire64},
	{"_atof_native", _entry__atof_native, _size__atof_native, _stack__atof_native, _pcsp__atof_native},
	{"_decimal_to_f64", _entry__decimal_to_f64, _size__decimal_to_f64, _stack__decimal_to_f64, _pcsp__decimal_to_f64},
	{"_left_shift", _entry__left_shift, _size__left_shift, _stack__left_shift, _pcsp__left_shift},
	{"_right_shift", _entry__right_shift, _size__right_shift, _stack__right_shift, _pcsp__right_shift},
	{"_vsigned", _entry__vsigned, _size__vsigned, _stack__vsigned, _pcsp__vsigned},
	{"_vstring", _entry__vstring, _size__vstring, _stack__vstring, _pcsp__vstring},
	{"_vunsigned", _entry__vunsigned, _size__vunsigned, _stack__vunsigned, _pcsp__vunsigned},
}
