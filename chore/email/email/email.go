package email

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"
)

// Config represents the SMTP configuration for email sending.
// These settings should be provided by the self-hosted instance administrator.
type Config struct {
	// Host is the SMTP server hostname (e.g., "smtp.gmail.com")
	Host string
	// Port is the SMTP server port (common: 587 for TLS, 465 for SSL, 25 for unencrypted)
	Port int
	// Username is the SMTP authentication username (usually the email address)
	Username string
	// Password is the SMTP authentication password or app-specific password
	Password string
	// FromEmail is the email address that will appear in the "From" field
	FromEmail string
	// FromName is the display name that will appear in the "From" field
	FromName string
	// UseTLS enables STARTTLS encryption (recommended for port 587)
	UseTLS bool
	// UseSSL enables SSL/TLS encryption (for port 465)
	UseSSL bool
}

// Message represents an email message to be sent.
type Message struct {
	To      []string // Required: recipient email addresses
	Cc      []string // Optional: carbon copy recipients
	Bcc     []string // Optional: blind carbon copy recipients
	Subject string   // Required: email subject
	Body    string   // Required: email body content
	IsHTML  bool     // Whether the body is HTML (default: false for plain text)
	ReplyTo string   // Optional: reply-to address
}

// Validate checks that the message has all required fields.
func (m *Message) Validate() error {
	if len(m.To) == 0 {
		return errors.New("at least one recipient is required")
	}
	if m.Subject == "" {
		return errors.New("subject is required")
	}
	if m.Body == "" {
		return errors.New("body is required")
	}
	return nil
}

// Format creates an RFC 5322 formatted email message.
func (m *Message) Format(fromEmail, fromName string) string {
	var sb strings.Builder
	fromEmail = sanitizeEmailHeaderValue(fromEmail)
	fromName = sanitizeEmailHeaderValue(fromName)
	to := sanitizeEmailHeaderValues(m.To)
	cc := sanitizeEmailHeaderValues(m.Cc)
	replyTo := sanitizeEmailHeaderValue(m.ReplyTo)
	subject := sanitizeEmailHeaderValue(m.Subject)

	// From header
	if fromName != "" {
		fmt.Fprintf(&sb, "From: %s <%s>\r\n", fromName, fromEmail)
	} else {
		fmt.Fprintf(&sb, "From: %s\r\n", fromEmail)
	}

	// To header
	fmt.Fprintf(&sb, "To: %s\r\n", strings.Join(to, ", "))

	// Cc header (optional)
	if len(cc) > 0 {
		fmt.Fprintf(&sb, "Cc: %s\r\n", strings.Join(cc, ", "))
	}

	// Reply-To header (optional)
	if replyTo != "" {
		fmt.Fprintf(&sb, "Reply-To: %s\r\n", replyTo)
	}

	// Subject header
	fmt.Fprintf(&sb, "Subject: %s\r\n", subject)

	// Date header (RFC 5322 format)
	fmt.Fprintf(&sb, "Date: %s\r\n", time.Now().Format(time.RFC1123Z))

	// MIME headers
	sb.WriteString("MIME-Version: 1.0\r\n")

	// Content-Type header
	if m.IsHTML {
		sb.WriteString("Content-Type: text/html; charset=utf-8\r\n")
	} else {
		sb.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	}

	// Empty line separating headers from body
	sb.WriteString("\r\n")

	// Body
	sb.WriteString(m.Body)

	return sb.String()
}

func sanitizeEmailHeaderValue(value string) string {
	value = strings.NewReplacer("\r", " ", "\n", " ").Replace(value)
	return strings.Join(strings.Fields(value), " ")
}

func sanitizeEmailHeaderValues(values []string) []string {
	sanitized := make([]string, 0, len(values))
	for _, value := range values {
		sanitized = append(sanitized, sanitizeEmailHeaderValue(value))
	}
	return sanitized
}

// GetAllRecipients returns all recipients (To, Cc, Bcc) as a single slice.
func (m *Message) GetAllRecipients() []string {
	var recipients []string
	recipients = append(recipients, m.To...)
	recipients = append(recipients, m.Cc...)
	recipients = append(recipients, m.Bcc...)
	return recipients
}

// Client represents an SMTP email client.
type Client struct {
	config *Config
}

// NewClient creates a new email client with the given configuration.
func NewClient(config *Config) *Client {
	return &Client{
		config: config,
	}
}

// Validate checks if the configuration is valid.
func (c *Config) Validate() error {
	if c.Host == "" {
		return errors.New("SMTP host is required")
	}
	if c.Port <= 0 || c.Port > 65535 {
		return errors.New("SMTP port must be between 1 and 65535")
	}
	if c.FromEmail == "" {
		return errors.New("from email is required")
	}
	return nil
}

// GetServerAddress returns the SMTP server address in the format "host:port".
func (c *Config) GetServerAddress() string { return fmt.Sprintf("%s:%d", c.Host, c.Port) }

const smtpOperationTimeout = 15 * time.Second

// validateConfig validates the client configuration.
func (c *Client) validateConfig() error {
	if c.config == nil {
		return errors.New("email configuration is required")
	}
	return c.config.Validate()
}

// createAuth creates an SMTP auth mechanism if credentials are provided.
func (c *Client) createAuth() smtp.Auth {
	if c.config.Username == "" && c.config.Password == "" {
		return nil
	}
	return smtp.PlainAuth("", c.config.Username, c.config.Password, c.config.Host)
}

// createTLSConfig creates a TLS configuration for secure connections.
func (c *Client) createTLSConfig() *tls.Config {
	return &tls.Config{
		ServerName: c.config.Host,
		MinVersion: tls.VersionTLS12,
	}
}

// Send sends an email message via SMTP.
func (c *Client) Send(message *Message) error {
	// Validate configuration
	if err := c.validateConfig(); err != nil {
		return fmt.Errorf("invalid email configuration: %w", err)
	}

	// Validate message
	if message == nil {
		return errors.New("message is required")
	}
	if err := message.Validate(); err != nil {
		return fmt.Errorf("invalid email message: %w", err)
	}

	// Format the message
	body := message.Format(c.config.FromEmail, c.config.FromName)

	// Get all recipients
	recipients := message.GetAllRecipients()

	// Create auth
	auth := c.createAuth()

	// Send based on encryption type
	if c.config.UseSSL {
		return c.sendWithSSL(auth, recipients, body)
	}
	return c.sendWithTLS(auth, recipients, body)
}

// sendWithTLS sends email using STARTTLS (port 587).
func (c *Client) sendWithTLS(auth smtp.Auth, recipients []string, body string) error {
	serverAddr := c.config.GetServerAddress()

	dialer := &net.Dialer{Timeout: smtpOperationTimeout}
	conn, err := dialer.Dial("tcp", serverAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %s: %w", serverAddr, err)
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(smtpOperationTimeout)); err != nil {
		return fmt.Errorf("failed to set SMTP connection deadline: %w", err)
	}

	client, err := smtp.NewClient(conn, c.config.Host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Quit()

	if c.config.UseTLS {
		if ok, _ := client.Extension("STARTTLS"); !ok {
			return errors.New("SMTP server does not support STARTTLS")
		}
		if err := client.StartTLS(c.createTLSConfig()); err != nil {
			return fmt.Errorf("failed to start SMTP STARTTLS: %w", err)
		}
	}

	return c.sendWithClient(client, auth, recipients, body)
}

// sendWithSSL sends email using SSL/TLS (port 465).
func (c *Client) sendWithSSL(auth smtp.Auth, recipients []string, body string) error {
	serverAddr := c.config.GetServerAddress()

	// Create TLS connection
	tlsConfig := c.createTLSConfig()
	dialer := &net.Dialer{Timeout: smtpOperationTimeout}
	conn, err := tls.DialWithDialer(dialer, "tcp", serverAddr, tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server with SSL: %s: %w", serverAddr, err)
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(smtpOperationTimeout)); err != nil {
		return fmt.Errorf("failed to set SMTP connection deadline: %w", err)
	}

	// Create SMTP client
	client, err := smtp.NewClient(conn, c.config.Host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Quit()

	return c.sendWithClient(client, auth, recipients, body)
}

func (c *Client) sendWithClient(
	client *smtp.Client,
	auth smtp.Auth,
	recipients []string,
	body string,
) error {
	if auth != nil {
		if err := client.Auth(auth); err != nil {
			return fmt.Errorf("SMTP authentication failed: %w", err)
		}
	}

	// Set sender
	if err := client.Mail(c.config.FromEmail); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}

	// Set recipients
	for _, recipient := range recipients {
		if err := client.Rcpt(recipient); err != nil {
			return fmt.Errorf("failed to set recipient: %s: %w", recipient, err)
		}
	}

	// Send message body
	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to send DATA command: %w", err)
	}

	if _, err := writer.Write([]byte(body)); err != nil {
		return fmt.Errorf("failed to write message body: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close message writer: %w", err)
	}

	return nil
}
