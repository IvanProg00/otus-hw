package memorystorage

import (
	"context"
	"fmt"
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
		storage        memoryStorage
		addEvent       storage.Event
		expectedEvents []storage.Event
	}{
		{
			addEvent: storage.Event{
				ID:          uuid.MustParse("476da6aa-8caf-406d-8242-139245204520"),
				Title:       "Flossi",
				Description: "Fix San",
				StartAt:     time.Date(1994, 14, 1, 2, 6, 28, 14, time.UTC),
				FinishAt:    time.Date(2009, 5, 28, 9, 9, 1, 5, time.UTC),
				UserID:      uuid.MustParse("0207d080-f49e-49bf-9dab-2165ffb9bed3"),
			},
			expectedEvents: []storage.Event{
				{
					ID:          uuid.MustParse("476da6aa-8caf-406d-8242-139245204520"),
					Title:       "Flossi",
					Description: "Fix San",
					StartAt:     time.Date(1994, 14, 1, 2, 6, 28, 14, time.UTC),
					FinishAt:    time.Date(2009, 5, 28, 9, 9, 1, 5, time.UTC),
					UserID:      uuid.MustParse("0207d080-f49e-49bf-9dab-2165ffb9bed3"),
				},
			},
		},
	}

	for i := range tests {
		require.NoError(tests[i].storage.CreateEvent(context.Background(), tests[i].addEvent))
		require.EqualValues(tests[i].expectedEvents, tests[i].storage.events)
	}
}

func TestUpdateEvent(t *testing.T) {
	require := require.New(t)
	tests := []struct {
		storage        memoryStorage
		updateEvent    storage.Event
		expectedEvents []storage.Event
	}{
		{
			storage: memoryStorage{
				events: []storage.Event{
					{
						ID:          uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
						Title:       "Flossi",
						Description: "Fix San",
						StartAt:     time.Date(1994, 14, 1, 2, 6, 28, 14, time.UTC),
						FinishAt:    time.Date(2009, 5, 28, 9, 9, 1, 5, time.UTC),
						UserID:      uuid.MustParse("5ea3b925-ec2b-47c4-a8e9-53ab20c15084"),
					},
				},
			},
			updateEvent: storage.Event{
				ID:          uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
				Title:       "Voyatouch",
				Description: "Bamity",
				StartAt:     time.Date(1987, 7, 26, 5, 4, 10, 15, time.UTC),
				FinishAt:    time.Date(2002, 11, 12, 6, 39, 51, 45, time.UTC),
				UserID:      uuid.MustParse("72b0b4fc-586c-439b-b97c-7783ac3e233b"),
			},
			expectedEvents: []storage.Event{
				{
					ID:          uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
					Title:       "Voyatouch",
					Description: "Bamity",
					StartAt:     time.Date(1987, 7, 26, 5, 4, 10, 15, time.UTC),
					FinishAt:    time.Date(2002, 11, 12, 6, 39, 51, 45, time.UTC),
					UserID:      uuid.MustParse("72b0b4fc-586c-439b-b97c-7783ac3e233b"),
				},
			},
		},
	}

	for i := range tests {
		require.NoError(tests[i].storage.UpdateEvent(context.Background(), tests[i].updateEvent))
		require.EqualValues(tests[i].expectedEvents, tests[i].storage.events)
	}
}

func TestUpdateEvent_notFound(t *testing.T) {
	require := require.New(t)
	tests := []struct {
		storage     memoryStorage
		updateEvent storage.Event
	}{
		{
			storage: memoryStorage{
				events: []storage.Event{
					{
						ID:          uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
						Title:       "Flossi",
						Description: "Fix San",
						StartAt:     time.Date(1994, 14, 1, 2, 6, 28, 14, time.UTC),
						FinishAt:    time.Date(2009, 5, 28, 9, 9, 1, 5, time.UTC),
						UserID:      uuid.MustParse("5ea3b925-ec2b-47c4-a8e9-53ab20c15084"),
					},
				},
			},
			updateEvent: storage.Event{
				ID:          uuid.MustParse("5871c266-e102-4a26-a3e8-a35757d93964"),
				Title:       "Voyatouch",
				Description: "Bamity",
				StartAt:     time.Date(1987, 7, 26, 5, 4, 10, 15, time.UTC),
				FinishAt:    time.Date(2002, 11, 12, 6, 39, 51, 45, time.UTC),
				UserID:      uuid.MustParse("72b0b4fc-586c-439b-b97c-7783ac3e233b"),
			},
		},
	}

	for i := range tests {
		expectedEvents := make([]storage.Event, len(tests[i].storage.events))
		copy(expectedEvents, tests[i].storage.events)
		require.ErrorIs(tests[i].storage.UpdateEvent(context.Background(), tests[i].updateEvent), errorsstorage.ErrNotFound)
		require.EqualValues(expectedEvents, tests[i].storage.events)
	}
}

func TestDeleteEvent(t *testing.T) {
	require := require.New(t)
	tests := []struct {
		storage        memoryStorage
		deleteID       uuid.UUID
		updateEvent    storage.Event
		expectedEvents []storage.Event
	}{
		{
			storage: memoryStorage{
				events: []storage.Event{
					{
						ID:          uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
						Title:       "Flossi",
						Description: "Fix San",
						StartAt:     time.Date(1994, 14, 1, 2, 6, 28, 14, time.UTC),
						FinishAt:    time.Date(2009, 5, 28, 9, 9, 1, 5, time.UTC),
						UserID:      uuid.MustParse("5ea3b925-ec2b-47c4-a8e9-53ab20c15084"),
					},
				},
			},
			deleteID:       uuid.MustParse("66240999-d75b-437a-8205-15d7bbd1213f"),
			expectedEvents: []storage.Event{},
		},
	}

	for i := range tests {
		require.NoError(tests[i].storage.DeleteEvent(context.Background(), tests[i].deleteID))
		require.EqualValues(tests[i].expectedEvents, tests[i].storage.events)
	}
}

func TestListByDayEvent(t *testing.T) {
	tests := []struct {
		storage        memoryStorage
		day            time.Time
		expectedEvents []storage.Event
	}{
		{
			storage: memoryStorage{
				events: []storage.Event{
					{
						ID:          uuid.MustParse("2965d713-e587-4a36-adc7-63afe2c06cb8"),
						Title:       "Christie",
						Description: "Enim aliquip eu commodo non dolor proident ullamco ex id sint eiusmod veniam.",
						StartAt:     time.Date(2015, 9, 10, 21, 44, 45, 49, time.UTC),
						FinishAt:    time.Date(2017, 3, 25, 18, 9, 1, 23, time.UTC),
						UserID:      uuid.MustParse("a1f4e1d8-e00b-4ebf-b008-42a926397cf5"),
					},
					{
						ID:          uuid.MustParse("a037df54-072f-4e11-a50b-9a33c71c74d5"),
						Title:       "Aristotle",
						Description: "Proident occaecat eu ipsum est consectetur ut minim consectetur sunt enim.",
						StartAt:     time.Date(2005, 4, 29, 0, 25, 50, 14, time.UTC),
						FinishAt:    time.Date(2009, 2, 25, 0, 38, 9, 5, time.UTC),
						UserID:      uuid.MustParse("2a830054-1867-4200-a612-b1dec5373ddb"),
					},
					{
						ID:          uuid.MustParse("3b371527-7727-470c-8edd-3a5e56a4890f"),
						Title:       "Truman",
						Description: "Duis reprehenderit ipsum pariatur minim.",
						StartAt:     time.Date(1985, 10, 6, 4, 46, 25, 38, time.UTC),
						FinishAt:    time.Date(2002, 5, 3, 8, 22, 19, 0, time.UTC),
						UserID:      uuid.MustParse("f8a32e15-63dc-4c93-be9d-b4d9891cfeb9"),
					},
					{
						ID:          uuid.MustParse("4bd8f2bc-6acc-4e29-aba6-fcd678f53cfe"),
						Title:       "Bartolomeo",
						Description: "Irure adipisicing id tempor in anim veniam occaecat.",
						StartAt:     time.Date(2005, 4, 29, 18, 14, 28, 42, time.UTC),
						FinishAt:    time.Date(2011, 11, 11, 6, 9, 3, 15, time.UTC),
						UserID:      uuid.MustParse("54347080-cde1-4fb6-977b-aad81c8dab28"),
					},
					{
						ID:          uuid.MustParse("a43848c9-1d09-4053-86e2-35aa58919f03"),
						Title:       "Cris",
						Description: "Tempor mollit et deserunt eu enim enim ullamco quis officia in dolor ea adipisicing.",
						StartAt:     time.Date(2005, 1, 29, 3, 16, 50, 14, time.UTC),
						FinishAt:    time.Date(2009, 7, 5, 4, 6, 10, 5, time.UTC),
						UserID:      uuid.MustParse("0dc8728f-3684-419a-bad3-3dbb7e082b43"),
					},
					{
						ID:          uuid.MustParse("c4292ab2-8920-421e-bc6c-3b815a2194cf"),
						Title:       "Bruno",
						Description: "Consectetur magna et minim aliquip irure tempor qui fugiat culpa consectetur.",
						StartAt:     time.Date(2008, 4, 28, 3, 16, 50, 14, time.UTC),
						FinishAt:    time.Date(2009, 7, 5, 4, 6, 10, 5, time.UTC),
						UserID:      uuid.MustParse("d70507ec-d509-4338-bd93-74002cfc63ce"),
					},
				},
			},
			day: time.Date(2005, 4, 29, 23, 16, 23, 0, time.UTC),
			expectedEvents: []storage.Event{
				{
					ID:          uuid.MustParse("a037df54-072f-4e11-a50b-9a33c71c74d5"),
					Title:       "Aristotle",
					Description: "Proident occaecat eu ipsum est consectetur ut minim consectetur sunt enim.",
					StartAt:     time.Date(2005, 4, 29, 0, 25, 50, 14, time.UTC),
					FinishAt:    time.Date(2009, 2, 25, 0, 38, 9, 5, time.UTC),
					UserID:      uuid.MustParse("2a830054-1867-4200-a612-b1dec5373ddb"),
				},
				{
					ID:          uuid.MustParse("4bd8f2bc-6acc-4e29-aba6-fcd678f53cfe"),
					Title:       "Bartolomeo",
					Description: "Irure adipisicing id tempor in anim veniam occaecat.",
					StartAt:     time.Date(2005, 4, 29, 18, 14, 28, 42, time.UTC),
					FinishAt:    time.Date(2011, 11, 11, 6, 9, 3, 15, time.UTC),
					UserID:      uuid.MustParse("54347080-cde1-4fb6-977b-aad81c8dab28"),
				},
			},
		},
	}

	for i := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			require := require.New(t)
			events, err := tests[i].storage.ListByDayEvent(context.Background(), tests[i].day)
			require.NoError(err)
			require.EqualValues(tests[i].expectedEvents, events)
		})
	}
}

func TestListByWeekEvent(t *testing.T) {
	require := require.New(t)
	tests := []struct {
		storage        memoryStorage
		week           time.Time
		expectedEvents []storage.Event
	}{
		{
			storage: memoryStorage{
				events: []storage.Event{
					{
						ID:          uuid.MustParse("dc4b884a-cfab-4c8f-bcfd-6d28e3ea9ed3"),
						Title:       "Aubert",
						Description: "Officia labore nisi consectetur proident ut sint mollit quis esse est duis nisi amet amet.",
						StartAt:     time.Date(2015, 4, 8, 21, 44, 45, 49, time.UTC),
						FinishAt:    time.Date(2017, 3, 25, 18, 9, 1, 23, time.UTC),
						UserID:      uuid.MustParse("d04dbaa7-b62a-4169-b003-e6bd4281fd63"),
					},
					{
						ID:          uuid.MustParse("1c8eec63-da3f-4471-bb5e-850e27f2aea7"),
						Title:       "Shepperd",
						Description: "Cupidatat ea commodo eu exercitation quis do fugiat quis et.",
						StartAt:     time.Date(2005, 4, 10, 0, 25, 50, 14, time.UTC),
						FinishAt:    time.Date(2009, 2, 25, 0, 38, 9, 5, time.UTC),
						UserID:      uuid.MustParse("6da73d09-6811-4d33-bc0c-2f9114b4194c"),
					},
					{
						ID:          uuid.MustParse("a3f713df-8069-41f3-a585-8c3a83effa14"),
						Title:       "Robbin",
						Description: "Elit dolore qui ex cupidatat eu id sit consectetur nisi cillum sit.",
						StartAt:     time.Date(2005, 4, 11, 4, 46, 25, 38, time.UTC),
						FinishAt:    time.Date(2019, 5, 3, 8, 22, 19, 0, time.UTC),
						UserID:      uuid.MustParse("9a0c2886-5048-4adc-8b32-fb07fdb43b72"),
					},
					{
						ID:          uuid.MustParse("f00e1d3b-17a8-4955-9c5b-04780fa3842a"),
						Title:       "Uriah",
						Description: "Quis aute ex dolor sint quis eu.",
						StartAt:     time.Date(2005, 4, 5, 18, 14, 28, 42, time.UTC),
						FinishAt:    time.Date(2011, 11, 11, 6, 9, 3, 15, time.UTC),
						UserID:      uuid.MustParse("3adf480b-875a-4c1d-9a4e-939eb47bbc11"),
					},
					{
						ID:          uuid.MustParse("3622a3d7-d153-4d59-89df-869cd6440211"),
						Title:       "Daffi",
						Description: "Ullamco reprehenderit culpa aute elit dolore et consectetur culpa.",
						StartAt:     time.Date(2005, 4, 4, 3, 16, 50, 14, time.UTC),
						FinishAt:    time.Date(2009, 7, 5, 4, 6, 10, 5, time.UTC),
						UserID:      uuid.MustParse("7045b6ac-dfa4-4aed-9105-acef7c0063de"),
					},
				},
			},
			week: time.Date(2005, 4, 4, 15, 16, 23, 0, time.UTC),
			expectedEvents: []storage.Event{
				{
					ID:          uuid.MustParse("1c8eec63-da3f-4471-bb5e-850e27f2aea7"),
					Title:       "Shepperd",
					Description: "Cupidatat ea commodo eu exercitation quis do fugiat quis et.",
					StartAt:     time.Date(2005, 4, 10, 0, 25, 50, 14, time.UTC),
					FinishAt:    time.Date(2009, 2, 25, 0, 38, 9, 5, time.UTC),
					UserID:      uuid.MustParse("6da73d09-6811-4d33-bc0c-2f9114b4194c"),
				},
				{
					ID:          uuid.MustParse("f00e1d3b-17a8-4955-9c5b-04780fa3842a"),
					Title:       "Uriah",
					Description: "Quis aute ex dolor sint quis eu.",
					StartAt:     time.Date(2005, 4, 5, 18, 14, 28, 42, time.UTC),
					FinishAt:    time.Date(2011, 11, 11, 6, 9, 3, 15, time.UTC),
					UserID:      uuid.MustParse("3adf480b-875a-4c1d-9a4e-939eb47bbc11"),
				},
				{
					ID:          uuid.MustParse("3622a3d7-d153-4d59-89df-869cd6440211"),
					Title:       "Daffi",
					Description: "Ullamco reprehenderit culpa aute elit dolore et consectetur culpa.",
					StartAt:     time.Date(2005, 4, 4, 3, 16, 50, 14, time.UTC),
					FinishAt:    time.Date(2009, 7, 5, 4, 6, 10, 5, time.UTC),
					UserID:      uuid.MustParse("7045b6ac-dfa4-4aed-9105-acef7c0063de"),
				},
			},
		},
	}

	for i := range tests {
		events, err := tests[i].storage.ListByWeekEvent(context.Background(), tests[i].week)
		require.NoError(err)
		require.EqualValues(tests[i].expectedEvents, events)
	}
}

func TestListByMonthEvent(t *testing.T) {
	tests := []struct {
		storage        memoryStorage
		day            time.Time
		expectedEvents []storage.Event
	}{
		{
			storage: memoryStorage{
				events: []storage.Event{
					{
						ID:          uuid.MustParse("88695f39-5621-4034-b64a-a3b3b5bb5402"),
						Title:       "Gabriel",
						Description: "Culpa excepteur reprehenderit veniam voluptate.",
						StartAt:     time.Date(2005, 5, 1, 21, 44, 45, 49, time.UTC),
						FinishAt:    time.Date(2017, 3, 25, 18, 9, 1, 23, time.UTC),
						UserID:      uuid.MustParse("8f09cf01-3d8f-49ae-96d7-fb398b497c57"),
					},
					{
						ID:          uuid.MustParse("c0db6973-7145-4119-890d-7f62662d5e86"),
						Title:       "Etta",
						Description: "Minim anim anim anim est non voluptate duis sit velit.",
						StartAt:     time.Date(2005, 4, 12, 0, 25, 50, 14, time.UTC),
						FinishAt:    time.Date(2009, 2, 25, 0, 38, 9, 5, time.UTC),
						UserID:      uuid.MustParse("7db77af6-4218-439b-8851-8cbafde47213"),
					},
					{
						ID:    uuid.MustParse("c8b69af1-2042-45e2-824a-45404dc578aa"),
						Title: "Lexine",
						Description: "Magna eiusmod laboris Lorem nostrud ad ex dolore pariatur" +
							" fugiat labore ullamco proident do aliqua.",
						StartAt:  time.Date(2004, 4, 15, 4, 46, 25, 38, time.UTC),
						FinishAt: time.Date(2002, 5, 3, 8, 22, 19, 0, time.UTC),
						UserID:   uuid.MustParse("719f59b6-3c79-4457-8551-ed2fcc5c332e"),
					},
					{
						ID:          uuid.MustParse("d7e1f028-e045-4b72-a76e-f4d6f8af12f1"),
						Title:       "Bride",
						Description: "Do sunt elit cupidatat irure adipisicing in ut.",
						StartAt:     time.Date(2005, 4, 1, 18, 14, 28, 42, time.UTC),
						FinishAt:    time.Date(2011, 11, 11, 6, 9, 3, 15, time.UTC),
						UserID:      uuid.MustParse("02174c96-9bd0-4825-9841-3515e5547e2d"),
					},
					{
						ID:          uuid.MustParse("2126d0ba-2a89-4ab8-bcfe-9c66c21b61f5"),
						Title:       "Pen",
						Description: "Quis elit culpa Lorem ad laboris voluptate.",
						StartAt:     time.Date(2005, 4, 30, 3, 16, 50, 14, time.UTC),
						FinishAt:    time.Date(2009, 7, 5, 4, 6, 10, 5, time.UTC),
						UserID:      uuid.MustParse("f857de76-4871-4fd4-add2-3191c19510cb"),
					},
					{
						ID:          uuid.MustParse("eb46f07f-6964-42b6-aa57-98238f77b4ee"),
						Title:       "Karina",
						Description: "Non eu quis veniam dolor exercitation dolore eu.",
						StartAt:     time.Date(2008, 4, 28, 3, 16, 50, 14, time.UTC),
						FinishAt:    time.Date(2009, 7, 5, 4, 6, 10, 5, time.UTC),
						UserID:      uuid.MustParse("88a419b0-9c4a-4a4e-899f-666c8124a599"),
					},
				},
			},
			day: time.Date(2005, 4, 1, 23, 16, 23, 0, time.UTC),
			expectedEvents: []storage.Event{
				{
					ID:          uuid.MustParse("c0db6973-7145-4119-890d-7f62662d5e86"),
					Title:       "Etta",
					Description: "Minim anim anim anim est non voluptate duis sit velit.",
					StartAt:     time.Date(2005, 4, 12, 0, 25, 50, 14, time.UTC),
					FinishAt:    time.Date(2009, 2, 25, 0, 38, 9, 5, time.UTC),
					UserID:      uuid.MustParse("7db77af6-4218-439b-8851-8cbafde47213"),
				},
				{
					ID:          uuid.MustParse("d7e1f028-e045-4b72-a76e-f4d6f8af12f1"),
					Title:       "Bride",
					Description: "Do sunt elit cupidatat irure adipisicing in ut.",
					StartAt:     time.Date(2005, 4, 1, 18, 14, 28, 42, time.UTC),
					FinishAt:    time.Date(2011, 11, 11, 6, 9, 3, 15, time.UTC),
					UserID:      uuid.MustParse("02174c96-9bd0-4825-9841-3515e5547e2d"),
				},
				{
					ID:          uuid.MustParse("2126d0ba-2a89-4ab8-bcfe-9c66c21b61f5"),
					Title:       "Pen",
					Description: "Quis elit culpa Lorem ad laboris voluptate.",
					StartAt:     time.Date(2005, 4, 30, 3, 16, 50, 14, time.UTC),
					FinishAt:    time.Date(2009, 7, 5, 4, 6, 10, 5, time.UTC),
					UserID:      uuid.MustParse("f857de76-4871-4fd4-add2-3191c19510cb"),
				},
			},
		},
	}

	for i := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			require := require.New(t)
			events, err := tests[i].storage.ListByMonthEvent(context.Background(), tests[i].day)
			require.NoError(err)
			require.EqualValues(tests[i].expectedEvents, events)
		})
	}
}
