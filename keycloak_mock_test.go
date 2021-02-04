package saramakeycloak

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/2phost/sarama-keycloak.Keycloak -o ./keycloak_mock_test.go

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/Nerzal/gocloak/v5"
	"github.com/gojuno/minimock/v3"
)

// KeycloakMock implements Keycloak
type KeycloakMock struct {
	t minimock.Tester

	funcLoginClient          func(clientID string, clientSecret string, realm string) (jp1 *gocloak.JWT, err error)
	inspectFuncLoginClient   func(clientID string, clientSecret string, realm string)
	afterLoginClientCounter  uint64
	beforeLoginClientCounter uint64
	LoginClientMock          mKeycloakMockLoginClient

	funcRefreshToken          func(refreshToken string, clientID string, clientSecret string, realm string) (jp1 *gocloak.JWT, err error)
	inspectFuncRefreshToken   func(refreshToken string, clientID string, clientSecret string, realm string)
	afterRefreshTokenCounter  uint64
	beforeRefreshTokenCounter uint64
	RefreshTokenMock          mKeycloakMockRefreshToken
}

// NewKeycloakMock returns a mock for Keycloak
func NewKeycloakMock(t minimock.Tester) *KeycloakMock {
	m := &KeycloakMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.LoginClientMock = mKeycloakMockLoginClient{mock: m}
	m.LoginClientMock.callArgs = []*KeycloakMockLoginClientParams{}

	m.RefreshTokenMock = mKeycloakMockRefreshToken{mock: m}
	m.RefreshTokenMock.callArgs = []*KeycloakMockRefreshTokenParams{}

	return m
}

type mKeycloakMockLoginClient struct {
	mock               *KeycloakMock
	defaultExpectation *KeycloakMockLoginClientExpectation
	expectations       []*KeycloakMockLoginClientExpectation

	callArgs []*KeycloakMockLoginClientParams
	mutex    sync.RWMutex
}

// KeycloakMockLoginClientExpectation specifies expectation struct of the Keycloak.LoginClient
type KeycloakMockLoginClientExpectation struct {
	mock    *KeycloakMock
	params  *KeycloakMockLoginClientParams
	results *KeycloakMockLoginClientResults
	Counter uint64
}

// KeycloakMockLoginClientParams contains parameters of the Keycloak.LoginClient
type KeycloakMockLoginClientParams struct {
	clientID     string
	clientSecret string
	realm        string
}

// KeycloakMockLoginClientResults contains results of the Keycloak.LoginClient
type KeycloakMockLoginClientResults struct {
	jp1 *gocloak.JWT
	err error
}

// Expect sets up expected params for Keycloak.LoginClient
func (mmLoginClient *mKeycloakMockLoginClient) Expect(clientID string, clientSecret string, realm string) *mKeycloakMockLoginClient {
	if mmLoginClient.mock.funcLoginClient != nil {
		mmLoginClient.mock.t.Fatalf("KeycloakMock.LoginClient mock is already set by Set")
	}

	if mmLoginClient.defaultExpectation == nil {
		mmLoginClient.defaultExpectation = &KeycloakMockLoginClientExpectation{}
	}

	mmLoginClient.defaultExpectation.params = &KeycloakMockLoginClientParams{clientID, clientSecret, realm}
	for _, e := range mmLoginClient.expectations {
		if minimock.Equal(e.params, mmLoginClient.defaultExpectation.params) {
			mmLoginClient.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLoginClient.defaultExpectation.params)
		}
	}

	return mmLoginClient
}

// Inspect accepts an inspector function that has same arguments as the Keycloak.LoginClient
func (mmLoginClient *mKeycloakMockLoginClient) Inspect(f func(clientID string, clientSecret string, realm string)) *mKeycloakMockLoginClient {
	if mmLoginClient.mock.inspectFuncLoginClient != nil {
		mmLoginClient.mock.t.Fatalf("Inspect function is already set for KeycloakMock.LoginClient")
	}

	mmLoginClient.mock.inspectFuncLoginClient = f

	return mmLoginClient
}

// Return sets up results that will be returned by Keycloak.LoginClient
func (mmLoginClient *mKeycloakMockLoginClient) Return(jp1 *gocloak.JWT, err error) *KeycloakMock {
	if mmLoginClient.mock.funcLoginClient != nil {
		mmLoginClient.mock.t.Fatalf("KeycloakMock.LoginClient mock is already set by Set")
	}

	if mmLoginClient.defaultExpectation == nil {
		mmLoginClient.defaultExpectation = &KeycloakMockLoginClientExpectation{mock: mmLoginClient.mock}
	}
	mmLoginClient.defaultExpectation.results = &KeycloakMockLoginClientResults{jp1, err}
	return mmLoginClient.mock
}

//Set uses given function f to mock the Keycloak.LoginClient method
func (mmLoginClient *mKeycloakMockLoginClient) Set(f func(clientID string, clientSecret string, realm string) (jp1 *gocloak.JWT, err error)) *KeycloakMock {
	if mmLoginClient.defaultExpectation != nil {
		mmLoginClient.mock.t.Fatalf("Default expectation is already set for the Keycloak.LoginClient method")
	}

	if len(mmLoginClient.expectations) > 0 {
		mmLoginClient.mock.t.Fatalf("Some expectations are already set for the Keycloak.LoginClient method")
	}

	mmLoginClient.mock.funcLoginClient = f
	return mmLoginClient.mock
}

// When sets expectation for the Keycloak.LoginClient which will trigger the result defined by the following
// Then helper
func (mmLoginClient *mKeycloakMockLoginClient) When(clientID string, clientSecret string, realm string) *KeycloakMockLoginClientExpectation {
	if mmLoginClient.mock.funcLoginClient != nil {
		mmLoginClient.mock.t.Fatalf("KeycloakMock.LoginClient mock is already set by Set")
	}

	expectation := &KeycloakMockLoginClientExpectation{
		mock:   mmLoginClient.mock,
		params: &KeycloakMockLoginClientParams{clientID, clientSecret, realm},
	}
	mmLoginClient.expectations = append(mmLoginClient.expectations, expectation)
	return expectation
}

// Then sets up Keycloak.LoginClient return parameters for the expectation previously defined by the When method
func (e *KeycloakMockLoginClientExpectation) Then(jp1 *gocloak.JWT, err error) *KeycloakMock {
	e.results = &KeycloakMockLoginClientResults{jp1, err}
	return e.mock
}

// LoginClient implements Keycloak
func (mmLoginClient *KeycloakMock) LoginClient(clientID string, clientSecret string, realm string) (jp1 *gocloak.JWT, err error) {
	mm_atomic.AddUint64(&mmLoginClient.beforeLoginClientCounter, 1)
	defer mm_atomic.AddUint64(&mmLoginClient.afterLoginClientCounter, 1)

	if mmLoginClient.inspectFuncLoginClient != nil {
		mmLoginClient.inspectFuncLoginClient(clientID, clientSecret, realm)
	}

	mm_params := &KeycloakMockLoginClientParams{clientID, clientSecret, realm}

	// Record call args
	mmLoginClient.LoginClientMock.mutex.Lock()
	mmLoginClient.LoginClientMock.callArgs = append(mmLoginClient.LoginClientMock.callArgs, mm_params)
	mmLoginClient.LoginClientMock.mutex.Unlock()

	for _, e := range mmLoginClient.LoginClientMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.jp1, e.results.err
		}
	}

	if mmLoginClient.LoginClientMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLoginClient.LoginClientMock.defaultExpectation.Counter, 1)
		mm_want := mmLoginClient.LoginClientMock.defaultExpectation.params
		mm_got := KeycloakMockLoginClientParams{clientID, clientSecret, realm}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmLoginClient.t.Errorf("KeycloakMock.LoginClient got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmLoginClient.LoginClientMock.defaultExpectation.results
		if mm_results == nil {
			mmLoginClient.t.Fatal("No results are set for the KeycloakMock.LoginClient")
		}
		return (*mm_results).jp1, (*mm_results).err
	}
	if mmLoginClient.funcLoginClient != nil {
		return mmLoginClient.funcLoginClient(clientID, clientSecret, realm)
	}
	mmLoginClient.t.Fatalf("Unexpected call to KeycloakMock.LoginClient. %v %v %v", clientID, clientSecret, realm)
	return
}

// LoginClientAfterCounter returns a count of finished KeycloakMock.LoginClient invocations
func (mmLoginClient *KeycloakMock) LoginClientAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLoginClient.afterLoginClientCounter)
}

// LoginClientBeforeCounter returns a count of KeycloakMock.LoginClient invocations
func (mmLoginClient *KeycloakMock) LoginClientBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLoginClient.beforeLoginClientCounter)
}

// Calls returns a list of arguments used in each call to KeycloakMock.LoginClient.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLoginClient *mKeycloakMockLoginClient) Calls() []*KeycloakMockLoginClientParams {
	mmLoginClient.mutex.RLock()

	argCopy := make([]*KeycloakMockLoginClientParams, len(mmLoginClient.callArgs))
	copy(argCopy, mmLoginClient.callArgs)

	mmLoginClient.mutex.RUnlock()

	return argCopy
}

// MinimockLoginClientDone returns true if the count of the LoginClient invocations corresponds
// the number of defined expectations
func (m *KeycloakMock) MinimockLoginClientDone() bool {
	for _, e := range m.LoginClientMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LoginClientMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLoginClientCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLoginClient != nil && mm_atomic.LoadUint64(&m.afterLoginClientCounter) < 1 {
		return false
	}
	return true
}

// MinimockLoginClientInspect logs each unmet expectation
func (m *KeycloakMock) MinimockLoginClientInspect() {
	for _, e := range m.LoginClientMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to KeycloakMock.LoginClient with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LoginClientMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLoginClientCounter) < 1 {
		if m.LoginClientMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to KeycloakMock.LoginClient")
		} else {
			m.t.Errorf("Expected call to KeycloakMock.LoginClient with params: %#v", *m.LoginClientMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLoginClient != nil && mm_atomic.LoadUint64(&m.afterLoginClientCounter) < 1 {
		m.t.Error("Expected call to KeycloakMock.LoginClient")
	}
}

type mKeycloakMockRefreshToken struct {
	mock               *KeycloakMock
	defaultExpectation *KeycloakMockRefreshTokenExpectation
	expectations       []*KeycloakMockRefreshTokenExpectation

	callArgs []*KeycloakMockRefreshTokenParams
	mutex    sync.RWMutex
}

// KeycloakMockRefreshTokenExpectation specifies expectation struct of the Keycloak.RefreshToken
type KeycloakMockRefreshTokenExpectation struct {
	mock    *KeycloakMock
	params  *KeycloakMockRefreshTokenParams
	results *KeycloakMockRefreshTokenResults
	Counter uint64
}

// KeycloakMockRefreshTokenParams contains parameters of the Keycloak.RefreshToken
type KeycloakMockRefreshTokenParams struct {
	refreshToken string
	clientID     string
	clientSecret string
	realm        string
}

// KeycloakMockRefreshTokenResults contains results of the Keycloak.RefreshToken
type KeycloakMockRefreshTokenResults struct {
	jp1 *gocloak.JWT
	err error
}

// Expect sets up expected params for Keycloak.RefreshToken
func (mmRefreshToken *mKeycloakMockRefreshToken) Expect(refreshToken string, clientID string, clientSecret string, realm string) *mKeycloakMockRefreshToken {
	if mmRefreshToken.mock.funcRefreshToken != nil {
		mmRefreshToken.mock.t.Fatalf("KeycloakMock.RefreshToken mock is already set by Set")
	}

	if mmRefreshToken.defaultExpectation == nil {
		mmRefreshToken.defaultExpectation = &KeycloakMockRefreshTokenExpectation{}
	}

	mmRefreshToken.defaultExpectation.params = &KeycloakMockRefreshTokenParams{refreshToken, clientID, clientSecret, realm}
	for _, e := range mmRefreshToken.expectations {
		if minimock.Equal(e.params, mmRefreshToken.defaultExpectation.params) {
			mmRefreshToken.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRefreshToken.defaultExpectation.params)
		}
	}

	return mmRefreshToken
}

// Inspect accepts an inspector function that has same arguments as the Keycloak.RefreshToken
func (mmRefreshToken *mKeycloakMockRefreshToken) Inspect(f func(refreshToken string, clientID string, clientSecret string, realm string)) *mKeycloakMockRefreshToken {
	if mmRefreshToken.mock.inspectFuncRefreshToken != nil {
		mmRefreshToken.mock.t.Fatalf("Inspect function is already set for KeycloakMock.RefreshToken")
	}

	mmRefreshToken.mock.inspectFuncRefreshToken = f

	return mmRefreshToken
}

// Return sets up results that will be returned by Keycloak.RefreshToken
func (mmRefreshToken *mKeycloakMockRefreshToken) Return(jp1 *gocloak.JWT, err error) *KeycloakMock {
	if mmRefreshToken.mock.funcRefreshToken != nil {
		mmRefreshToken.mock.t.Fatalf("KeycloakMock.RefreshToken mock is already set by Set")
	}

	if mmRefreshToken.defaultExpectation == nil {
		mmRefreshToken.defaultExpectation = &KeycloakMockRefreshTokenExpectation{mock: mmRefreshToken.mock}
	}
	mmRefreshToken.defaultExpectation.results = &KeycloakMockRefreshTokenResults{jp1, err}
	return mmRefreshToken.mock
}

//Set uses given function f to mock the Keycloak.RefreshToken method
func (mmRefreshToken *mKeycloakMockRefreshToken) Set(f func(refreshToken string, clientID string, clientSecret string, realm string) (jp1 *gocloak.JWT, err error)) *KeycloakMock {
	if mmRefreshToken.defaultExpectation != nil {
		mmRefreshToken.mock.t.Fatalf("Default expectation is already set for the Keycloak.RefreshToken method")
	}

	if len(mmRefreshToken.expectations) > 0 {
		mmRefreshToken.mock.t.Fatalf("Some expectations are already set for the Keycloak.RefreshToken method")
	}

	mmRefreshToken.mock.funcRefreshToken = f
	return mmRefreshToken.mock
}

// When sets expectation for the Keycloak.RefreshToken which will trigger the result defined by the following
// Then helper
func (mmRefreshToken *mKeycloakMockRefreshToken) When(refreshToken string, clientID string, clientSecret string, realm string) *KeycloakMockRefreshTokenExpectation {
	if mmRefreshToken.mock.funcRefreshToken != nil {
		mmRefreshToken.mock.t.Fatalf("KeycloakMock.RefreshToken mock is already set by Set")
	}

	expectation := &KeycloakMockRefreshTokenExpectation{
		mock:   mmRefreshToken.mock,
		params: &KeycloakMockRefreshTokenParams{refreshToken, clientID, clientSecret, realm},
	}
	mmRefreshToken.expectations = append(mmRefreshToken.expectations, expectation)
	return expectation
}

// Then sets up Keycloak.RefreshToken return parameters for the expectation previously defined by the When method
func (e *KeycloakMockRefreshTokenExpectation) Then(jp1 *gocloak.JWT, err error) *KeycloakMock {
	e.results = &KeycloakMockRefreshTokenResults{jp1, err}
	return e.mock
}

// RefreshToken implements Keycloak
func (mmRefreshToken *KeycloakMock) RefreshToken(refreshToken string, clientID string, clientSecret string, realm string) (jp1 *gocloak.JWT, err error) {
	mm_atomic.AddUint64(&mmRefreshToken.beforeRefreshTokenCounter, 1)
	defer mm_atomic.AddUint64(&mmRefreshToken.afterRefreshTokenCounter, 1)

	if mmRefreshToken.inspectFuncRefreshToken != nil {
		mmRefreshToken.inspectFuncRefreshToken(refreshToken, clientID, clientSecret, realm)
	}

	mm_params := &KeycloakMockRefreshTokenParams{refreshToken, clientID, clientSecret, realm}

	// Record call args
	mmRefreshToken.RefreshTokenMock.mutex.Lock()
	mmRefreshToken.RefreshTokenMock.callArgs = append(mmRefreshToken.RefreshTokenMock.callArgs, mm_params)
	mmRefreshToken.RefreshTokenMock.mutex.Unlock()

	for _, e := range mmRefreshToken.RefreshTokenMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.jp1, e.results.err
		}
	}

	if mmRefreshToken.RefreshTokenMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRefreshToken.RefreshTokenMock.defaultExpectation.Counter, 1)
		mm_want := mmRefreshToken.RefreshTokenMock.defaultExpectation.params
		mm_got := KeycloakMockRefreshTokenParams{refreshToken, clientID, clientSecret, realm}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRefreshToken.t.Errorf("KeycloakMock.RefreshToken got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRefreshToken.RefreshTokenMock.defaultExpectation.results
		if mm_results == nil {
			mmRefreshToken.t.Fatal("No results are set for the KeycloakMock.RefreshToken")
		}
		return (*mm_results).jp1, (*mm_results).err
	}
	if mmRefreshToken.funcRefreshToken != nil {
		return mmRefreshToken.funcRefreshToken(refreshToken, clientID, clientSecret, realm)
	}
	mmRefreshToken.t.Fatalf("Unexpected call to KeycloakMock.RefreshToken. %v %v %v %v", refreshToken, clientID, clientSecret, realm)
	return
}

// RefreshTokenAfterCounter returns a count of finished KeycloakMock.RefreshToken invocations
func (mmRefreshToken *KeycloakMock) RefreshTokenAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRefreshToken.afterRefreshTokenCounter)
}

// RefreshTokenBeforeCounter returns a count of KeycloakMock.RefreshToken invocations
func (mmRefreshToken *KeycloakMock) RefreshTokenBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRefreshToken.beforeRefreshTokenCounter)
}

// Calls returns a list of arguments used in each call to KeycloakMock.RefreshToken.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRefreshToken *mKeycloakMockRefreshToken) Calls() []*KeycloakMockRefreshTokenParams {
	mmRefreshToken.mutex.RLock()

	argCopy := make([]*KeycloakMockRefreshTokenParams, len(mmRefreshToken.callArgs))
	copy(argCopy, mmRefreshToken.callArgs)

	mmRefreshToken.mutex.RUnlock()

	return argCopy
}

// MinimockRefreshTokenDone returns true if the count of the RefreshToken invocations corresponds
// the number of defined expectations
func (m *KeycloakMock) MinimockRefreshTokenDone() bool {
	for _, e := range m.RefreshTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RefreshTokenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRefreshTokenCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRefreshToken != nil && mm_atomic.LoadUint64(&m.afterRefreshTokenCounter) < 1 {
		return false
	}
	return true
}

// MinimockRefreshTokenInspect logs each unmet expectation
func (m *KeycloakMock) MinimockRefreshTokenInspect() {
	for _, e := range m.RefreshTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to KeycloakMock.RefreshToken with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RefreshTokenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRefreshTokenCounter) < 1 {
		if m.RefreshTokenMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to KeycloakMock.RefreshToken")
		} else {
			m.t.Errorf("Expected call to KeycloakMock.RefreshToken with params: %#v", *m.RefreshTokenMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRefreshToken != nil && mm_atomic.LoadUint64(&m.afterRefreshTokenCounter) < 1 {
		m.t.Error("Expected call to KeycloakMock.RefreshToken")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *KeycloakMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockLoginClientInspect()

		m.MinimockRefreshTokenInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *KeycloakMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *KeycloakMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockLoginClientDone() &&
		m.MinimockRefreshTokenDone()
}
