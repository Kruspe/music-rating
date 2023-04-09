package api_test_helper

import (
	"io"
	"net/http"
)

const TestToken = "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJNZXRhbExvdmVyNjY2IiwiaWF0IjoxNTE2MjM5MDIyfQ.JZ3R_3it-1K9ttH5NA80fpIsBhnW6DNsIzwB2zEFRmo7hgE-HhW3jJbArXNS0fw2Pcj-xrU-DMF8KoLr8_EJh2XdTDjaRqz859p0RJ1gPLovsQ8N1HeqeQXKi2mwDJe2rZhWILHdWZ1zmduCY5fF8jUYyBIqLRh1B44L_CBlgeEejKoJfw7V3WoZhxdLeW8SlS2PQ7kN0XIyzm-_TPq1j5QnNHRnXRIh8V7o9rBtdM7PVGEFTpzb1jC6bZ3W-aHuZEWk5e1kRTV8IOXiLf-xtPQ42Hn4j2F27mDg0h2PsgVWmNjr2eqc9y0izps-rmoXHnzmBzvbtGS2yytEFw_WAA"

func NewAuthenticatedRequest(method, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	request.Header.Set("authorization", TestToken)
	return request
}
