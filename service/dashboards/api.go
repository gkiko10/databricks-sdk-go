// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Genie, Lakeview, etc.
package dashboards

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type GenieInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockGenieInterface instead.
	WithImpl(impl GenieService) GenieInterface

	// Impl returns low-level Genie API implementation
	// Deprecated: use MockGenieInterface instead.
	Impl() GenieService

	// WaitGetMessageGenieCompleted repeatedly calls [GenieAPI.GetMessage] and waits to reach COMPLETED state
	WaitGetMessageGenieCompleted(ctx context.Context, conversationId string, messageId string, spaceId string,
		timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error)

	// Create conversation message.
	//
	// Create new message in [conversation](:method:genie/startconversation). The AI
	// response uses all previously created messages in the conversation to respond.
	CreateMessage(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieMessage], error)

	// Calls [GenieAPIInterface.CreateMessage] and waits to reach COMPLETED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[GenieMessage](60*time.Minute) functional option.
	//
	// Deprecated: use [GenieAPIInterface.CreateMessage].Get() or [GenieAPIInterface.WaitGetMessageGenieCompleted]
	CreateMessageAndWait(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error)

	// Execute SQL query in a conversation message.
	//
	// Execute the SQL query in the message.
	ExecuteMessageQuery(ctx context.Context, request ExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get conversation message.
	//
	// Get message from conversation.
	GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error)

	// Get conversation message.
	//
	// Get message from conversation.
	GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieMessage, error)

	// Get conversation message SQL query result.
	//
	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY`.
	GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get conversation message SQL query result.
	//
	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY`.
	GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieGetMessageQueryResultResponse, error)

	// Start conversation.
	//
	// Start a new conversation.
	StartConversation(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieStartConversationResponse], error)

	// Calls [GenieAPIInterface.StartConversation] and waits to reach COMPLETED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[GenieMessage](60*time.Minute) functional option.
	//
	// Deprecated: use [GenieAPIInterface.StartConversation].Get() or [GenieAPIInterface.WaitGetMessageGenieCompleted]
	StartConversationAndWait(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error)
}

func NewGenie(client *client.DatabricksClient) *GenieAPI {
	return &GenieAPI{
		GenieService: &genieImpl{
			client: client,
		},
	}
}

// Genie provides a no-code experience for business users, powered by AI/BI.
// Analysts set up spaces that business users can use to ask questions using
// natural language. Genie uses data registered to Unity Catalog and requires at
// least CAN USE permission on a Pro or Serverless SQL warehouse. Also,
// Databricks Assistant must be enabled.
type GenieAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(GenieService)
	GenieService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockGenieInterface instead.
func (a *GenieAPI) WithImpl(impl GenieService) GenieInterface {
	a.GenieService = impl
	return a
}

// Impl returns low-level Genie API implementation
// Deprecated: use MockGenieInterface instead.
func (a *GenieAPI) Impl() GenieService {
	return a.GenieService
}

// WaitGetMessageGenieCompleted repeatedly calls [GenieAPI.GetMessage] and waits to reach COMPLETED state
func (a *GenieAPI) WaitGetMessageGenieCompleted(ctx context.Context, conversationId string, messageId string, spaceId string,
	timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[GenieMessage](ctx, timeout, func() (*GenieMessage, *retries.Err) {
		genieMessage, err := a.GetMessage(ctx, GenieGetConversationMessageRequest{
			ConversationId: conversationId,
			MessageId:      messageId,
			SpaceId:        spaceId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(genieMessage)
		}
		status := genieMessage.Status
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case MessageStatusCompleted: // target state
			return genieMessage, nil
		case MessageStatusFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				MessageStatusCompleted, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetMessageGenieCompleted is a wrapper that calls [GenieAPI.WaitGetMessageGenieCompleted] and waits to reach COMPLETED state.
type WaitGetMessageGenieCompleted[R any] struct {
	Response       *R
	ConversationId string `json:"conversation_id"`
	MessageId      string `json:"message_id"`
	SpaceId        string `json:"space_id"`
	Poll           func(time.Duration, func(*GenieMessage)) (*GenieMessage, error)
	callback       func(*GenieMessage)
	timeout        time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetMessageGenieCompleted[R]) OnProgress(callback func(*GenieMessage)) *WaitGetMessageGenieCompleted[R] {
	w.callback = callback
	return w
}

// Get the GenieMessage with the default timeout of 20 minutes.
func (w *WaitGetMessageGenieCompleted[R]) Get() (*GenieMessage, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the GenieMessage with custom timeout.
func (w *WaitGetMessageGenieCompleted[R]) GetWithTimeout(timeout time.Duration) (*GenieMessage, error) {
	return w.Poll(timeout, w.callback)
}

// Create conversation message.
//
// Create new message in [conversation](:method:genie/startconversation). The AI
// response uses all previously created messages in the conversation to respond.
func (a *GenieAPI) CreateMessage(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieMessage], error) {
	genieMessage, err := a.GenieService.CreateMessage(ctx, genieCreateConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetMessageGenieCompleted[GenieMessage]{
		Response:       genieMessage,
		ConversationId: genieCreateConversationMessageRequest.ConversationId,
		MessageId:      genieMessage.Id,
		SpaceId:        genieCreateConversationMessageRequest.SpaceId,
		Poll: func(timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error) {
			return a.WaitGetMessageGenieCompleted(ctx, genieCreateConversationMessageRequest.ConversationId, genieMessage.Id, genieCreateConversationMessageRequest.SpaceId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [GenieAPI.CreateMessage] and waits to reach COMPLETED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GenieMessage](60*time.Minute) functional option.
//
// Deprecated: use [GenieAPI.CreateMessage].Get() or [GenieAPI.WaitGetMessageGenieCompleted]
func (a *GenieAPI) CreateMessageAndWait(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error) {
	wait, err := a.CreateMessage(ctx, genieCreateConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[GenieMessage]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *GenieMessage) {
		for _, o := range options {
			o(&retries.Info[GenieMessage]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Get conversation message.
//
// Get message from conversation.
func (a *GenieAPI) GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieMessage, error) {
	return a.GenieService.GetMessage(ctx, GenieGetConversationMessageRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
	})
}

// Get conversation message SQL query result.
//
// Get the result of SQL query if the message has a query attachment. This is
// only available if a message has a query attachment and the message status is
// `EXECUTING_QUERY`.
func (a *GenieAPI) GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.GenieService.GetMessageQueryResult(ctx, GenieGetMessageQueryResultRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
	})
}

// Start conversation.
//
// Start a new conversation.
func (a *GenieAPI) StartConversation(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieStartConversationResponse], error) {
	genieStartConversationResponse, err := a.GenieService.StartConversation(ctx, genieStartConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetMessageGenieCompleted[GenieStartConversationResponse]{
		Response:       genieStartConversationResponse,
		ConversationId: genieStartConversationResponse.ConversationId,
		MessageId:      genieStartConversationResponse.MessageId,
		SpaceId:        genieStartConversationMessageRequest.SpaceId,
		Poll: func(timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error) {
			return a.WaitGetMessageGenieCompleted(ctx, genieStartConversationResponse.ConversationId, genieStartConversationResponse.MessageId, genieStartConversationMessageRequest.SpaceId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [GenieAPI.StartConversation] and waits to reach COMPLETED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GenieMessage](60*time.Minute) functional option.
//
// Deprecated: use [GenieAPI.StartConversation].Get() or [GenieAPI.WaitGetMessageGenieCompleted]
func (a *GenieAPI) StartConversationAndWait(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error) {
	wait, err := a.StartConversation(ctx, genieStartConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[GenieMessage]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *GenieMessage) {
		for _, o := range options {
			o(&retries.Info[GenieMessage]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

type LakeviewInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockLakeviewInterface instead.
	WithImpl(impl LakeviewService) LakeviewInterface

	// Impl returns low-level Lakeview API implementation
	// Deprecated: use MockLakeviewInterface instead.
	Impl() LakeviewService

	// Create dashboard.
	//
	// Create a draft dashboard.
	Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error)

	// Create dashboard schedule.
	CreateSchedule(ctx context.Context, request CreateScheduleRequest) (*Schedule, error)

	// Create schedule subscription.
	CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error)

	// Delete dashboard schedule.
	DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error

	// Delete dashboard schedule.
	DeleteScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) error

	// Delete schedule subscription.
	DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) error

	// Delete schedule subscription.
	DeleteSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) error

	// Get dashboard.
	//
	// Get a draft dashboard.
	Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// Get dashboard.
	//
	// Get a draft dashboard.
	GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error)

	// Get published dashboard.
	//
	// Get the current published dashboard.
	GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error)

	// Get published dashboard.
	//
	// Get the current published dashboard.
	GetPublishedByDashboardId(ctx context.Context, dashboardId string) (*PublishedDashboard, error)

	// Get dashboard schedule.
	GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error)

	// Get dashboard schedule.
	GetScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*Schedule, error)

	// Get schedule subscription.
	GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error)

	// Get schedule subscription.
	GetSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*Subscription, error)

	// List dashboards.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListDashboardsRequest) listing.Iterator[Dashboard]

	// List dashboards.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error)

	// List dashboard schedules.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSchedules(ctx context.Context, request ListSchedulesRequest) listing.Iterator[Schedule]

	// List dashboard schedules.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSchedulesAll(ctx context.Context, request ListSchedulesRequest) ([]Schedule, error)

	// List dashboard schedules.
	ListSchedulesByDashboardId(ctx context.Context, dashboardId string) (*ListSchedulesResponse, error)

	// List schedule subscriptions.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) listing.Iterator[Subscription]

	// List schedule subscriptions.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSubscriptionsAll(ctx context.Context, request ListSubscriptionsRequest) ([]Subscription, error)

	// List schedule subscriptions.
	ListSubscriptionsByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*ListSubscriptionsResponse, error)

	// Migrate dashboard.
	//
	// Migrates a classic SQL dashboard to Lakeview.
	Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error)

	// Publish dashboard.
	//
	// Publish the current draft dashboard.
	Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error)

	// Trash dashboard.
	//
	// Trash a dashboard.
	Trash(ctx context.Context, request TrashDashboardRequest) error

	// Trash dashboard.
	//
	// Trash a dashboard.
	TrashByDashboardId(ctx context.Context, dashboardId string) error

	// Unpublish dashboard.
	//
	// Unpublish the dashboard.
	Unpublish(ctx context.Context, request UnpublishDashboardRequest) error

	// Unpublish dashboard.
	//
	// Unpublish the dashboard.
	UnpublishByDashboardId(ctx context.Context, dashboardId string) error

	// Update dashboard.
	//
	// Update a draft dashboard.
	Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error)

	// Update dashboard schedule.
	UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error)
}

func NewLakeview(client *client.DatabricksClient) *LakeviewAPI {
	return &LakeviewAPI{
		LakeviewService: &lakeviewImpl{
			client: client,
		},
	}
}

// These APIs provide specific management operations for Lakeview dashboards.
// Generic resource management can be done with Workspace API (import, export,
// get-status, list, delete).
type LakeviewAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(LakeviewService)
	LakeviewService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockLakeviewInterface instead.
func (a *LakeviewAPI) WithImpl(impl LakeviewService) LakeviewInterface {
	a.LakeviewService = impl
	return a
}

// Impl returns low-level Lakeview API implementation
// Deprecated: use MockLakeviewInterface instead.
func (a *LakeviewAPI) Impl() LakeviewService {
	return a.LakeviewService
}

// Delete dashboard schedule.
func (a *LakeviewAPI) DeleteScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) error {
	return a.LakeviewService.DeleteSchedule(ctx, DeleteScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Delete schedule subscription.
func (a *LakeviewAPI) DeleteSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) error {
	return a.LakeviewService.DeleteSubscription(ctx, DeleteSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// Get dashboard.
//
// Get a draft dashboard.
func (a *LakeviewAPI) GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.LakeviewService.Get(ctx, GetDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get published dashboard.
//
// Get the current published dashboard.
func (a *LakeviewAPI) GetPublishedByDashboardId(ctx context.Context, dashboardId string) (*PublishedDashboard, error) {
	return a.LakeviewService.GetPublished(ctx, GetPublishedDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get dashboard schedule.
func (a *LakeviewAPI) GetScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*Schedule, error) {
	return a.LakeviewService.GetSchedule(ctx, GetScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Get schedule subscription.
func (a *LakeviewAPI) GetSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*Subscription, error) {
	return a.LakeviewService.GetSubscription(ctx, GetSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// List dashboards.
//
// This method is generated by Databricks SDK Code Generator.
func (a *LakeviewAPI) List(ctx context.Context, request ListDashboardsRequest) listing.Iterator[Dashboard] {

	getNextPage := func(ctx context.Context, req ListDashboardsRequest) (*ListDashboardsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.LakeviewService.List(ctx, req)
	}
	getItems := func(resp *ListDashboardsResponse) []Dashboard {
		return resp.Dashboards
	}
	getNextReq := func(resp *ListDashboardsResponse) *ListDashboardsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List dashboards.
//
// This method is generated by Databricks SDK Code Generator.
func (a *LakeviewAPI) ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Dashboard](ctx, iterator)
}

// List dashboard schedules.
//
// This method is generated by Databricks SDK Code Generator.
func (a *LakeviewAPI) ListSchedules(ctx context.Context, request ListSchedulesRequest) listing.Iterator[Schedule] {

	getNextPage := func(ctx context.Context, req ListSchedulesRequest) (*ListSchedulesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.LakeviewService.ListSchedules(ctx, req)
	}
	getItems := func(resp *ListSchedulesResponse) []Schedule {
		return resp.Schedules
	}
	getNextReq := func(resp *ListSchedulesResponse) *ListSchedulesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List dashboard schedules.
//
// This method is generated by Databricks SDK Code Generator.
func (a *LakeviewAPI) ListSchedulesAll(ctx context.Context, request ListSchedulesRequest) ([]Schedule, error) {
	iterator := a.ListSchedules(ctx, request)
	return listing.ToSlice[Schedule](ctx, iterator)
}

// List dashboard schedules.
func (a *LakeviewAPI) ListSchedulesByDashboardId(ctx context.Context, dashboardId string) (*ListSchedulesResponse, error) {
	return a.LakeviewService.ListSchedules(ctx, ListSchedulesRequest{
		DashboardId: dashboardId,
	})
}

// List schedule subscriptions.
//
// This method is generated by Databricks SDK Code Generator.
func (a *LakeviewAPI) ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) listing.Iterator[Subscription] {

	getNextPage := func(ctx context.Context, req ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.LakeviewService.ListSubscriptions(ctx, req)
	}
	getItems := func(resp *ListSubscriptionsResponse) []Subscription {
		return resp.Subscriptions
	}
	getNextReq := func(resp *ListSubscriptionsResponse) *ListSubscriptionsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List schedule subscriptions.
//
// This method is generated by Databricks SDK Code Generator.
func (a *LakeviewAPI) ListSubscriptionsAll(ctx context.Context, request ListSubscriptionsRequest) ([]Subscription, error) {
	iterator := a.ListSubscriptions(ctx, request)
	return listing.ToSlice[Subscription](ctx, iterator)
}

// List schedule subscriptions.
func (a *LakeviewAPI) ListSubscriptionsByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*ListSubscriptionsResponse, error) {
	return a.LakeviewService.ListSubscriptions(ctx, ListSubscriptionsRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Trash dashboard.
//
// Trash a dashboard.
func (a *LakeviewAPI) TrashByDashboardId(ctx context.Context, dashboardId string) error {
	return a.LakeviewService.Trash(ctx, TrashDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Unpublish dashboard.
//
// Unpublish the dashboard.
func (a *LakeviewAPI) UnpublishByDashboardId(ctx context.Context, dashboardId string) error {
	return a.LakeviewService.Unpublish(ctx, UnpublishDashboardRequest{
		DashboardId: dashboardId,
	})
}
