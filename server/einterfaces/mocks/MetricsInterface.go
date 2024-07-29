// Code generated by mockery v2.42.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	logr "github.com/mattermost/logr/v2"
	mock "github.com/stretchr/testify/mock"

	model "github.com/mattermost/mattermost/server/public/model"

	sql "database/sql"
)

// MetricsInterface is an autogenerated mock type for the MetricsInterface type
type MetricsInterface struct {
	mock.Mock
}

// AddMemCacheHitCounter provides a mock function with given fields: cacheName, amount
func (_m *MetricsInterface) AddMemCacheHitCounter(cacheName string, amount float64) {
	_m.Called(cacheName, amount)
}

// AddMemCacheMissCounter provides a mock function with given fields: cacheName, amount
func (_m *MetricsInterface) AddMemCacheMissCounter(cacheName string, amount float64) {
	_m.Called(cacheName, amount)
}

// DecrementHTTPWebSockets provides a mock function with given fields: originClient
func (_m *MetricsInterface) DecrementHTTPWebSockets(originClient string) {
	_m.Called(originClient)
}

// DecrementJobActive provides a mock function with given fields: jobType
func (_m *MetricsInterface) DecrementJobActive(jobType string) {
	_m.Called(jobType)
}

// DecrementWebSocketBroadcastBufferSize provides a mock function with given fields: hub, amount
func (_m *MetricsInterface) DecrementWebSocketBroadcastBufferSize(hub string, amount float64) {
	_m.Called(hub, amount)
}

// DecrementWebSocketBroadcastUsersRegistered provides a mock function with given fields: hub, amount
func (_m *MetricsInterface) DecrementWebSocketBroadcastUsersRegistered(hub string, amount float64) {
	_m.Called(hub, amount)
}

// GetLoggerMetricsCollector provides a mock function with given fields:
func (_m *MetricsInterface) GetLoggerMetricsCollector() logr.MetricsCollector {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLoggerMetricsCollector")
	}

	var r0 logr.MetricsCollector
	if rf, ok := ret.Get(0).(func() logr.MetricsCollector); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logr.MetricsCollector)
		}
	}

	return r0
}

// IncrementChannelIndexCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementChannelIndexCounter() {
	_m.Called()
}

// IncrementClientLongTasks provides a mock function with given fields: platform, agent, inc
func (_m *MetricsInterface) IncrementClientLongTasks(platform string, agent string, inc float64) {
	_m.Called(platform, agent, inc)
}

// IncrementClusterEventType provides a mock function with given fields: eventType
func (_m *MetricsInterface) IncrementClusterEventType(eventType model.ClusterEvent) {
	_m.Called(eventType)
}

// IncrementClusterRequest provides a mock function with given fields:
func (_m *MetricsInterface) IncrementClusterRequest() {
	_m.Called()
}

// IncrementEtagHitCounter provides a mock function with given fields: route
func (_m *MetricsInterface) IncrementEtagHitCounter(route string) {
	_m.Called(route)
}

// IncrementEtagMissCounter provides a mock function with given fields: route
func (_m *MetricsInterface) IncrementEtagMissCounter(route string) {
	_m.Called(route)
}

// IncrementFileIndexCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementFileIndexCounter() {
	_m.Called()
}

// IncrementFilesSearchCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementFilesSearchCounter() {
	_m.Called()
}

// IncrementHTTPError provides a mock function with given fields:
func (_m *MetricsInterface) IncrementHTTPError() {
	_m.Called()
}

// IncrementHTTPRequest provides a mock function with given fields:
func (_m *MetricsInterface) IncrementHTTPRequest() {
	_m.Called()
}

// IncrementHTTPWebSockets provides a mock function with given fields: originClient
func (_m *MetricsInterface) IncrementHTTPWebSockets(originClient string) {
	_m.Called(originClient)
}

// IncrementJobActive provides a mock function with given fields: jobType
func (_m *MetricsInterface) IncrementJobActive(jobType string) {
	_m.Called(jobType)
}

// IncrementLogin provides a mock function with given fields:
func (_m *MetricsInterface) IncrementLogin() {
	_m.Called()
}

// IncrementLoginFail provides a mock function with given fields:
func (_m *MetricsInterface) IncrementLoginFail() {
	_m.Called()
}

// IncrementMemCacheHitCounter provides a mock function with given fields: cacheName
func (_m *MetricsInterface) IncrementMemCacheHitCounter(cacheName string) {
	_m.Called(cacheName)
}

// IncrementMemCacheHitCounterSession provides a mock function with given fields:
func (_m *MetricsInterface) IncrementMemCacheHitCounterSession() {
	_m.Called()
}

// IncrementMemCacheInvalidationCounter provides a mock function with given fields: cacheName
func (_m *MetricsInterface) IncrementMemCacheInvalidationCounter(cacheName string) {
	_m.Called(cacheName)
}

// IncrementMemCacheInvalidationCounterSession provides a mock function with given fields:
func (_m *MetricsInterface) IncrementMemCacheInvalidationCounterSession() {
	_m.Called()
}

// IncrementMemCacheMissCounter provides a mock function with given fields: cacheName
func (_m *MetricsInterface) IncrementMemCacheMissCounter(cacheName string) {
	_m.Called(cacheName)
}

// IncrementMemCacheMissCounterSession provides a mock function with given fields:
func (_m *MetricsInterface) IncrementMemCacheMissCounterSession() {
	_m.Called()
}

// IncrementNotificationAckCounter provides a mock function with given fields: notificationType, platform
func (_m *MetricsInterface) IncrementNotificationAckCounter(notificationType model.NotificationType, platform string) {
	_m.Called(notificationType, platform)
}

// IncrementNotificationCounter provides a mock function with given fields: notificationType, platform
func (_m *MetricsInterface) IncrementNotificationCounter(notificationType model.NotificationType, platform string) {
	_m.Called(notificationType, platform)
}

// IncrementNotificationErrorCounter provides a mock function with given fields: notificationType, errorReason, platform
func (_m *MetricsInterface) IncrementNotificationErrorCounter(notificationType model.NotificationType, errorReason model.NotificationReason, platform string) {
	_m.Called(notificationType, errorReason, platform)
}

// IncrementNotificationNotSentCounter provides a mock function with given fields: notificationType, notSentReason, platform
func (_m *MetricsInterface) IncrementNotificationNotSentCounter(notificationType model.NotificationType, notSentReason model.NotificationReason, platform string) {
	_m.Called(notificationType, notSentReason, platform)
}

// IncrementNotificationSuccessCounter provides a mock function with given fields: notificationType, platform
func (_m *MetricsInterface) IncrementNotificationSuccessCounter(notificationType model.NotificationType, platform string) {
	_m.Called(notificationType, platform)
}

// IncrementNotificationUnsupportedCounter provides a mock function with given fields: notificationType, notSentReason, platform
func (_m *MetricsInterface) IncrementNotificationUnsupportedCounter(notificationType model.NotificationType, notSentReason model.NotificationReason, platform string) {
	_m.Called(notificationType, notSentReason, platform)
}

// IncrementPostBroadcast provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostBroadcast() {
	_m.Called()
}

// IncrementPostCreate provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostCreate() {
	_m.Called()
}

// IncrementPostFileAttachment provides a mock function with given fields: count
func (_m *MetricsInterface) IncrementPostFileAttachment(count int) {
	_m.Called(count)
}

// IncrementPostIndexCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostIndexCounter() {
	_m.Called()
}

// IncrementPostSentEmail provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostSentEmail() {
	_m.Called()
}

// IncrementPostSentPush provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostSentPush() {
	_m.Called()
}

// IncrementPostsSearchCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostsSearchCounter() {
	_m.Called()
}

// IncrementRemoteClusterConnStateChangeCounter provides a mock function with given fields: remoteID, online
func (_m *MetricsInterface) IncrementRemoteClusterConnStateChangeCounter(remoteID string, online bool) {
	_m.Called(remoteID, online)
}

// IncrementRemoteClusterMsgErrorsCounter provides a mock function with given fields: remoteID, timeout
func (_m *MetricsInterface) IncrementRemoteClusterMsgErrorsCounter(remoteID string, timeout bool) {
	_m.Called(remoteID, timeout)
}

// IncrementRemoteClusterMsgReceivedCounter provides a mock function with given fields: remoteID
func (_m *MetricsInterface) IncrementRemoteClusterMsgReceivedCounter(remoteID string) {
	_m.Called(remoteID)
}

// IncrementRemoteClusterMsgSentCounter provides a mock function with given fields: remoteID
func (_m *MetricsInterface) IncrementRemoteClusterMsgSentCounter(remoteID string) {
	_m.Called(remoteID)
}

// IncrementSharedChannelsSyncCounter provides a mock function with given fields: remoteID
func (_m *MetricsInterface) IncrementSharedChannelsSyncCounter(remoteID string) {
	_m.Called(remoteID)
}

// IncrementUserIndexCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementUserIndexCounter() {
	_m.Called()
}

// IncrementWebSocketBroadcast provides a mock function with given fields: eventType
func (_m *MetricsInterface) IncrementWebSocketBroadcast(eventType model.WebsocketEventType) {
	_m.Called(eventType)
}

// IncrementWebSocketBroadcastBufferSize provides a mock function with given fields: hub, amount
func (_m *MetricsInterface) IncrementWebSocketBroadcastBufferSize(hub string, amount float64) {
	_m.Called(hub, amount)
}

// IncrementWebSocketBroadcastUsersRegistered provides a mock function with given fields: hub, amount
func (_m *MetricsInterface) IncrementWebSocketBroadcastUsersRegistered(hub string, amount float64) {
	_m.Called(hub, amount)
}

// IncrementWebhookPost provides a mock function with given fields:
func (_m *MetricsInterface) IncrementWebhookPost() {
	_m.Called()
}

// IncrementWebsocketEvent provides a mock function with given fields: eventType
func (_m *MetricsInterface) IncrementWebsocketEvent(eventType model.WebsocketEventType) {
	_m.Called(eventType)
}

// IncrementWebsocketReconnectEvent provides a mock function with given fields: eventType
func (_m *MetricsInterface) IncrementWebsocketReconnectEvent(eventType string) {
	_m.Called(eventType)
}

// ObserveAPIEndpointDuration provides a mock function with given fields: endpoint, method, statusCode, originClient, pageLoadContext, elapsed
func (_m *MetricsInterface) ObserveAPIEndpointDuration(endpoint string, method string, statusCode string, originClient string, pageLoadContext string, elapsed float64) {
	_m.Called(endpoint, method, statusCode, originClient, pageLoadContext, elapsed)
}

// ObserveClientChannelSwitchDuration provides a mock function with given fields: platform, agent, elapsed
func (_m *MetricsInterface) ObserveClientChannelSwitchDuration(platform string, agent string, elapsed float64) {
	_m.Called(platform, agent, elapsed)
}

// ObserveClientCumulativeLayoutShift provides a mock function with given fields: platform, agent, elapsed
func (_m *MetricsInterface) ObserveClientCumulativeLayoutShift(platform string, agent string, elapsed float64) {
	_m.Called(platform, agent, elapsed)
}

// ObserveClientFirstContentfulPaint provides a mock function with given fields: platform, agent, elapsed
func (_m *MetricsInterface) ObserveClientFirstContentfulPaint(platform string, agent string, elapsed float64) {
	_m.Called(platform, agent, elapsed)
}

// ObserveClientInteractionToNextPaint provides a mock function with given fields: platform, agent, interaction, elapsed
func (_m *MetricsInterface) ObserveClientInteractionToNextPaint(platform string, agent string, interaction string, elapsed float64) {
	_m.Called(platform, agent, interaction, elapsed)
}

// ObserveClientLargestContentfulPaint provides a mock function with given fields: platform, agent, region, elapsed
func (_m *MetricsInterface) ObserveClientLargestContentfulPaint(platform string, agent string, region string, elapsed float64) {
	_m.Called(platform, agent, region, elapsed)
}

// ObserveClientPageLoadDuration provides a mock function with given fields: platform, agent, elapsed
func (_m *MetricsInterface) ObserveClientPageLoadDuration(platform string, agent string, elapsed float64) {
	_m.Called(platform, agent, elapsed)
}

// ObserveClientRHSLoadDuration provides a mock function with given fields: platform, agent, elapsed
func (_m *MetricsInterface) ObserveClientRHSLoadDuration(platform string, agent string, elapsed float64) {
	_m.Called(platform, agent, elapsed)
}

// ObserveClientTeamSwitchDuration provides a mock function with given fields: platform, agent, elapsed
func (_m *MetricsInterface) ObserveClientTeamSwitchDuration(platform string, agent string, elapsed float64) {
	_m.Called(platform, agent, elapsed)
}

// ObserveClientTimeToFirstByte provides a mock function with given fields: platform, agent, elapsed
func (_m *MetricsInterface) ObserveClientTimeToFirstByte(platform string, agent string, elapsed float64) {
	_m.Called(platform, agent, elapsed)
}

// ObserveClusterRequestDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObserveClusterRequestDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObserveEnabledUsers provides a mock function with given fields: users
func (_m *MetricsInterface) ObserveEnabledUsers(users int64) {
	_m.Called(users)
}

// ObserveFilesSearchDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObserveFilesSearchDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObserveGlobalThreadsLoadDuration provides a mock function with given fields: platform, agent, elapsed
func (_m *MetricsInterface) ObserveGlobalThreadsLoadDuration(platform string, agent string, elapsed float64) {
	_m.Called(platform, agent, elapsed)
}

// ObserveMobileClientChannelSwitchDuration provides a mock function with given fields: platform, elapsed
func (_m *MetricsInterface) ObserveMobileClientChannelSwitchDuration(platform string, elapsed float64) {
	_m.Called(platform, elapsed)
}

// ObserveMobileClientLoadDuration provides a mock function with given fields: platform, elapsed
func (_m *MetricsInterface) ObserveMobileClientLoadDuration(platform string, elapsed float64) {
	_m.Called(platform, elapsed)
}

// ObserveMobileClientTeamSwitchDuration provides a mock function with given fields: platform, elapsed
func (_m *MetricsInterface) ObserveMobileClientTeamSwitchDuration(platform string, elapsed float64) {
	_m.Called(platform, elapsed)
}

// ObservePluginAPIDuration provides a mock function with given fields: pluginID, apiName, success, elapsed
func (_m *MetricsInterface) ObservePluginAPIDuration(pluginID string, apiName string, success bool, elapsed float64) {
	_m.Called(pluginID, apiName, success, elapsed)
}

// ObservePluginHookDuration provides a mock function with given fields: pluginID, hookName, success, elapsed
func (_m *MetricsInterface) ObservePluginHookDuration(pluginID string, hookName string, success bool, elapsed float64) {
	_m.Called(pluginID, hookName, success, elapsed)
}

// ObservePluginMultiHookDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObservePluginMultiHookDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObservePluginMultiHookIterationDuration provides a mock function with given fields: pluginID, elapsed
func (_m *MetricsInterface) ObservePluginMultiHookIterationDuration(pluginID string, elapsed float64) {
	_m.Called(pluginID, elapsed)
}

// ObservePostsSearchDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObservePostsSearchDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObserveRemoteClusterClockSkew provides a mock function with given fields: remoteID, skew
func (_m *MetricsInterface) ObserveRemoteClusterClockSkew(remoteID string, skew float64) {
	_m.Called(remoteID, skew)
}

// ObserveRemoteClusterPingDuration provides a mock function with given fields: remoteID, elapsed
func (_m *MetricsInterface) ObserveRemoteClusterPingDuration(remoteID string, elapsed float64) {
	_m.Called(remoteID, elapsed)
}

// ObserveSharedChannelsQueueSize provides a mock function with given fields: size
func (_m *MetricsInterface) ObserveSharedChannelsQueueSize(size int64) {
	_m.Called(size)
}

// ObserveSharedChannelsSyncCollectionDuration provides a mock function with given fields: remoteID, elapsed
func (_m *MetricsInterface) ObserveSharedChannelsSyncCollectionDuration(remoteID string, elapsed float64) {
	_m.Called(remoteID, elapsed)
}

// ObserveSharedChannelsSyncCollectionStepDuration provides a mock function with given fields: remoteID, step, elapsed
func (_m *MetricsInterface) ObserveSharedChannelsSyncCollectionStepDuration(remoteID string, step string, elapsed float64) {
	_m.Called(remoteID, step, elapsed)
}

// ObserveSharedChannelsSyncSendDuration provides a mock function with given fields: remoteID, elapsed
func (_m *MetricsInterface) ObserveSharedChannelsSyncSendDuration(remoteID string, elapsed float64) {
	_m.Called(remoteID, elapsed)
}

// ObserveSharedChannelsSyncSendStepDuration provides a mock function with given fields: remoteID, step, elapsed
func (_m *MetricsInterface) ObserveSharedChannelsSyncSendStepDuration(remoteID string, step string, elapsed float64) {
	_m.Called(remoteID, step, elapsed)
}

// ObserveSharedChannelsTaskInQueueDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObserveSharedChannelsTaskInQueueDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObserveStoreMethodDuration provides a mock function with given fields: method, success, elapsed
func (_m *MetricsInterface) ObserveStoreMethodDuration(method string, success string, elapsed float64) {
	_m.Called(method, success, elapsed)
}

// Register provides a mock function with given fields:
func (_m *MetricsInterface) Register() {
	_m.Called()
}

// RegisterDBCollector provides a mock function with given fields: db, name
func (_m *MetricsInterface) RegisterDBCollector(db *sql.DB, name string) {
	_m.Called(db, name)
}

// SetReplicaLagAbsolute provides a mock function with given fields: node, value
func (_m *MetricsInterface) SetReplicaLagAbsolute(node string, value float64) {
	_m.Called(node, value)
}

// SetReplicaLagTime provides a mock function with given fields: node, value
func (_m *MetricsInterface) SetReplicaLagTime(node string, value float64) {
	_m.Called(node, value)
}

// UnregisterDBCollector provides a mock function with given fields: db, name
func (_m *MetricsInterface) UnregisterDBCollector(db *sql.DB, name string) {
	_m.Called(db, name)
}

// NewMetricsInterface creates a new instance of MetricsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMetricsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MetricsInterface {
	mock := &MetricsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}