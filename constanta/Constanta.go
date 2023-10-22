package constanta

import "time"

// ======== LOG CONSTANTA
const LogLevelNotSet = 0
const LogLevelDebug = 10
const LogLevelInfo = 20
const LogLevelWarn = 30
const LogLevelError = 40
const LogLevelCritical = 50

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
