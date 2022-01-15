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
	require := require.New(t)
	tests := []struct {
		storage        Storage
		deleteId       uuid.UUID
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
			deleteId:       uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
			expectedEvents: []storage.Event{},
		},
	}

	for i := range tests {
		require.NoError(tests[i].storage.DeleteEvent(tests[i].deleteId))
		require.EqualValues(tests[i].storage.events, tests[i].expectedEvents)
	}
}

func TestListByDayEvent(t *testing.T) {
	require := require.New(t)
	tests := []struct {
		storage        Storage
		day            time.Time
		expectedEvents []storage.Event
	}{
		{
			storage: Storage{
				events: []storage.Event{
					{
						ID:             uuid.MustParse("2965d713-e587-4a36-adc7-63afe2c06cb8"),
						Title:          "Christie",
						Description:    "Enim aliquip eu commodo non dolor proident ullamco ex id sint eiusmod veniam.",
						DateTime:       time.Date(2015, 9, 10, 21, 44, 45, 49, time.UTC),
						FinishDateTime: time.Date(2017, 3, 25, 18, 9, 01, 23, time.UTC),
						UserID:         uuid.MustParse("a1f4e1d8-e00b-4ebf-b008-42a926397cf5"),
					},
					{
						ID:             uuid.MustParse("a037df54-072f-4e11-a50b-9a33c71c74d5"),
						Title:          "Aristotle",
						Description:    "Proident occaecat eu ipsum est consectetur ut minim consectetur sunt enim.",
						DateTime:       time.Date(2005, 4, 29, 0, 25, 50, 14, time.UTC),
						FinishDateTime: time.Date(2009, 2, 25, 0, 38, 9, 05, time.UTC),
						UserID:         uuid.MustParse("2a830054-1867-4200-a612-b1dec5373ddb"),
					},
					{
						ID:             uuid.MustParse("3b371527-7727-470c-8edd-3a5e56a4890f"),
						Title:          "Truman",
						Description:    "Duis reprehenderit ipsum pariatur minim.",
						DateTime:       time.Date(1985, 10, 6, 4, 46, 25, 38, time.UTC),
						FinishDateTime: time.Date(2002, 5, 3, 8, 22, 19, 00, time.UTC),
						UserID:         uuid.MustParse("f8a32e15-63dc-4c93-be9d-b4d9891cfeb9"),
					},
					{
						ID:             uuid.MustParse("4bd8f2bc-6acc-4e29-aba6-fcd678f53cfe"),
						Title:          "Bartolomeo",
						Description:    "Irure adipisicing id tempor in anim veniam occaecat.",
						DateTime:       time.Date(2005, 4, 29, 18, 14, 28, 42, time.UTC),
						FinishDateTime: time.Date(2011, 11, 11, 6, 9, 3, 15, time.UTC),
						UserID:         uuid.MustParse("54347080-cde1-4fb6-977b-aad81c8dab28"),
					},
					{
						ID:             uuid.MustParse("a43848c9-1d09-4053-86e2-35aa58919f03"),
						Title:          "Cris",
						Description:    "Tempor mollit et deserunt eu enim enim ullamco quis officia in dolor ea adipisicing.",
						DateTime:       time.Date(2005, 1, 29, 3, 16, 50, 14, time.UTC),
						FinishDateTime: time.Date(2009, 07, 5, 4, 6, 10, 05, time.UTC),
						UserID:         uuid.MustParse("0dc8728f-3684-419a-bad3-3dbb7e082b43"),
					},
				},
			},
			day: time.Date(2005, 4, 29, 23, 16, 23, 0, time.UTC),
			expectedEvents: []storage.Event{
				{
					ID:             uuid.MustParse("a037df54-072f-4e11-a50b-9a33c71c74d5"),
					Title:          "Aristotle",
					Description:    "Proident occaecat eu ipsum est consectetur ut minim consectetur sunt enim.",
					DateTime:       time.Date(2005, 4, 29, 0, 25, 50, 14, time.UTC),
					FinishDateTime: time.Date(2009, 2, 25, 0, 38, 9, 05, time.UTC),
					UserID:         uuid.MustParse("2a830054-1867-4200-a612-b1dec5373ddb"),
				},
				{
					ID:             uuid.MustParse("4bd8f2bc-6acc-4e29-aba6-fcd678f53cfe"),
					Title:          "Bartolomeo",
					Description:    "Irure adipisicing id tempor in anim veniam occaecat.",
					DateTime:       time.Date(2005, 4, 29, 18, 14, 28, 42, time.UTC),
					FinishDateTime: time.Date(2011, 11, 11, 6, 9, 3, 15, time.UTC),
					UserID:         uuid.MustParse("54347080-cde1-4fb6-977b-aad81c8dab28"),
				},
			},
		},
	}

	for i := range tests {
		events, err := tests[i].storage.ListByDayEvent(tests[i].day)
		require.NoError(err)
		require.EqualValues(events, tests[i].expectedEvents)
	}
}

func TestListByWeekEvent(t *testing.T) {
}

func TestListByMonthEvent(t *testing.T) {
}
