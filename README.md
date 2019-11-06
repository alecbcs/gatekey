# GateKey
GateKey is a simple Go web server which generates and authenticates one time passwords for external services.

## Outline:
### Structure:
- Database
- Auth

### Function:
- Accept Generate Key Request from Builder & Generate One Time Auth Token for Builder VM.
- Pass token back to Builder.
- Accept https requests from builder VMs with one time tokens.
- If token is valid pass finished command to builder.
- Remove token from database.
