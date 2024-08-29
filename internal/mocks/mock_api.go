// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/favonia/cloudflare-ddns/internal/api (interfaces: Handle)
//
// Generated by this command:
//
//	mockgen -typed -destination=../mocks/mock_api.go -package=mocks . Handle
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	netip "net/netip"
	reflect "reflect"

	api "github.com/favonia/cloudflare-ddns/internal/api"
	domain "github.com/favonia/cloudflare-ddns/internal/domain"
	ipnet "github.com/favonia/cloudflare-ddns/internal/ipnet"
	pp "github.com/favonia/cloudflare-ddns/internal/pp"
	gomock "go.uber.org/mock/gomock"
)

// MockHandle is a mock of Handle interface.
type MockHandle struct {
	ctrl     *gomock.Controller
	recorder *MockHandleMockRecorder
}

// MockHandleMockRecorder is the mock recorder for MockHandle.
type MockHandleMockRecorder struct {
	mock *MockHandle
}

// NewMockHandle creates a new mock instance.
func NewMockHandle(ctrl *gomock.Controller) *MockHandle {
	mock := &MockHandle{ctrl: ctrl}
	mock.recorder = &MockHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandle) EXPECT() *MockHandleMockRecorder {
	return m.recorder
}

// ClearWAFListAsync mocks base method.
func (m *MockHandle) ClearWAFListAsync(arg0 context.Context, arg1 pp.PP, arg2 api.WAFList, arg3 bool) (bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearWAFListAsync", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ClearWAFListAsync indicates an expected call of ClearWAFListAsync.
func (mr *MockHandleMockRecorder) ClearWAFListAsync(arg0, arg1, arg2, arg3 any) *HandleClearWAFListAsyncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearWAFListAsync", reflect.TypeOf((*MockHandle)(nil).ClearWAFListAsync), arg0, arg1, arg2, arg3)
	return &HandleClearWAFListAsyncCall{Call: call}
}

// HandleClearWAFListAsyncCall wrap *gomock.Call
type HandleClearWAFListAsyncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleClearWAFListAsyncCall) Return(arg0, arg1 bool) *HandleClearWAFListAsyncCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleClearWAFListAsyncCall) Do(f func(context.Context, pp.PP, api.WAFList, bool) (bool, bool)) *HandleClearWAFListAsyncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleClearWAFListAsyncCall) DoAndReturn(f func(context.Context, pp.PP, api.WAFList, bool) (bool, bool)) *HandleClearWAFListAsyncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateRecord mocks base method.
func (m *MockHandle) CreateRecord(arg0 context.Context, arg1 pp.PP, arg2 ipnet.Type, arg3 domain.Domain, arg4 netip.Addr, arg5 api.TTL, arg6 bool, arg7 string) (api.ID, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecord", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(api.ID)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CreateRecord indicates an expected call of CreateRecord.
func (mr *MockHandleMockRecorder) CreateRecord(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 any) *HandleCreateRecordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecord", reflect.TypeOf((*MockHandle)(nil).CreateRecord), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	return &HandleCreateRecordCall{Call: call}
}

// HandleCreateRecordCall wrap *gomock.Call
type HandleCreateRecordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleCreateRecordCall) Return(arg0 api.ID, arg1 bool) *HandleCreateRecordCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleCreateRecordCall) Do(f func(context.Context, pp.PP, ipnet.Type, domain.Domain, netip.Addr, api.TTL, bool, string) (api.ID, bool)) *HandleCreateRecordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleCreateRecordCall) DoAndReturn(f func(context.Context, pp.PP, ipnet.Type, domain.Domain, netip.Addr, api.TTL, bool, string) (api.ID, bool)) *HandleCreateRecordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateWAFListItems mocks base method.
func (m *MockHandle) CreateWAFListItems(arg0 context.Context, arg1 pp.PP, arg2 api.WAFList, arg3 []netip.Prefix, arg4 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWAFListItems", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CreateWAFListItems indicates an expected call of CreateWAFListItems.
func (mr *MockHandleMockRecorder) CreateWAFListItems(arg0, arg1, arg2, arg3, arg4 any) *HandleCreateWAFListItemsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWAFListItems", reflect.TypeOf((*MockHandle)(nil).CreateWAFListItems), arg0, arg1, arg2, arg3, arg4)
	return &HandleCreateWAFListItemsCall{Call: call}
}

// HandleCreateWAFListItemsCall wrap *gomock.Call
type HandleCreateWAFListItemsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleCreateWAFListItemsCall) Return(arg0 bool) *HandleCreateWAFListItemsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleCreateWAFListItemsCall) Do(f func(context.Context, pp.PP, api.WAFList, []netip.Prefix, string) bool) *HandleCreateWAFListItemsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleCreateWAFListItemsCall) DoAndReturn(f func(context.Context, pp.PP, api.WAFList, []netip.Prefix, string) bool) *HandleCreateWAFListItemsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteRecord mocks base method.
func (m *MockHandle) DeleteRecord(arg0 context.Context, arg1 pp.PP, arg2 ipnet.Type, arg3 domain.Domain, arg4 api.ID, arg5 bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecord", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeleteRecord indicates an expected call of DeleteRecord.
func (mr *MockHandleMockRecorder) DeleteRecord(arg0, arg1, arg2, arg3, arg4, arg5 any) *HandleDeleteRecordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecord", reflect.TypeOf((*MockHandle)(nil).DeleteRecord), arg0, arg1, arg2, arg3, arg4, arg5)
	return &HandleDeleteRecordCall{Call: call}
}

// HandleDeleteRecordCall wrap *gomock.Call
type HandleDeleteRecordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleDeleteRecordCall) Return(arg0 bool) *HandleDeleteRecordCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleDeleteRecordCall) Do(f func(context.Context, pp.PP, ipnet.Type, domain.Domain, api.ID, bool) bool) *HandleDeleteRecordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleDeleteRecordCall) DoAndReturn(f func(context.Context, pp.PP, ipnet.Type, domain.Domain, api.ID, bool) bool) *HandleDeleteRecordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteWAFListItems mocks base method.
func (m *MockHandle) DeleteWAFListItems(arg0 context.Context, arg1 pp.PP, arg2 api.WAFList, arg3 []api.ID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWAFListItems", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeleteWAFListItems indicates an expected call of DeleteWAFListItems.
func (mr *MockHandleMockRecorder) DeleteWAFListItems(arg0, arg1, arg2, arg3 any) *HandleDeleteWAFListItemsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWAFListItems", reflect.TypeOf((*MockHandle)(nil).DeleteWAFListItems), arg0, arg1, arg2, arg3)
	return &HandleDeleteWAFListItemsCall{Call: call}
}

// HandleDeleteWAFListItemsCall wrap *gomock.Call
type HandleDeleteWAFListItemsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleDeleteWAFListItemsCall) Return(arg0 bool) *HandleDeleteWAFListItemsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleDeleteWAFListItemsCall) Do(f func(context.Context, pp.PP, api.WAFList, []api.ID) bool) *HandleDeleteWAFListItemsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleDeleteWAFListItemsCall) DoAndReturn(f func(context.Context, pp.PP, api.WAFList, []api.ID) bool) *HandleDeleteWAFListItemsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EnsureWAFList mocks base method.
func (m *MockHandle) EnsureWAFList(arg0 context.Context, arg1 pp.PP, arg2 api.WAFList, arg3 string) (api.ID, bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureWAFList", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(api.ID)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// EnsureWAFList indicates an expected call of EnsureWAFList.
func (mr *MockHandleMockRecorder) EnsureWAFList(arg0, arg1, arg2, arg3 any) *HandleEnsureWAFListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureWAFList", reflect.TypeOf((*MockHandle)(nil).EnsureWAFList), arg0, arg1, arg2, arg3)
	return &HandleEnsureWAFListCall{Call: call}
}

// HandleEnsureWAFListCall wrap *gomock.Call
type HandleEnsureWAFListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleEnsureWAFListCall) Return(arg0 api.ID, arg1, arg2 bool) *HandleEnsureWAFListCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleEnsureWAFListCall) Do(f func(context.Context, pp.PP, api.WAFList, string) (api.ID, bool, bool)) *HandleEnsureWAFListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleEnsureWAFListCall) DoAndReturn(f func(context.Context, pp.PP, api.WAFList, string) (api.ID, bool, bool)) *HandleEnsureWAFListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListRecords mocks base method.
func (m *MockHandle) ListRecords(arg0 context.Context, arg1 pp.PP, arg2 ipnet.Type, arg3 domain.Domain) ([]api.Record, bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRecords", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]api.Record)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// ListRecords indicates an expected call of ListRecords.
func (mr *MockHandleMockRecorder) ListRecords(arg0, arg1, arg2, arg3 any) *HandleListRecordsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecords", reflect.TypeOf((*MockHandle)(nil).ListRecords), arg0, arg1, arg2, arg3)
	return &HandleListRecordsCall{Call: call}
}

// HandleListRecordsCall wrap *gomock.Call
type HandleListRecordsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleListRecordsCall) Return(arg0 []api.Record, arg1, arg2 bool) *HandleListRecordsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleListRecordsCall) Do(f func(context.Context, pp.PP, ipnet.Type, domain.Domain) ([]api.Record, bool, bool)) *HandleListRecordsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleListRecordsCall) DoAndReturn(f func(context.Context, pp.PP, ipnet.Type, domain.Domain) ([]api.Record, bool, bool)) *HandleListRecordsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListWAFListItems mocks base method.
func (m *MockHandle) ListWAFListItems(arg0 context.Context, arg1 pp.PP, arg2 api.WAFList) ([]api.WAFListItem, bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWAFListItems", arg0, arg1, arg2)
	ret0, _ := ret[0].([]api.WAFListItem)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// ListWAFListItems indicates an expected call of ListWAFListItems.
func (mr *MockHandleMockRecorder) ListWAFListItems(arg0, arg1, arg2 any) *HandleListWAFListItemsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWAFListItems", reflect.TypeOf((*MockHandle)(nil).ListWAFListItems), arg0, arg1, arg2)
	return &HandleListWAFListItemsCall{Call: call}
}

// HandleListWAFListItemsCall wrap *gomock.Call
type HandleListWAFListItemsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleListWAFListItemsCall) Return(arg0 []api.WAFListItem, arg1, arg2 bool) *HandleListWAFListItemsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleListWAFListItemsCall) Do(f func(context.Context, pp.PP, api.WAFList) ([]api.WAFListItem, bool, bool)) *HandleListWAFListItemsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleListWAFListItemsCall) DoAndReturn(f func(context.Context, pp.PP, api.WAFList) ([]api.WAFListItem, bool, bool)) *HandleListWAFListItemsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateRecord mocks base method.
func (m *MockHandle) UpdateRecord(arg0 context.Context, arg1 pp.PP, arg2 ipnet.Type, arg3 domain.Domain, arg4 api.ID, arg5 netip.Addr) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecord", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateRecord indicates an expected call of UpdateRecord.
func (mr *MockHandleMockRecorder) UpdateRecord(arg0, arg1, arg2, arg3, arg4, arg5 any) *HandleUpdateRecordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecord", reflect.TypeOf((*MockHandle)(nil).UpdateRecord), arg0, arg1, arg2, arg3, arg4, arg5)
	return &HandleUpdateRecordCall{Call: call}
}

// HandleUpdateRecordCall wrap *gomock.Call
type HandleUpdateRecordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleUpdateRecordCall) Return(arg0 bool) *HandleUpdateRecordCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleUpdateRecordCall) Do(f func(context.Context, pp.PP, ipnet.Type, domain.Domain, api.ID, netip.Addr) bool) *HandleUpdateRecordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleUpdateRecordCall) DoAndReturn(f func(context.Context, pp.PP, ipnet.Type, domain.Domain, api.ID, netip.Addr) bool) *HandleUpdateRecordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
