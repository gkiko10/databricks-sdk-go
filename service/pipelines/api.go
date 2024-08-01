// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// The Delta Live Tables API allows you to create, edit, delete, start, and view details about pipelines.
package pipelines

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type PipelinesInterface interface {

	// WaitGetPipelineIdle repeatedly calls [PipelinesAPI.Get] and waits to reach IDLE state
	WaitGetPipelineIdle(ctx context.Context, pipelineId string,
		timeout time.Duration, callback func(*GetPipelineResponse)) (*GetPipelineResponse, error)

	// WaitGetPipelineRunning repeatedly calls [PipelinesAPI.Get] and waits to reach RUNNING state
	WaitGetPipelineRunning(ctx context.Context, pipelineId string,
		timeout time.Duration, callback func(*GetPipelineResponse)) (*GetPipelineResponse, error)

	// Create a pipeline.
	//
	// Creates a new data processing pipeline based on the requested configuration.
	// If successful, this method returns the ID of the new pipeline.
	Create(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error)

	// Delete a pipeline.
	//
	// Deletes a pipeline.
	Delete(ctx context.Context, request DeletePipelineRequest) error

	// Delete a pipeline.
	//
	// Deletes a pipeline.
	DeleteByPipelineId(ctx context.Context, pipelineId string) error

	// Get a pipeline.
	Get(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error)

	// Get a pipeline.
	GetByPipelineId(ctx context.Context, pipelineId string) (*GetPipelineResponse, error)

	// Get pipeline permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetPipelinePermissionLevelsRequest) (*GetPipelinePermissionLevelsResponse, error)

	// Get pipeline permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevelsByPipelineId(ctx context.Context, pipelineId string) (*GetPipelinePermissionLevelsResponse, error)

	// Get pipeline permissions.
	//
	// Gets the permissions of a pipeline. Pipelines can inherit permissions from
	// their root object.
	GetPermissions(ctx context.Context, request GetPipelinePermissionsRequest) (*PipelinePermissions, error)

	// Get pipeline permissions.
	//
	// Gets the permissions of a pipeline. Pipelines can inherit permissions from
	// their root object.
	GetPermissionsByPipelineId(ctx context.Context, pipelineId string) (*PipelinePermissions, error)

	// Get a pipeline update.
	//
	// Gets an update from an active pipeline.
	GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error)

	// Get a pipeline update.
	//
	// Gets an update from an active pipeline.
	GetUpdateByPipelineIdAndUpdateId(ctx context.Context, pipelineId string, updateId string) (*GetUpdateResponse, error)

	// List pipeline events.
	//
	// Retrieves events for a pipeline.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListPipelineEvents(ctx context.Context, request ListPipelineEventsRequest) listing.Iterator[PipelineEvent]

	// List pipeline events.
	//
	// Retrieves events for a pipeline.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListPipelineEventsAll(ctx context.Context, request ListPipelineEventsRequest) ([]PipelineEvent, error)

	// List pipeline events.
	//
	// Retrieves events for a pipeline.
	ListPipelineEventsByPipelineId(ctx context.Context, pipelineId string) (*ListPipelineEventsResponse, error)

	// List pipelines.
	//
	// Lists pipelines defined in the Delta Live Tables system.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListPipelines(ctx context.Context, request ListPipelinesRequest) listing.Iterator[PipelineStateInfo]

	// List pipelines.
	//
	// Lists pipelines defined in the Delta Live Tables system.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListPipelinesAll(ctx context.Context, request ListPipelinesRequest) ([]PipelineStateInfo, error)

	// PipelineStateInfoNameToPipelineIdMap calls [PipelinesAPI.ListPipelinesAll] and creates a map of results with [PipelineStateInfo].Name as key and [PipelineStateInfo].PipelineId as value.
	//
	// Returns an error if there's more than one [PipelineStateInfo] with the same .Name.
	//
	// Note: All [PipelineStateInfo] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	PipelineStateInfoNameToPipelineIdMap(ctx context.Context, request ListPipelinesRequest) (map[string]string, error)

	// GetByName calls [PipelinesAPI.PipelineStateInfoNameToPipelineIdMap] and returns a single [PipelineStateInfo].
	//
	// Returns an error if there's more than one [PipelineStateInfo] with the same .Name.
	//
	// Note: All [PipelineStateInfo] instances are loaded into memory before returning matching by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByName(ctx context.Context, name string) (*PipelineStateInfo, error)

	// List pipeline updates.
	//
	// List updates for an active pipeline.
	ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error)

	// List pipeline updates.
	//
	// List updates for an active pipeline.
	ListUpdatesByPipelineId(ctx context.Context, pipelineId string) (*ListUpdatesResponse, error)

	// Set pipeline permissions.
	//
	// Sets permissions on a pipeline. Pipelines can inherit permissions from their
	// root object.
	SetPermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error)

	// Start a pipeline.
	//
	// Starts a new update for the pipeline. If there is already an active update
	// for the pipeline, the request will fail and the active update will remain
	// running.
	StartUpdate(ctx context.Context, request StartUpdate) (*StartUpdateResponse, error)

	// Stop a pipeline.
	//
	// Stops the pipeline by canceling the active update. If there is no active
	// update for the pipeline, this request is a no-op.
	Stop(ctx context.Context, stopRequest StopRequest) (*WaitGetPipelineIdle[struct{}], error)

	// Calls [PipelinesAPIInterface.Stop] and waits to reach IDLE state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[GetPipelineResponse](60*time.Minute) functional option.
	//
	// Deprecated: use [PipelinesAPIInterface.Stop].Get() or [PipelinesAPIInterface.WaitGetPipelineIdle]
	StopAndWait(ctx context.Context, stopRequest StopRequest, options ...retries.Option[GetPipelineResponse]) (*GetPipelineResponse, error)

	// Edit a pipeline.
	//
	// Updates a pipeline with the supplied configuration.
	Update(ctx context.Context, request EditPipeline) error

	// Update pipeline permissions.
	//
	// Updates the permissions on a pipeline. Pipelines can inherit permissions from
	// their root object.
	UpdatePermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error)
}

func NewPipelines(client *client.DatabricksClient) *PipelinesAPI {
	return &PipelinesAPI{
		pipelinesImpl: pipelinesImpl{
			client: client,
		},
	}
}

// The Delta Live Tables API allows you to create, edit, delete, start, and view
// details about pipelines.
//
// Delta Live Tables is a framework for building reliable, maintainable, and
// testable data processing pipelines. You define the transformations to perform
// on your data, and Delta Live Tables manages task orchestration, cluster
// management, monitoring, data quality, and error handling.
//
// Instead of defining your data pipelines using a series of separate Apache
// Spark tasks, Delta Live Tables manages how your data is transformed based on
// a target schema you define for each processing step. You can also enforce
// data quality with Delta Live Tables expectations. Expectations allow you to
// define expected data quality and specify how to handle records that fail
// those expectations.
type PipelinesAPI struct {
	pipelinesImpl
}

// WaitGetPipelineIdle repeatedly calls [PipelinesAPI.Get] and waits to reach IDLE state
func (a *PipelinesAPI) WaitGetPipelineIdle(ctx context.Context, pipelineId string,
	timeout time.Duration, callback func(*GetPipelineResponse)) (*GetPipelineResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[GetPipelineResponse](ctx, timeout, func() (*GetPipelineResponse, *retries.Err) {
		getPipelineResponse, err := a.Get(ctx, GetPipelineRequest{
			PipelineId: pipelineId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(getPipelineResponse)
		}
		status := getPipelineResponse.State
		statusMessage := getPipelineResponse.Cause
		switch status {
		case PipelineStateIdle: // target state
			return getPipelineResponse, nil
		case PipelineStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				PipelineStateIdle, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetPipelineIdle is a wrapper that calls [PipelinesAPI.WaitGetPipelineIdle] and waits to reach IDLE state.
type WaitGetPipelineIdle[R any] struct {
	Response   *R
	PipelineId string `json:"pipeline_id"`
	Poll       func(time.Duration, func(*GetPipelineResponse)) (*GetPipelineResponse, error)
	callback   func(*GetPipelineResponse)
	timeout    time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetPipelineIdle[R]) OnProgress(callback func(*GetPipelineResponse)) *WaitGetPipelineIdle[R] {
	w.callback = callback
	return w
}

// Get the GetPipelineResponse with the default timeout of 20 minutes.
func (w *WaitGetPipelineIdle[R]) Get() (*GetPipelineResponse, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the GetPipelineResponse with custom timeout.
func (w *WaitGetPipelineIdle[R]) GetWithTimeout(timeout time.Duration) (*GetPipelineResponse, error) {
	return w.Poll(timeout, w.callback)
}

// WaitGetPipelineRunning repeatedly calls [PipelinesAPI.Get] and waits to reach RUNNING state
func (a *PipelinesAPI) WaitGetPipelineRunning(ctx context.Context, pipelineId string,
	timeout time.Duration, callback func(*GetPipelineResponse)) (*GetPipelineResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[GetPipelineResponse](ctx, timeout, func() (*GetPipelineResponse, *retries.Err) {
		getPipelineResponse, err := a.Get(ctx, GetPipelineRequest{
			PipelineId: pipelineId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(getPipelineResponse)
		}
		status := getPipelineResponse.State
		statusMessage := getPipelineResponse.Cause
		switch status {
		case PipelineStateRunning: // target state
			return getPipelineResponse, nil
		case PipelineStateFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				PipelineStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetPipelineRunning is a wrapper that calls [PipelinesAPI.WaitGetPipelineRunning] and waits to reach RUNNING state.
type WaitGetPipelineRunning[R any] struct {
	Response   *R
	PipelineId string `json:"pipeline_id"`
	Poll       func(time.Duration, func(*GetPipelineResponse)) (*GetPipelineResponse, error)
	callback   func(*GetPipelineResponse)
	timeout    time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetPipelineRunning[R]) OnProgress(callback func(*GetPipelineResponse)) *WaitGetPipelineRunning[R] {
	w.callback = callback
	return w
}

// Get the GetPipelineResponse with the default timeout of 20 minutes.
func (w *WaitGetPipelineRunning[R]) Get() (*GetPipelineResponse, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the GetPipelineResponse with custom timeout.
func (w *WaitGetPipelineRunning[R]) GetWithTimeout(timeout time.Duration) (*GetPipelineResponse, error) {
	return w.Poll(timeout, w.callback)
}

// Delete a pipeline.
//
// Deletes a pipeline.
func (a *PipelinesAPI) DeleteByPipelineId(ctx context.Context, pipelineId string) error {
	return a.pipelinesImpl.Delete(ctx, DeletePipelineRequest{
		PipelineId: pipelineId,
	})
}

// Get a pipeline.
func (a *PipelinesAPI) GetByPipelineId(ctx context.Context, pipelineId string) (*GetPipelineResponse, error) {
	return a.pipelinesImpl.Get(ctx, GetPipelineRequest{
		PipelineId: pipelineId,
	})
}

// Get pipeline permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *PipelinesAPI) GetPermissionLevelsByPipelineId(ctx context.Context, pipelineId string) (*GetPipelinePermissionLevelsResponse, error) {
	return a.pipelinesImpl.GetPermissionLevels(ctx, GetPipelinePermissionLevelsRequest{
		PipelineId: pipelineId,
	})
}

// Get pipeline permissions.
//
// Gets the permissions of a pipeline. Pipelines can inherit permissions from
// their root object.
func (a *PipelinesAPI) GetPermissionsByPipelineId(ctx context.Context, pipelineId string) (*PipelinePermissions, error) {
	return a.pipelinesImpl.GetPermissions(ctx, GetPipelinePermissionsRequest{
		PipelineId: pipelineId,
	})
}

// Get a pipeline update.
//
// Gets an update from an active pipeline.
func (a *PipelinesAPI) GetUpdateByPipelineIdAndUpdateId(ctx context.Context, pipelineId string, updateId string) (*GetUpdateResponse, error) {
	return a.pipelinesImpl.GetUpdate(ctx, GetUpdateRequest{
		PipelineId: pipelineId,
		UpdateId:   updateId,
	})
}

// List pipeline events.
//
// Retrieves events for a pipeline.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) ListPipelineEvents(ctx context.Context, request ListPipelineEventsRequest) listing.Iterator[PipelineEvent] {

	getNextPage := func(ctx context.Context, req ListPipelineEventsRequest) (*ListPipelineEventsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.pipelinesImpl.ListPipelineEvents(ctx, req)
	}
	getItems := func(resp *ListPipelineEventsResponse) []PipelineEvent {
		return resp.Events
	}
	getNextReq := func(resp *ListPipelineEventsResponse) *ListPipelineEventsRequest {
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

// List pipeline events.
//
// Retrieves events for a pipeline.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) ListPipelineEventsAll(ctx context.Context, request ListPipelineEventsRequest) ([]PipelineEvent, error) {
	iterator := a.ListPipelineEvents(ctx, request)
	return listing.ToSliceN[PipelineEvent, int](ctx, iterator, request.MaxResults)

}

// List pipeline events.
//
// Retrieves events for a pipeline.
func (a *PipelinesAPI) ListPipelineEventsByPipelineId(ctx context.Context, pipelineId string) (*ListPipelineEventsResponse, error) {
	return a.pipelinesImpl.ListPipelineEvents(ctx, ListPipelineEventsRequest{
		PipelineId: pipelineId,
	})
}

// List pipelines.
//
// Lists pipelines defined in the Delta Live Tables system.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) ListPipelines(ctx context.Context, request ListPipelinesRequest) listing.Iterator[PipelineStateInfo] {

	getNextPage := func(ctx context.Context, req ListPipelinesRequest) (*ListPipelinesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.pipelinesImpl.ListPipelines(ctx, req)
	}
	getItems := func(resp *ListPipelinesResponse) []PipelineStateInfo {
		return resp.Statuses
	}
	getNextReq := func(resp *ListPipelinesResponse) *ListPipelinesRequest {
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

// List pipelines.
//
// Lists pipelines defined in the Delta Live Tables system.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) ListPipelinesAll(ctx context.Context, request ListPipelinesRequest) ([]PipelineStateInfo, error) {
	iterator := a.ListPipelines(ctx, request)
	return listing.ToSliceN[PipelineStateInfo, int](ctx, iterator, request.MaxResults)

}

// PipelineStateInfoNameToPipelineIdMap calls [PipelinesAPI.ListPipelinesAll] and creates a map of results with [PipelineStateInfo].Name as key and [PipelineStateInfo].PipelineId as value.
//
// Returns an error if there's more than one [PipelineStateInfo] with the same .Name.
//
// Note: All [PipelineStateInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) PipelineStateInfoNameToPipelineIdMap(ctx context.Context, request ListPipelinesRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListPipelinesAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.PipelineId
	}
	return mapping, nil
}

// GetByName calls [PipelinesAPI.PipelineStateInfoNameToPipelineIdMap] and returns a single [PipelineStateInfo].
//
// Returns an error if there's more than one [PipelineStateInfo] with the same .Name.
//
// Note: All [PipelineStateInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PipelinesAPI) GetByName(ctx context.Context, name string) (*PipelineStateInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListPipelinesAll(ctx, ListPipelinesRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]PipelineStateInfo{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("PipelineStateInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of PipelineStateInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// List pipeline updates.
//
// List updates for an active pipeline.
func (a *PipelinesAPI) ListUpdatesByPipelineId(ctx context.Context, pipelineId string) (*ListUpdatesResponse, error) {
	return a.pipelinesImpl.ListUpdates(ctx, ListUpdatesRequest{
		PipelineId: pipelineId,
	})
}

// Stop a pipeline.
//
// Stops the pipeline by canceling the active update. If there is no active
// update for the pipeline, this request is a no-op.
func (a *PipelinesAPI) Stop(ctx context.Context, stopRequest StopRequest) (*WaitGetPipelineIdle[struct{}], error) {
	err := a.pipelinesImpl.Stop(ctx, stopRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetPipelineIdle[struct{}]{

		PipelineId: stopRequest.PipelineId,
		Poll: func(timeout time.Duration, callback func(*GetPipelineResponse)) (*GetPipelineResponse, error) {
			return a.WaitGetPipelineIdle(ctx, stopRequest.PipelineId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [PipelinesAPI.Stop] and waits to reach IDLE state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetPipelineResponse](60*time.Minute) functional option.
//
// Deprecated: use [PipelinesAPI.Stop].Get() or [PipelinesAPI.WaitGetPipelineIdle]
func (a *PipelinesAPI) StopAndWait(ctx context.Context, stopRequest StopRequest, options ...retries.Option[GetPipelineResponse]) (*GetPipelineResponse, error) {
	wait, err := a.Stop(ctx, stopRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[GetPipelineResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *GetPipelineResponse) {
		for _, o := range options {
			o(&retries.Info[GetPipelineResponse]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}
