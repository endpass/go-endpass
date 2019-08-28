package endpass

import (
	"fmt"
	"io"
	"net/http"
)

type Document struct {
	ID               string `json:"id"`
	CreatedAt        int64  `json:"createdAt"`
	Status           string `json:"status"`
	DocumentType     string `json:"documentType"`
	Description      string `json:"description"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Number           string `json:"number"`
	DateOfBirth      int64  `json:"dateOfBirth"`
	DateOfIssue      int64  `json:"dateOfIssue"`
	DateOfExpiry     int64  `json:"dateOfExpiry"`
	IssuingCountry   string `json:"issuingCountry"`
	IssuingAuthority string `json:"issuingAuthority"`
	IssuingPlace     string `json:"issuingPlace"`
	Address          string `json:"address"`
}

func (c *Client) Documents() ([]*Document, error) {
	resp := make([]*Document, 0, 10)
	r, err := c.Get("/documents")
	if err != nil {
		return nil, err
	}
	err = c.parseResponse(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Document(id string) (*Document, error) {
	resp := &Document{}
	r, err := c.Get(fmt.Sprintf("/documents/%s", id))
	if err != nil {
		return nil, err
	}
	err = c.parseResponse(r, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) DocumentFile(id string) (io.ReadCloser, error) {
	r, err := c.Get(fmt.Sprintf("/documents/%s/file", id))
	if err != nil {
		return nil, err
	}
	return r.Body, nil
}

func (c *Client) DocumentFrontFile(id string) (io.ReadCloser, error) {
	return c.documentFrontBackFile(id, "front")
}

func (c *Client) DocumentBackFile(id string) (io.ReadCloser, error) {
	return c.documentFrontBackFile(id, "back")
}

func (c *Client) documentFrontBackFile(id string, documentSide string) (io.ReadCloser, error) {
	switch documentSide {
	case "front":
	case "back":
	default:
		return nil, fmt.Errorf(
			"document side must be \"front\" or \"back\", got \"%s\"",
			documentSide,
		)
	}
	resp, err := c.Get(fmt.Sprintf("/documents/%s/%s/file", id, documentSide))
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		return resp.Body, nil
	case http.StatusNoContent:
		return nil, ErrFileNotUploaded
	default:
		return nil, responseToError(resp)
	}
}
