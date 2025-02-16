// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package queues

import (
	"go.temporal.io/server/service/history/tasks"
)

// TODO: Following is a copy of existing metrics scope and name definition
// for using the new MetricsHandler interface. We should probably move them
// to a common place once all the metrics refactoring work is done.

// Metric names
const (
	TaskRequests = "task_requests"

	TaskFailures                    = "task_errors"
	TaskDiscarded                   = "task_errors_discarded"
	TaskSkipped                     = "task_skipped"
	TaskAttempt                     = "task_attempt"
	TaskStandbyRetryCounter         = "task_errors_standby_retry_counter"
	TaskWorkflowBusyCounter         = "task_errors_workflow_busy"
	TaskNotActiveCounter            = "task_errors_not_active_counter"
	TaskLoadLatency                 = "task_latency_load"                     // latency from task generation to task loading (persistence scheduleToStart)
	TaskScheduleLatency             = "task_latency_schedule"                 // latency from task submission to in-memory queue to processing (in-memory scheduleToStart)
	TaskProcessingLatency           = "task_latency_processing"               // latency for processing task one time
	TaskNoUserProcessingLatency     = "task_latency_processing_nouserlatency" // same as TaskProcessingLatency, but excludes workflow lock latency
	TaskLatency                     = "task_latency"                          // task in-memory latency across multiple attempts
	TaskNoUserLatency               = "task_latency_nouserlatency"            // same as TaskLatency, but excludes workflow lock latency
	TaskQueueLatency                = "task_latency_queue"                    // task e2e latency
	TaskNoUserQueueLatency          = "task_latency_queue_nouserlatency"      // same as TaskQueueLatency, but excludes workflow lock latency
	TaskUserLatency                 = "task_latency_userlatency"              // workflow lock latency across multiple attempts
	TaskReschedulerPendingTasks     = "task_rescheduler_pending_tasks"
	TaskThrottledCounter            = "task_throttled_counter"
	TasksDependencyTaskNotCompleted = "task_dependency_task_not_completed"

	TaskBatchCompleteCounter    = "task_batch_complete_counter"
	NewTimerNotifyCounter       = "new_timer_notifications"
	AckLevelUpdateCounter       = "ack_level_update"
	AckLevelUpdateFailedCounter = "ack_level_update_failed"
	PendingTasksCounter         = "pending_tasks"

	QueueScheduleLatency      = "queue_latency_schedule" // latency for scheduling 100 tasks in one task channel
	QueueReaderCountHistogram = "queue_reader_count"
	QueueSliceCountHistogram  = "queue_slice_count"
	QueueActionCounter        = "queue_actions"
)

// Operation tag value for queue processors
const (
	OperationTimerQueueProcessor           = "TimerQueueProcessor"
	OperationTimerActiveQueueProcessor     = "TimerActiveQueueProcessor"
	OperationTimerStandbyQueueProcessor    = "TimerStandbyQueueProcessor"
	OperationTransferQueueProcessor        = "TransferQueueProcessor"
	OperationTransferActiveQueueProcessor  = "TransferActiveQueueProcessor"
	OperationTransferStandbyQueueProcessor = "TransferStandbyQueueProcessor"
	OperationVisibilityQueueProcessor      = "VisibilityQueueProcessor"
)

// Task type tag value for active and standby tasks
const (
	TaskTypeTransferActiveTaskActivity             = "TransferActiveTaskActivity"
	TaskTypeTransferActiveTaskWorkflowTask         = "TransferActiveTaskWorkflowTask"
	TaskTypeTransferActiveTaskCloseExecution       = "TransferActiveTaskCloseExecution"
	TaskTypeTransferActiveTaskCancelExecution      = "TransferActiveTaskCancelExecution"
	TaskTypeTransferActiveTaskSignalExecution      = "TransferActiveTaskSignalExecution"
	TaskTypeTransferActiveTaskStartChildExecution  = "TransferActiveTaskStartChildExecution"
	TaskTypeTransferActiveTaskResetWorkflow        = "TransferActiveTaskResetWorkflow"
	TaskTypeTransferStandbyTaskActivity            = "TransferStandbyTaskActivity"
	TaskTypeTransferStandbyTaskWorkflowTask        = "TransferStandbyTaskWorkflowTask"
	TaskTypeTransferStandbyTaskCloseExecution      = "TransferStandbyTaskCloseExecution"
	TaskTypeTransferStandbyTaskCancelExecution     = "TransferStandbyTaskCancelExecution"
	TaskTypeTransferStandbyTaskSignalExecution     = "TransferStandbyTaskSignalExecution"
	TaskTypeTransferStandbyTaskStartChildExecution = "TransferStandbyTaskStartChildExecution"
	TaskTypeTransferStandbyTaskResetWorkflow       = "TransferStandbyTaskResetWorkflow"
	TaskTypeVisibilityTaskStartExecution           = "VisibilityTaskStartExecution"
	TaskTypeVisibilityTaskUpsertExecution          = "VisibilityTaskUpsertExecution"
	TaskTypeVisibilityTaskCloseExecution           = "VisibilityTaskCloseExecution"
	TaskTypeVisibilityTaskDeleteExecution          = "VisibilityTaskDeleteExecution"
	TaskTypeTimerActiveTaskActivityTimeout         = "TimerActiveTaskActivityTimeout"
	TaskTypeTimerActiveTaskWorkflowTaskTimeout     = "TimerActiveTaskWorkflowTaskTimeout"
	TaskTypeTimerActiveTaskUserTimer               = "TimerActiveTaskUserTimer"
	TaskTypeTimerActiveTaskWorkflowTimeout         = "TimerActiveTaskWorkflowTimeout"
	TaskTypeTimerActiveTaskActivityRetryTimer      = "TimerActiveTaskActivityRetryTimer"
	TaskTypeTimerActiveTaskWorkflowBackoffTimer    = "TimerActiveTaskWorkflowBackoffTimer"
	TaskTypeTimerActiveTaskDeleteHistoryEvent      = "TimerActiveTaskDeleteHistoryEvent"
	TaskTypeTimerStandbyTaskActivityTimeout        = "TimerStandbyTaskActivityTimeout"
	TaskTypeTimerStandbyTaskWorkflowTaskTimeout    = "TimerStandbyTaskWorkflowTaskTimeout"
	TaskTypeTimerStandbyTaskUserTimer              = "TimerStandbyTaskUserTimer"
	TaskTypeTimerStandbyTaskWorkflowTimeout        = "TimerStandbyTaskWorkflowTimeout"
	TaskTypeTimerStandbyTaskActivityRetryTimer     = "TimerStandbyTaskActivityRetryTimer"
	TaskTypeTimerStandbyTaskWorkflowBackoffTimer   = "TimerStandbyTaskWorkflowBackoffTimer"
	TaskTypeTimerStandbyTaskDeleteHistoryEvent     = "TimerStandbyTaskDeleteHistoryEvent"
)

func GetActiveTransferTaskTypeTagValue(
	task tasks.Task,
) string {
	switch task.(type) {
	case *tasks.ActivityTask:
		return TaskTypeTransferActiveTaskActivity
	case *tasks.WorkflowTask:
		return TaskTypeTransferActiveTaskWorkflowTask
	case *tasks.CloseExecutionTask:
		return TaskTypeTransferActiveTaskCloseExecution
	case *tasks.CancelExecutionTask:
		return TaskTypeTransferActiveTaskCancelExecution
	case *tasks.SignalExecutionTask:
		return TaskTypeTransferActiveTaskSignalExecution
	case *tasks.StartChildExecutionTask:
		return TaskTypeTransferActiveTaskStartChildExecution
	case *tasks.ResetWorkflowTask:
		return TaskTypeTransferActiveTaskResetWorkflow
	default:
		return ""
	}
}

func GetStandbyTransferTaskTypeTagValue(
	task tasks.Task,
) string {
	switch task.(type) {
	case *tasks.ActivityTask:
		return TaskTypeTransferStandbyTaskActivity
	case *tasks.WorkflowTask:
		return TaskTypeTransferStandbyTaskWorkflowTask
	case *tasks.CloseExecutionTask:
		return TaskTypeTransferStandbyTaskCloseExecution
	case *tasks.CancelExecutionTask:
		return TaskTypeTransferStandbyTaskCancelExecution
	case *tasks.SignalExecutionTask:
		return TaskTypeTransferStandbyTaskSignalExecution
	case *tasks.StartChildExecutionTask:
		return TaskTypeTransferStandbyTaskStartChildExecution
	case *tasks.ResetWorkflowTask:
		return TaskTypeTransferStandbyTaskResetWorkflow
	default:
		return ""
	}
}

func GetActiveTimerTaskTypeTagValue(
	task tasks.Task,
) string {
	switch task.(type) {
	case *tasks.WorkflowTaskTimeoutTask:
		return TaskTypeTimerActiveTaskWorkflowTaskTimeout
	case *tasks.ActivityTimeoutTask:
		return TaskTypeTimerActiveTaskActivityTimeout
	case *tasks.UserTimerTask:
		return TaskTypeTimerActiveTaskUserTimer
	case *tasks.WorkflowTimeoutTask:
		return TaskTypeTimerActiveTaskWorkflowTimeout
	case *tasks.DeleteHistoryEventTask:
		return TaskTypeTimerActiveTaskDeleteHistoryEvent
	case *tasks.ActivityRetryTimerTask:
		return TaskTypeTimerActiveTaskActivityRetryTimer
	case *tasks.WorkflowBackoffTimerTask:
		return TaskTypeTimerActiveTaskWorkflowBackoffTimer
	default:
		return ""
	}
}

func GetStandbyTimerTaskTypeTagValue(
	task tasks.Task,
) string {
	switch task.(type) {
	case *tasks.WorkflowTaskTimeoutTask:
		return TaskTypeTimerStandbyTaskWorkflowTaskTimeout
	case *tasks.ActivityTimeoutTask:
		return TaskTypeTimerStandbyTaskActivityTimeout
	case *tasks.UserTimerTask:
		return TaskTypeTimerStandbyTaskUserTimer
	case *tasks.WorkflowTimeoutTask:
		return TaskTypeTimerStandbyTaskWorkflowTimeout
	case *tasks.DeleteHistoryEventTask:
		return TaskTypeTimerStandbyTaskDeleteHistoryEvent
	case *tasks.ActivityRetryTimerTask:
		return TaskTypeTimerStandbyTaskActivityRetryTimer
	case *tasks.WorkflowBackoffTimerTask:
		return TaskTypeTimerStandbyTaskWorkflowBackoffTimer
	default:
		return ""
	}
}

func GetVisibilityTaskTypeTagValue(
	task tasks.Task,
) string {
	switch task.(type) {
	case *tasks.StartExecutionVisibilityTask:
		return TaskTypeVisibilityTaskStartExecution
	case *tasks.UpsertExecutionVisibilityTask:
		return TaskTypeVisibilityTaskUpsertExecution
	case *tasks.CloseExecutionVisibilityTask:
		return TaskTypeVisibilityTaskCloseExecution
	case *tasks.DeleteExecutionVisibilityTask:
		return TaskTypeVisibilityTaskDeleteExecution
	default:
		return ""
	}
}
