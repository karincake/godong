# Cors Manager Wrapper

This is a simple wrapper for managing cors.

Pre-requisites:
- pacakge apem (github.com/karincake/apem)

What to do:
- Make sure to use apem as config manager and app start
- Set the yaml config usigg and create the `corsCfg` consisting of
  - `allowedOrigins`
  - `allowedMethod`
- wrap `http.NewServeMux()` instance with `SetLog()`.
