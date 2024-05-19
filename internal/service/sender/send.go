package sender

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/Waramoto/hryvnia-svc/internal/service/handlers"
	"github.com/Waramoto/hryvnia-svc/internal/types"
)

func (s *Sender) Send(_ context.Context) error {
	emailConfig := s.config.SenderConfig().Email

	currentRate, err := handlers.GetCurrentUAHRate()
	if err != nil {
		return fmt.Errorf("failed to get current rate: %w", err)
	}

	subscribers, err := s.db.New().Select()
	if err != nil {
		return fmt.Errorf("failed to select subscribers: %w", err)
	}

	receivers := make([]string, 0, len(subscribers))
	for _, subscriber := range subscribers {
		isTimeToSend := time.Now().After(subscriber.LastSend.Add(s.config.SenderConfig().Period))
		if subscriber.Status == types.StatusNotSent || isTimeToSend {
			receivers = append(receivers, subscriber.Email)
		}
	}

	to := strings.Join(receivers, ", ")
	msg := []byte(fmt.Sprintf("From: %s\r\n", emailConfig.From) +
		fmt.Sprintf("To: %s\r\n", to) +
		"Subject: UAH exchange rate\r\n" +
		"\r\n" +
		fmt.Sprintf("Current USD to UAH exchange rate: %.2f\r\n", currentRate))

	err = smtp.SendMail(
		fmt.Sprintf("%s:%s", emailConfig.Host, emailConfig.Port),
		s.auth,
		emailConfig.From,
		receivers,
		msg,
	)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	err = s.db.New().FilterByEmails(receivers...).UpdateLastSend(time.Now())
	if err != nil {
		return fmt.Errorf("failed to update last send time: %w", err)
	}

	err = s.db.New().FilterByEmails(receivers...).UpdateStatus(types.StatusSent)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	return nil
}
