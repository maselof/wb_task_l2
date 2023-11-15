// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package bolt

import (
	"calendar/event"
	"sync"
	"time"
)

// Ensure, that EventRepositoryMock does implement event.EventRepository.
// If this is not the case, regenerate this file with moq.
var _ event.EventRepository = &EventRepositoryMock{}

// EventRepositoryMock is a mock implementation of event.EventRepository.
//
// 	func TestSomethingThatUsesEventRepository(t *testing.T) {
//
// 		// make and configure a mocked event.EventRepository
// 		mockedEventRepository := &EventRepositoryMock{
// 			CreateFunc: func(user_id uint64, e event.Event) (event.Event, error) {
// 				panic("mock out the Create method")
// 			},
// 			DeleteFunc: func(user_id uint64, event_id uint64) error {
// 				panic("mock out the Delete method")
// 			},
// 			GetForDayFunc: func(user_id uint64, day time.Time) ([]event.Event, error) {
// 				panic("mock out the GetForDay method")
// 			},
// 			GetForMonthFunc: func(user_id uint64, month time.Time) ([]event.Event, error) {
// 				panic("mock out the GetForMonth method")
// 			},
// 			GetForWeekFunc: func(user_id uint64, week time.Time) ([]event.Event, error) {
// 				panic("mock out the GetForWeek method")
// 			},
// 			UpdateFunc: func(user_id uint64, e event.Event) error {
// 				panic("mock out the Update method")
// 			},
// 		}
//
// 		// use mockedEventRepository in code that requires event.EventRepository
// 		// and then make assertions.
//
// 	}
type EventRepositoryMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(user_id uint64, e event.Event) (event.Event, error)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(user_id uint64, event_id uint64) error

	// GetForDayFunc mocks the GetForDay method.
	GetForDayFunc func(user_id uint64, day time.Time) ([]event.Event, error)

	// GetForMonthFunc mocks the GetForMonth method.
	GetForMonthFunc func(user_id uint64, month time.Time) ([]event.Event, error)

	// GetForWeekFunc mocks the GetForWeek method.
	GetForWeekFunc func(user_id uint64, week time.Time) ([]event.Event, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(user_id uint64, e event.Event) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// User_id is the user_id argument value.
			User_id uint64
			// E is the e argument value.
			E event.Event
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// User_id is the user_id argument value.
			User_id uint64
			// Event_id is the event_id argument value.
			Event_id uint64
		}
		// GetForDay holds details about calls to the GetForDay method.
		GetForDay []struct {
			// User_id is the user_id argument value.
			User_id uint64
			// Day is the day argument value.
			Day time.Time
		}
		// GetForMonth holds details about calls to the GetForMonth method.
		GetForMonth []struct {
			// User_id is the user_id argument value.
			User_id uint64
			// Month is the month argument value.
			Month time.Time
		}
		// GetForWeek holds details about calls to the GetForWeek method.
		GetForWeek []struct {
			// User_id is the user_id argument value.
			User_id uint64
			// Week is the week argument value.
			Week time.Time
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// User_id is the user_id argument value.
			User_id uint64
			// E is the e argument value.
			E event.Event
		}
	}
	lockCreate      sync.RWMutex
	lockDelete      sync.RWMutex
	lockGetForDay   sync.RWMutex
	lockGetForMonth sync.RWMutex
	lockGetForWeek  sync.RWMutex
	lockUpdate      sync.RWMutex
}

// Create calls CreateFunc.
func (mock *EventRepositoryMock) Create(user_id uint64, e event.Event) (event.Event, error) {
	if mock.CreateFunc == nil {
		panic("EventRepositoryMock.CreateFunc: method is nil but EventRepository.Create was just called")
	}
	callInfo := struct {
		User_id uint64
		E       event.Event
	}{
		User_id: user_id,
		E:       e,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(user_id, e)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedEventRepository.CreateCalls())
func (mock *EventRepositoryMock) CreateCalls() []struct {
	User_id uint64
	E       event.Event
} {
	var calls []struct {
		User_id uint64
		E       event.Event
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *EventRepositoryMock) Delete(user_id uint64, event_id uint64) error {
	if mock.DeleteFunc == nil {
		panic("EventRepositoryMock.DeleteFunc: method is nil but EventRepository.Delete was just called")
	}
	callInfo := struct {
		User_id  uint64
		Event_id uint64
	}{
		User_id:  user_id,
		Event_id: event_id,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(user_id, event_id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedEventRepository.DeleteCalls())
func (mock *EventRepositoryMock) DeleteCalls() []struct {
	User_id  uint64
	Event_id uint64
} {
	var calls []struct {
		User_id  uint64
		Event_id uint64
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// GetForDay calls GetForDayFunc.
func (mock *EventRepositoryMock) GetForDay(user_id uint64, day time.Time) ([]event.Event, error) {
	if mock.GetForDayFunc == nil {
		panic("EventRepositoryMock.GetForDayFunc: method is nil but EventRepository.GetForDay was just called")
	}
	callInfo := struct {
		User_id uint64
		Day     time.Time
	}{
		User_id: user_id,
		Day:     day,
	}
	mock.lockGetForDay.Lock()
	mock.calls.GetForDay = append(mock.calls.GetForDay, callInfo)
	mock.lockGetForDay.Unlock()
	return mock.GetForDayFunc(user_id, day)
}

// GetForDayCalls gets all the calls that were made to GetForDay.
// Check the length with:
//     len(mockedEventRepository.GetForDayCalls())
func (mock *EventRepositoryMock) GetForDayCalls() []struct {
	User_id uint64
	Day     time.Time
} {
	var calls []struct {
		User_id uint64
		Day     time.Time
	}
	mock.lockGetForDay.RLock()
	calls = mock.calls.GetForDay
	mock.lockGetForDay.RUnlock()
	return calls
}

// GetForMonth calls GetForMonthFunc.
func (mock *EventRepositoryMock) GetForMonth(user_id uint64, month time.Time) ([]event.Event, error) {
	if mock.GetForMonthFunc == nil {
		panic("EventRepositoryMock.GetForMonthFunc: method is nil but EventRepository.GetForMonth was just called")
	}
	callInfo := struct {
		User_id uint64
		Month   time.Time
	}{
		User_id: user_id,
		Month:   month,
	}
	mock.lockGetForMonth.Lock()
	mock.calls.GetForMonth = append(mock.calls.GetForMonth, callInfo)
	mock.lockGetForMonth.Unlock()
	return mock.GetForMonthFunc(user_id, month)
}

// GetForMonthCalls gets all the calls that were made to GetForMonth.
// Check the length with:
//     len(mockedEventRepository.GetForMonthCalls())
func (mock *EventRepositoryMock) GetForMonthCalls() []struct {
	User_id uint64
	Month   time.Time
} {
	var calls []struct {
		User_id uint64
		Month   time.Time
	}
	mock.lockGetForMonth.RLock()
	calls = mock.calls.GetForMonth
	mock.lockGetForMonth.RUnlock()
	return calls
}

// GetForWeek calls GetForWeekFunc.
func (mock *EventRepositoryMock) GetForWeek(user_id uint64, week time.Time) ([]event.Event, error) {
	if mock.GetForWeekFunc == nil {
		panic("EventRepositoryMock.GetForWeekFunc: method is nil but EventRepository.GetForWeek was just called")
	}
	callInfo := struct {
		User_id uint64
		Week    time.Time
	}{
		User_id: user_id,
		Week:    week,
	}
	mock.lockGetForWeek.Lock()
	mock.calls.GetForWeek = append(mock.calls.GetForWeek, callInfo)
	mock.lockGetForWeek.Unlock()
	return mock.GetForWeekFunc(user_id, week)
}

// GetForWeekCalls gets all the calls that were made to GetForWeek.
// Check the length with:
//     len(mockedEventRepository.GetForWeekCalls())
func (mock *EventRepositoryMock) GetForWeekCalls() []struct {
	User_id uint64
	Week    time.Time
} {
	var calls []struct {
		User_id uint64
		Week    time.Time
	}
	mock.lockGetForWeek.RLock()
	calls = mock.calls.GetForWeek
	mock.lockGetForWeek.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *EventRepositoryMock) Update(user_id uint64, e event.Event) error {
	if mock.UpdateFunc == nil {
		panic("EventRepositoryMock.UpdateFunc: method is nil but EventRepository.Update was just called")
	}
	callInfo := struct {
		User_id uint64
		E       event.Event
	}{
		User_id: user_id,
		E:       e,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(user_id, e)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedEventRepository.UpdateCalls())
func (mock *EventRepositoryMock) UpdateCalls() []struct {
	User_id uint64
	E       event.Event
} {
	var calls []struct {
		User_id uint64
		E       event.Event
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
