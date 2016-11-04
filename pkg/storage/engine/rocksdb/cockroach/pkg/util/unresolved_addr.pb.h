// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: cockroach/pkg/util/unresolved_addr.proto

#ifndef PROTOBUF_cockroach_2fpkg_2futil_2funresolved_5faddr_2eproto__INCLUDED
#define PROTOBUF_cockroach_2fpkg_2futil_2funresolved_5faddr_2eproto__INCLUDED

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 3001000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 3001000 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/message_lite.h>
#include <google/protobuf/repeated_field.h>
#include <google/protobuf/extension_set.h>
// @@protoc_insertion_point(includes)

namespace cockroach {
namespace util {

// Internal implementation detail -- do not call these.
void protobuf_AddDesc_cockroach_2fpkg_2futil_2funresolved_5faddr_2eproto();
void protobuf_InitDefaults_cockroach_2fpkg_2futil_2funresolved_5faddr_2eproto();
void protobuf_AssignDesc_cockroach_2fpkg_2futil_2funresolved_5faddr_2eproto();
void protobuf_ShutdownFile_cockroach_2fpkg_2futil_2funresolved_5faddr_2eproto();

class UnresolvedAddr;

// ===================================================================

class UnresolvedAddr : public ::google::protobuf::MessageLite /* @@protoc_insertion_point(class_definition:cockroach.util.UnresolvedAddr) */ {
 public:
  UnresolvedAddr();
  virtual ~UnresolvedAddr();

  UnresolvedAddr(const UnresolvedAddr& from);

  inline UnresolvedAddr& operator=(const UnresolvedAddr& from) {
    CopyFrom(from);
    return *this;
  }

  inline const ::std::string& unknown_fields() const {
    return _unknown_fields_.GetNoArena(
        &::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }

  inline ::std::string* mutable_unknown_fields() {
    return _unknown_fields_.MutableNoArena(
        &::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }

  static const UnresolvedAddr& default_instance();

  static const UnresolvedAddr* internal_default_instance();

  void Swap(UnresolvedAddr* other);

  // implements Message ----------------------------------------------

  inline UnresolvedAddr* New() const { return New(NULL); }

  UnresolvedAddr* New(::google::protobuf::Arena* arena) const;
  void CheckTypeAndMergeFrom(const ::google::protobuf::MessageLite& from);
  void CopyFrom(const UnresolvedAddr& from);
  void MergeFrom(const UnresolvedAddr& from);
  void Clear();
  bool IsInitialized() const;

  size_t ByteSizeLong() const;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input);
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const;
  void DiscardUnknownFields();
  int GetCachedSize() const { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const;
  void InternalSwap(UnresolvedAddr* other);
  void UnsafeMergeFrom(const UnresolvedAddr& from);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return _arena_ptr_;
  }
  inline ::google::protobuf::Arena* MaybeArenaPtr() const {
    return _arena_ptr_;
  }
  public:

  ::std::string GetTypeName() const;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // optional string network_field = 1;
  bool has_network_field() const;
  void clear_network_field();
  static const int kNetworkFieldFieldNumber = 1;
  const ::std::string& network_field() const;
  void set_network_field(const ::std::string& value);
  void set_network_field(const char* value);
  void set_network_field(const char* value, size_t size);
  ::std::string* mutable_network_field();
  ::std::string* release_network_field();
  void set_allocated_network_field(::std::string* network_field);

  // optional string address_field = 2;
  bool has_address_field() const;
  void clear_address_field();
  static const int kAddressFieldFieldNumber = 2;
  const ::std::string& address_field() const;
  void set_address_field(const ::std::string& value);
  void set_address_field(const char* value);
  void set_address_field(const char* value, size_t size);
  ::std::string* mutable_address_field();
  ::std::string* release_address_field();
  void set_allocated_address_field(::std::string* address_field);

  // @@protoc_insertion_point(class_scope:cockroach.util.UnresolvedAddr)
 private:
  inline void set_has_network_field();
  inline void clear_has_network_field();
  inline void set_has_address_field();
  inline void clear_has_address_field();

  ::google::protobuf::internal::ArenaStringPtr _unknown_fields_;
  ::google::protobuf::Arena* _arena_ptr_;

  ::google::protobuf::internal::HasBits<1> _has_bits_;
  mutable int _cached_size_;
  ::google::protobuf::internal::ArenaStringPtr network_field_;
  ::google::protobuf::internal::ArenaStringPtr address_field_;
  friend void  protobuf_InitDefaults_cockroach_2fpkg_2futil_2funresolved_5faddr_2eproto_impl();
  friend void  protobuf_AddDesc_cockroach_2fpkg_2futil_2funresolved_5faddr_2eproto_impl();
  friend void protobuf_AssignDesc_cockroach_2fpkg_2futil_2funresolved_5faddr_2eproto();
  friend void protobuf_ShutdownFile_cockroach_2fpkg_2futil_2funresolved_5faddr_2eproto();

  void InitAsDefaultInstance();
};
extern ::google::protobuf::internal::ExplicitlyConstructed<UnresolvedAddr> UnresolvedAddr_default_instance_;

// ===================================================================


// ===================================================================

#if !PROTOBUF_INLINE_NOT_IN_HEADERS
// UnresolvedAddr

// optional string network_field = 1;
inline bool UnresolvedAddr::has_network_field() const {
  return (_has_bits_[0] & 0x00000001u) != 0;
}
inline void UnresolvedAddr::set_has_network_field() {
  _has_bits_[0] |= 0x00000001u;
}
inline void UnresolvedAddr::clear_has_network_field() {
  _has_bits_[0] &= ~0x00000001u;
}
inline void UnresolvedAddr::clear_network_field() {
  network_field_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  clear_has_network_field();
}
inline const ::std::string& UnresolvedAddr::network_field() const {
  // @@protoc_insertion_point(field_get:cockroach.util.UnresolvedAddr.network_field)
  return network_field_.GetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void UnresolvedAddr::set_network_field(const ::std::string& value) {
  set_has_network_field();
  network_field_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cockroach.util.UnresolvedAddr.network_field)
}
inline void UnresolvedAddr::set_network_field(const char* value) {
  set_has_network_field();
  network_field_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cockroach.util.UnresolvedAddr.network_field)
}
inline void UnresolvedAddr::set_network_field(const char* value, size_t size) {
  set_has_network_field();
  network_field_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cockroach.util.UnresolvedAddr.network_field)
}
inline ::std::string* UnresolvedAddr::mutable_network_field() {
  set_has_network_field();
  // @@protoc_insertion_point(field_mutable:cockroach.util.UnresolvedAddr.network_field)
  return network_field_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* UnresolvedAddr::release_network_field() {
  // @@protoc_insertion_point(field_release:cockroach.util.UnresolvedAddr.network_field)
  clear_has_network_field();
  return network_field_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void UnresolvedAddr::set_allocated_network_field(::std::string* network_field) {
  if (network_field != NULL) {
    set_has_network_field();
  } else {
    clear_has_network_field();
  }
  network_field_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), network_field);
  // @@protoc_insertion_point(field_set_allocated:cockroach.util.UnresolvedAddr.network_field)
}

// optional string address_field = 2;
inline bool UnresolvedAddr::has_address_field() const {
  return (_has_bits_[0] & 0x00000002u) != 0;
}
inline void UnresolvedAddr::set_has_address_field() {
  _has_bits_[0] |= 0x00000002u;
}
inline void UnresolvedAddr::clear_has_address_field() {
  _has_bits_[0] &= ~0x00000002u;
}
inline void UnresolvedAddr::clear_address_field() {
  address_field_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  clear_has_address_field();
}
inline const ::std::string& UnresolvedAddr::address_field() const {
  // @@protoc_insertion_point(field_get:cockroach.util.UnresolvedAddr.address_field)
  return address_field_.GetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void UnresolvedAddr::set_address_field(const ::std::string& value) {
  set_has_address_field();
  address_field_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cockroach.util.UnresolvedAddr.address_field)
}
inline void UnresolvedAddr::set_address_field(const char* value) {
  set_has_address_field();
  address_field_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cockroach.util.UnresolvedAddr.address_field)
}
inline void UnresolvedAddr::set_address_field(const char* value, size_t size) {
  set_has_address_field();
  address_field_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cockroach.util.UnresolvedAddr.address_field)
}
inline ::std::string* UnresolvedAddr::mutable_address_field() {
  set_has_address_field();
  // @@protoc_insertion_point(field_mutable:cockroach.util.UnresolvedAddr.address_field)
  return address_field_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* UnresolvedAddr::release_address_field() {
  // @@protoc_insertion_point(field_release:cockroach.util.UnresolvedAddr.address_field)
  clear_has_address_field();
  return address_field_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void UnresolvedAddr::set_allocated_address_field(::std::string* address_field) {
  if (address_field != NULL) {
    set_has_address_field();
  } else {
    clear_has_address_field();
  }
  address_field_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), address_field);
  // @@protoc_insertion_point(field_set_allocated:cockroach.util.UnresolvedAddr.address_field)
}

inline const UnresolvedAddr* UnresolvedAddr::internal_default_instance() {
  return &UnresolvedAddr_default_instance_.get();
}
#endif  // !PROTOBUF_INLINE_NOT_IN_HEADERS

// @@protoc_insertion_point(namespace_scope)

}  // namespace util
}  // namespace cockroach

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_cockroach_2fpkg_2futil_2funresolved_5faddr_2eproto__INCLUDED
