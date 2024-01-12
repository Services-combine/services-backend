// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	domain "github.com/b0shka/services/internal/domain"
	auth "github.com/b0shka/services/internal/domain/auth"
	folders "github.com/b0shka/services/internal/domain/folders"
	settings "github.com/b0shka/services/internal/domain/settings"
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuth) Login(ctx *gin.Context, inp auth.LoginInput) (auth.LoginOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, inp)
	ret0, _ := ret[0].(auth.LoginOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthMockRecorder) Login(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuth)(nil).Login), ctx, inp)
}

// RefreshToken mocks base method.
func (m *MockAuth) RefreshToken(ctx context.Context, inp auth.RefreshTokenInput) (auth.RefreshTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", ctx, inp)
	ret0, _ := ret[0].(auth.RefreshTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockAuthMockRecorder) RefreshToken(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockAuth)(nil).RefreshToken), ctx, inp)
}

// MockSettings is a mock of Settings interface.
type MockSettings struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsMockRecorder
}

// MockSettingsMockRecorder is the mock recorder for MockSettings.
type MockSettingsMockRecorder struct {
	mock *MockSettings
}

// NewMockSettings creates a new mock instance.
func NewMockSettings(ctrl *gomock.Controller) *MockSettings {
	mock := &MockSettings{ctrl: ctrl}
	mock.recorder = &MockSettingsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettings) EXPECT() *MockSettingsMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSettings) Get(ctx context.Context) (settings.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(settings.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSettingsMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSettings)(nil).Get), ctx)
}

// Save mocks base method.
func (m *MockSettings) Save(ctx context.Context, inp settings.SaveSettingsInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockSettingsMockRecorder) Save(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockSettings)(nil).Save), ctx, inp)
}

// MockFolders is a mock of Folders interface.
type MockFolders struct {
	ctrl     *gomock.Controller
	recorder *MockFoldersMockRecorder
}

// MockFoldersMockRecorder is the mock recorder for MockFolders.
type MockFoldersMockRecorder struct {
	mock *MockFolders
}

// NewMockFolders creates a new mock instance.
func NewMockFolders(ctrl *gomock.Controller) *MockFolders {
	mock := &MockFolders{ctrl: ctrl}
	mock.recorder = &MockFoldersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFolders) EXPECT() *MockFoldersMockRecorder {
	return m.recorder
}

// ChangeChat mocks base method.
func (m *MockFolders) ChangeChat(ctx context.Context, folderID primitive.ObjectID, chat string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeChat", ctx, folderID, chat)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeChat indicates an expected call of ChangeChat.
func (mr *MockFoldersMockRecorder) ChangeChat(ctx, folderID, chat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeChat", reflect.TypeOf((*MockFolders)(nil).ChangeChat), ctx, folderID, chat)
}

// ChangeGroups mocks base method.
func (m *MockFolders) ChangeGroups(ctx context.Context, folderID primitive.ObjectID, groups []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeGroups", ctx, folderID, groups)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeGroups indicates an expected call of ChangeGroups.
func (mr *MockFoldersMockRecorder) ChangeGroups(ctx, folderID, groups interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeGroups", reflect.TypeOf((*MockFolders)(nil).ChangeGroups), ctx, folderID, groups)
}

// ChangeMessage mocks base method.
func (m *MockFolders) ChangeMessage(ctx context.Context, folderID primitive.ObjectID, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeMessage", ctx, folderID, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeMessage indicates an expected call of ChangeMessage.
func (mr *MockFoldersMockRecorder) ChangeMessage(ctx, folderID, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMessage", reflect.TypeOf((*MockFolders)(nil).ChangeMessage), ctx, folderID, message)
}

// ChangeUsernames mocks base method.
func (m *MockFolders) ChangeUsernames(ctx context.Context, folderID primitive.ObjectID, usernames []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUsernames", ctx, folderID, usernames)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUsernames indicates an expected call of ChangeUsernames.
func (mr *MockFoldersMockRecorder) ChangeUsernames(ctx, folderID, usernames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUsernames", reflect.TypeOf((*MockFolders)(nil).ChangeUsernames), ctx, folderID, usernames)
}

// Create mocks base method.
func (m *MockFolders) Create(ctx context.Context, folder domain.Folder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, folder)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockFoldersMockRecorder) Create(ctx, folder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFolders)(nil).Create), ctx, folder)
}

// Delete mocks base method.
func (m *MockFolders) Delete(ctx context.Context, folderID primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, folderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFoldersMockRecorder) Delete(ctx, folderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFolders)(nil).Delete), ctx, folderID)
}

// GetAllDataFolderById mocks base method.
func (m *MockFolders) GetAllDataFolderById(ctx context.Context, folderID primitive.ObjectID) (folders.GetFolderOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDataFolderById", ctx, folderID)
	ret0, _ := ret[0].(folders.GetFolderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDataFolderById indicates an expected call of GetAllDataFolderById.
func (mr *MockFoldersMockRecorder) GetAllDataFolderById(ctx, folderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDataFolderById", reflect.TypeOf((*MockFolders)(nil).GetAllDataFolderById), ctx, folderID)
}

// GetFolderById mocks base method.
func (m *MockFolders) GetFolderById(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFolderById", ctx, folderID)
	ret0, _ := ret[0].(domain.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFolderById indicates an expected call of GetFolderById.
func (mr *MockFoldersMockRecorder) GetFolderById(ctx, folderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFolderById", reflect.TypeOf((*MockFolders)(nil).GetFolderById), ctx, folderID)
}

// GetFolders mocks base method.
func (m *MockFolders) GetFolders(ctx context.Context) (folders.GetFoldersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFolders", ctx)
	ret0, _ := ret[0].(folders.GetFoldersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFolders indicates an expected call of GetFolders.
func (mr *MockFoldersMockRecorder) GetFolders(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFolders", reflect.TypeOf((*MockFolders)(nil).GetFolders), ctx)
}

// GetFoldersByPath mocks base method.
func (m *MockFolders) GetFoldersByPath(ctx context.Context, path string) ([]domain.FolderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFoldersByPath", ctx, path)
	ret0, _ := ret[0].([]domain.FolderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFoldersByPath indicates an expected call of GetFoldersByPath.
func (mr *MockFoldersMockRecorder) GetFoldersByPath(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFoldersByPath", reflect.TypeOf((*MockFolders)(nil).GetFoldersByPath), ctx, path)
}

// GetFoldersMove mocks base method.
func (m *MockFolders) GetFoldersMove(ctx context.Context, folderID primitive.ObjectID) ([]domain.AccountDataMove, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFoldersMove", ctx, folderID)
	ret0, _ := ret[0].([]domain.AccountDataMove)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFoldersMove indicates an expected call of GetFoldersMove.
func (mr *MockFoldersMockRecorder) GetFoldersMove(ctx, folderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFoldersMove", reflect.TypeOf((*MockFolders)(nil).GetFoldersMove), ctx, folderID)
}

// LaunchInviting mocks base method.
func (m *MockFolders) LaunchInviting(ctx context.Context, folderID primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LaunchInviting", ctx, folderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// LaunchInviting indicates an expected call of LaunchInviting.
func (mr *MockFoldersMockRecorder) LaunchInviting(ctx, folderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LaunchInviting", reflect.TypeOf((*MockFolders)(nil).LaunchInviting), ctx, folderID)
}

// LaunchMailingGroups mocks base method.
func (m *MockFolders) LaunchMailingGroups(ctx context.Context, folderID primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LaunchMailingGroups", ctx, folderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// LaunchMailingGroups indicates an expected call of LaunchMailingGroups.
func (mr *MockFoldersMockRecorder) LaunchMailingGroups(ctx, folderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LaunchMailingGroups", reflect.TypeOf((*MockFolders)(nil).LaunchMailingGroups), ctx, folderID)
}

// LaunchMailingUsernames mocks base method.
func (m *MockFolders) LaunchMailingUsernames(ctx context.Context, folderID primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LaunchMailingUsernames", ctx, folderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// LaunchMailingUsernames indicates an expected call of LaunchMailingUsernames.
func (mr *MockFoldersMockRecorder) LaunchMailingUsernames(ctx, folderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LaunchMailingUsernames", reflect.TypeOf((*MockFolders)(nil).LaunchMailingUsernames), ctx, folderID)
}

// Move mocks base method.
func (m *MockFolders) Move(ctx context.Context, folderID primitive.ObjectID, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", ctx, folderID, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Move indicates an expected call of Move.
func (mr *MockFoldersMockRecorder) Move(ctx, folderID, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockFolders)(nil).Move), ctx, folderID, path)
}

// Rename mocks base method.
func (m *MockFolders) Rename(ctx context.Context, folderID primitive.ObjectID, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", ctx, folderID, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rename indicates an expected call of Rename.
func (mr *MockFoldersMockRecorder) Rename(ctx, folderID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockFolders)(nil).Rename), ctx, folderID, name)
}

// MockAccounts is a mock of Accounts interface.
type MockAccounts struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsMockRecorder
}

// MockAccountsMockRecorder is the mock recorder for MockAccounts.
type MockAccountsMockRecorder struct {
	mock *MockAccounts
}

// NewMockAccounts creates a new mock instance.
func NewMockAccounts(ctrl *gomock.Controller) *MockAccounts {
	mock := &MockAccounts{ctrl: ctrl}
	mock.recorder = &MockAccountsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccounts) EXPECT() *MockAccountsMockRecorder {
	return m.recorder
}

// CheckBlock mocks base method.
func (m *MockAccounts) CheckBlock(ctx context.Context, folderID primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBlock", ctx, folderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckBlock indicates an expected call of CheckBlock.
func (mr *MockAccountsMockRecorder) CheckBlock(ctx, folderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBlock", reflect.TypeOf((*MockAccounts)(nil).CheckBlock), ctx, folderID)
}

// CheckingUniqueness mocks base method.
func (m *MockAccounts) CheckingUniqueness(ctx context.Context, phone string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckingUniqueness", ctx, phone)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckingUniqueness indicates an expected call of CheckingUniqueness.
func (mr *MockAccountsMockRecorder) CheckingUniqueness(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckingUniqueness", reflect.TypeOf((*MockAccounts)(nil).CheckingUniqueness), ctx, phone)
}

// Create mocks base method.
func (m *MockAccounts) Create(ctx context.Context, accountCreate domain.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, accountCreate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAccountsMockRecorder) Create(ctx, accountCreate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccounts)(nil).Create), ctx, accountCreate)
}

// Delete mocks base method.
func (m *MockAccounts) Delete(ctx context.Context, accountID primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAccountsMockRecorder) Delete(ctx, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccounts)(nil).Delete), ctx, accountID)
}

// GenerateInterval mocks base method.
func (m *MockAccounts) GenerateInterval(ctx context.Context, folderID primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateInterval", ctx, folderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateInterval indicates an expected call of GenerateInterval.
func (mr *MockAccountsMockRecorder) GenerateInterval(ctx, folderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateInterval", reflect.TypeOf((*MockAccounts)(nil).GenerateInterval), ctx, folderID)
}

// JoinGroup mocks base method.
func (m *MockAccounts) JoinGroup(ctx context.Context, folderID primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JoinGroup", ctx, folderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// JoinGroup indicates an expected call of JoinGroup.
func (mr *MockAccountsMockRecorder) JoinGroup(ctx, folderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JoinGroup", reflect.TypeOf((*MockAccounts)(nil).JoinGroup), ctx, folderID)
}

// Update mocks base method.
func (m *MockAccounts) Update(ctx context.Context, account domain.AccountUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAccountsMockRecorder) Update(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccounts)(nil).Update), ctx, account)
}