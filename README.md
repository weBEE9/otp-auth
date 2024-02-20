# OTP Auth

Backend of OTP(one time password) auth service

## APIs

- Generating OTP
```sh
curl  -X POST \
  '{your_host}/api/v1/otp' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "phone_number": "{given_phone_number}"
}'
```

- Verifying OTP
```sh
curl  -X POST \
  '{your_host}/api/v1/otp/verify' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "phone_number": "{given_phone_number}",
  "otp": "{given_OTP}"
}'
```

### Rate Limiting

APIs above will applied a rate limiting middleware as a basic security method, you can use env variable `APP_RATELIMITER_RATE_PER_MIN` and `APP_RATELIMITER_BUCKET_SIZE`
to configure the middleware.
See `.env.example` file for more information.

## TODOs
- [x] TTL(time to live) of OTP