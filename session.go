package ace

import (
	"github.com/plimble/sessions"
)

//SessionOptions session options
type SessionOptions struct {
	Path   string
	Domain string
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

//Session use session middleware
func (a *Ace) Session(store sessions.Store, options *SessionOptions) {
	var sessionOptions *sessions.Options

	if options != nil {
		sessionOptions = &sessions.Options{
			Path:     options.Path,
			Domain:   options.Domain,
			MaxAge:   options.MaxAge,
			Secure:   options.Secure,
			HttpOnly: options.HTTPOnly,
		}
	}

	manager := sessions.New(10000, store, sessionOptions)

	a.Use(func(c *C) {
		c.Sessions = manager.GetSessions(c.Request)
		defer manager.Close(c.Sessions)

		c.Writer.Before(func(ResponseWriter) {
			c.Sessions.Save(c.Writer)
		})

		c.Next()
	})
}
