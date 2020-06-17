// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: error_code.proto

package errorCodeProto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrorCode int32

const (
	ErrorCode_startId ErrorCode = 0
	// 客户端请求进入捕鱼房间时,服务端发现已经在房间中了
	ErrorCode_already_in_fish_room_when_req_enter_fish_room ErrorCode = 1
	// 客户端请求离开捕鱼房间时,服务端发现还未进入房间
	ErrorCode_has_not_in_fish_room_when_req_leave_fish_room ErrorCode = 2
	// 客户端请求子弹时,服务端发现还未进入房间
	ErrorCode_has_not_in_fish_room_when_player_shoot ErrorCode = 3
	// 客户端击中鱼时,服务端发现还未进入房间
	ErrorCode_has_not_in_fish_room_when_hit_shoot ErrorCode = 4
	// 客户端请求子弹时,金币不够
	ErrorCode_not_enought_gold_coin_when_player_shoot ErrorCode = 5
	// 客户端击中鱼时上传的子弹 id 错误
	ErrorCode_bullet_id_error_when_player_hit_shoot ErrorCode = 6
	// 客户端进入房间时,上传的房间 id 错误
	ErrorCode_room_id_error_when_player_enter_fishRoom ErrorCode = 7
	// 客户端请求进入牛牛房间时,服务端发现已经在牛牛房间中了
	ErrorCode_already_in_niuniu_room_when_req_enter_niuniu_room ErrorCode = 8
	// 客户端请求离开牛牛房间时,服务端发现还未进入房间
	ErrorCode_has_not_in_niuniu_room_when_req_leave_niuniu_room ErrorCode = 9
	// 客户端进行牛牛操作时,发现不在房间中
	ErrorCode_has_not_in_niuniu_room_when_play ErrorCode = 10
	// 请求进入大厅的某个牛牛房间时,发现房间不存在
	ErrorCode_niuniu_public_room_is_not_exist_when_request_enter_room ErrorCode = 11
	// 牛牛断线重连时,上发的 room id 有误
	ErrorCode_niuniu_room_id_error_when_reconnect ErrorCode = 12
	// 牛牛, 提示玩家在游戏中不能离开房间
	ErrorCode_can_not_leave_niuniu_room_when_playing ErrorCode = 13
	//玩家在某个房间中，无法新建房间
	ErrorCode_can_not_build_room_when_in_room ErrorCode = 14
	//自建niuniu房时获取roomid写redis错误
	ErrorCode_redis_error_when_self_build_niuniu_room ErrorCode = 15
	//自建niuniu房存入房间信息时写redis错误
	ErrorCode_redis_error_when_self_build_niuniu_room_save_room_info ErrorCode = 16
	//自建牛牛房的参数超出限制
	ErrorCode_can_not_self_build_room_with_bad_request ErrorCode = 17
	//不允许中途加入的自建房
	ErrorCode_can_not_enter_self_build_room ErrorCode = 18
	// 登录失败, redis setnx 出错
	ErrorCode_redis_setnx_has_error_when_login ErrorCode = 150
	// 登录失败, 因为有多人同时登录, 被其他人挤下线了
	ErrorCode_other_already_login ErrorCode = 151
	// 登录时数据库出错
	ErrorCode_mysql_has_error_when_login ErrorCode = 152
	// 登录时从 redis 中通过 wxid 查询 uid 出错
	ErrorCode_redis_get_uid_by_wxid_error_when_login ErrorCode = 153
	// 牛牛结算时,写 mysql 出现错误,导致输赢无效
	ErrorCode_write_mysql_error_when_niuniu_settlement ErrorCode = 154
	// 登录时从 redis 中通过 uid 查询玩家基本信息出错
	ErrorCode_redis_get_player_base_info_by_uid_error_when_login ErrorCode = 155
	// 创建角色时, 初始化角色数据失败, 可能是从 redis set ( key = digital_id:id )获取到空字符串,或者字符串不是数字
	ErrorCode_lobby_init_new_account_fail ErrorCode = 156
	// 创建角色时, 数据库执行事务出错
	ErrorCode_mysql_has_error_when_create_account ErrorCode = 157
	// 用户输入邀请码时,查询数据库表 agent 出现错误
	ErrorCode_mysql_query_agent_has_error_when_player_input_invite_code ErrorCode = 158
	//
	// 用户输入邀请码时,redis setnx 获取锁失败,表明有其他设备正在使用同一个账号输入邀请码
	ErrorCode_redis_setnx_fail_when_player_input_invite_code ErrorCode = 160
	// 用户输入邀请码时,redis setnx 操作 redis 出现错误
	ErrorCode_redis_setnx_has_error_when_player_input_invite_code ErrorCode = 161
	// 无效的邀请码
	ErrorCode_invalid_invite_code ErrorCode = 162
	// 用户输入邀请码时,写数据库表 subordinate 出现错误
	ErrorCode_mysql_insert_subordinate_has_error_when_player_input_invite_code ErrorCode = 163
	// 已经输入过邀请码了, 不能重复输入邀请码
	ErrorCode_can_not_repeat_input_invite_code ErrorCode = 164
	// 用户输入邀请码时,写数据库表 agent 出现错误
	ErrorCode_mysql_insert_agent_has_error_when_player_input_invite_code ErrorCode = 165
)

var ErrorCode_name = map[int32]string{
	0:   "startId",
	1:   "already_in_fish_room_when_req_enter_fish_room",
	2:   "has_not_in_fish_room_when_req_leave_fish_room",
	3:   "has_not_in_fish_room_when_player_shoot",
	4:   "has_not_in_fish_room_when_hit_shoot",
	5:   "not_enought_gold_coin_when_player_shoot",
	6:   "bullet_id_error_when_player_hit_shoot",
	7:   "room_id_error_when_player_enter_fishRoom",
	8:   "already_in_niuniu_room_when_req_enter_niuniu_room",
	9:   "has_not_in_niuniu_room_when_req_leave_niuniu_room",
	10:  "has_not_in_niuniu_room_when_play",
	11:  "niuniu_public_room_is_not_exist_when_request_enter_room",
	12:  "niuniu_room_id_error_when_reconnect",
	13:  "can_not_leave_niuniu_room_when_playing",
	14:  "can_not_build_room_when_in_room",
	15:  "redis_error_when_self_build_niuniu_room",
	16:  "redis_error_when_self_build_niuniu_room_save_room_info",
	17:  "can_not_self_build_room_with_bad_request",
	18:  "can_not_enter_self_build_room",
	150: "redis_setnx_has_error_when_login",
	151: "other_already_login",
	152: "mysql_has_error_when_login",
	153: "redis_get_uid_by_wxid_error_when_login",
	154: "write_mysql_error_when_niuniu_settlement",
	155: "redis_get_player_base_info_by_uid_error_when_login",
	156: "lobby_init_new_account_fail",
	157: "mysql_has_error_when_create_account",
	158: "mysql_query_agent_has_error_when_player_input_invite_code",
	160: "redis_setnx_fail_when_player_input_invite_code",
	161: "redis_setnx_has_error_when_player_input_invite_code",
	162: "invalid_invite_code",
	163: "mysql_insert_subordinate_has_error_when_player_input_invite_code",
	164: "can_not_repeat_input_invite_code",
	165: "mysql_insert_agent_has_error_when_player_input_invite_code",
}

var ErrorCode_value = map[string]int32{
	"startId": 0,
	"already_in_fish_room_when_req_enter_fish_room":                    1,
	"has_not_in_fish_room_when_req_leave_fish_room":                    2,
	"has_not_in_fish_room_when_player_shoot":                           3,
	"has_not_in_fish_room_when_hit_shoot":                              4,
	"not_enought_gold_coin_when_player_shoot":                          5,
	"bullet_id_error_when_player_hit_shoot":                            6,
	"room_id_error_when_player_enter_fishRoom":                         7,
	"already_in_niuniu_room_when_req_enter_niuniu_room":                8,
	"has_not_in_niuniu_room_when_req_leave_niuniu_room":                9,
	"has_not_in_niuniu_room_when_play":                                 10,
	"niuniu_public_room_is_not_exist_when_request_enter_room":          11,
	"niuniu_room_id_error_when_reconnect":                              12,
	"can_not_leave_niuniu_room_when_playing":                           13,
	"can_not_build_room_when_in_room":                                  14,
	"redis_error_when_self_build_niuniu_room":                          15,
	"redis_error_when_self_build_niuniu_room_save_room_info":           16,
	"can_not_self_build_room_with_bad_request":                         17,
	"can_not_enter_self_build_room":                                    18,
	"redis_setnx_has_error_when_login":                                 150,
	"other_already_login":                                              151,
	"mysql_has_error_when_login":                                       152,
	"redis_get_uid_by_wxid_error_when_login":                           153,
	"write_mysql_error_when_niuniu_settlement":                         154,
	"redis_get_player_base_info_by_uid_error_when_login":               155,
	"lobby_init_new_account_fail":                                      156,
	"mysql_has_error_when_create_account":                              157,
	"mysql_query_agent_has_error_when_player_input_invite_code":        158,
	"redis_setnx_fail_when_player_input_invite_code":                   160,
	"redis_setnx_has_error_when_player_input_invite_code":              161,
	"invalid_invite_code":                                              162,
	"mysql_insert_subordinate_has_error_when_player_input_invite_code": 163,
	"can_not_repeat_input_invite_code":                                 164,
	"mysql_insert_agent_has_error_when_player_input_invite_code":       165,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c5513ac0a8e17e40, []int{0}
}

type S2CErrorCode struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CErrorCode) Reset()         { *m = S2CErrorCode{} }
func (m *S2CErrorCode) String() string { return proto.CompactTextString(m) }
func (*S2CErrorCode) ProtoMessage()    {}
func (*S2CErrorCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5513ac0a8e17e40, []int{0}
}
func (m *S2CErrorCode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2CErrorCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2CErrorCode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *S2CErrorCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CErrorCode.Merge(m, src)
}
func (m *S2CErrorCode) XXX_Size() int {
	return m.Size()
}
func (m *S2CErrorCode) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CErrorCode.DiscardUnknown(m)
}

var xxx_messageInfo_S2CErrorCode proto.InternalMessageInfo

func (m *S2CErrorCode) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterEnum("errorCodeProto.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterType((*S2CErrorCode)(nil), "errorCodeProto.s2cErrorCode")
}

func init() { proto.RegisterFile("error_code.proto", fileDescriptor_c5513ac0a8e17e40) }

var fileDescriptor_c5513ac0a8e17e40 = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdb, 0x4e, 0x14, 0x4b,
	0x14, 0xa5, 0xce, 0xe1, 0x72, 0xd8, 0x70, 0x38, 0x75, 0xca, 0x17, 0xa2, 0x71, 0x18, 0x41, 0x60,
	0x04, 0x21, 0x01, 0xa2, 0x28, 0x26, 0x6a, 0x34, 0x3c, 0xf8, 0x66, 0xf8, 0x81, 0x4a, 0x5f, 0xf6,
	0x4c, 0x57, 0xd2, 0x54, 0x0d, 0x55, 0xd5, 0x0c, 0xf3, 0x23, 0x0a, 0xde, 0xe2, 0xf5, 0x5f, 0x7c,
	0xf4, 0x13, 0x0c, 0xfe, 0x88, 0xa9, 0xae, 0x69, 0xa6, 0x07, 0x7b, 0x10, 0xdf, 0x66, 0x6a, 0xaf,
	0xb5, 0xd7, 0xda, 0x6b, 0xef, 0xcc, 0x00, 0x45, 0xad, 0x95, 0xe6, 0x91, 0x8a, 0x71, 0xbd, 0xad,
	0x95, 0x55, 0x6c, 0x26, 0x7f, 0x79, 0xaa, 0x62, 0x7c, 0xee, 0xbe, 0xcf, 0xcf, 0xc3, 0xb4, 0xd9,
	0x8c, 0x76, 0x8b, 0x47, 0xc6, 0x60, 0xd4, 0xa1, 0x67, 0x49, 0x9d, 0x34, 0xc6, 0xf6, 0xf2, 0xcf,
	0x2b, 0x27, 0xd3, 0x30, 0x79, 0x46, 0x63, 0x53, 0x30, 0x61, 0x6c, 0xa0, 0xed, 0xb3, 0x98, 0x8e,
	0xb0, 0x0d, 0x58, 0x0b, 0x52, 0x8d, 0x41, 0xdc, 0xe5, 0x42, 0xf2, 0xa6, 0x30, 0x09, 0xd7, 0x4a,
	0xed, 0xf3, 0x4e, 0x82, 0x92, 0x6b, 0x3c, 0xe0, 0x28, 0x2d, 0xea, 0x7e, 0x81, 0x12, 0x47, 0x49,
	0x02, 0xc3, 0xa5, 0xb2, 0x43, 0x28, 0x29, 0x06, 0x87, 0x58, 0xa2, 0xfc, 0xc5, 0x56, 0x60, 0x69,
	0x38, 0xa5, 0x9d, 0x06, 0x5d, 0xd4, 0xdc, 0x24, 0x4a, 0x59, 0xfa, 0x37, 0x5b, 0x86, 0x85, 0xe1,
	0xd8, 0x44, 0xd8, 0x1e, 0x70, 0x94, 0xad, 0xc2, 0xb2, 0x03, 0xa1, 0x54, 0x59, 0x2b, 0xb1, 0xbc,
	0xa5, 0xd2, 0x98, 0x47, 0x4a, 0xc8, 0x8a, 0xae, 0x63, 0xec, 0x16, 0x2c, 0x86, 0x59, 0x9a, 0xa2,
	0xe5, 0x22, 0xe6, 0x3e, 0xd4, 0x32, 0xac, 0xdf, 0x77, 0x9c, 0xdd, 0x86, 0x46, 0x2e, 0x58, 0x09,
	0xec, 0xc7, 0xb1, 0xe7, 0x46, 0x9b, 0x60, 0x77, 0x60, 0xa3, 0x14, 0xa0, 0x14, 0x99, 0x14, 0x59,
	0x65, 0x84, 0xa5, 0x12, 0xfd, 0xc7, 0xd1, 0x4a, 0x53, 0x56, 0xd2, 0x7c, 0x8c, 0x65, 0xda, 0x24,
	0xbb, 0x09, 0xf5, 0x8b, 0x68, 0xce, 0x24, 0x05, 0xf6, 0x00, 0xb6, 0x7b, 0xa5, 0x76, 0x16, 0xa6,
	0x22, 0xf2, 0x08, 0xe1, 0x79, 0x78, 0x24, 0x8c, 0x3d, 0x53, 0xc9, 0xd0, 0xd8, 0x9e, 0xc1, 0x5c,
	0x62, 0xca, 0xe5, 0x5f, 0xee, 0x3b, 0x98, 0x82, 0xc6, 0x48, 0x49, 0x89, 0x91, 0xa5, 0xd3, 0x6e,
	0xa9, 0x51, 0x20, 0xf3, 0x9e, 0xbf, 0x58, 0xed, 0xdb, 0x11, 0xb2, 0x45, 0xff, 0x65, 0x0b, 0x30,
	0x57, 0x60, 0xc3, 0x4c, 0xa4, 0x71, 0x09, 0x24, 0xa4, 0x57, 0x9e, 0x71, 0x0b, 0xd5, 0x18, 0x0b,
	0x53, 0x16, 0x34, 0x98, 0x36, 0x7b, 0x94, 0x72, 0x12, 0xff, 0xb1, 0x1d, 0xb8, 0x7b, 0x49, 0x30,
	0x37, 0xce, 0x9c, 0x1f, 0x46, 0x36, 0x15, 0xa5, 0x6e, 0xc3, 0x85, 0x9b, 0x12, 0xc5, 0x5b, 0x12,
	0x36, 0xe1, 0x61, 0x10, 0x17, 0xd9, 0xd0, 0xff, 0xd9, 0x0d, 0xb8, 0x5e, 0xa0, 0x7d, 0x50, 0xe7,
	0x38, 0x94, 0xb1, 0x45, 0xa8, 0x7b, 0x33, 0x06, 0xad, 0x3c, 0xe2, 0x6e, 0x45, 0x25, 0x5b, 0xa9,
	0x6a, 0x09, 0x49, 0x5f, 0x10, 0x36, 0x0b, 0x57, 0x94, 0x4d, 0x50, 0xf3, 0xe2, 0x62, 0x7c, 0xe5,
	0x25, 0x61, 0x73, 0x70, 0x75, 0xbf, 0x6b, 0x0e, 0xd2, 0x6a, 0xea, 0x31, 0x61, 0xab, 0xb0, 0xe4,
	0x15, 0x5a, 0x68, 0x79, 0x26, 0x62, 0x1e, 0x76, 0x79, 0xe7, 0x68, 0x70, 0x39, 0x1e, 0x7c, 0x42,
	0xd8, 0x1a, 0x34, 0x3a, 0x5a, 0x58, 0xe4, 0xbe, 0x67, 0x09, 0xd2, 0x8b, 0xc5, 0xa0, 0xb5, 0x29,
	0xee, 0xa3, 0xb4, 0xf4, 0x15, 0x61, 0xdb, 0xb0, 0xd9, 0xef, 0xdd, 0xbb, 0xf3, 0x30, 0x30, 0x98,
	0xc7, 0xe5, 0x84, 0xb2, 0x2a, 0x9d, 0xd7, 0x84, 0xd5, 0xe1, 0x5a, 0xaa, 0xc2, 0xd0, 0x5d, 0xbe,
	0xb0, 0x5c, 0x62, 0x87, 0x07, 0x51, 0xa4, 0x32, 0x69, 0x79, 0x33, 0x10, 0x29, 0x7d, 0x43, 0x58,
	0x03, 0x16, 0x2a, 0xe7, 0x8a, 0x34, 0x06, 0x16, 0x0b, 0x38, 0x7d, 0x4b, 0xd8, 0x43, 0xb8, 0xef,
	0x91, 0x07, 0x19, 0xea, 0x2e, 0x0f, 0x5a, 0x28, 0xed, 0x79, 0x56, 0xcf, 0x9b, 0x90, 0xed, 0xcc,
	0xdd, 0xff, 0xa1, 0x1b, 0xd1, 0xfd, 0xc0, 0xd1, 0x77, 0x84, 0x6d, 0xc1, 0x7a, 0x79, 0x05, 0xce,
	0xc0, 0x6f, 0x48, 0xef, 0x09, 0xbb, 0x07, 0x5b, 0x17, 0xec, 0x6d, 0x28, 0xf3, 0x43, 0xbe, 0x4a,
	0x21, 0x0f, 0x83, 0x54, 0xc4, 0x03, 0x95, 0x8f, 0x84, 0xed, 0xc2, 0x63, 0x3f, 0x88, 0x90, 0x06,
	0xb5, 0xe5, 0x26, 0x0b, 0x95, 0x8e, 0x85, 0x74, 0xf3, 0x5e, 0x56, 0xe0, 0x13, 0x71, 0x27, 0x55,
	0x5c, 0x9d, 0xc6, 0x36, 0x06, 0xb6, 0x02, 0xf6, 0x99, 0xb0, 0x47, 0xb0, 0x33, 0xa0, 0xf6, 0x67,
	0xb9, 0x7d, 0x21, 0x4f, 0xe8, 0xd7, 0xd3, 0x1a, 0xf9, 0x76, 0x5a, 0x23, 0xdf, 0x4f, 0x6b, 0xe4,
	0xf8, 0x47, 0x6d, 0x24, 0x1c, 0xcf, 0xff, 0x68, 0xb6, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x3e,
	0x34, 0x12, 0xcc, 0x7c, 0x06, 0x00, 0x00,
}

func (m *S2CErrorCode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2CErrorCode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *S2CErrorCode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Code != 0 {
		i = encodeVarintErrorCode(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintErrorCode(dAtA []byte, offset int, v uint64) int {
	offset -= sovErrorCode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *S2CErrorCode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovErrorCode(uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovErrorCode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErrorCode(x uint64) (n int) {
	return sovErrorCode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *S2CErrorCode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrorCode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: s2cErrorCode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: s2cErrorCode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrorCode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipErrorCode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrorCode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthErrorCode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipErrorCode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrorCode
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrorCode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrorCode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthErrorCode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErrorCode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErrorCode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErrorCode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrorCode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErrorCode = fmt.Errorf("proto: unexpected end of group")
)
