// Code generated by protoc-gen-go. DO NOT EDIT.
// source: post/postpb/post.proto

package postpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

// defenition of post
type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Image                string   `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Tag                  []string `protobuf:"bytes,6,rep,name=tag,proto3" json:"tag,omitempty"`
	UserId               string   `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_2037b85c7cbf25b6, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Post) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Post) GetTag() []string {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *Post) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// create a new post
type CreatePostRequest struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRequest) Reset()         { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()    {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2037b85c7cbf25b6, []int{1}
}

func (m *CreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRequest.Unmarshal(m, b)
}
func (m *CreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRequest.Marshal(b, m, deterministic)
}
func (m *CreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRequest.Merge(m, src)
}
func (m *CreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostRequest.Size(m)
}
func (m *CreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRequest proto.InternalMessageInfo

func (m *CreatePostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type CreatePostResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostResponse) Reset()         { *m = CreatePostResponse{} }
func (m *CreatePostResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePostResponse) ProtoMessage()    {}
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2037b85c7cbf25b6, []int{2}
}

func (m *CreatePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostResponse.Unmarshal(m, b)
}
func (m *CreatePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostResponse.Marshal(b, m, deterministic)
}
func (m *CreatePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostResponse.Merge(m, src)
}
func (m *CreatePostResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePostResponse.Size(m)
}
func (m *CreatePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostResponse proto.InternalMessageInfo

func (m *CreatePostResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// delete a post
type DeletePostRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostRequest) Reset()         { *m = DeletePostRequest{} }
func (m *DeletePostRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePostRequest) ProtoMessage()    {}
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2037b85c7cbf25b6, []int{3}
}

func (m *DeletePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostRequest.Unmarshal(m, b)
}
func (m *DeletePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostRequest.Marshal(b, m, deterministic)
}
func (m *DeletePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostRequest.Merge(m, src)
}
func (m *DeletePostRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePostRequest.Size(m)
}
func (m *DeletePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostRequest proto.InternalMessageInfo

func (m *DeletePostRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeletePostResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostResponse) Reset()         { *m = DeletePostResponse{} }
func (m *DeletePostResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePostResponse) ProtoMessage()    {}
func (*DeletePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2037b85c7cbf25b6, []int{4}
}

func (m *DeletePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostResponse.Unmarshal(m, b)
}
func (m *DeletePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostResponse.Marshal(b, m, deterministic)
}
func (m *DeletePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostResponse.Merge(m, src)
}
func (m *DeletePostResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePostResponse.Size(m)
}
func (m *DeletePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostResponse proto.InternalMessageInfo

func (m *DeletePostResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// get all posts by user_id
type GetAllPostsByUserIDRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllPostsByUserIDRequest) Reset()         { *m = GetAllPostsByUserIDRequest{} }
func (m *GetAllPostsByUserIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllPostsByUserIDRequest) ProtoMessage()    {}
func (*GetAllPostsByUserIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2037b85c7cbf25b6, []int{5}
}

func (m *GetAllPostsByUserIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPostsByUserIDRequest.Unmarshal(m, b)
}
func (m *GetAllPostsByUserIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPostsByUserIDRequest.Marshal(b, m, deterministic)
}
func (m *GetAllPostsByUserIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPostsByUserIDRequest.Merge(m, src)
}
func (m *GetAllPostsByUserIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllPostsByUserIDRequest.Size(m)
}
func (m *GetAllPostsByUserIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPostsByUserIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPostsByUserIDRequest proto.InternalMessageInfo

func (m *GetAllPostsByUserIDRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetAllPostsByUserIDResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllPostsByUserIDResponse) Reset()         { *m = GetAllPostsByUserIDResponse{} }
func (m *GetAllPostsByUserIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllPostsByUserIDResponse) ProtoMessage()    {}
func (*GetAllPostsByUserIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2037b85c7cbf25b6, []int{6}
}

func (m *GetAllPostsByUserIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPostsByUserIDResponse.Unmarshal(m, b)
}
func (m *GetAllPostsByUserIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPostsByUserIDResponse.Marshal(b, m, deterministic)
}
func (m *GetAllPostsByUserIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPostsByUserIDResponse.Merge(m, src)
}
func (m *GetAllPostsByUserIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllPostsByUserIDResponse.Size(m)
}
func (m *GetAllPostsByUserIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPostsByUserIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPostsByUserIDResponse proto.InternalMessageInfo

func (m *GetAllPostsByUserIDResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

// Search Posts
type SearchPostsRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	How                  string   `protobuf:"bytes,2,opt,name=how,proto3" json:"how,omitempty"`
	Keywords             []string `protobuf:"bytes,3,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPostsRequest) Reset()         { *m = SearchPostsRequest{} }
func (m *SearchPostsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchPostsRequest) ProtoMessage()    {}
func (*SearchPostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2037b85c7cbf25b6, []int{7}
}

func (m *SearchPostsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPostsRequest.Unmarshal(m, b)
}
func (m *SearchPostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPostsRequest.Marshal(b, m, deterministic)
}
func (m *SearchPostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPostsRequest.Merge(m, src)
}
func (m *SearchPostsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchPostsRequest.Size(m)
}
func (m *SearchPostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPostsRequest proto.InternalMessageInfo

func (m *SearchPostsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SearchPostsRequest) GetHow() string {
	if m != nil {
		return m.How
	}
	return ""
}

func (m *SearchPostsRequest) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

type SearchPostsResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPostsResponse) Reset()         { *m = SearchPostsResponse{} }
func (m *SearchPostsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchPostsResponse) ProtoMessage()    {}
func (*SearchPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2037b85c7cbf25b6, []int{8}
}

func (m *SearchPostsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPostsResponse.Unmarshal(m, b)
}
func (m *SearchPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPostsResponse.Marshal(b, m, deterministic)
}
func (m *SearchPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPostsResponse.Merge(m, src)
}
func (m *SearchPostsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchPostsResponse.Size(m)
}
func (m *SearchPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPostsResponse proto.InternalMessageInfo

func (m *SearchPostsResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

// scraping by URL
type PostURLRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostURLRequest) Reset()         { *m = PostURLRequest{} }
func (m *PostURLRequest) String() string { return proto.CompactTextString(m) }
func (*PostURLRequest) ProtoMessage()    {}
func (*PostURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2037b85c7cbf25b6, []int{9}
}

func (m *PostURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostURLRequest.Unmarshal(m, b)
}
func (m *PostURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostURLRequest.Marshal(b, m, deterministic)
}
func (m *PostURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostURLRequest.Merge(m, src)
}
func (m *PostURLRequest) XXX_Size() int {
	return xxx_messageInfo_PostURLRequest.Size(m)
}
func (m *PostURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostURLRequest proto.InternalMessageInfo

func (m *PostURLRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type PostResponse struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Image                string   `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostResponse) Reset()         { *m = PostResponse{} }
func (m *PostResponse) String() string { return proto.CompactTextString(m) }
func (*PostResponse) ProtoMessage()    {}
func (*PostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2037b85c7cbf25b6, []int{10}
}

func (m *PostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostResponse.Unmarshal(m, b)
}
func (m *PostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostResponse.Marshal(b, m, deterministic)
}
func (m *PostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostResponse.Merge(m, src)
}
func (m *PostResponse) XXX_Size() int {
	return xxx_messageInfo_PostResponse.Size(m)
}
func (m *PostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostResponse proto.InternalMessageInfo

func (m *PostResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PostResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PostResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PostResponse) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func init() {
	proto.RegisterType((*Post)(nil), "postpb.Post")
	proto.RegisterType((*CreatePostRequest)(nil), "postpb.CreatePostRequest")
	proto.RegisterType((*CreatePostResponse)(nil), "postpb.CreatePostResponse")
	proto.RegisterType((*DeletePostRequest)(nil), "postpb.DeletePostRequest")
	proto.RegisterType((*DeletePostResponse)(nil), "postpb.DeletePostResponse")
	proto.RegisterType((*GetAllPostsByUserIDRequest)(nil), "postpb.GetAllPostsByUserIDRequest")
	proto.RegisterType((*GetAllPostsByUserIDResponse)(nil), "postpb.GetAllPostsByUserIDResponse")
	proto.RegisterType((*SearchPostsRequest)(nil), "postpb.SearchPostsRequest")
	proto.RegisterType((*SearchPostsResponse)(nil), "postpb.SearchPostsResponse")
	proto.RegisterType((*PostURLRequest)(nil), "postpb.PostURLRequest")
	proto.RegisterType((*PostResponse)(nil), "postpb.PostResponse")
}

func init() { proto.RegisterFile("post/postpb/post.proto", fileDescriptor_2037b85c7cbf25b6) }

var fileDescriptor_2037b85c7cbf25b6 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0xeb, 0x38, 0x2f, 0xcf, 0x33, 0x6e, 0x2b, 0x3a, 0xad, 0xca, 0xe2, 0x5e, 0xac, 0xed,
	0x85, 0x53, 0x90, 0x8a, 0x7a, 0xe0, 0x84, 0x5a, 0x82, 0x0a, 0x12, 0x07, 0x94, 0xaa, 0x17, 0x38,
	0x80, 0x1b, 0x8f, 0x92, 0x15, 0x6e, 0xd6, 0x78, 0x37, 0x44, 0xfd, 0x06, 0x7c, 0x0d, 0xbe, 0x29,
	0x9a, 0x75, 0xfd, 0x86, 0x43, 0x2a, 0x2e, 0x89, 0xe7, 0xed, 0x3f, 0x33, 0x9e, 0x9f, 0x0c, 0xc7,
	0x99, 0x36, 0xf6, 0x05, 0xff, 0x64, 0xb7, 0xee, 0x6f, 0x9c, 0xe5, 0xda, 0x6a, 0x1c, 0x16, 0x2e,
	0xf9, 0xcb, 0x83, 0xfe, 0x47, 0x6d, 0x2c, 0xee, 0x43, 0x4f, 0x25, 0xc2, 0x8b, 0xbc, 0xe7, 0xff,
	0x4f, 0x7b, 0x2a, 0xc1, 0x27, 0xe0, 0xaf, 0xf2, 0x54, 0xf4, 0x9c, 0x83, 0x1f, 0xf1, 0x08, 0x06,
	0x56, 0xd9, 0x94, 0x84, 0xef, 0x7c, 0x85, 0x81, 0x11, 0x04, 0x09, 0x99, 0x59, 0xae, 0x32, 0xab,
	0xf4, 0x52, 0xf4, 0x5d, 0xac, 0xe9, 0xe2, 0x3a, 0x75, 0x17, 0xcf, 0x49, 0x0c, 0x8a, 0x3a, 0x67,
	0xb0, 0xbe, 0x8d, 0xe7, 0x62, 0x18, 0xf9, 0xac, 0x6f, 0xe3, 0x39, 0x3e, 0x85, 0xd1, 0xca, 0x50,
	0xfe, 0x45, 0x25, 0x62, 0xe4, 0x32, 0x87, 0x6c, 0xbe, 0x4f, 0xe4, 0x39, 0x1c, 0xbc, 0xc9, 0x29,
	0xb6, 0xc4, 0x83, 0x4e, 0xe9, 0xfb, 0x8a, 0x8c, 0xc5, 0x08, 0xfa, 0xbc, 0x82, 0x9b, 0x38, 0x38,
	0xdb, 0x1d, 0x17, 0xfb, 0x8c, 0x5d, 0x8a, 0x8b, 0xc8, 0x31, 0x60, 0xb3, 0xcc, 0x64, 0x7a, 0x69,
	0x08, 0x05, 0x8c, 0xee, 0xc8, 0x18, 0x9e, 0xa7, 0x58, 0xb6, 0x34, 0xe5, 0x29, 0x1c, 0x4c, 0x28,
	0xa5, 0x76, 0x9b, 0x3f, 0x5e, 0x0b, 0x8b, 0x36, 0x93, 0x1e, 0x15, 0x3d, 0x87, 0xf0, 0x8a, 0xec,
	0x45, 0x9a, 0x72, 0xbe, 0xb9, 0xbc, 0xbf, 0xe1, 0x95, 0x26, 0xa5, 0x7a, 0x63, 0x65, 0xaf, 0xb5,
	0xf2, 0x05, 0x9c, 0x6c, 0x2c, 0x7b, 0xe8, 0x27, 0x61, 0xc0, 0x2b, 0x1a, 0xe1, 0x45, 0x7e, 0x67,
	0xfb, 0x22, 0x24, 0x3f, 0x03, 0x5e, 0x53, 0x9c, 0xcf, 0x16, 0x4e, 0xe2, 0xb1, 0x8e, 0x7c, 0x8f,
	0x85, 0x5e, 0x97, 0xf7, 0x5e, 0xe8, 0x35, 0x86, 0xf0, 0xdf, 0x37, 0xba, 0x5f, 0xeb, 0x3c, 0x31,
	0xc2, 0x77, 0x67, 0xaa, 0x6c, 0xf9, 0x0a, 0x0e, 0x5b, 0xe2, 0xff, 0x30, 0x97, 0x84, 0x7d, 0x36,
	0x6f, 0xa6, 0x1f, 0xca, 0x99, 0x1e, 0x50, 0xf3, 0x2a, 0xd4, 0xe4, 0x12, 0x76, 0x5b, 0xef, 0xb7,
	0x93, 0x51, 0xc3, 0xd8, 0xdb, 0x02, 0xa3, 0xbf, 0x05, 0xc6, 0x7e, 0x03, 0xc6, 0xb3, 0x9f, 0x3e,
	0x04, 0xdc, 0xf0, 0x9a, 0xf2, 0x1f, 0x6a, 0x46, 0xf8, 0x16, 0xa0, 0x46, 0x07, 0x9f, 0x95, 0x6b,
	0x74, 0x28, 0x0c, 0xc3, 0x4d, 0xa1, 0x62, 0x68, 0xb9, 0xc3, 0x32, 0x35, 0x2c, 0xb5, 0x4c, 0x87,
	0xb2, 0x5a, 0xa6, 0xcb, 0x96, 0xdc, 0xc1, 0xaf, 0x70, 0xb8, 0x01, 0x06, 0x94, 0x65, 0xd1, 0xdf,
	0x01, 0x0b, 0x4f, 0xb7, 0xe6, 0x54, 0x1d, 0xde, 0x41, 0xd0, 0x38, 0x27, 0x56, 0xe3, 0x74, 0x01,
	0x0a, 0x4f, 0x36, 0xc6, 0x2a, 0xa5, 0xd7, 0xb0, 0x77, 0x45, 0x96, 0xbd, 0x13, 0xb2, 0xb1, 0x4a,
	0xf1, 0xb8, 0xc9, 0x40, 0x7d, 0xf4, 0xf0, 0xa8, 0xc5, 0x46, 0x25, 0x70, 0xb9, 0xf7, 0x29, 0x68,
	0x7c, 0xb2, 0x6e, 0x87, 0xee, 0x73, 0xf5, 0xf2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xcd,
	0xe8, 0xe9, 0xc8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostServiceClient interface {
	// create a new post
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	// delete a post
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	// get all posts by user_id
	GetAllPostsByUserID(ctx context.Context, in *GetAllPostsByUserIDRequest, opts ...grpc.CallOption) (*GetAllPostsByUserIDResponse, error)
	// Search Posts
	SearchPosts(ctx context.Context, in *SearchPostsRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error)
	// scraping by URL
	GetPostDetail(ctx context.Context, in *PostURLRequest, opts ...grpc.CallOption) (*PostResponse, error)
}

type postServiceClient struct {
	cc *grpc.ClientConn
}

func NewPostServiceClient(cc *grpc.ClientConn) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllPostsByUserID(ctx context.Context, in *GetAllPostsByUserIDRequest, opts ...grpc.CallOption) (*GetAllPostsByUserIDResponse, error) {
	out := new(GetAllPostsByUserIDResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/GetAllPostsByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) SearchPosts(ctx context.Context, in *SearchPostsRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error) {
	out := new(SearchPostsResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/SearchPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostDetail(ctx context.Context, in *PostURLRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/GetPostDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
type PostServiceServer interface {
	// create a new post
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	// delete a post
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	// get all posts by user_id
	GetAllPostsByUserID(context.Context, *GetAllPostsByUserIDRequest) (*GetAllPostsByUserIDResponse, error)
	// Search Posts
	SearchPosts(context.Context, *SearchPostsRequest) (*SearchPostsResponse, error)
	// scraping by URL
	GetPostDetail(context.Context, *PostURLRequest) (*PostResponse, error)
}

func RegisterPostServiceServer(s *grpc.Server, srv PostServiceServer) {
	s.RegisterService(&_PostService_serviceDesc, srv)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllPostsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPostsByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllPostsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/GetAllPostsByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllPostsByUserID(ctx, req.(*GetAllPostsByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_SearchPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).SearchPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/SearchPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).SearchPosts(ctx, req.(*SearchPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/GetPostDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostDetail(ctx, req.(*PostURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "postpb.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
		{
			MethodName: "GetAllPostsByUserID",
			Handler:    _PostService_GetAllPostsByUserID_Handler,
		},
		{
			MethodName: "SearchPosts",
			Handler:    _PostService_SearchPosts_Handler,
		},
		{
			MethodName: "GetPostDetail",
			Handler:    _PostService_GetPostDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post/postpb/post.proto",
}
