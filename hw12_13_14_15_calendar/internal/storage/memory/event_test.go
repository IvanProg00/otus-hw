package memorystorage

import (
	"testing"
	"time"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage"
	errorsstorage "github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateEvent(t *testing.T) {
	require := require.New(t)
	tests := []struct {
		storage        Storage
		addEvent       storage.Event
		expectedEvents []storage.Event
	}{
		{
			addEvent: storage.Event{
				ID:             uuid.MustParse("476da6aa-8caf-406d-8242-139245204520"),
				Title:          "Flossi",
				Description:    "Fix San",
				DateTime:       time.Date(1994, 14, 1, 2, 6, 28, 14, time.UTC),
				FinishDateTime: time.Date(2009, 5, 28, 9, 9, 01, 05, time.UTC),
				UserID:         uuid.MustParse("0207d080-f49e-49bf-9dab-2165ffb9bed3"),
			},
			expectedEvents: []storage.Event{
				{
					ID:             uuid.MustParse("476da6aa-8caf-406d-8242-139245204520"),
					Title:          "Flossi",
					Description:    "Fix San",
					DateTime:       time.Date(1994, 14, 1, 2, 6, 28, 14, time.UTC),
					FinishDateTime: time.Date(2009, 5, 28, 9, 9, 01, 05, time.UTC),
					UserID:         uuid.MustParse("0207d080-f49e-49bf-9dab-2165ffb9bed3"),
				},
			},
		},
	}

	for i := range tests {
		require.NoError(tests[i].storage.CreateEvent(tests[i].addEvent))
		require.EqualValues(tests[i].storage.events, tests[i].expectedEvents)
	}
}

func TestUpdateEvent(t *testing.T) {
	require := require.New(t)
	tests := []struct {
		storage        Storage
		updateId       uuid.UUID
		updateEvent    storage.Event
		expectedEvents []storage.Event
	}{
		{
			storage: Storage{
				events: []storage.Event{
					{
						ID:             uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
						Title:          "Flossi",
						Description:    "Fix San",
						DateTime:       time.Date(1994, 14, 1, 2, 6, 28, 14, time.UTC),
						FinishDateTime: time.Date(2009, 5, 28, 9, 9, 01, 05, time.UTC),
						UserID:         uuid.MustParse("5ea3b925-ec2b-47c4-a8e9-53ab20c15084"),
					},
				},
			},
			updateId: uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
			updateEvent: storage.Event{
				ID:             uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
				Title:          "Voyatouch",
				Description:    "Bamity",
				DateTime:       time.Date(1987, 7, 26, 5, 4, 10, 15, time.UTC),
				FinishDateTime: time.Date(2002, 11, 12, 6, 39, 51, 45, time.UTC),
				UserID:         uuid.MustParse("72b0b4fc-586c-439b-b97c-7783ac3e233b"),
			},
			expectedEvents: []storage.Event{
				{
					ID:             uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
					Title:          "Voyatouch",
					Description:    "Bamity",
					DateTime:       time.Date(1987, 7, 26, 5, 4, 10, 15, time.UTC),
					FinishDateTime: time.Date(2002, 11, 12, 6, 39, 51, 45, time.UTC),
					UserID:         uuid.MustParse("72b0b4fc-586c-439b-b97c-7783ac3e233b"),
				},
			},
		},
	}

	for i := range tests {
		require.NoError(tests[i].storage.UpdateEvent(tests[i].updateId, tests[i].updateEvent))
		require.EqualValues(tests[i].storage.events, tests[i].expectedEvents)
	}
}

func TestUpdateEvent_notFound(t *testing.T) {
	require := require.New(t)
	tests := []struct {
		storage     Storage
		updateId    uuid.UUID
		updateEvent storage.Event
	}{
		{
			storage: Storage{
				events: []storage.Event{
					{
						ID:             uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
						Title:          "Flossi",
						Description:    "Fix San",
						DateTime:       time.Date(1994, 14, 1, 2, 6, 28, 14, time.UTC),
						FinishDateTime: time.Date(2009, 5, 28, 9, 9, 01, 05, time.UTC),
						UserID:         uuid.MustParse("5ea3b925-ec2b-47c4-a8e9-53ab20c15084"),
					},
				},
			},
			updateId: uuid.MustParse("5871c266-e102-4a26-a3e8-a35757d93964"),
			updateEvent: storage.Event{
				ID:             uuid.MustParse("5871c266-e102-4a26-a3e8-a35757d93964"),
				Title:          "Voyatouch",
				Description:    "Bamity",
				DateTime:       time.Date(1987, 7, 26, 5, 4, 10, 15, time.UTC),
				FinishDateTime: time.Date(2002, 11, 12, 6, 39, 51, 45, time.UTC),
				UserID:         uuid.MustParse("72b0b4fc-586c-439b-b97c-7783ac3e233b"),
			},
		},
	}

	for i := range tests {
		expectedEvents := make([]storage.Event, len(tests[i].storage.events))
		copy(expectedEvents, tests[i].storage.events)
		require.ErrorIs(tests[i].storage.UpdateEvent(tests[i].updateId, tests[i].updateEvent), errorsstorage.ErrNotFound)
		require.EqualValues(tests[i].storage.events, expectedEvents)
	}
}

func TestDeleteEvent(t *testing.T) {
}

func TestListForDayEvent(t *testing.T) {
}

func TestListForWeekEvent(t *testing.T) {
}

func TestListForMonthEvent(t *testing.T) {
}
