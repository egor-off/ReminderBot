package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"src/lib/e"
	"strconv"
)

const (
	getUpdatesMethod = "getUpdates"
	sendMessageMethod = "sendMessage"
	errMsg = "can't send message"
	errMarshalJSON = "cannot Marshal json"
	errUnmarshalJSON = "can't unmarshal json"
)

type Client struct {
	host		string
	basePath	string
	client		http.Client
}

func New(host string, token string) *Client {
	return &Client{
		host:	host,
		basePath: newBasePath(token),
		client:	http.Client{},
	}
}

func NewMessage(chatID int, text string, buttons *InlineKeyboardMarkup) *Message {
	return &Message{
		Chat: Chat{
			ID: chatID,
		},
		Text: text,
		Buttons: buttons,
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, e.Wrap("can't do request", err)
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, e.Wrap(errUnmarshalJSON, err)
	}

	return res.Result, nil
}

func (c *Client) SendMessage(message Message) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(message.Chat.ID))
	q.Add("text", message.Text)

	// me don't like it
	if message.Buttons != nil {
		s, err := prepareJSON(message.Buttons)
		if err != nil {
			return err
		}
		q.Add("reply_markup", s)
	}

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap(errMsg, err)
	}
	return nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	const errMsg = "can't do request"
	u := url.URL{
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)

	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	defer func() {_ = resp.Body.Close()}()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	return body, nil
}

func prepareJSON(buttons *InlineKeyboardMarkup) (string, error) {
	b, err := json.Marshal(*buttons)
	if err != nil {
		return "", e.Wrap(errMarshalJSON, err)
	}
	return string(b), nil
}
