package mailpit

import (
	"context"
	"fmt"
	"journey/internal/pgstore"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wneessen/go-mail"
)

type store interface {
	GetTrip(context.Context, uuid.UUID) (pgstore.Trip, error)
}

type Mailpit struct {
	store store
}

func NewMailpit(pool *pgxpool.Pool) Mailpit {
	return Mailpit{pgstore.New(pool)}
}

func (mp Mailpit) SendConfirmTripEmailToTripOwner(tripID uuid.UUID) error {
	ctx := context.Background()
	trip, err := mp.store.GetTrip(ctx, tripID)

	if err != nil {
		return fmt.Errorf("mailpit: failed to get trip for SendConfirmTripEmailToTripOwner: %w", err)
	}

	msg := mail.NewMsg()
	if err := msg.From("mailpit@journey.com"); err != nil {
		return fmt.Errorf("mailpit: failed to set From in email SendConfirmEmailToTripOwner: %w", err)
	}

	if err := msg.To(trip.OwnerEmail); err != nil {
		return fmt.Errorf("mailpit: failed to set To in email SendConfirmEmailToTripOwner: %w", err)
	}

	msg.Subject("Confirme sua viagem")
	msg.SetBodyString(mail.TypeTextPlain, fmt.Sprintf(`
		Olá, %s!

		A sua viagem para %s que começa no dia %s precisa ser confirmada.
		Clique no botão abaixo para confirmar. 
	`,
		trip.OwnerName, trip.Destination, trip.StartsAt.Time.Format(time.DateOnly),
	))

	client, err := mail.NewClient("localhost", mail.WithTLSPortPolicy(mail.NoTLS), mail.WithPort(1025))

	if err != nil {
		return fmt.Errorf("mailpit: failed to create email client for SendConfirmEmailToTripOwner: %w", err)
	}

	if err := client.DialAndSend(msg); err != nil {
		return fmt.Errorf("mailpit: failed to send email SendConfirmEmailToTripOwner: %w", err)
	}

	return nil
}

func (mp Mailpit) SendConfirmTripEmailToTripParticipant(tripID uuid.UUID, email string) error {
	ctx := context.Background()
	trip, err := mp.store.GetTrip(ctx, tripID)

	if err != nil {
		return fmt.Errorf("mailpit: failed to get trip for SendConfirmTripEmailToTripParticipant: %w", err)
	}

	msg := mail.NewMsg()
	if err := msg.From("mailpit@journey.com"); err != nil {
		return fmt.Errorf("mailpit: failed to set From in email SendConfirmEmailToTripParticipant: %w", err)
	}

	if err := msg.To(email); err != nil {
		return fmt.Errorf("mailpit: failed to set To in email SendConfirmTripEmailToTripParticipant: %w", err)
	}

	msg.Subject("Confirme sua viagem")
	msg.SetBodyString(mail.TypeTextPlain, fmt.Sprintf(`
		Saudações,

		A sua viagem para %s que começa no dia %s precisa ser confirmada.
		Clique no botão abaixo para confirmar. 
	`,
		trip.Destination, trip.StartsAt.Time.Format(time.DateOnly),
	))

	client, err := mail.NewClient("mailpit", mail.WithTLSPortPolicy(mail.NoTLS), mail.WithPort(1025))

	if err != nil {
		return fmt.Errorf("mailpit: failed to create email client for SendConfirmEmailToTripOwner: %w", err)
	}

	if err := client.DialAndSend(msg); err != nil {
		return fmt.Errorf("mailpit: failed to send email SendConfirmEmailToTripOwner: %w", err)
	}

	return nil
}

func (mp Mailpit) SendTripConfirmationEmailToParticipants(tripID uuid.UUID, email string) error {
	ctx := context.Background()
	trip, err := mp.store.GetTrip(ctx, tripID)
	if err != nil {
		return fmt.Errorf("mailpit: failed to get trip for SendTripConfirmationToParticipants: %w", err)
	}

	msg := mail.NewMsg()
	if err := msg.From("mailpit@journey.com"); err != nil {
		return fmt.Errorf("mailpit: failed to send email for SendTripConfirmationToParticipants: %w", err)
	}

	if err := msg.To(email); err != nil {
		return fmt.Errorf("mailpit: failed to attach participant email for SendTripConfirmationToInvitators: %w", err)
	}

	msg.Subject("Viagem confirmada!")
	msg.SetBodyString(mail.TypeTextPlain, fmt.Sprintf(`
		Saudações,

		a sua viagem para %s que começa no dia %s foi confirmada por %s.
	`, trip.Destination, trip.StartsAt.Time.Format(time.DateOnly), trip.OwnerName,
	))

	client, err := mail.NewClient("localhost", mail.WithTLSPortPolicy(mail.NoTLS), mail.WithPort(1025))
	if err != nil {
		return fmt.Errorf("mailpit: failed to create email client for SendTripConfirmationToInvitators: %w", err)
	}

	if err := client.DialAndSend(msg); err != nil {
		return fmt.Errorf("mailpit: failed to send message to participant email for SendTripConfirmationToInvitators: %w", err)
	}

	return nil
}
