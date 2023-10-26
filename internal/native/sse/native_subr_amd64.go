//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	"github.com/bytedance/sonic/loader"
)

const (
	_entry__f32toa                 = 31344
	_entry__f64toa                 = 128
	_entry__format_significand     = 36272
	_entry__format_integer         = 3280
	_entry__fsm_exec               = 18832
	_entry__advance_string         = 15024
	_entry__advance_string_default = 37808
	_entry__do_skip_number         = 21376
	_entry__get_by_path            = 26768
	_entry__skip_one_fast          = 22896
	_entry__unescape               = 38752
	_entry__unhex16_is             = 9584
	_entry__html_escape            = 9776
	_entry__i64toa                 = 3712
	_entry__u64toa                 = 3984
	_entry__lspace                 = 16
	_entry__quote                  = 5472
	_entry__skip_array             = 18800
	_entry__skip_number            = 22496
	_entry__skip_object            = 21024
	_entry__skip_one               = 22672
	_entry__unquote                = 7248
	_entry__validate_one           = 22720
	_entry__validate_utf8          = 30096
	_entry__validate_utf8_fast     = 30784
	_entry__value                  = 13072
	_entry__vnumber                = 16400
	_entry__atof_eisel_lemire64    = 11072
	_entry__atof_native            = 12464
	_entry__decimal_to_f64         = 11472
	_entry__left_shift             = 36752
	_entry__right_shift            = 37296
	_entry__vsigned                = 18048
	_entry__vstring                = 14848
	_entry__vunsigned              = 18416
)

const (
	_stack__f32toa                 = 64
	_stack__f64toa                 = 80
	_stack__format_significand     = 24
	_stack__format_integer         = 16
	_stack__fsm_exec               = 160
	_stack__advance_string         = 72
	_stack__advance_string_default = 56
	_stack__do_skip_number         = 32
	_stack__get_by_path            = 264
	_stack__skip_one_fast          = 136
	_stack__unescape               = 64
	_stack__unhex16_is             = 8
	_stack__html_escape            = 64
	_stack__i64toa                 = 16
	_stack__u64toa                 = 8
	_stack__lspace                 = 8
	_stack__quote                  = 80
	_stack__skip_array             = 168
	_stack__skip_number            = 88
	_stack__skip_object            = 168
	_stack__skip_one               = 168
	_stack__unquote                = 112
	_stack__validate_one           = 168
	_stack__validate_utf8          = 48
	_stack__validate_utf8_fast     = 24
	_stack__value                  = 352
	_stack__vnumber                = 264
	_stack__atof_eisel_lemire64    = 40
	_stack__atof_native            = 144
	_stack__decimal_to_f64         = 88
	_stack__left_shift             = 32
	_stack__right_shift            = 16
	_stack__vsigned                = 16
	_stack__vstring                = 128
	_stack__vunsigned              = 24
)

const (
	_size__f32toa                 = 3696
	_size__f64toa                 = 3152
	_size__format_significand     = 480
	_size__format_integer         = 432
	_size__fsm_exec               = 1656
	_size__advance_string         = 1328
	_size__advance_string_default = 944
	_size__do_skip_number         = 908
	_size__get_by_path            = 3328
	_size__skip_one_fast          = 3348
	_size__unescape               = 704
	_size__unhex16_is             = 128
	_size__html_escape            = 1296
	_size__i64toa                 = 272
	_size__u64toa                 = 1440
	_size__lspace                 = 96
	_size__quote                  = 1760
	_size__skip_array             = 32
	_size__skip_number            = 160
	_size__skip_object            = 32
	_size__skip_one               = 32
	_size__unquote                = 2336
	_size__validate_one           = 48
	_size__validate_utf8          = 688
	_size__validate_utf8_fast     = 544
	_size__value                  = 1268
	_size__vnumber                = 1648
	_size__atof_eisel_lemire64    = 400
	_size__atof_native            = 608
	_size__decimal_to_f64         = 992
	_size__left_shift             = 544
	_size__right_shift            = 480
	_size__vsigned                = 368
	_size__vstring                = 128
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
		{3638, 64},
		{3642, 48},
		{3643, 40},
		{3645, 32},
		{3647, 24},
		{3649, 16},
		{3651, 8},
		{3652, 0},
		{3682, 64},
	}
	_pcsp__f64toa = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{3033, 56},
		{3037, 48},
		{3038, 40},
		{3040, 32},
		{3042, 24},
		{3044, 16},
		{3046, 8},
		{3047, 0},
		{3138, 56},
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
		{1317, 88},
		{1321, 48},
		{1322, 40},
		{1324, 32},
		{1326, 24},
		{1328, 16},
		{1330, 8},
		{1331, 0},
		{1656, 88},
	}
	_pcsp__advance_string = [][2]uint32{
		{14, 0},
		{18, 8},
		{20, 16},
		{22, 24},
		{24, 32},
		{26, 40},
		{27, 48},
		{587, 72},
		{591, 48},
		{592, 40},
		{594, 32},
		{596, 24},
		{598, 16},
		{600, 8},
		{601, 0},
		{1325, 72},
	}
	_pcsp__advance_string_default = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{402, 56},
		{406, 48},
		{407, 40},
		{409, 32},
		{411, 24},
		{413, 16},
		{415, 8},
		{416, 0},
		{936, 56},
	}
	_pcsp__do_skip_number = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{771, 32},
		{772, 24},
		{774, 16},
		{776, 8},
		{777, 0},
		{908, 32},
	}
	_pcsp__get_by_path = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{3278, 104},
		{3282, 48},
		{3283, 40},
		{3285, 32},
		{3287, 24},
		{3289, 16},
		{3291, 8},
		{3292, 0},
		{3317, 104},
	}
	_pcsp__skip_one_fast = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{600, 136},
		{604, 48},
		{605, 40},
		{607, 32},
		{609, 24},
		{611, 16},
		{613, 8},
		{614, 0},
		{3348, 136},
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
		{1281, 64},
		{1285, 48},
		{1286, 40},
		{1288, 32},
		{1290, 24},
		{1292, 16},
		{1294, 8},
		{1296, 0},
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
		{608, 0},
		{882, 8},
		{976, 0},
		{1434, 8},
		{1436, 0},
	}
	_pcsp__lspace = [][2]uint32{
		{1, 0},
		{85, 8},
		{87, 0},
	}
	_pcsp__quote = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{1701, 80},
		{1705, 48},
		{1706, 40},
		{1708, 32},
		{1710, 24},
		{1712, 16},
		{1714, 8},
		{1715, 0},
		{1750, 80},
	}
	_pcsp__skip_array = [][2]uint32{
		{1, 0},
		{26, 8},
		{32, 0},
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
		{26, 8},
		{32, 0},
	}
	_pcsp__skip_one = [][2]uint32{
		{1, 0},
		{26, 8},
		{32, 0},
	}
	_pcsp__unquote = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{1614, 104},
		{1618, 48},
		{1619, 40},
		{1621, 32},
		{1623, 24},
		{1625, 16},
		{1627, 8},
		{1628, 0},
		{2329, 104},
	}
	_pcsp__validate_one = [][2]uint32{
		{1, 0},
		{31, 8},
		{37, 0},
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
		{247, 24},
		{251, 16},
		{252, 8},
		{253, 0},
		{527, 24},
		{531, 16},
		{532, 8},
		{534, 0},
	}
	_pcsp__value = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{439, 88},
		{443, 48},
		{444, 40},
		{446, 32},
		{448, 24},
		{450, 16},
		{452, 8},
		{453, 0},
		{1268, 88},
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
		{1642, 120},
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
		{951, 56},
		{955, 48},
		{956, 40},
		{958, 32},
		{960, 24},
		{962, 16},
		{964, 8},
		{965, 0},
		{977, 56},
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
		{111, 16},
		{112, 8},
		{113, 0},
		{124, 16},
		{125, 8},
		{126, 0},
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
		{72, 24},
		{73, 16},
		{75, 8},
		{76, 0},
		{87, 24},
		{88, 16},
		{90, 8},
		{91, 0},
		{114, 24},
		{115, 16},
		{117, 8},
		{118, 0},
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
