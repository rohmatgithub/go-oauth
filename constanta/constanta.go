package constanta

import "time"

// ======== LOG CONSTANTA
const LogLevelTrace = 0
const LogLevelDebug = 1
const LogLevelInfo = 2
const LogLevelWarn = 3
const LogLevelError = 4
const LogLevelFata = 5
const LogLevelPanic = 6

// --------------------------- Header Request Constanta ------------------------------------
const RequestIDConstanta = "X-Request-ID"
const IPAddressConstanta = "X-Forwarded-For"
const SourceConstanta = "X-Source"
const TokenHeaderNameConstanta = "Authorization"
const ApplicationContextConstanta = "application_context"

// --------------------------------- Expired Time Constanta ---------------------------------------------------------
const ExpiredAuthCodeConstanta = 10 * time.Minute
const ExpiredJWTCodeConstanta = 12 * time.Hour
const TimeLockOutConstanta = 5 * time.Minute
const DefaultApplicationsLanguage = "en-US"
