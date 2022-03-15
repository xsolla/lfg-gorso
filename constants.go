package gorso

import "time"

const DEFAULT_TIMEOUT = 10 * time.Second

// Scope is a requested data scope
type Scope string

const (
	// OpenID grants authorization flow
	ScopeOpenID Scope = "openid"
	// CPID grants LoL & LoR info, such as active LoL region
	ScopeCPID Scope = "cpid"
	// Allows refresh tokens to be used to retrieve new AccessTokens
	// that have access to the /userinfo endpoint
	ScopeOfflineAccess Scope = "offline_access"
	// Yet unknown
	ScopeAccount Scope = "account"
	// Returns the email of the account
	// Seems not to work with response_type code
	ScopeEmail Scope = "email"
	// Yet Unknown
	ScopeProfile Scope = "profile"
)

// Scope is a requested data scope
type TokenType string

const (
	// Bearer means the entire token should be provided
	TokenTypeBearer TokenType = "Bearer"
)
