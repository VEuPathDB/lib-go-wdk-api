package wdk

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"github.com/Foxcapades/Go-ChainRequest"
	"github.com/Foxcapades/Go-ChainRequest/simple"
)

var (
	unmarshaler = simple.UnmarshallerFunc(json.Unmarshal)
	marshaler   = simple.MarshallerFunc(json.Marshal)
)

func prepGet(url string, props *apiProps) creq.Request {
	return prepCreq(simple.GetRequest(url), props)
}

func prepPost(url string, props *apiProps) creq.Request {
	return prepCreq(simple.PostRequest(url), props)
}

func prepCreq(req creq.Request, props *apiProps) creq.Request {
	defer func() {
		logger.WithFields(log.Fields{
			"url":           req.GetUrl(),
			"method":        req.GetMethod(),
			"authToken":     props.authToken,
			"shareSessions": props.oneSession,
			"sessionId":     props.sessionId,
			"cookies":       req.GetCookies(),
		}).Debug("Prepared HTTP request")
	}()

	if props.authToken != "" {
		req.AddCookie(&http.Cookie{Name: cookieAuthToken, Value: props.authToken})
	}

	if props.oneSession {
		req.AddCookie(&http.Cookie{Name: cookieSessionId, Value: props.sessionId})
	}

	return req
}

func getSessionId(url string, props apiProps) string {
	props.oneSession = false
	res := prepGet(url, &props).
		Submit()

	cookie := res.MustGetCookie(cookieSessionId)

	if cookie == nil {
		hmm := res.MustGetHeader("Set-Cookie")
		pos := strings.Index(hmm, cookieSessionId)
		if pos == -1 {
			logger.Panic("could not retrieve session id cookie")
		}

		return hmm[pos+len(cookieSessionId)+1 : strings.IndexByte(hmm[pos:], ';')]
	}

	return cookie.Value
}
