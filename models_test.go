package main

import (
	"testing"
	"time"

	"github.com/bootdotdev/learn-cicd-starter/internal/database"
)

func TestDatabaseUserToUser(t *testing.T) {
	timeStr := "2023-01-01T12:00:00Z"
	dbUser := database.User{
		ID:        "test-id",
		CreatedAt: timeStr,
		UpdatedAt: timeStr,
		Name:      "Test User",
		ApiKey:    "test-api-key",
	}

	user, err := databaseUserToUser(dbUser)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.ID != dbUser.ID {
		t.Errorf("expected ID %s, got %s", dbUser.ID, user.ID)
	}
	if user.Name != dbUser.Name {
		t.Errorf("expected Name %s, got %s", dbUser.Name, user.Name)
	}
	if user.ApiKey != dbUser.ApiKey {
		t.Errorf("expected ApiKey %s, got %s", dbUser.ApiKey, user.ApiKey)
	}

	expectedTime, _ := time.Parse(time.RFC3339, timeStr)
	if !user.CreatedAt.Equal(expectedTime) {
		t.Errorf("expected CreatedAt %v, got %v", expectedTime, user.CreatedAt)
	}
}

func TestDatabaseNoteToNote(t *testing.T) {
	timeStr := "2023-01-01T12:00:00Z"
	dbNote := database.Note{
		ID:        "note-id",
		CreatedAt: timeStr,
		UpdatedAt: timeStr,
		Note:      "Test note",
		UserID:    "user-id",
	}

	note, err := databaseNoteToNote(dbNote)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if note.ID != dbNote.ID {
		t.Errorf("expected ID %s, got %s", dbNote.ID, note.ID)
	}
	if note.Note != dbNote.Note {
		t.Errorf("expected Note %s, got %s", dbNote.Note, note.Note)
	}
	if note.UserID != dbNote.UserID {
		t.Errorf("expected UserID %s, got %s", dbNote.UserID, note.UserID)
	}
}

func TestDatabasePostsToPosts(t *testing.T) {
	timeStr := "2023-01-01T12:00:00Z"
	dbNotes := []database.Note{
		{
			ID:        "note-1",
			CreatedAt: timeStr,
			UpdatedAt: timeStr,
			Note:      "First note",
			UserID:    "user-id",
		},
		{
			ID:        "note-2",
			CreatedAt: timeStr,
			UpdatedAt: timeStr,
			Note:      "Second note",
			UserID:    "user-id",
		},
	}

	notes, err := databasePostsToPosts(dbNotes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(notes) != len(dbNotes) {
		t.Errorf("expected %d notes, got %d", len(dbNotes), len(notes))
	}

	for i, note := range notes {
		if note.ID != dbNotes[i].ID {
			t.Errorf("expected note %d ID %s, got %s", i, dbNotes[i].ID, note.ID)
		}
	}
}
